package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Pepengu/ComputerShop/computershop"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Input file is not provided")
		return
	}

	if len(os.Args) > 2 {
		fmt.Println("Too many arguments provided")
		return
	}

	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("File couldn't be opened")
		return
	}

	inputScanner := bufio.NewScanner(inputFile)
	inputScanner.Split(bufio.ScanLines)

	var shop computershop.ComputerShop
	shop.Init(inputScanner)
	fmt.Println(shop.Schedule.Open.Format("15:04"))

	for inputScanner.Scan() {
		shop.RunAction(inputScanner.Text())
	}

	fmt.Println(shop.Schedule.Close.Format("15:04"))
	shop.Close()
}
