package testmod

import "fmt"

//Приветствие
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s! Как дела?", name)
}
