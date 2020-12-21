package main

import (
	"log"

	"github.com/VolantMQ/vlapi/vlauth"
	"github.com/VolantMQ/volantmq/auth"
	//jwt "github.com/dgrijalva/jwt-go"
)

// internalAuth allows all connections
type internalAuth struct{}

func (a *internalAuth) Password(clientid, user, pass string) error {
	return vlauth.StatusAllow
}
func (a *internalAuth) ACL(clientId, user, topic string, access vlauth.AccessType) error {
	return vlauth.StatusAllow
}
func (a *internalAuth) Shutdown() error {
	return nil
}

func RegisterAuthManagers() {
	err := auth.Register("internal", &internalAuth{})
	if err != nil {
		log.Fatalf("Could not register internal auth: %v", err)
	}
	//err = auth.Register("gcp", &gcpAuth{})
	//if err != nil {
	//	log.Fatalf("Could not register gcp auth: %v", err)
	//}
}

func NewInternalAuthManager() *auth.Manager {
	internalAuthManager, err := auth.NewManager([]string{"internal"}, false)
	if err != nil {
		log.Fatalf("Could not create internal auth manager: %v", err)
	}

	return internalAuthManager
}

/*
func NewGCPAuth() *auth.Manager {
	gcpAuthManager, err := auth.NewManager([]string{"gcp"}, false)
	if err != nil {
		log.Fatal("Could not create gcp auth manager: %v", err)
	}

	return gcpAuthManager
}

type gcpAuth struct{}

func (a *gcpAuth) Password(clientid, user, pass string) error {
	if user != "unused" {
		return vlauth.StatusDeny
	}

	if clientid != "projects/auto-fleet-mgnt/locations/europe-west1/registries/fleet-registry/devices/etest" {
		log.Printf("Unknown clientid: %v", clientid)
		return vlauth.StatusDeny
	}

	keyData, err := ioutil.ReadFile("rsa_cert.pem")
	if err != nil {
		log.Printf("Could not read cert file: %v", err)
		return err
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		log.Printf("Could not parse cert file: %v", err)
		return err
	}

	token, err := jwt.Parse(pass, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != "RS256" {
			log.Printf("Invalid signing method: %v", token.Method.Alg())
			return nil, errors.New("auth: invalid signing method")
		}
		return key, nil
	})
	if err != nil {
		log.Print(err)
		return vlauth.StatusDeny
	}

	claims := token.Claims.(jwt.MapClaims)
	if !claims.VerifyAudience("auto-fleet-mgnt", true) {
		log.Print("Invalid audience")
		return vlauth.StatusDeny
	}

	t := time.Now()
	if !claims.VerifyExpiresAt(t.Unix(), true) {
		log.Print("Invalid expires at")
		return vlauth.StatusDeny
	}
	if !claims.VerifyIssuedAt(t.Unix(), true) {
		log.Print("Invalid issued at")
		return vlauth.StatusDeny
	}
	log.Print("Allowing")

	return vlauth.StatusAllow
}
func (a *gcpAuth) ACL(clientId, user, topic string, access vlauth.AccessType) error {
	//expClientId := fmt.Sprintf(
	//	"projects/%s/locations/%s/registries/%s/devices/%s",
	//	*project, *region, *registry, *device)

	//log.Printf("Checking ACL: %v %v %v %v", clientId, user, topic, access)
	//if user == "unused" &&
	//	topic == fmt.Sprintf("/devices/%s/%s", *device, topicType) &&
	//	clientId == expClientId {
	//	return auth.StatusAllow
	//}

	//panic("ACL failed")
	return vlauth.StatusDeny
}
func (a *gcpAuth) Shutdown() error {
	return nil
}
*/
