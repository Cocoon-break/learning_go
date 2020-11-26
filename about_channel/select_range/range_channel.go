package select_range

import (
	"fmt"
	"strconv"
)

func makeCakeAndSend(cs chan string, count int, suffix string) {
	for i := 1; i <= count; i++ {
		cakeName := "Strawberry Cake " + strconv.Itoa(i) + suffix
		cs <- cakeName //send a strawberry cake
	}
}

func receiveCakeAndPack(cs chan string) {
	for s := range cs {
		fmt.Println("Packing received cake: ", s)
	}
}

func rangeChannel() {
	cs := make(chan string)
	go makeCakeAndSend(cs, 5, " first")
	go receiveCakeAndPack(cs)
	go makeCakeAndSend(cs, 5, " second")

	select {}
}
