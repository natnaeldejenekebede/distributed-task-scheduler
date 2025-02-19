package main

import (
	"flag"
	"log"

	"github.com/natnaeldejenekebede/distributed-task-scheduler/pkg/common"
	"github.com/natnaeldejenekebede/distributed-task-scheduler/pkg/scheduler"
)

var (
	schedulerPort = flag.String("scheduler_port", ":8081", "Port on which the Scheduler serves requests.")
)

func main() {
	dbConnectionString := common.GetDBConnectionString()
	schedulerServer := scheduler.NewServer(*schedulerPort, dbConnectionString)
	err := schedulerServer.Start()
	if err != nil {
		log.Fatalf("Error while starting server: %+v", err)
	}
}
