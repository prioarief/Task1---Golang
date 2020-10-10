package main

import (
	"fmt"
)

func cetakGambar(length int) {
	template := ""
	median := (length / 2) + 1

	for i := 1; i <= length; i++ {
		for k := 1; k <= length; k++ {
			if i == median {
				template += "*"
				template += " "
			} else {
				if k == 1 || k == length {
					template += "*"
					template += " "
				} else {
					template += "="
					template += " "
				}
			}
		}
		fmt.Println(template)
		template = ""
	}
}

func main() {
	cetakGambar(5)
}
