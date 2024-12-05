package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func InputOption() int {

	var option int
	var err error

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("Input the option - 1: Encode - 2: Decode")

		if scanner.Scan() {

			option, err = strconv.Atoi(scanner.Text())

			if err != nil {

				fmt.Println("Input incorrect")
				continue
			} else {

				break
			}

		}

	}

	return option
}

func InputText() string {

	var text string
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("Input the text:")
		if scanner.Scan() {

			text = scanner.Text()

			if len(text) == 0 {

				fmt.Println("Input a valid string")
				continue
			} else {

				break
			}
		}

	}

	return text
}
