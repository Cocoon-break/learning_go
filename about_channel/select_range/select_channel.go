package select_range

import (
	"log"
	"time"
)

//非阻塞，打印
func NoBlockChannel() {
	ch := make(chan int)
	select {
	case i := <-ch:
		println(i)
	default:
		println("default")
	}
}

// 阻塞
func BlockChannel() {
	ch := make(chan int)
	select {
	case i := <-ch:
		println(i)
	}
	println("end")
}

// 多个case 就绪，随机执行
func randomSelect() {
	ch := make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()

	for {
		select {
		case <-ch:
			println("case1")
		case <-ch:
			println("case2")
		}
	}
}

func selectClosedChannel() {
	stopCh := make(chan struct{})
	for i := 0; i < 2; i++ {
		go func() {
			for {
				select {
				case <-stopCh: //这个会执行一次 _, ok := <-c1 ok=false
					log.Println("-----read closed channel-----")
					return
					// default: //开启default就不会阻塞
				}
			}
		}()
	}
	close(stopCh)
	select {}
}
func selectNilChannel() {
	stopCh := make(chan struct{})
	stopCh = nil
	for i := 0; i < 2; i++ {
		go func() {
			for {
				select {
				case <-stopCh: //channel 为nil，select 会直接忽略，永远不会执行
					log.Println("-----read closed channel-----")
					return
				default:
				}
				log.Println("********")
			}
		}()
	}
	select {}
}
