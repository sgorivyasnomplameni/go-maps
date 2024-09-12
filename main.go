package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["apple"] = 1
	m["banana"] = 2
	m["cherry"] = 3

	fmt.Println("Выводим все пары ключ-значение:")
	for key, value := range m {
		fmt.Printf("Ключ: %s, Значение: %d\n", key, value)
	}

	m["banana"] = 4

	value, ok := m["banana"]
	if ok {
		fmt.Printf("\nКлюч banana найден. Его значение - %d\n", value)
	} else {
		fmt.Printf("\nКлюч не найден\n")
	}

	fmt.Println("\nУдаление ключа 'cherry'")
	delete(m, "cherry")
	fmt.Println("Вывод оставшихся пар \"ключ-значение\"")
	for key, value := range m {
		fmt.Printf("Ключ: %s, Значение: %d\n", key, value)
	}

	value, ok = m["cherry"]
	if ok {
		fmt.Printf("\nЗначение для ключа 'cherry': %d\n", value)
	} else {
		fmt.Println("\nКлюч 'cherry' не найден")
	}

}
