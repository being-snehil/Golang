package main 

import "fmt"

func main(){
	nums := []int{1, 2, 3, 4, 5}

	for i , num := range nums {
		fmt.Println(i,num);
	}

	m := map[string]int{"price": 40, "phones": 3, "address": 50}

	for k , v := range m {
		fmt.Println(k,v);
	}

	// unicode code point rune
	// starting byte of rune
	// 300 -> 1 byte , 2 byte
	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}
}