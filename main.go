package main

import (
	"fmt"
	"github.com/brankad/api"
	"github.com/julienschmidt/httprouter"
	"github.com/kardianos/service"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	serviceIsRunning bool
	programIsRunning bool
	writingSync    sync.Mutex
)

const serviceName = "Distance public transport"
const serviceDescription = "Simple service for TX"

type program struct{}

func (p program) Start(s service.Service) error {
	fmt.Println(s.String() + " started")
	writingSync.Lock()
	serviceIsRunning = true
	writingSync.Unlock()
	go p.run()
	return nil
}

func (p program) Stop(s service.Service) error {
	writingSync.Lock()
	serviceIsRunning = false
	writingSync.Unlock()
	for programIsRunning {
		fmt.Println(s.String() + " stopping...")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(s.String() + " stopped")
	return nil
}

func (p program) run() {
	router := httprouter.New()
	//timer := sse.New()
	router.ServeFiles("/ui/js/*filepath", http.Dir("ui/js"))
	router.ServeFiles("/ui/css/*filepath", http.Dir("ui/css"))
	router.GET("/", serveHomepage)

	router.POST("/v1/get_distance", api.GetDistance)

	//router.Handler("GET", "/v1/time", timer)
	//go streamTime(timer)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Problem starting web server: " + err.Error())
		os.Exit(-1)
	}
}

func main() {
	serviceConfig := &service.Config{
		Name:        serviceName,
		DisplayName: serviceName,
		Description: serviceDescription,
	}
	prg := &program{}
	s, err := service.New(prg, serviceConfig)
	if err != nil {
		fmt.Println("Cannot create the service: " + err.Error())
	}
	err = s.Run()
	if err != nil {
		fmt.Println("Cannot start the service: " + err.Error())
	}
}
