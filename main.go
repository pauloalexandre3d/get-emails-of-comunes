package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	regions := getRegions()
	regions.showMenu()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

}
