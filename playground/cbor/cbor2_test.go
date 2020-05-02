package cbor

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/2tvenom/cbor"
)

type Image struct {
	Name    string
	Content []byte
}

func TestCborEncode2(t *testing.T) {
	// Read origin file
	b := new(bytes.Buffer)
	err := getImageBytes("./on.png", b)
	if err != nil {
		fmt.Println(err)
		return
	}
	origin := b.Bytes()

	//image := &Image{Name:"on",Content:origin}

	// Encode
	fmt.Println("--------- Encode ---------")

	var buf bytes.Buffer
	encoder := cbor.NewEncoder(&buf)
	ok, error := encoder.Marshal(origin)

	//check binary string
	if !ok {
		fmt.Errorf("Error decoding %s", error)
	} else {
		fmt.Println("Size: ", len(buf.Bytes()))
	}

	// Decode
	//fmt.Println("--------- Decode ---------",buf.Bytes())
	//var img []byte
	//ok, err = encoder.Unmarshal(buf.Bytes(), &img)
	//if !ok {
	//	fmt.Printf("Error Unmarshal %s", err)
	//	return
	//}
	////output
	//fmt.Printf("%v", img)

}
