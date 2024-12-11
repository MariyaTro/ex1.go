package main

import "fmt"

func main() {
	variable8 := int8(127)
	variable16 := int16(16383)

	fmt.Println("Приведення типів\n")

	fmt.Printf("variable8         = %-5d = %.16b\n", variable8, variable8)
	fmt.Printf("variable16        = %-5d = %.16b\n", variable16, variable16)
	fmt.Printf("uint16(variable8) = %-5d = %.16b\n", uint16(variable8), uint16(variable8))
	fmt.Printf("uint8(variable16) = %-5d = %.16b\n", uint8(variable16), uint8(variable16))

	// Завдання.
	// 1. Створіть 2 змінні різних типів. Виконайте арифметичні операції. Результат вивести в консоль

	// Виконання

	a := int16(100)
	b := int32(200)

	fmt.Println("\nАрифметичні операції:")
	fmt.Printf("a + int16(b) = %d\n", a+int16(b))
	fmt.Printf("b + int32(a) = %d\n", b+int32(a))
}
