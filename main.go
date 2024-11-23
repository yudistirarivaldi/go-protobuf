package main

import (
	"fmt"
	"log"
	"my-protobuf/jobsearch"
	"time"
)

type logWriter struct{}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05" + " " + string(bytes)))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	// basic.BasicHello()
	// basic.BasicUser()
	// basic.ProtoToJsonUser()
	// basic.JsonToProtoUser()
	// basic.BasicUserGroup()
	jobsearch.JobSearchCancdidate()
	jobsearch.JobSearchCancdidate()
}