package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func	getMultRes(v string) int {
	reee := regexp.MustCompile("[0-9]+")
	resss := reee.FindAllString(v, 2)
	nb1, err := strconv.Atoi(resss[0])
	if err != nil {
		log.Fatal(err)
	}
	nb2, err := strconv.Atoi(resss[1])
	if err != nil {
		log.Fatal(err)
	}
	final := nb1 * nb2
	return final
}

func main() {
	var	final int = 0
	var dont bool = false
	scanner, file := getFileScan()
	defer file.Close()

	re := regexp.MustCompile(`don\'t\(\)|do\(\)|mul\([0-9]+,[0-9]+\)`)
	for scanner.Scan() {
		res := re.FindAllString(scanner.Text(), -1)
		for _, v := range res {
			if v == "don't()"{
				dont = true
				continue
			} else if v  == "do()"{
				dont = false
				continue
			}
			if (dont == false) {
				final += getMultRes(v)
			}
		}
	}
	fmt.Println(final)
}
