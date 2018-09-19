package http

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"testing"
)

func TestUploadFormFile(t *testing.T) {
	var filePath = "/Users/bruce/Desktop/HVAC-CoolMasterNet.yml"
	var url = "http://localhost:48081/api/v1/deviceprofile/uploadfile"

	// Retch file
	fmt.Println("Read file: ", filepath.Base(filePath))
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}

	// create form data
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	formFile, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		t.Fatal(err)
	}
	if _, err = io.Copy(formFile, bytes.NewReader(yamlFile)); err != nil {
		t.Fatal(err)
	}
	writer.Close()

	// create http post request
	req, err := http.NewRequest(http.MethodPost, url, &body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

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
