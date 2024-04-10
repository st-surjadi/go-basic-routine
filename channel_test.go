package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Steven"
		fmt.Println("Data sent from the function")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(2 * time.Second)
}

// CHANNEL AS PARAMETER
func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(2 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Steven"
}

// CHANNEL AS PARAMETER ONLY IN/OUT
func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Steven"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

// CHANNEL BUFFER
func TestChannelBuffer(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Steven"
		channel <- "Sean"
		channel <- "Surjadi"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}

// CHANNEL RANGE
func TestChannelRange(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Data-" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Done")
}

// CHANNEL SELECT
func TestChannelSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data channel 1: ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data channel 2: ", data)
			counter++
		default:
			fmt.Println("Waiting for another data.")
		}

		if counter == 2 {
			break
		}
	}
}
