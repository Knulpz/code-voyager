package concurrency

import "fmt"

// PrintNumAndAlphabet 交替打印数字和字母
func PrintNumAndAlphabet() {
	number, letter, done := make(chan bool), make(chan bool), make(chan bool)

	// number
	go func() {
		i := 1
		for {
			if <-number {
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()

	// letter
	go func() {
		i := 'A'
		for {
			if <-letter {
				fmt.Print(string(i))
				if i == 'Z' {
					done <- true
					return
				}
				i++
				number <- true
			}
		}
	}()

	number <- true
	if <-done {
		return
	}
}
