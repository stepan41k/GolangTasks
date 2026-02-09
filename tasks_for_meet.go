// TODO: 1
package main

import (
	"log"
	"time"
)

func main(){
	ch := make(chan int, 1)

	for i := 0; i < 10; i++ {
		select {
			case x := <-ch:
				print(x)
			case ch <- i:
		}
	}
}
//0, 2, 4, 6, 8

//TODO: 2
package main

func main() {
	nums := []int{1, 2, 3}
	addNum(nums[0:2])
	fmt.Println(nums) // ?
	addNums(nums[0:2])
	fmt.Println(nums) // ?
}

func addNum(nums []int) {
	nums = append(nums, 4)
}

func addNums(nums []int) {
	nums = append(nums, 5, 6)
}

//1) 1 2 4
//2) 1 2 4

//TODO: 3
package main

import "fmt"

func main() {
	lst := []string{"a", "b", "c", "d"}

	for k, v := range lst {
		if k == 0 {
			lst = []string{"aa", "bb", "cc", "dd"}
		}

		fmt.Println(v)
	}
}

//TODO: 4
func main() {
	lst := []string{"a", "b", "c", "d"}

	for k, v := range lst {
		if k == 0 {
			lst[3] = "z"
		}

	fmt.Println(v)
	}
}

//TODO: 5
package main

func main() {
	var urls = []string{
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"https://ya.ru",
		"http://ya.ru",
		"http://ёёёё",
	}
}

// 1. Поочередно выполнит http запросы по предложенному списку ссылок
// • в случае получения http-кода ответа на запрос "200 OK" печатаем на экране "адрес url -
// ok"
// • в случае получения http-кода ответа на запрос отличного от "200 OK" либо в случае
// ошибки печатаем на экране "адрес url - not ok"

// 2. Модифицируйте программу таким образом, чтобы использовались каналы для
// коммуникации основного потока с горутинами. Пример:
// • Запросы по списку выполняются в горутинах.
// • Печать результатов на экран происходит в основном потоке

// 3. Модифицируйте программу таким образом, чтобы нигде не использовалась длина
// слайса урлов. Считайте, что урлы приходят из внешнего источника. Сколько их будет
// заранее - неизвестно. Предложите идиоматичный вариант, как ваша программа будет
// узнавать об окончании списка и передавать сигнал об окончании действий далее.
// 4. (необязательно, можно обсудить устно, чтобы убедиться, что кандидат понимает идею
// контекста, либо предложить как домашнее задание) Модифицируйте программу таким
// образом, что бы при получении 2 первых ответов с "200 OK" остальные запросы штатно
// прерывались.
// При этом необходимо напечатать на экране сообщение о завершении запроса.
// 5. (необязательно, можно обсудить устно) Предложите отрефакторить код. Какие тесты
// кандидат написал бы к этому коду?
// Предложите написать код теста и интерфейсы, для которых будут генериться моки. (Как
// показывает практика это самая сложная часть задачи)

//TODO: 6
func main() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}

//TODO: 7
// Написать код функции, которая делает merge N каналов. Весь входной поток
// перенаправляется в один канал.
func merge(cs ...<-chan int) <-chan int {

}

//TODO: 8
func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		fmt.Println(n)
	}
}

//TODO: 9
package main

import "fmt"

func main() {
	m := map[string]int{"a":1,"b":2,"c":3}

	for a, b := range m {
		fmt.Println(a, b)
	}
}

//TODO: 10
package main

import "fmt"

func main() {
	a := map[*int]int{}
	x := 10
	a[&x] = 0
	e, ok := a[&x]
	fmt.Println(e, ok)
}

//TODO: 11
package main

type Foo struct{}

func (f *Foo) A() {}
func (f *Foo) B() {}
func (f *Foo) C() {}

type AB interface {
	A()
	B()
}

type BC interface {
	B()
	C()
}

func main() {
	var f AB = &Foo{}
	y := f.(BC) // сработает ли такой type-assertion?
	y.A() 		// а этот вызов?
	
	_ = y
}

//TODO: 12
func accum() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

//TODO: 13
package main

import (
	"fmt"
	"time"
)

func main() {
	x := make(map[int]int, 1)

	go func() { x[1] = 2 }()
	go func() { x[1] = 7 }()
	go func() { x[1] = 10 }()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("x[1] =", x[1])
}

//TODO: 14
package main

import (
	"time"
)

func main() {
	timeStart := time.Now()

	_, _ = <-worker(), <-worker()
	println(int(time.Since(timeStart).Seconds()))
}

func worker() chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()
	return ch
}

//6

//TODO: 15
package main

func main() {
	s := "test"
	println(s[0]) // ?
	s[0] = "R"

	utf
	println(s) // ?
}


//TODO: 16
//Что выведет программа?
package main

import (
	"fmt"
)

func a() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)
	y := append(x, 3)
	z := append(x, 4)
	fmt.Println(y, z)
}
func main() {
	a()
}
//[0 1 2 4] [0 1 2 4]


//TODO: 17
package main

import (
	"fmt"
)

type Example struct {
	Value string
}

type MyInterface interface {}

func example1() MyInterface {
	var e *Example
	return e
}

func example2() MyInterface {
	return nil
}

func main() {
	fmt.Println(example1() == example2())
}


//TODO: 18
package main

import (
	"fmt"
)

type myError struct {
	code int
}

func (e myError) Error() string {
	log.Log
	return fmt.Sprintf("code: %d", e.code)
}

func run() error {
	select {
	case <-time.After():
	}
}

func main() {
	err := run()
	if err != nil {
		fmt.Println("failed to run, error: ", err)
	} else {
		fmt.Println("success")
	}
}


//TODO: 19
package main

func createUser() *User {
    u := User{Name: "Alice"}
    return &u
}

func main() {
    _ = createUser()
}
//В куче (Heap)


//TODO: 20
import "fmt"

func main() {
    n := 42
    fmt.Println(n)
}
//В куче (Heap).

//TODO: 21
func incrementer() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    inc := incrementer()
    inc()
}
//Heap

//TODO: 22
func makeSlice(size int) {
    data := make([]int, size)
    data[0] = 100
}

func main() {
    makeSlice(10)
}
//Heap

//TODO: 23
func makeSlice(size int) {
    data := make([]int, 5)
    data[0] = 100
}

func main() {
    makeSlice(10)
}

//TODO: 24
package main

import (
	"sync"
	"time"
)

func worker() chan int {
   ch := make(chan int)

   go func() {
    	time.Sleep(3 * time.Second)
    	ch <- 42
   }()

   return ch
}

func main() {
	timeStart := time.Now()
	wg := sync.WaitGroup{}

	wg.Add(2)
	for range 2 {
		go func() {
			defer wg.Done()
			_ = <-worker()
		}()
	}

	
	wg.Wait()


   println(int(time.Since(timeStart).Seconds()))
}

//TODO: 25
package main

import (
	"fmt"
)

type MyError struct {}

func (m MyError) Error() string {
	return "MyError"
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
   var err *MyError
   errorHandler(err)

   err = &MyError{}
   errorHandler(err)
}


//TODO: 26
package main

import (
	"fmt"
)

// input: ["a", "bb", "bb", "aa", "a", "a"]
// output: ["a", "bb"]

func printRepeats(arr []string) {
	m := map[string]uint8{}
	out := make([]string, 0, len(arr))

	for _, v := range arr {
		if val, ok := m[v]; ok && val == 1 {
			m[v] = 2
			out = append(out, v)
		} else {
			m[v] = 1
		}
	}

	fmt.Println(out)
}

func main () {
	arr := []string{"a", "bb", "bb", "aa", "a", "a"}

	printRepeats(arr)
}

