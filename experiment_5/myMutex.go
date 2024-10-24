package main
import (
    "fmt"
    "sync"
)

var (
    mutex   sync.Mutex
    balance int
)

// 运行在main之前
func init() {
    balance = 1000
	fmt.Println("init")
}
func deposit(value int, wg *sync.WaitGroup) {
    mutex.Lock()     //加锁
    fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
    balance += value
    mutex.Unlock()      //释放锁
    wg.Done()       //执行完毕，协程计数-1
}

func withdraw(value int, wg *sync.WaitGroup) {
    mutex.Lock()
    fmt.Printf("Withdrawing %d from account with balance: %d\n", value, balance)
    balance -= value
    mutex.Unlock()
    wg.Done()
}

func main() {
    fmt.Println("Go Mutex Example")

    var wg sync.WaitGroup
    wg.Add(2)      //协程计数+2
    go withdraw(700, &wg)
    go deposit(500, &wg)
    wg.Wait()     //当协程计数为0时退出
    fmt.Printf("New Balance %d\n", balance)
}
