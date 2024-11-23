package jobsearch

import (
	"encoding/json"
	"log"
	"my-protobuf/protogen/basic"
	"my-protobuf/protogen/dummy"
	"my-protobuf/protogen/jobsearch"
)

func JobSearchSoftware() {
    js := jobsearch.JobSoftware{
        JobSoftwareId: 777,
        Application: &basic.Application{
            Version: "1.0.0",
            Name: "The Amazing Proto",
            Platforms: []string{"Mac", "Linux", "Windows"},
        },
    }

    jsonBytes, _ := json.Marshal(&js)
    log.Println("Software : ", string(jsonBytes))
}

func JobSearchCancdidate() {
    jc := jobsearch.JobCandidate{
        JobCandidateId: 888,
        Application: &dummy.Application{
            ApplicationId: 887,
            ApplicantFullName: "shazam",
            Phone: "999-999-999",
            Email: "shazam@gmail.com",
        },
    }

    jsonBytes, _ := json.Marshal(&jc)
    log.Println("Candidate : ", string(jsonBytes))
}