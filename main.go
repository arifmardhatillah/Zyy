package main

import "fmt"

func ArrayMarge(arrayA, ArrayB []string) []string {

}

func main() {
	fmt.Println(ArrayMarge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "stave", "geese"}))

	fmt.Println(ArrayMarge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	fmt.Println(ArrayMarge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimatsu", "alisa", "law"}))

	fmt.Println(ArrayMarge([]string{}, []string{"devil jin", "sergei"}))

	fmt.Println(ArrayMarge([]string{"hwoarang"}, []string{}))

	fmt.Println(ArrayMarge([]string{}, []string{}))
}
