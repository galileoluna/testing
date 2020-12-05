package services
import (
	"github.com/galileoluna/testing/src/api/utils/sort"

)

const(
	privateConst = "private"
	Public = "public"
)

func Sort(elements []int){
	if len(elements) <= 10000{
		sort.BubbleSort(elements)
		return
	}

	sort.Sort(elements)
}