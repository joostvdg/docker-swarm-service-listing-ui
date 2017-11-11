package main

import (
	"./api"
	"./webserver"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("=============================================")
	fmt.Println("=============================================")
	fmt.Println("=== Docker Swarm Service Lister UI 1.0.5 ====")
	fmt.Println("=============================================")

	apiProtocol := "http"
	if len(os.Getenv("API_PROTOCOL")) > 0 {
		apiProtocol = os.Getenv("API_PROTOCOL")
	}

	apiHost := "127.0.0.1"
	if len(os.Getenv("API_HOST")) > 0 {
		apiHost = os.Getenv("API_HOST")
	}

	apiPort := "7777"
	if len(os.Getenv("API_PORT")) > 0 {
		apiPort = os.Getenv("API_PORT")
	}

	apiUrl := fmt.Sprintf("%s://%s:%s/stacks", apiProtocol, apiHost, apiPort)
	serverPort := "8087"
	if len(os.Getenv("SERVER_PORT")) > 0 {
		serverPort = os.Getenv("SERVER_PORT")
	}

	fmt.Printf("=== POLLING API @%s\n", apiUrl)
	fmt.Printf("=== STARTING WEB SERVER @%s\n", serverPort)
	fmt.Println("=============================================")

	stacks := api.GetStacks(apiUrl)
	webserverData := &webserver.WebserverData{Stacks: stacks, Title: "Service Listing"}

	c := make(chan bool)
	go webserver.StartServer(serverPort, webserverData, c)
	fmt.Println("> Started the web server, now polling swarm")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	for i := 1; ; i++ { // this is still infinite
		t := time.NewTicker(time.Second * 30)
		select {
		case <-stop:
			fmt.Println("> Shutting down polling")
			break
		case <-t.C:
			fmt.Println("  > Updating Stacks")
			webserverData.UpdateStacks(api.GetStacks(apiUrl))
			continue
		}
		break // only reached if the quitCh case happens
	}
	fmt.Println("> Shutting down webserver")
	c <- true
	if b := <-c; b {
		fmt.Println("> Webserver shut down")
	}
	fmt.Println("> Shut down app")
}
