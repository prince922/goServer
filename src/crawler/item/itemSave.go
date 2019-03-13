package item

import "fmt"

func ItemSave() chan interface{} {
	out := make(chan interface{})

	go func() {
		for {
			item := <-out

			fmt.Println(item)
		}

	}()

	return out
}
