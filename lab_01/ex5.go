package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Синоніми цілих типів\n")

	fmt.Println("byte    - int8")
	fmt.Println("rune    - int32")
	fmt.Println("int     - int32 або int64, в залежності від платформи")
	fmt.Println("uint    - uint32 або uint64, в залежності від платформи")

	// Завдання.

	// Виконання
	arch := runtime.GOARCH
	if arch == "amd64" {
		fmt.Println("Платформа 64-розрядна")
	} else if arch == "386" {
		fmt.Println("Платформа 32-розрядна")
	} else {
		fmt.Println("Не вдалося визначити розрядність платформи.")
	}
}
