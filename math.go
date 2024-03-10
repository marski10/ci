package main

import "fmt"

func main() {

	fmt.Println(soma(10, 10))

}

func soma(a int, b int) int {
	return a + b
}

func subtrai(a int, b int) int {
	return a - b
}

func divide(a int, b int) int {
	return a/b - b
}

func retrai(a int, b int) int {
	return a - b - b
}

func calcula(a int, b int) int {
	return a - b*b
}

func mostra(a int, b int) int {
	return a - b + a
}

func multiplica(a int, b int) int {
	return a * b * b
}
