package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner((os.Stdin))
	template := "Received"

	for {
		fmt.Println("Waiting for input....:")
		scanner.Scan()
		text := fmt.Sprintf("\"%s\\n\"", scanner.Text())
		output := fmt.Sprintf("%s %s", template, text)
		fmt.Println(output)
	}
}
