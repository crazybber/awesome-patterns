package parse

import (
	"fmt"
	"strconv"
	"testing"
)

func TestParseInt(t *testing.T) {
	var _, err = strconv.ParseInt("127", 0, 8)
	if err != nil {
		fmt.Println(err)
	}
}
