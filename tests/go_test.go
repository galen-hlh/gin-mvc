package tests

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestStructToMapViaJson(t *testing.T) {
	type HttpResponse struct {
		Code  int32                  `json:"code"`
		Msg   string                 `json:"msg"`
		Data  map[string]interface{} `json:"data"`
		Trace []string               `json:"trace"`
	}
	t1 := time.Now()

	m := make(map[string]interface{})
	response := &HttpResponse{
		Code:  100,
		Msg:   "not",
		Data:  nil,
		Trace: []string{"1", "2", "3"},
	}
	j, _ := json.Marshal(response)
	_ = json.Unmarshal(j, &m)

	fmt.Println(m)
	fmt.Printf("duration:%d", time.Now().Sub(t1))
}

//struct转MAP 通过反射
func TestStructToMapViaReflect(t *testing.T) {
	type HttpResponse struct {
		Code  int32                  `json:"code"`
		Msg   string                 `json:"msg"`
		Data  map[string]interface{} `json:"data"`
		Trace []string               `json:"trace"`
	}
	t1 := time.Now()

	m := make(map[string]interface{})
	response := &HttpResponse{
		Code:  100,
		Msg:   "not",
		Data:  nil,
		Trace: []string{"1", "2", "3"},
	}
	elem := reflect.ValueOf(&response).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		m[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	fmt.Println(m)
	fmt.Printf("duration:%d", time.Now().Sub(t1))
}

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
