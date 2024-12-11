package main

import "fmt"

func main() {
	var first, second bool
	var third bool = true
	fourth := !third
	var fifth = true

	fmt.Println("first  = ", first)       // false (значення за замовчуванням для типу bool)
	fmt.Println("second = ", second)      // false (значення за замовчуванням)
	fmt.Println("third  = ", third)       // true (ініціалізоване як true)
	fmt.Println("fourth = ", fourth)      // false (результат логічного заперечення !third)
	fmt.Println("fifth  = ", fifth, "\n") // true (ініціалізоване як true)

	// Оператор `!`  інвертує значення булевої змінної
	fmt.Println("!true  = ", !true)        // false (значення true інвертується у false)
	fmt.Println("!false = ", !false, "\n") // true (значення false інвертується у true)

	/*Оператор `&&` :
	true — якщо обидва операнди true
	false — якщо хоча б один операнд false

	*/
	fmt.Println("true && true   = ", true && true)         // true (обидва операнди true, тому результат true)
	fmt.Println("true && false  = ", true && false)        // false (один з операндів false, тому результат false)
	fmt.Println("false && false = ", false && false, "\n") // false (обидва операнди false, тому результат false)

	/*
		Оператор || (або):
		true — якщо хоча б один true
		false — коли обидва операнди  false
	*/
	fmt.Println("true || true   = ", true || true)         // true (обидва операнди true, тому результат true)
	fmt.Println("true || false  = ", true || false)        // true (один з операндів true, тому результат true)
	fmt.Println("false || false = ", false || false, "\n") // false (обидва операнди false, тому результат false)

	/*
		< — true, якщо ліве значення менше за праве.
		> — true, якщо ліве значення більше за праве.
		<= — true, якщо ліве менше або дорівнює правому.
		>= — true, якщо ліве більше або дорівнює правому.
		== — true, якщо значення рівні.
		!= — true, якщо значення не рівні.
	*/

	fmt.Println("2 < 3  = ", 2 < 3)        // true (2 менше за 3)
	fmt.Println("2 > 3  = ", 2 > 3)        // false (2 не більше за 3)
	fmt.Println("3 < 3  = ", 3 < 3)        // false (3 не менше за 3)
	fmt.Println("3 <= 3 = ", 3 <= 3)       // true (3 дорівнює 3)
	fmt.Println("3 > 3  = ", 3 > 3)        // false (3 не більше за 3)
	fmt.Println("3 >= 3 = ", 3 >= 3)       // true (3 дорівнює 3)
	fmt.Println("2 == 3 = ", 2 == 3)       // false (2 не дорівнює 3)
	fmt.Println("3 == 3 = ", 3 == 3)       // true (3 дорівнює 3)
	fmt.Println("2 != 3 = ", 2 != 3)       // true (2 не дорівнює 3)
	fmt.Println("3 != 3 = ", 3 != 3, "\n") // false (3 дорівнює 3, тому результат false)

	// Завдання.
	// 1. Пояснити результати операцій
}
