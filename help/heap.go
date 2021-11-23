package help

func Print() {
}
type ArrayHeapCell struct {
	Val int
	Idx int
}

type ArrayHeapCellList []ArrayHeapCell

func (a *ArrayHeapCellList) Push(x interface{}) {
	*a = append(*a, x.(ArrayHeapCell))
}

func (a *ArrayHeapCellList) Pop() interface{} {
	v := (*a)[len(*a)-1]
	*a=(*a)[:len(*a)-1]
	return v
}

func (a *ArrayHeapCellList) Len() int {
	return len(*a)
}

func (a *ArrayHeapCellList) Less(i, j int) bool {
	if (*a)[i].Val == (*a)[j].Val {
		return (*a)[i].Idx < (*a)[j].Idx
	}
	return (*a)[i].Val < (*a)[j].Val
}

func (a *ArrayHeapCellList) Swap(i, j int) {
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
}





