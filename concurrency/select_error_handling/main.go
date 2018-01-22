package main

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	if err := selectWithErrorDemo(); err != nil {
		spew.Dump(err)
	}
	time.Sleep(time.Second * 5)
}

func selectWithErrorDemo() error {
	// start update race and market
	count, errChan, doneChan := 2, make(chan error), make(chan bool)

	go func() {
		// update race
		if err := doFirstThing(); err != nil {
			errChan <- err
		} else {
			doneChan <- true
			fmt.Println("Finish First Go Routine")
		}

	}()

	go func() {
		// update market
		if err := doSecondThing(); err != nil {
			errChan <- err
		} else {
			doneChan <- true
			fmt.Println("Finish Second Go Routine")
		}
	}()

	for i := 0; i < count; i++ {
		select {
		case err := <-errChan:
			return err
		case d := <-doneChan:
			spew.Dump(i, d)
		}
	}
	return nil
}

func doFirstThing() error {
	time.Sleep(time.Second * 2)
	fmt.Println("Exectue Do First")
	return nil

}

func doSecondThing() error {
	time.Sleep(time.Second * 1)
	fmt.Println("Exectue Do Second")
	// return errors.New("Error In Second")
	return nil
}
