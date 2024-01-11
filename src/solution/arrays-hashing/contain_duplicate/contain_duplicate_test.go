package main

import "testing"


func TestContainDuplicate(t *testing.T) {
    input := [][]int{
        {1,2,3,1},
        {1,2,3,4},
        {1,1,1,3,3,4,3,2,4,2},
    }

    output := []bool{true, false, true}

    for i := 0; i < 3; i++ {
        actual := ContainsDuplicate(input[i])
        if actual != output[i] {
            t.Fatalf(`Expected %t but get %t for test case %v`, output[i], 
            actual, input[i])
        }
    }
}

func TestContainDuplicateAlt1(t *testing.T) {
    input := [][]int{
        {1,2,3,1},
        {1,2,3,4},
        {1,1,1,3,3,4,3,2,4,2},
    }

    output := []bool{true, false, true}

    for i := 0; i < 3; i++ {
        actual := ContainsDuplicateAlt1(input[i])
        if actual != output[i] {
            t.Fatalf(`Expected %t but get %t for test case %v`, output[i], 
            actual, input[i])
        }
    }
}
