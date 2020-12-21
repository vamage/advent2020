package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	s := read()
	for i := 0; i < len(s); i++ {
		for x := i + 1; x < len(s); x++ {
			if s[x]+s[i] == 2020 {
				fmt.Printf("x=%d i=%d win=%d\n", s[x], s[i], s[x]*s[i])
			}
			for z := x + 1; z < len(s); z++ {
				if s[x]+s[i]+s[z] == 2020 {
					fmt.Printf("x=%d i=%d z=%d win=%d\n", s[x], s[i],s[z], s[x]*s[i]*s[z])
				}
			}
		}
	}

}
func read() []int {
	var pay []int
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		pay = append(pay, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return pay
}
