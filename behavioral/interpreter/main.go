/**
 * @author  : Jagepard <jagepard@yandex.ru>
 * @license https://mit-license.org/ MIT
 */

package main

func main() {
	interpreter := Interpreter{registry: make(map[int]Album)}

	interpreter.addAlbumToRegistry(Album{name: "Untouchables", author: "Korn"})
	interpreter.addAlbumToRegistry(Album{name: "Adrenaline", author: "Deftones"})
	interpreter.interpret("album 2")
	interpreter.interpret("album author 2")
	interpreter.interpret("album author 1")
	interpreter.interpret("author 1")
}
