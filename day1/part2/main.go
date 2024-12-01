// ************************************************************************** //
//                                                                            //
//                                                        :::      ::::::::   //
//   main.go                                            :+:      :+:    :+:   //
//                                                    +:+ +:+         +:+     //
//   By: adjoly <adjoly@student.42angouleme.fr>     +#+  +:+       +#+        //
//                                                +#+#+#+#+#+   +#+           //
//   Created: 2024/12/01 14:55:41 by adjoly            #+#    #+#             //
//   Updated: 2024/12/01 16:05:02 by adjoly           ###   ########.fr       //
//                                                                            //
// ************************************************************************** //

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getOccureence(list[] int, nb int) int {
	res := 0

	for _, v := range list {
		if (v == nb) {
			res++
		}
	}
	fmt.Println(res)
	return res
}

func main() {
	file, err := os.Open("../input.txt")
	if (err != nil) {
		log.Fatal(err)
	}
	defer	file.Close()

	scanner := bufio.NewScanner(file)
	list1 := []int{}
	list2 := []int{}
	for (scanner.Scan()) {
		buf := strings.Split(scanner.Text(), "   ")
		int1, err := strconv.Atoi(buf[0])
		if (err != nil) {
			log.Fatal(err);
		}
		list1 = append(list1, int1)
		int2, err := strconv.Atoi(buf[1])
		if (err != nil) {
			log.Fatal(err);
		}
		list2 = append(list2, int2)
	}
	final := []int{}
	for _, v := range list1 {
		final = append(final, v * getOccureence(list2, v))
	}

	result := 0
	for _, v := range final {
		result += v
	}
	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
