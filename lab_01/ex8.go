package main

//Импорт нескольких пакетов
import (
	"fmt"
	"math"
)

func main() {
	var defaultFloat float32
	var defaultDouble float64 = 5.5

	fmt.Println("defaultfloat       = ", defaultFloat)
	fmt.Printf("defaultDouble (%T) = %f\n\n", defaultDouble, defaultDouble)

	fmt.Println("MAX float32        = ", math.MaxFloat32)
	fmt.Println("MIN float32        = ", math.SmallestNonzeroFloat32, "\n")

	fmt.Println("MAX float64        = ", math.MaxFloat64)
	fmt.Println("MIN float64        = ", math.SmallestNonzeroFloat64, "\n")

	// Завдання.
	// 1. Створіть змінні різних типів, використовуючи короткий запис та ініціалізацію за замовчуванням. Результат вивести в консоль

	// Виконання
	
	floatVar := float32(3.14)
	doubleVar := float64(2.718)
	defaultInitVar := 1.41

	fmt.Printf("floatVar: %f (type %T)\n", floatVar, floatVar)
	fmt.Printf("doubleVar: %f (type %T)\n", doubleVar, doubleVar)
	fmt.Printf("defaultInitVar: %f (type %T)\n", defaultInitVar, defaultInitVar)
}
