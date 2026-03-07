package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"syncLearn/utils"
	"time"
)

func myGoroutine1(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Println(id)
}

type Config struct{}

var instance *Config
var once sync.Once

// 利用 sync.Once 确保无论多少个 Goroutine 同时调用 InitConfig，初始化逻辑都只会执行一次
func InitConfig() *Config {
	once.Do(func() {
		instance = &Config{}
	})
	return instance
}

// 错误写法如下，在并发环境下，会产生 race condition，创建多个 Config 实例
func InitConfigError() *Config {
	if instance == nil {
		instance = &Config{}
	}
	return instance
}

func add(wg *sync.WaitGroup, num *int) {
	defer wg.Done()
	*num += 1
}

func addUnSafe() {
	// 同时开启 10000 个 goroutine 对 num 进行递增，最后发现 num 不等于 10000
	var wg sync.WaitGroup

	n := 10 * 10 * 10 * 10
	wg.Add(n)
	num := 0
	for range n {
		go add(&wg, &num)
	}
	wg.Wait()
	fmt.Println(num)
	fmt.Println(num == n)
}

// 使用 mutex 保证并发安全
func addMutex(num *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	defer wg.Done()
	*num += 1
	mu.Unlock()
}

func addSafeMutex() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	n := 10 * 10 * 10 * 10
	wg.Add(n)

	num := 0

	for range n {
		go addMutex(&num, &mu, &wg)
	}
	wg.Wait()

	fmt.Println(num)
	fmt.Println(num == n)
}

// 读写锁
func rwMutex() {
	// 读读锁可以一起加
	// 读写锁互斥
	// 写写锁互斥
	var mr sync.RWMutex
	mr.RLock() // 读锁
	fmt.Println("a")
	mr.RUnlock()
	mr.Lock() // 写锁
	fmt.Println("b")
	mr.Unlock()
}

// 死锁 1
func deadLock1() {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	deadLock1_sub(&mu)
}

func deadLock1_sub(mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
}

// 死锁 2 循环等待
func deadLock2() {
	var mu1, mu2 sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		mu1.Lock()
		defer mu1.Unlock()
		time.Sleep(100 * time.Millisecond)
		mu2.Lock()
		defer mu2.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu2.Lock()
		defer mu2.Unlock()
		time.Sleep(100 * time.Millisecond)
		mu1.Lock()
		defer mu1.Unlock()
	}()

	wg.Wait()

}

func addAtomic() {
	var wg sync.WaitGroup
	var num int32 = 0
	n := 10 * 10 * 10 * 10
	wg.Add(n)
	for range n {
		go func() {
			defer wg.Done()
			atomic.AddInt32(&num, 1)
		}()
	}
	wg.Wait()
	fmt.Println(num == int32(n))
}

// 使用 sync.pool 减少对象的创建次数，从而提高运行效率
func myLog(msg string, p *sync.Pool) {

	b := p.Get().(*bytes.Buffer)
	b.Reset() // 从 pool 中取出对象时，内容可能未初始化，需要调用 Reset 方法进行初始化

	b.WriteString("LOG: ")
	b.WriteString(msg)
	fmt.Println(b.String())

	p.Put(b)
}

func poolUsage() {

	p := sync.Pool{
		New: func() interface{} {
			fmt.Println("创建一个新 buffer")
			return new(bytes.Buffer)
		},
	}

	for i := range 5 {
		myLog(fmt.Sprintf("这是第 %d 条消息", i), &p)
	}

	/*
		输出如下，只用了一个 bytes.Buffer 对象

		创建一个新 buffer
		LOG: 这是第 0 条消息
		LOG: 这是第 1 条消息
		LOG: 这是第 2 条消息
		LOG: 这是第 3 条消息
		LOG: 这是第 4 条消息
	*/

}

func main() {
	fmt.Println(utils.Sum(1, 2))

	// 使用 channel 等待所有 goroutine 执行完毕
	ch1 := make(chan struct{}, 10) // 10 个 goroutine

	for i := range 10 {
		go func(i int) {
			fmt.Println(i)
			ch1 <- struct{}{}
		}(i)
	}

	for range 10 {
		<-ch1
	}

	fmt.Println("all 10 goroutines end")

	// 使用 WaitGroup 实现等待
	var wg sync.WaitGroup
	wg.Add(10)
	for i := range 10 {
		go myGoroutine1(&wg, i)
	}
	wg.Wait()
	fmt.Println("all 10 goroutines end")

	// sync.Once

	// sync.Lock
	addUnSafe()
	addSafeMutex()

	// deadLock1()
	// deadLock2()

	// sync.Map
	// map 不是并发安全的
	// 要保证并发安全，一种方法是加锁，另一种方式是使用 sync.Map
	var m sync.Map
	m.Store("name", "tutu")
	m.Store("age", 18)

	age, _ := m.Load("age")
	fmt.Println(age)

	m.Range(func(k, v interface{}) bool {
		fmt.Printf("key is %v, value is %v\n", k, v)
		return true
	})

	m.Delete("age")
	age, ok := m.Load("age")
	if !ok {
		fmt.Println("age is not in map")
	}

	m.LoadOrStore("name", "tutu2")
	name, _ := m.Load("name")
	fmt.Println(name)

	// atomic
	/*
		func AddT(addr *T, delta T)(new T)
		func StoreT(addr *T, val T)
		func LoadT(addr *T) (val T)
		func SwapT(addr *T, new T) (old T)
		func CompareAndSwapT(addr *T, old, new T) (swapped bool)
		T的类型是int32、int64、uint32、uint64和uintptr中的任意一种
	*/
	addAtomic()

	poolUsage()

}
