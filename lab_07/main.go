package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
)

func calculateArrayTasks(arr []float64) (float64, float64) {
	if len(arr) == 0 {
		return 0, 0
	}

	var positiveSum float64 = 0
	for _, num := range arr {
		if num > 0 {
			positiveSum += num
		}
	}

	var maxAbsIndex, minAbsIndex int
	maxAbs := math.Abs(arr[0])
	minAbs := math.Abs(arr[0])

	for i, num := range arr {
		currAbs := math.Abs(num)
		if currAbs > maxAbs {
			maxAbs = currAbs
			maxAbsIndex = i
		}
		if currAbs < minAbs {
			minAbs = currAbs
			minAbsIndex = i
		}
	}

	var productBetween float64 = 1

	if minAbsIndex > maxAbsIndex {
		minAbsIndex, maxAbsIndex = maxAbsIndex, minAbsIndex
	}

	for i := minAbsIndex + 1; i < maxAbsIndex; i++ {
		productBetween *= arr[i]
	}

	return positiveSum, productBetween
}

func parseNumbers(numbersStr string) ([]float64, error) {
	if numbersStr == "" {
		return nil, fmt.Errorf("порожній масив")
	}

	strArr := strings.Split(numbersStr, ",")
	arr := make([]float64, len(strArr))

	for i, strNum := range strArr {
		num, err := strconv.ParseFloat(strings.TrimSpace(strNum), 64)
		if err != nil {
			return nil, fmt.Errorf("некоректне число: %s", strNum)
		}
		arr[i] = num
	}

	return arr, nil
}

func handleArrayCalculations(w http.ResponseWriter, r *http.Request) {
	var arr []float64
	var err error
	var method string

	if r.Method == "GET" {
		method = "GET"
		numbersStr := r.URL.Query().Get("numbers")
		arr, err = parseNumbers(numbersStr)
	}

	if r.Method == "POST" {
		method = "POST"
		err = r.ParseForm()
		if err == nil {
			numbersStr := r.FormValue("numbers")
			arr, err = parseNumbers(numbersStr)
		}
	}

	fmt.Fprintf(w, `<!DOCTYPE html>
    <html>
    <head>
        <title>Обчислення масиву</title>
        <style>
            body { font-family: Arial; max-width: 600px; margin: 0 auto; padding: 20px; }
        </style>
    </head>
    <body>
        <h2>Обчислення масиву</h2>
        <p>Метод: %s</p>
    `, method)

	if err != nil {
		fmt.Fprintf(w, "<p style='color:red'>Помилка: %v</p>", err)
	} else if len(arr) > 0 {
		positiveSum, productBetween := calculateArrayTasks(arr)

		fmt.Fprintf(w, `
            <h3>Результати:</h3>
            <p>Вхідний масив: %v</p>
            <p>Сума додатних елементів: %.2f</p>
            <p>Добуток між max і min за модулем: %.2f</p>
        `, arr, positiveSum, productBetween)
	}

	fmt.Fprintf(w, `
        <h3>GET запит:</h3>
        <form method="get">
            <input type="text" name="numbers" placeholder="Числа через кому">
            <input type="submit" value="Обчислити GET">
        </form>

        <h3>POST запит:</h3>
        <form method="post">
            <input type="text" name="numbers" placeholder="Числа через кому">
            <input type="submit" value="Обчислити POST">
        </form>
    </body>
    </html>`)
}

func main() {
	http.HandleFunc("/", handleArrayCalculations)
	fmt.Println("Сервер запущено на localhost:8080")
	http.ListenAndServe(":8080", nil)
}
