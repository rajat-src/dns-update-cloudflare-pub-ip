package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Payload struct{
	Type string `json:"type"`
	Name string `json:"name"`
	Content string `json:"content"`
	Ttl int `json:"ttl"`
	Proxied bool `json:"proxied"`
}

func main(){
	fmt.Println("Starting cloudflare ip update service")
	apiKey :="25QWbZEvVMsfWP7OcWVnCL53wmq3C6XKe6F-KnVD"
	response, err := http.Get("https://httpbin.org/ip")
    if err != nil {
        fmt.Printf("Error while retriving the ISP public IP %s\n", err)
	}
	url := "https://api.cloudflare.com/client/v4/zones/812b6ad358a70c68b2fba05dc606584b/dns_records/bbcd840c37c1f359bc41a2069ecce975"
	method := "PUT"
	pay := Payload{
		type: "A",
		name: "test",
		content: "1.1.1.1",
		ttl: 120,
		proxied: false
	}
	
	
}