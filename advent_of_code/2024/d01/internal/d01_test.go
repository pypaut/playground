package d01

import (
	"reflect"
	"testing"
)

func TestParseLists(t *testing.T) {
    l1, l2 := ParseLists("../input_test")

    expectedL1 := []int{3, 4, 2, 1, 3, 3}
    expectedL2 := []int{4, 3, 5, 3, 9, 3}

    if !reflect.DeepEqual(l1, expectedL1) {
        t.Fatalf("expected %v, got %v", expectedL1, l1)
    }

    if !reflect.DeepEqual(l2, expectedL2) {
        t.Fatalf("expected %v, got %v", expectedL2, l2)
    }
}

func TestFindMin(t *testing.T) {

    cases := []struct{
        List []int
        ExpectedValue int
        ExpectedIndex int
    }{
        {
            []int{9, 2, 1, 4, 56, 0},
            0,
            5,
        },
        {
            []int{9, 2, 1, 4, 56, 0, 231},
            0,
            5,
        },
        {
            []int{9, 2, -1, 4, 56, 0, 231},
            -1,
            2,
        },
    }

    for _, c := range cases {
        value, index, err := FindMin(c.List)
        if err != nil {
            t.Fatalf("%v", err)
        }

        if value != c.ExpectedValue {
            t.Fatalf("value is %d, expected %d", value, c.ExpectedValue)
        }

        if index != c.ExpectedIndex {
            t.Fatalf("index is %d, expected %d", index, c.ExpectedIndex)
        }
    }
}

func TestRemoveElementAt(t *testing.T) {
    list := []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    
    newList, err := RemoveElementAt(list, 2)
    if err != nil {
        t.Fatal(err)
    }
    expectedList := []int{-1, 0, 2, 3, 4, 5, 6, 7, 8, 9}

    if !reflect.DeepEqual(newList, expectedList) {
        t.Fatalf("expected %v, got %v", expectedList, list)
    }
}

func TestComputeTotalDistance(t *testing.T) {
    l1, l2 := ParseLists("../input_test")
    
    distance, err := ComputeTotalDistance(l1, l2)
    if err != nil {
        t.Fatal(err)
    }

    expectedDistance := 11
    if distance != expectedDistance {
        t.Fatalf("expected %d, got %d", expectedDistance, distance)
    }
}

func TestComputeSimilarityScore(t *testing.T) {
    l1, l2 := ParseLists("../input_test")
    
    score := ComputeSimilarityScore(l1, l2)

    expectedScore := 31
    if score != expectedScore {
        t.Fatalf("expected %d, got %d", expectedScore, score)
    }
}
