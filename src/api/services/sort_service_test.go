package services

import (
	"testing"

)

func TestConstant(t *testing.T){
	if privateConst != "private"{
		t.Error("privateCOnst should be 'private'")
	}
}

func  TestSort(t *testing.T)  {
	elements := GetElements(10)

	Sort(elements)

	if elements[0] != 0{
		t.Error("primer elem debe ser 9",elements[0])
	}
}

func  TestMoreSort(t *testing.T)  {
	elements := GetElements(10001)

	Sort(elements)

	if elements[0] != 0{
		t.Error("primer elem debe ser 9",elements[0])
	}
}

func GetElements(n int) []int {
	result := make([]int, n)
	j := 0
	for i := n - 1; i > 0; i-- {
		result[j] = i
		j++
	}
	return result
}