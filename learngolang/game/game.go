package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	number := rand.Intn(100) + 1
	win := false

	for i := 0; i < 10; i++ {
		fmt.Print("Сколько бутылок пива выпьешь?  ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		bottle, err := strconv.Atoi(input)

		if err != nil {
			log.Fatal(err)
		}

		if bottle > number {
			fmt.Println("Слишком много")
		} else if bottle < number {
			fmt.Println("Фу маловато")
		} else {
			fmt.Println("Вот тут заебись")
			win = true
			break
		}
	}
	if win {
		fmt.Println("You WIN")
	} else {
		fmt.Println("Проебав сорри")
	}
}
