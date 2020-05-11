/**
 * @author  : Jagepard <jagepard@yandex.ru>
 * @license https://mit-license.org/ MIT
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Interpreter is ...
type Interpreter struct {
	registry map[int]Album
}

var i int = 1

func (interpreter *Interpreter) addAlbumToRegistry(album Album) {
	interpreter.registry[i] = album
	i++
}

func (interpreter *Interpreter) interpret(input string) {
	var exploded = strings.Split(input, " ")

	for i := range exploded {
		if interpreter.isNumeric(exploded[i]) {
			number, err := strconv.Atoi(exploded[i])
			if err != nil {
				fmt.Println(err)
			}

			interpreter.getDataFromRegistry(exploded, interpreter.registry[number])
		}
	}
}

func (interpreter *Interpreter) getDataFromRegistry(exploded []string, album Album) {
	var output strings.Builder

	for i := range exploded {
		if exploded[i] == "author" {
			output.WriteString(album.author)
			output.WriteString(" ")
		}

		if exploded[i] == "album" {
			output.WriteString(album.name)
			output.WriteString(" ")
		}
	}

	fmt.Println(output.String())
}

func (interpreter *Interpreter) isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
