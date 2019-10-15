package tests

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"runtime"
	"sync"
	"testing"
)

// 通过位运算交换两个变量的值
func TestExchangePosition(t *testing.T) {
	a, b := 1, 2

	// ^ 按位异或  两对应的二进位相异时，结果为1
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("a:%d,b:%d \n", a, b)
}

// golang 的异常恢复
func TestPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			info := fmt.Sprintf("==> %s\n", string(buf[:n]))
			fmt.Println(info) // 将 interface{} 转型为具体类型。
		}
	}()

	panic("i am error")
}

// 雪花
func TestSnowFlake(t *testing.T) {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate a snowflake ID.
	id := node.Generate()

	// Print out the ID in a few different ways.
	fmt.Printf("Int64  ID: %d\n", id)
	fmt.Printf("String ID: %s\n", id)
	fmt.Printf("Base2  ID: %s\n", id.Base2())
	fmt.Printf("Base64 ID: %s\n", id.Base64())

	// Print out the ID's timestamp
	fmt.Printf("ID Time  : %d\n", id.Time())

	// Print out the ID's node number
	fmt.Printf("ID Node  : %d\n", id.Node())

	// Print out the ID's sequence number
	fmt.Printf("ID Step  : %d\n", id.Step())

	// Generate and print, all in one.
	fmt.Printf("ID       : %d\n", node.Generate().Int64())

}

// 多线程顺序打印1-100
func TestGoroutine(t *testing.T) {
	var i = 1
	var lock = sync.Mutex{}
	var group = sync.WaitGroup{}

	for n := 1; n <= 4; n++ {
		group.Add(1)
		go func() {
			for {
				lock.Lock()
				if i >= 101 {
					lock.Unlock()
					group.Add(-1)
					break
				}
				fmt.Println(i)
				i++
				lock.Unlock()
			}
		}()
	}

	group.Wait()
}
