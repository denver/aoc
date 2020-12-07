package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type item struct {
	line string
}

func main() {

	var index int
	var trees int

	var items []item

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		var x item
		s := scanner.Text()
		x.line = s
		fmt.Println(x)
		items = append(items, x)
	}

	for i, item := range items {

		a := item.line

		for index > len(a) {
			a += a
		}

		if a[index] == 35 {
			fmt.Println("tree")
			trees++
		}
		fmt.Println("trees", trees, "index", index, "line", i, a[index])
		fmt.Println("starting line", item.line)
		fmt.Println("after line", a)

		index += 3
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
