package cbor

import (
	"bytes"
	"fmt"
	"testing"

	"go.mozilla.org/cose"
)

func TestCborEncode3(t *testing.T) {
	// Read origin file
	b := new(bytes.Buffer)
	err := getImageBytes("./on.png", b)
	if err != nil {
		fmt.Println(err)
		return
	}
	origin := b.Bytes()

	//image := &Image{Name:"on",Content:origin}
	fmt.Println("Size: ", len(origin))
	// Encode
	fmt.Println("--------- Encode ---------")

	encoded, error := cose.Marshal(origin)

	if error != nil {
		fmt.Errorf("Error decoding %s", error)
	} else {
		fmt.Println("Size: ", len(encoded))
		//fmt.Println("Content: ",string(encoded))
	}

	// Decode
	fmt.Println("--------- Decode ---------")
	//var img Image
	unmarshal, err := cose.Unmarshal(encoded)
	if err != nil {
		fmt.Printf("Error Unmarshal %s", err)
		return
	} else {
		fmt.Println("Size: ", len(unmarshal.([]byte)))
		//fmt.Println("Content: ",unmarshal.([]byte))
	}
	writeToFile("3", unmarshal.([]byte))

}
