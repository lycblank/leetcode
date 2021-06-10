package main

import "sort"

func main() {

}

func stoneGameVI(aliceValues []int, bobValues []int) int {
	n := len(aliceValues)
	sum := make([]*DataValue, n)
	for i:=0;i<n;i++{
		sum[i] = &DataValue{
			Val:aliceValues[i]+bobValues[i],
			Idx:i,
		}
	}
	sort.Sort(DataValueList(sum))

	suma, sumb := 0, 0
	for i:=0;i<n;i++ {
		if i&1 == 0 {
			suma += aliceValues[sum[i].Idx]
		} else {
			sumb += bobValues[sum[i].Idx]
		}
	}
	if suma == sumb {
		return 0
	}
	if suma > sumb {
		return 1
	}
	return -1
}

type DataValue struct {
	Val int
	Idx int
}

type DataValueList []*DataValue
func (dv DataValueList)Len() int {
	return len(dv)
}

func (dv DataValueList) Less(i, j int) bool {
	return dv[i].Val > dv[j].Val
}

func (dv DataValueList) Swap(i, j int) {
	dv[i], dv[j] = dv[j], dv[i]
}

