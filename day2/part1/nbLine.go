// ************************************************************************** //
//                                                                            //
//                                                        :::      ::::::::   //
//   nbLine.go                                          :+:      :+:    :+:   //
//                                                    +:+ +:+         +:+     //
//   By: adjoly <adjoly@student.42angouleme.fr>     +#+  +:+       +#+        //
//                                                +#+#+#+#+#+   +#+           //
//   Created: 2024/12/02 12:06:47 by adjoly            #+#    #+#             //
//   Updated: 2024/12/02 12:06:49 by adjoly           ###   ########.fr       //
//                                                                            //
// ************************************************************************** //

package main

import (
	"log"
	"strconv"
	"strings"
)

func	getNbrInline(line string) []int {
	var	nbr []int

	split := strings.Split(line, " ")
	for _, v := range split {
		n, err := strconv.Atoi(v)
		if err != nil { log.Fatal(err) }
		nbr = append(nbr, n)
	}
	return nbr
}
