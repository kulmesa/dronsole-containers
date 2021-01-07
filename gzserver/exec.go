package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func logPipe(logger *log.Logger, pipe io.ReadCloser) {
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		logger.Print(scanner.Text())
	}
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
	go logPipe(log.New(os.Stdout, logPrefix, log.LstdFlags), stdout)
	go logPipe(log.New(os.Stderr, logPrefix, log.LstdFlags), stderr)
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	return cmd, nil
}
