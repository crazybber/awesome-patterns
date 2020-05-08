package main

import "testing"
import "sync"

func TestStringorError_Execute(t *testing.T) {
	future := &MaybeString{}
	t.Run("Success Result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)
		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})
		future.Execute(func() (string, error) {
			return "Hello World", nil
		})
		wg.Wait()
	})
	t.Run("Failed Result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)
		future.Success(func(s string) {
			t.Fail()
			wg.Done()

		}).Fail(func(e error) {
			t.Log(e.Error())
			wg.Done()
		})
		future.Execute(func() (string, error) {
			return "", nil
		})
		wg.Wait()
	})
}
