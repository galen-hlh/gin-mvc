package tests

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

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
