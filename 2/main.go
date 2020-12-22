package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var regex = regexp.MustCompilePOSIX(`^([[:digit:]]+)-([[:digit:]]+)[[:space:]]([[:alpha:]]):[[:space:]]([[:alpha:]]+)$$`)

func main() {
	read()
	read2()
}
func read() {

	valid := 0
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		line := regex.FindStringSubmatch(s)
		min, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(line[2])
		if err != nil {
			log.Fatal(err)
		}
		char := line[3]
		password := line[4]
		finder := regexp.MustCompile(char)
		count := finder.FindAllStringIndex(password, -1)
		if len(count) >= min && len(count) <= max {
			valid++

		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("valid passwords: ", valid)
}
func read2() {

	valid := 0
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		line := regex.FindStringSubmatch(s)
		min, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(line[2])
		if err != nil {
			log.Fatal(err)
		}
		char := line[3]
		password := line[4]
		finder := regexp.MustCompile(fmt.Sprintf("^.{%d}(.).{%d}(.)", min-1, (min - min)))
		found := finder.FindStringSubmatch(password)
		fmt.Printf("min: %d,max:%d,found:%v,char %s, regex: %s, password %s\n", min, max, found, char, fmt.Sprintf("^.{%d}(.).{%d}(.)", min-1, max-min-1), password)
		if len(found) == 3 && (strings.Compare(char, found[1]) > 0 || strings.Compare(char, found[2]) > 0) {
			valid++
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("valid passwords: ", valid)
}
