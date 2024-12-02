// ************************************************************************** //
//                                                                            //
//                                                        :::      ::::::::   //
//   getFile.go                                         :+:      :+:    :+:   //
//                                                    +:+ +:+         +:+     //
//   By: adjoly <adjoly@student.42angouleme.fr>     +#+  +:+       +#+        //
//                                                +#+#+#+#+#+   +#+           //
//   Created: 2024/12/02 12:06:50 by adjoly            #+#    #+#             //
//   Updated: 2024/12/02 12:06:54 by adjoly           ###   ########.fr       //
//                                                                            //
// ************************************************************************** //

package main

import (
	"log"
	"os"
	"bufio"
)

func	getFileScan() (*bufio.Scanner, *os.File) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	return scanner, file
}
