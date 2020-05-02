package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPostYamlString(t *testing.T) {
	var filePath = "/Users/bruce/Desktop/HVAC-CoolMasterNet.yml"
	var url = "http://localhost:48081/api/v1/deviceprofile/upload"

	// read file to byte
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(yamlFile))

	// create http post request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(yamlFile))
	if err != nil {
		t.Fatal(err)
	}

	// submit request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// check response
	fmt.Println("== upload finish ==")
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.StatusCode)
	fmt.Println(res.Header)
	res.Body.Close()
	fmt.Println(string(resBody))
}
