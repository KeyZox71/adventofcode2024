package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func main() {
	var	final int = 0
	scanner, file := getFileScan()
	defer file.Close()

	re := regexp.MustCompile("mul\\([0-9]+,[0-9]+\\)")
	reee := regexp.MustCompile("[0-9]+")
	for scanner.Scan() {
		res := re.FindAllString(scanner.Text(), -1)
		for _, v := range res {
			resss := reee.FindAllString(v, 2)
			nb1, err := strconv.Atoi(resss[0])
			if err != nil {
				log.Fatal(err)
			}
			nb2, err := strconv.Atoi(resss[1])
			if err != nil {
				log.Fatal(err)
			}
			final += nb1 * nb2
		}
	}
	fmt.Println(final)
}
