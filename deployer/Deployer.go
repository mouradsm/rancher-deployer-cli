package deployer

import (
	"crypto/tls"
	"log"
	"net/http"
)

func Deploy(InsecureSkipVerify bool) {

	// create key and secret

	// create a http instance and check ssl verify

	// get the cluster an check if exists

	// get the projet and check if exists

	// get namespace and check if it exists

	// get workload

	// check if a service with the same name of the workload exists, if no, create it

	// if worload exists just update or create if not
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: InsecureSkipVerify}

	cluster, err := http.Get("https://localhost/v3")

	log.Println(cluster)

	if err != nil {
		log.Fatal(err)
	}

	if cluster == nil {
		log.Fatal("[FATAL] The cluster does not exist! ")
	}
}
