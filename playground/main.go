package main

func main() {

}

type Dog struct {
	Name string
}

func (d *Dog) GetName() string {
	return d.Name
}
