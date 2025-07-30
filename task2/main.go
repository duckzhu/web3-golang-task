package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//指针1
	a := 10
	func1(&a)
	fmt.Printf("a的值%v，类型为%T \n", a, a)
	//指针2
	arr := []int{1, 2, 3}
	func2(&arr)
	fmt.Println("arr改变后", arr)
	//协程1
	goroutinetest()
	//协程2
	funcarr := []func(){test3, test3, test3}
	goroutinetest2(funcarr)
	//面向对象1
	interfaceTest()
	//面向对象2
	structTest()
	//管道通信
	channelTest()
	//锁机制1
	testSync()
	//锁机制2
	atomicTest()

}

//题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。

func func1(a *int) {
	*a = *a + 10
}

// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。 int数组a  b:=&a取的是a的指针，b类型为*[]int, *b是指针指向内存地址的引用变量，也就是数组a
func func2(arr *[]int) {
	b := *arr
	for i := 0; i < len(b); i++ {
		b[i] *= 2
	}
}

// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
var wg sync.WaitGroup

func test1() {
	var arr []int
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			arr = append(arr, i)
		}
	}
	fmt.Println("偶数", arr)
	time.Sleep(time.Millisecond * 10)
	wg.Done()
}
func test2() {
	var arr []int
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			arr = append(arr, i)
		}
	}
	fmt.Println("奇数", arr)
	time.Sleep(time.Millisecond * 10)
	wg.Done()
}
func goroutinetest() {
	wg.Add(1)
	go test2()
	wg.Add(1)
	go test1()
	wg.Wait()
}

// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。
var wg2 sync.WaitGroup

func test3() {
	start := time.Now()
	// var a int = 0
	// for i := 0; i < 99; i++ {
	// 	a += i
	// }
	rangomint := rand.Intn(20)
	time.Sleep(time.Millisecond * time.Duration(rangomint))
	cost := time.Since(start)
	fmt.Println("任务完成,耗时：", cost)
	wg2.Done()

}
func goroutinetest2(arr []func()) {
	for _, f := range arr {
		wg2.Add(1)
		go f()
	}
	wg2.Wait()
}

// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。
type Shape interface {
	Area()
	Perchan()
}

type Rectangle struct {
}

type Circle struct {
}

func (r Rectangle) Area() {
	fmt.Println("结构体Rectangle的函数Area实现方法")
}
func (r Circle) Area() {
	fmt.Println("结构体Circle的函数Area实现方法")
}
func (r Rectangle) Perchan() {
	fmt.Println("结构体Rectangle的函数Perchan实现方法")
}
func (r Circle) Perchan() {
	fmt.Println("结构体Circle的函数Perchan实现方法")
}
func interfaceTest() {
	var s Shape
	s = Rectangle{}
	s.Area()
	s.Perchan()
	s = Circle{}
	s.Area()
	s.Perchan()

}

// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// 考察点 ：组合的使用、方法接收者
type Person struct {
	Name string
	Age  int
}
type Employee struct {
	EmployeeID int
	Person
}

func (e Employee) PrintInfo() {
	fmt.Printf("员工信息%#v \n", e)
}

func structTest() {
	e := Employee{
		EmployeeID: 1,
		Person: Person{
			Name: "张三",
			Age:  18,
		},
	}

	e.PrintInfo()

}

// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。
var wg3 sync.WaitGroup

func test4(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("管道中放入：", i)
		time.Sleep(time.Millisecond * 100)

	}
	close(ch)
	wg3.Done()
}
func test5(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("从管道中取出：", <-ch)
	}
	wg3.Done()
}
func channelTest() {
	ch := make(chan int, 10)
	wg3.Add(1)
	go test4(ch)
	wg3.Add(1)
	go test5(ch)
	wg3.Wait()
}

// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。

var mutex sync.Mutex
var wg4 sync.WaitGroup
var count = 0

func testSync() {
	for i := 0; i < 10; i++ {

		wg4.Add(1)
		go test6()
	}
	wg4.Wait()
	fmt.Println("count值：", count)
}
func test6() {
	mutex.Lock()
	for i := 0; i < 1000; i++ {
		count++
	}
	wg4.Done()
	mutex.Unlock()
}

// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全
var counter int64
var wg5 sync.WaitGroup

func test7() {
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&counter, 1)
	}

	wg5.Done()
}

func atomicTest() {
	for i := 0; i < 10; i++ {
		wg5.Add(1)
		go test7()
	}
	wg5.Wait()
	fmt.Println("counter值为：", counter)
}
