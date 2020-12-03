package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type item struct {
	one      string
	two      string
	minMax   string
	letter   string
	password string
}

func main() {

	var items []item

	var valid int

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		var x item
		s := scanner.Text()
		//input = append(input, s)
		a := strings.Split(s, " ")

		minMax := strings.Split(a[0], "-")
		x.one = minMax[0]
		x.two = minMax[1]
		x.letter = string(a[1][0])
		x.password = string(a[2])
		fmt.Println(x)
		items = append(items, x)
	}

	fmt.Println("items:", items)

	for _, item := range items {
		var positionOneTrue bool
		var positionTwoTrue bool
		fmt.Println("starting on:", item)
		fmt.Println("itemOne", item.one)
		fmt.Println("itemLetter", item.letter)
		index1, _ := strconv.Atoi(item.one)
		index2, _ := strconv.Atoi(item.two)

		if string(item.password[index1-1]) == item.letter {
			fmt.Println("match on 1")
			positionOneTrue = true
		} else {
			fmt.Println("no match on 1")
		}
		if string(item.password[index2-1]) == item.letter {
			fmt.Println("match on 2")
			positionTwoTrue = true
		} else {
			fmt.Println("no match on 2")
		}
		if positionOneTrue && positionTwoTrue {
			fmt.Println("both match ... not valid")
		} else if positionOneTrue || positionTwoTrue {
			fmt.Println("valid", item)
			valid++
		}

		// for _, password := range item.password {
		// 	fmt.Println("password", string(password))
		// 	fmt.Println("should match letter", item.letter)
		// 	if string(password) == item.letter {
		// 	}
		// }
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(items)
	fmt.Println("Valid:", valid)
}
