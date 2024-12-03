package main

import (
    "fmt"

	d01 "d01/internal"
)


func main() {
    l1, l2 := d01.ParseLists("input")
    if len(l1) != len(l2) {
        panic("l1 and l2 have different len")
    }

    // Part 1
    // totalDistance, err := d01.ComputeTotalDistance(l1, l2)
    // CheckErr(err)
    // fmt.Printf("Total distance: %d\n", totalDistance)

    // Part 2
    similarityScore := d01.ComputeSimilarityScore(l1, l2)
    fmt.Printf("Similarity score: %d\n", similarityScore)

    return
}

func CheckErr(err error) {
    if err != nil {
        panic(err)
    }
}
