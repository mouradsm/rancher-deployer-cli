package deployer

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

type HttpResponse struct {
	Data string `json:"data"`
}

func Deploy(params map[string]string, insecureSkipVerify bool) {

	// get the project and check if exists

	// get namespace and check if it exists

	// get workload

	// check if a service with the same name of the workload exists, if no, create it

	// if worload exists just update or create if not
	c := httpClient()

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: insecureSkipVerify}

	bearerToken := fmt.Sprintf("Bearer %s:%s", params["rancher_key"], params["rancher_secret"])

	baseUrl := params["rancher_url"] + "/v3"

	clusterUrl := baseUrl + "/clusters?name=" + params["cluster"]

	clusterResponse := sendRequest(c, http.MethodGet, clusterUrl, bearerToken, nil)

	cluster := gjson.Get(string(clusterResponse), "data")
	if len(cluster.Array()) == 0 {
		log.Fatal("[FATAL] The cluster does not exist! ")
	}

	clusterId := gjson.Get(cluster.Array()[0].String(), "id")

	projectUrl := baseUrl + "/projects?name=" + params["project"]

	projectResponse := sendRequest(c, http.MethodGet, projectUrl, bearerToken, nil)

	project := gjson.Get(string(projectResponse), "data.0")
	if len(project.Array()) == 0 {
		log.Fatal("[FATAL] The project does not exist!")
	}

	projectId := gjson.Get(project.String(), "id")

	namespaceUrl := fmt.Sprintf("%v/clusters/%v/namespaces?name=%v", baseUrl, clusterId, params["namespace"])

	namespaceResponse := sendRequest(c, http.MethodGet, namespaceUrl, bearerToken, nil)

	namespace := gjson.Get(string(namespaceResponse), "data")

	if len(namespace.Array()) == 0 {
		log.Fatal("[FATAL] The namespace does not exist!")
	}

	workloadUrl := fmt.Sprintf("%v/projects/%v/workloads", baseUrl, projectId)
	//serviceUrl := fmt.Sprintf("%v/projects/%v/workloads", baseUrl, projectId)

	workloadResponse := sendRequest(c, http.MethodGet, workloadUrl, bearerToken, nil)
	targetWorkload := gjson.Get(string(workloadResponse), "data.#(name=="+params["deployment"]+")")

	workloadHttpMethod := http.MethodPost

	if targetWorkload.Type == gjson.Null {
		workloadHttpMethod = http.MethodPut
	}

	// /workloadRequesBody := map[string]string{}

	sendRequest(c, workloadHttpMethod, workloadUrl, bearerToken, make(map[string]string))

}

func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}

	return client
}

func sendRequest(client *http.Client, method string, url string, authorizationToken string, requestBody map[string]string) []byte {

	endpoint := url

	jsonData, err := json.Marshal(requestBody)

	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

	request, err := http.NewRequest(method, endpoint, bytes.NewBuffer(jsonData))
	request.Header.Add("Authorization", authorizationToken)

	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}

	return body

}
