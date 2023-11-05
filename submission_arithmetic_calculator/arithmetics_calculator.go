package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO Problem 1: Buat fungsi penjumlahan
func Add(a int, b int) int {
	// Tulis jawabanmu disini
	return a + b
}

// TODO Problem 2: Buat fungsi nilai maksimum
func max(a, b int) int {
	// Tulis jawabanmu disini
	if (a > b){
		return a
	}
	return b
}

// TODO Problem 3: Buat fungsi pencarian angka
func getNumber(s string) int {
	// Tulis jawabanmu disini
	num, _ := strconv.Atoi(s)
	return num
}

func ArithmeticArranger(problems []string, val ...bool) string {

	// ! DON'T CHANGE THIS CODE
	the_solutions := false
	if len(val) > 0 && val[0] {
		the_solutions = true
	}

	// TODO Problem and Testing 1: Buat batasan dengan 5 input data
	arrangedProblems := ""
	// Tulis jawabanmu disini {
		if len(problems) > 5 {
			arrangedProblems = "Error: Too many problems." // Tulis return berdasarkan rules README.md
			return arrangedProblems
		}

	// TODO Problem dan Testing 2: Buat batasan hanya menggunakan operator "+" dan "-"
	operations := make([]string, len(problems))
	for i, problem := range problems {
		parts := strings.Fields(problem)
		operations[i] = parts[1]
	}
	uniqueOperations := make(map[string]bool)
	for _, op := range operations {
		uniqueOperations[op] = true
	}
	// Tulis jawabanmu disini 
	if len(uniqueOperations) != 2 || !uniqueOperations["+"] || !uniqueOperations["-"] {
		arrangedProblems = "Error: Operator must be '+' or '-'." // Tulis return berdasarkan rules README.md
		return arrangedProblems
	}

	// TODO Problem dan Testing 3: Buat batasan dimana hanya terdapat angka dan tidak boleh lebih dari 4 angka
	numbers := []string{}
	for _, problem := range problems {
		parts := strings.Fields(problem)
		numbers = append(numbers, parts[0], parts[2])
	}
	for _, number := range numbers {
		// * Tulis jawabanmu disini 
		if _, err := strconv.Atoi(number); err != nil {
			arrangedProblems = "Error: Numbers must only contain digits." // Tulis jawabanmu berdasarkan rules README.md
			return arrangedProblems
		}
		// * Tulis jawabanmu disini 
		if len(number) > 4 {
			arrangedProblems = "Error: Numbers cannot be more than four digits." // Tulis jawabanmu berdasarkan rules README.md
			return arrangedProblems
		}
	}

	// TODO Problem dan Testing 4: Buat output sesuai dengan format yang berlaku. Lihat di README.md
	topRow := ""
	dashes := ""
	values := []int{}
	solutions := ""
	for i := 0; i < len(numbers); i += 2 {
		width := max(len(numbers[i]), len(numbers[i+1])) + 2
		topRow += fmt.Sprintf("%*s", width, numbers[i]) // Tulis jawabanmu disini
		dashes += strings.Repeat("-", width) // Tulis jawabanmu disini
		value, _ := strconv.Atoi(numbers[i])
		if operations[i/2] == "+" {
			value += getNumber(numbers[i+1])
		} else {
			value -= getNumber(numbers[i+1])
		}
		values = append(values, value) // Tulis jawabanmu disini
		solutions += fmt.Sprintf("%*d", width, value) // Tulis jawabanmu disini
		if i != len(numbers)-2 {
			topRow += "    "
			dashes += "    "
			solutions += "    "
		}
	}

	bottomRow := ""
	for i := 1; i < len(numbers); i += 2 {
		width := max(len(numbers[i-1]), len(numbers[i])) + 1
		bottomRow += fmt.Sprintf("%s%*s", operations[i/2], width, numbers[i])
		if i != len(numbers)-1 {
			bottomRow += "    "
		}
	}

	if the_solutions {
		arrangedProblems = strings.Join([]string{topRow, bottomRow, dashes, solutions}, "\n") // Tulis jawabanmu disini
	} else {
		arrangedProblems = strings.Join([]string{topRow, bottomRow, dashes}, "\n") // Tulis jawabanmu disini
	}
	return arrangedProblems
}
