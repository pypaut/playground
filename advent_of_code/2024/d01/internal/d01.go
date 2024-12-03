package d01

import (
	"errors"
	"math"
	"os"
	"strconv"
	"strings"
)

func ParseLists(inputFile string) (l1, l2 []int) {
    data, err := os.ReadFile(inputFile)
    if err != nil {
        panic(err)
    }

    lines := strings.Split(string(data), "\n")

    for _, l := range lines {
        l = strings.TrimSpace(l)
        l = strings.ReplaceAll(l, "   ", " ")

        lineIds := strings.Split(l, " ")
        if len(lineIds) < 2 || lineIds[0] == "" || lineIds[1] == "" {
            continue
        }

        id1, err := strconv.Atoi(lineIds[0])
        if err != nil {
            panic(err)
        }

        id2, err := strconv.Atoi(lineIds[1])
        if err != nil {
            panic(err)
        }

        l1 = append(l1, id1)
        l2 = append(l2, id2)
    }

    return
}

func FindMin(list []int) (value, index int, err error) {
    if len(list) == 0 {
        return 0, 0, errors.New("list is empty")
    }

    value = list[0]

    for i, e := range list {
        if e < value {
            value = e
            index = i
        }
    }

    return
}

func RemoveElementAt(list []int, index int) ([]int, error) {
    if (index < 0 || index > len(list) - 1) {
        return nil, errors.New("index error")
    }

    newList := append(list[:index], list[index+1:]...)
    return newList, nil
}

func ComputeTotalDistance(l1 []int, l2 []int) (distance int, err error) {
    for (len(l1) > 0) {
        min1, index1, err := FindMin(l1)
        if err != nil {
            return 0, err
        }
    
        min2, index2, err := FindMin(l2)
        if err != nil {
            return 0, err
        }

        distance += int(math.Abs(float64(min1) - float64(min2)))

        l1, err = RemoveElementAt(l1, index1)
        if err != nil {
            return 0, err
        }

        l2, err = RemoveElementAt(l2, index2)
        if err != nil {
            return 0, err
        }
    }

    return
}

func ComputeSimilarityScore(l1, l2 []int) (score int) {
    for _, e1 := range l1 {
        nbOccurrences := 0
        for _, e2 := range l2 {
            if e1 == e2 {
                nbOccurrences += 1
            }
        }

        score += e1 * nbOccurrences
    }

    return
}
