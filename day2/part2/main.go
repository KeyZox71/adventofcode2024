// ************************************************************************** //
//                                                                            //
//                                                        :::      ::::::::   //
//   main.go                                            :+:      :+:    :+:   //
//                                                    +:+ +:+         +:+     //
//   By: adjoly <adjoly@student.42angouleme.fr>     +#+  +:+       +#+        //
//                                                +#+#+#+#+#+   +#+           //
//   Created: 2024/12/02 11:28:48 by adjoly            #+#    #+#             //
//   Updated: 2024/12/02 13:14:24 by adjoly           ###   ########.fr       //
//                                                                            //
// ************************************************************************** //

package main

import (
	"fmt"
	"math"
)

func	isTheNextSafe(nbr []int, i int, sign bool) bool {
	v := nbr[i + 2] - nbr[i]
	if v > 0 && sign == true {
		return false
	} else if v < 0 && sign == false {
		return false
	} else if math.Abs(float64(v)) < 1 || math.Abs(float64(v)) > 3 {
		return false 
	}
	return true
}

func	isSafe(line string) bool {
	var	sign bool
	nbr := getNbrInline(line)

	if nbr[0] < nbr[1] {
		sign = false
	} else {
		sign = true
	}
	for i := 0; i < len(nbr) - 1; i++ {
		v := nbr[i + 1] - nbr[i]
		if v > 0 && sign == true {
			if (i < len(nbr) - 2) {
				return isTheNextSafe(nbr, i, sign)
			}
			return false
		} else if v < 0 && sign == false {
			if (i < len(nbr) - 2) {
				return isTheNextSafe(nbr, i, sign)
			}
			return false
		} else if math.Abs(float64(v)) < 1 || math.Abs(float64(v)) > 3 {
			if (i < len(nbr) - 2) {
				return isTheNextSafe(nbr, i, sign)
			}
			return false
		}
	}
	return true
}

func	main() {
	var	i int = 0

	scanner, file := getFileScan()
	defer file.Close()

	for scanner.Scan() {
		if (isSafe(scanner.Text())) {
			i++;
		}
	}
	fmt.Println(i)
}
