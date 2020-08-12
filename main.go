package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

)

//Payload defines the json request object
type Payload struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
	Proxied bool   `json:"proxied"`
}

//IpAPI defines the json response object for Origin IP api
type IpAPI struct {
	Origin string `json:"origin"`
}

func main() {

	fmt.Println("Starting CloudFlare ip update service")
	response, _ := http.Get("https://httpbin.org/ip")
	var responseIPJson IpAPI
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		print("Error parsing response from service")
	} else {
		json.Unmarshal(body, &responseIPJson)
	}
	url := "https://api.cloudflare.com/client/v4/zones/812b6ad358a70c68b2fba05dc606584b/dns_records/501e2b36c6822aca18f7a02cb1c3c58c"
	payLoad := Payload{
		Type:    "A",
		Name:    "office",
		Content: responseIPJson.Origin,
		TTL:     120,
		Proxied: false,
	}
		updateCloudflareRecord(payLoad, url)
}

func updateCloudflareRecord(payload Payload, url string) {
	jsonPayload, _ := json.Marshal(payload)
	apiKey := "25QWbZEvVMsfWP7OcWVnCL53wmq3C6XKe6F-KnVD"
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonPayload))
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print("something wrong while sending req " + err.Error())
	} else {
		res, _ := ioutil.ReadAll(res.Body)
		print(string(res))
	}


}
