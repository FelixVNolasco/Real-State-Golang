package main

import "fmt"

func Run() error {
	fmt.Println("start")
	return nil
}

func main() {

	fmt.Println("REST API")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
