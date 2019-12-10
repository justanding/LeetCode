package twoCitySchedCost_test

import (
	"encoding/json"
	"sort"
	"testing"
)

func TestTwoCitySchedCost(t *testing.T) {
	s := [][]int{
		[]int{10, 20},
		[]int{30, 200},
		[]int{400, 50},
		[]int{30, 20},
	}
	t.Log(twoCitySchedCost(s))

	s2 := json2Slice("[[945,563],[598,753],[558,341],[372,54],[39,522],[249,459],[536,264],[491,125],[367,118],[34,665],[472,410],[109,995],[147,436],[814,112],[45,545],[561,308],[491,504],[113,548],[626,104],[556,206],[538,592],[250,460],[718,134],[809,221],[893,641],[404,964],[980,751],[111,935]]")
	t.Log(twoCitySchedCost(s2))
}

/*
 * https://leetcode-cn.com/problems/two-city-scheduling/
 * 采用路程差，先假设所有都去A，然后排序路程差，找到路程差最大的N个
 */
func twoCitySchedCost(costs [][]int) int {
	l := len(costs)
	total := 0
	costDiff := make([]int, l)
	for i := 0; i < l; i++ {
		diff := costs[i][0] - costs[i][1]
		costDiff[i] = diff
		total += costs[i][0]
	}

	sort.Ints(costDiff)
	for i := l / 2; i < l; i++ {
		total -= costDiff[i]
	}

	return total
}

func json2Slice(data string) [][]int {
	var slice [][]int
	err := json.Unmarshal([]byte(data), &slice)
	if err != nil {
		return nil
	}
	return slice
}
