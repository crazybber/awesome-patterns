package main

import "fmt"

func main() {
	// unbuf()
	// buf()
	nBuf()
}

/*
上面的代码虽然可以正确同步，但是对管道的缓存大小太敏感：如果管道有缓存的话，就无
法保证能main退出之前后台线程能正常打印了。更好的做法是将管道的发送和接收方向调换
一下，这样可以避免同步事件受管道缓存大小的影响
*/
func unbuf() {
	done := make(chan int)
	go func() {
		fmt.Println("Hello World")
		<-done
	}()
	done <- 1
}

func buf() {
	done := make(chan int, 1) // 带缓存的管道
	go func() {
		fmt.Println("你好, 世界")
		done <- 1
	}()
	<-done
}

func nBuf() {
	done := make(chan int, 10) // 带 10 个缓存
	// 开N个后台打印线程
	for i := 0; i < cap(done); i++ {
		go func(i int) {
			fmt.Println("你好, 世界", i)
			done <- i
		}(i)
	}
	// 等待N个后台线程完成
	for i := 0; i < cap(done); i++ {
		<-done
	}
}
