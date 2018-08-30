package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/chrobid/queue"
)

func main() {
	cuteQueue := queue.NewQueue()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("So...you have a queue.")
		fmt.Println("Watcha wanna do?")
		fmt.Println("	1) Push")
		fmt.Println("	2) Pop")
		fmt.Println("	3) Tell me if it's empty")
		fmt.Println("	4) Show me the queue")
		fmt.Println("	5) Quit")
		fmt.Println("Choice:")
		scanner.Scan()
		fmt.Println("")

		switch scanner.Text() {
		case "1":
			fmt.Println("Push what?:")
			scanner.Scan()
			q, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Panic("Did you enter something other than numbers? Ya gotta enter only numbers.\n", err)
			}
			cuteQueue.Push(q)
		case "2":
			if cuteQueue.IsEmpty() {
				fmt.Println("The queue is empty.")
			} else {
				fmt.Println("(My lipgloss is) Poppin':", cuteQueue.Pop())
			}
		case "3":
			status := `not empty.`
			if cuteQueue.IsEmpty() {
				status = `empty.`
			}
			fmt.Println("It's def", status)
		case "4":
			fmt.Printf("Behold! The queue:  tail-> %v <-head\n", cuteQueue.ToArray())
		case "5":
			fmt.Println("Bye!")
			return
		}
	}
}
