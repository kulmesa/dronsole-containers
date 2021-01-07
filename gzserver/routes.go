package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
)

func registerRoutes(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, "/simulation/start", startSimulationHandler)
	router.HandlerFunc(http.MethodGet, "/simulation/stop", stopSimulationHandler)

	router.HandlerFunc(http.MethodGet, "/simulation/drones", listDronesHandler)
	router.HandlerFunc(http.MethodPost, "/simulation/drones", createDroneHandler)
	router.HandlerFunc(http.MethodDelete, "/simulation/drones/:id", deleteDroneHandler)
}

type Drone struct {
	Location string
}

var (
	gzserverCmd *exec.Cmd
	drones      map[string]*Drone = make(map[string]*Drone)
)

func startSimulationHandler(w http.ResponseWriter, r *http.Request) {
	if gzserverCmd != nil {
		log.Printf("Simulation already running")
		http.Error(w, "Simulation already running", http.StatusBadRequest)
		return
	}

	var err error
	gzserverCmd, err = startCommandWithLogging("gzserver", "bash", "-c", "/gzserver-api/scripts/launch-gzserver.sh")
	if err != nil {
		log.Fatal(err)
	}
}
func stopSimulationHandler(w http.ResponseWriter, r *http.Request) {
	if gzserverCmd == nil {
		log.Printf("Simulation not running")
		http.Error(w, "Simulation not running", http.StatusBadRequest)
		return
	}
	syscall.Kill(-gzserverCmd.Process.Pid, syscall.SIGKILL)
	log.Printf("gzserver killed")
	gzserverCmd = nil
}

func listDronesHandler(w http.ResponseWriter, r *http.Request) {
	type drone struct {
		DeviceID      string `json:"device_id"`
		DroneLocation string `json:"drone_location"`
	}

	droneList := make([]drone, 0)
	for id, d := range drones {
		droneList = append(droneList, drone{
			DeviceID:      id,
			DroneLocation: d.Location,
		})
	}

	writeJSON(w, droneList)
}

func createDroneHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		DroneLocation  string `json:"drone_location"`
		DeviceID       string `json:"device_id"`
		MAVLinkAddress string `json:"mavlink_address"`
		MAVLinkUDPPort int32  `json:"mavlink_udp_port"`
		MAVLinkTCPPort int32  `json:"mavlink_tcp_port"`
		PosX           int32  `json:"pos_x"`
		PosY           int32  `json:"pos_y"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Printf("Could not decode body: %v", err)
		http.Error(w, "Malformatted body", http.StatusBadRequest)
		return
	}

	_, ok := drones[requestBody.DeviceID]
	if ok {
		log.Printf("Request to add drone with device id already in use")
		http.Error(w, "DeviceID already in use", http.StatusBadRequest)
		return
	}

	ips, err := net.LookupIP(requestBody.MAVLinkAddress)
	if err != nil {
		log.Printf("Could not lookup mavlink IP '%s': %v", requestBody.MAVLinkAddress, err)
		http.Error(w, "Cloud not lookup mavlink IP", http.StatusInternalServerError)
		return
	}
	if len(ips) == 0 {
		log.Printf("Could not lookup mavlink IP '%s': %v", requestBody.MAVLinkAddress, err)
		http.Error(w, "Cloud not lookup mavlink IP", http.StatusInternalServerError)
		return
	}
	command := fmt.Sprintf("/gzserver-api/scripts/spawn-drone.sh %s %d %d %s %d %d",
		ips[0].String(),
		requestBody.MAVLinkUDPPort,
		requestBody.MAVLinkTCPPort,
		requestBody.DeviceID,
		requestBody.PosX,
		requestBody.PosY)

	// add drone model and connect it to the mavlink
	droneSpawnCmd, err := startCommandWithLogging("drone", "bash", "-c", command)
	if err != nil {
		log.Printf("Could not spawn drone model: %v", err)
		http.Error(w, "Could not spawn drone model", http.StatusInternalServerError)
		return
	}

	done := make(chan error)
	go func() { done <- droneSpawnCmd.Wait() }()
	select {
	case <-time.After(20 * time.Second):
		log.Printf("Spawn timeout after 20 seconds")
		droneSpawnCmd.Process.Kill()
	case err := <-done:
		if err != nil {
			log.Printf("Spawn failed: %v", err)
			http.Error(w, "Spawn failed", http.StatusInternalServerError)
		}
	}

	drones[requestBody.DeviceID] = &Drone{
		Location: requestBody.DroneLocation,
	}
}
func deleteDroneHandler(w http.ResponseWriter, r *http.Request) {
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		log.Printf("Could not marshal data to json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
