package main

// # test update
import (
	"fmt"

	"github.com/brentgroves/filter_private"
)

func main() {
	fmt.Println("Go private module Example")

	intArray1 := []int{1, 2, 3, 4, 5, 5}
	intArray2 := []int{0, 9, 45, 5}
	result := filter_private.Filter(intArray1, intArray2)
	fmt.Println(result)
}
