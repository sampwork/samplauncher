package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func StartServer() {
	cmd := exec.Command("samp-server.exe")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out.String())

	time.Sleep(200 * time.Millisecond)

	fmt.Println("Server With Port:", GetServerPort(configFile), "started")
}

func UsageText() {
	fmt.Println("Usage: samp-server.exe [ports ...]. (Example: samp-server.exe 7771 7772 7773)")
}
