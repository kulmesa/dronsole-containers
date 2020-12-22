package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"syscall"
)

func logPipe(pre string, pipe io.ReadCloser) {
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		log.Printf("%s: %s", pre, scanner.Text())
	}
	log.Printf("%s: --DONE--", pre)
}

func startCommandWithLogging(logPrefix string, name string, arg ...string) (*exec.Cmd, error) {
	cmd := exec.Command(name, arg...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}
	go logPipe(fmt.Sprintf("%s out", logPrefix), stdout)
	go logPipe(fmt.Sprintf("%s err", logPrefix), stderr)
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	return cmd, nil
}
