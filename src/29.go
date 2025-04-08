package main

import "fmt"

func main() {
    // Define a slice of integers
    numbers := []int{1, 2, 3, 4, 5}
    
    // Print each element of the slice
    for _, num := range numbers {
        fmt.Println(num)
    }
}
