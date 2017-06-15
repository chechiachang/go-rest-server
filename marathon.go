package main

import(
	"github.com/gambol99/go-marathon"
	"log"
	"net/http"
	"time"
	"net"
	"crypto/tls"
)

func createClient(){
	marathonURL := "http://192.168.65.111:20950"
	config := marathon.NewDefaultConfig()
	config.URL = marathonURL
	config.HTTPClient = &http.Client{
		Timeout: time.Duration(10) * time.Second,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 10 * time.Second,
				KeepAlive: 10 * time.Second,
			}).Dial,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	client, err := marathon.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create a client for marathon, error: %s", err)
	}

	client.Info()

	//applications, err := client.Applications()
	//if err != nil {
	//	log.Fatalf("Failed to list applications")
	//}
	//
	//log.Printf("Found %d applications running", len(applications.Apps))
	//for _, application := range applications.Apps {
	//	log.Printf("Application: %s", application)
	//	details, err := client.Application(application.ID)
	//	if err != nil{
	//		log.Fatal(err)
	//	}
	//
	//	if details.Tasks != nil && len(details.Tasks) > 0 {
	//		for _, task := range details.Tasks {
	//			log.Printf("task: %s", task)
	//		}
	//		// check the health of the application
	//		health, err := client.ApplicationOK(details.ID)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		log.Printf("Application: %s, healthy: %t", details.ID, health)
	//	}
	//}
}

func createApplication(){
	//log.Printf("Deploying a new application")
	//application := marathon.NewDockerApplication().
	//	Name(applicationName).
	//	CPU(0.1).
	//	Memory(64).
	//	Storage(0.0).
	//	Count(2).
	//	AddArgs("/usr/sbin/apache2ctl", "-D", "FOREGROUND").
	//	AddEnv("NAME", "frontend_http").
	//	AddEnv("SERVICE_80_NAME", "test_http").
	//	CheckHTTP("/health", 10, 5)
	//
	//application.
	//Container.Docker.Container("quay.io/gambol99/apache-php:latest").
	//	Bridged().
	//	Expose(80).
	//	Expose(443)
	//
	//if _, err := client.CreateApplication(application); err != nil {
	//	log.Fatalf("Failed to create application: %s, error: %s", application, err)
	//} else {
	//	log.Printf("Created the application: %s", application)
	//}
}

