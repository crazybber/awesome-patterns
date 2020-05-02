package cbor

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"testing"

	"github.com/ugorji/go/codec"
)

func TestCborEncode(t *testing.T) {
	// Read origin file
	buf := new(bytes.Buffer)
	err := getImageBytes("./on.png", buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	origin := buf.Bytes()

	fmt.Println("size:", len(origin))
	writeToFile("1", origin)

	// Encode
	encoded, err := encodeBinaryValue(buf.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("size:", len(encoded))
	writeToFile("2", encoded)

	// Decode
	decoded, err := decodeBinaryValue(encoded)
	fmt.Println("size:", len(decoded))
	writeToFile("3", decoded)
}

func encodeBinaryValue(b []byte) ([]byte, error) {
	var ch codec.CborHandle
	buf := new(bytes.Buffer)

	enc := codec.NewEncoder(buf, &ch)
	err := enc.Encode(&b)

	return buf.Bytes(), err
}

func decodeBinaryValue(b []byte) ([]byte, error) {
	var ch codec.CborHandle
	var decoded []byte
	var bufReader = bufio.NewReader(bytes.NewReader(b))
	var dec = codec.NewDecoder(bufReader, &ch)
	var err = dec.Decode(&decoded)
	return decoded, err
}

func getImageBytes(imgFile string, buf *bytes.Buffer) error {
	// Read existing image from file
	img, err := os.Open(imgFile)
	if err != nil {
		return err
	}
	defer img.Close()

	// TODO: Attach MediaType property, determine if decoding
	//  early is required (to optimize edge processing)

	// Expect "png" or "jpeg" image type
	imageData, imageType, err := image.Decode(img)
	if err != nil {
		return err
	}
	// Finished with file. Reset file pointer
	img.Seek(0, 0)
	if imageType == "jpeg" {
		err = jpeg.Encode(buf, imageData, nil)
		if err != nil {
			return err
		}
	} else if imageType == "png" {
		err = png.Encode(buf, imageData)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeToFile(fileName string, b []byte) {
	f, _ := os.Create(fileName)
	n2, err := f.Write(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("wrote file %d bytes\n", n2)
}
