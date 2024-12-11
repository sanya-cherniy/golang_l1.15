package main

var justString string

// Длина строки может быть меньше 100 вследствие чего произойдет паника, нужно добавить проверку длины строки
func someFunc() {
	v := createHugeString(1 << 10)
	if len(v) >= 100 {
		justString = v[:100]
	}

}

func main() {
	someFunc()
}
