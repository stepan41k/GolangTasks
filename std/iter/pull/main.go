package main

import (
	"fmt"
	"iter"
)

func Zip[T1, T2 any](seq1 iter.Seq[T1], seq2 iter.Seq[T2]) iter.Seq2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		next1, stop1 := iter.Pull(seq1)
		defer stop1()

		next2, stop2 := iter.Pull(seq2)
		defer stop2()

		for {
			v1, ok1 := next1()
			v2, ok2 := next2()

			if !ok1 || !ok2 {
				return
			}

			if !yield(v1, v2) {
				return
			}
		}
	}
}

func main() {
	seq := func(yield func(int) bool) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("  (Генерация %d)\n", i)
			if !yield(i) {
				return
			}
		}
	}

	// Превращаем его в Pull-итератор
	// next — возвращает (значение, bool), где bool=false означает конец данных
	// stop — функция для очистки ресурсов (обязательно вызывать!)
	next, stop := iter.Pull(seq)
	defer stop()

	fmt.Println("Вытягиваем значения вручную")

	val, ok := next()
	if ok {
		fmt.Println("Получено:", val)
	}

	// Оставшиеся в for
	for {
		val, ok = next()

		if !ok {
			break
		}

		fmt.Println("Получено:", val)
	}
}
