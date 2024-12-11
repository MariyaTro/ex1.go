package main

import "fmt"

func main() {
	// Ініціалізація змінних
	var userinit8 uint8 = 1
	var userinit16 uint16 = 2
	var userinit64 int64 = -3
	var userautoinit = -4 // Такий варіант ініціалізації також можливий

	fmt.Println("Values: ", userinit8, userinit16, userinit64, userautoinit, "\n")

	// Короткий запис оголошення змінної
	// тільки для нових змінних
	intVar := 10

	fmt.Printf("Value = %d Type = %T\n", intVar, intVar)

	// Завдання.
	// 1. Вивести типи всіх змінних
	// 2. Присвоїти змінній intVar змінні userinit16 і userautoinit. Результат вивести в консоль.

	// Виконання

	// 1. Вивести типи всіх змінних
	fmt.Printf("userinit8 type: %T\n", userinit8)
	fmt.Printf("userinit16 type: %T\n", userinit16)
	fmt.Printf("userinit64 type: %T\n", userinit64)
	fmt.Printf("userautoinit type: %T\n", userautoinit)
	fmt.Printf("intVar type: %T\n", intVar)

	// 2. Присвоїти змінній intVar змінні userinit16 і userautoinit.
	intVar = int(userinit16)
	fmt.Printf("intVar after userinit16: %d\n", intVar)
	intVar = userautoinit
	fmt.Printf("intVar after userautoinit: %d\n", intVar)
}
