package main

import (
	"fmt"
	"lab_02/mathutils"
)

func main() {

	min := mathutils.FindMin(10.5, -2.3, 4.7)
	fmt.Printf("Мінімальне значення: %.2f\n", min)

	avg := mathutils.CalculateAverage(10.5, -2.3, 4.7)
	fmt.Printf("Середнє значення: %.2f\n", avg)
	
	x := mathutils.SolveLinearEquation(2.0, -4.0) // 2x - 4 = 0
	fmt.Printf("Розв'язок рівняння: x = %.2f\n", x)
}
