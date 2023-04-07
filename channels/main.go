package main

import "github.com/AvishkaUmesh/Golang-Basics/channels/helpers"

const numPool = 1000

func calcValue(iChan chan int) {
	randomNum := helpers.RandomNumber(numPool)
	iChan <- randomNum

}

func main() {
	intChannel := make(chan int)
	defer close(intChannel)

	// this is a go routine
	go calcValue(intChannel)

	num := <-intChannel
	println(num)

}
