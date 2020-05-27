package xml

import (
	"fmt"
	xj "github.com/basgys/goxml2json"
	"github.com/buger/jsonparser"
	"os"
	"testing"
)

type Person struct {
	Name string `xml:"name"`
}

func TestParse(t *testing.T) {
	xmlFile, err := os.Open("2018Q1.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer xmlFile.Close()
	json, err := xj.Convert(xmlFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(json.String())

	//err = jsonparser.ObjectEach(json.Bytes(), objectEach,"xbrl")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	// 營業收入
	var revenue string
	_, err = jsonparser.ArrayEach(json.Bytes(), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		revenue, err = jsonparser.GetString(value, "#content")
		fmt.Println(revenue)
	}, "xbrl", "Revenue")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func parse(b []byte) {
	err := jsonparser.ObjectEach(b, objectEach, "xbrl")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func objectEach(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
	fmt.Printf("[Key]: [%s]\n Value: '%s'\n Type: %s\n", string(key), string(value), dataType)
	return nil
}

func arrayEach(value []byte, dataType jsonparser.ValueType, offset int, err error) {
	fmt.Printf("Array Value: '%s'\n Type: %s\n", string(value), dataType)
}
