package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"k8s.io/utils/env"
)

var envVars = map[string]bool{"CONNECTION_STRING": true, "DB_USER": true, "DB_PASS": true, "DB_QA_USER": true}
var totalCount int = 0

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Radix app</h1>")
	count, _ := env.GetInt("COUNT", 1200)
	for i := 0; i < count; i++ {
		totalCount += 1
		log.Infof("Log line #%d ABCDEFGHIKLMNOPQRSTVXYZ ABCDEFGHIKLMNOPQRSTVXYZ ABCDEFGHIKLMNOPQRSTVXYZ ABCDEFGHIKLMNOPQRSTVXYZ ABCDEFGHIKLMNOPQRSTVXYZ ABCDEFGHIKLMNOPQRSTVXYZ ABCDEFGHIKLMNOPQRSTVXYZ",
			totalCount)
	}
}

func main() {
	fmt.Println("Start server")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8082", nil))
	fmt.Println("Stopped server")
}
