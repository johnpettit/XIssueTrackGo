package main

import (
	"github.com/johnpettit/XIssueTrackGo/restapi"
	"log"
)

func main() {
	log.Print("Johnny is running the XIssueTracker")
	restapi.Start()
}
