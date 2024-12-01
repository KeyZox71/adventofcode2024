// ************************************************************************** //
//                                                                            //
//                                                        :::      ::::::::   //
//   main.go                                            :+:      :+:    :+:   //
//                                                    +:+ +:+         +:+     //
//   By: adjoly <adjoly@student.42angouleme.fr>     +#+  +:+       +#+        //
//                                                +#+#+#+#+#+   +#+           //
//   Created: 2024/12/01 14:55:41 by adjoly            #+#    #+#             //
//   Updated: 2024/12/01 15:50:17 by adjoly           ###   ########.fr       //
//                                                                            //
// ************************************************************************** //

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

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
	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})


	final := []int{}
	for	i := 0; i < len(list1); i++ {
		dist := list2[i] - list1[i]
		final = append(final, int(math.Abs(float64(dist))))
	}
	result := 0
	for i := 0; i < len(final); i++ {
		result += final[i]
	}
	//for v := range final {
	//	result += v
	//}
	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
