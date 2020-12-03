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
	min      string
	max      string
	minMax   string
	letter   string
	password string
}

func main() {

	var items []item

	var valid int

	//var input []string

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
		x.min = minMax[0]
		x.max = minMax[1]
		x.letter = string(a[1][0])
		x.password = string(a[2])
		fmt.Println(x)
		items = append(items, x)
	}

	fmt.Println("items:", items)

	for _, item := range items {
		fmt.Println("starting on:", item)
		var count int
		for _, password := range item.password {
			fmt.Println("password", string(password))
			fmt.Println("should match letter", item.letter)
			if string(password) == item.letter {
				count++
				fmt.Println("match", count)
			}
		}

		min, _ := strconv.Atoi(item.min)
		max, _ := strconv.Atoi(item.max)
		fmt.Println("min:", min, "max:", max, "count:", count)
		if (min <= count) && (count <= max) {
			fmt.Print("valid++:", item)
			valid++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(items)
	fmt.Println("Valid:", valid)
}
