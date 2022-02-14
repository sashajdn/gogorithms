package graphs

// NestedInteger ...
type NestedInteger struct{}

func (n NestedInteger) IsInteger() bool           { return true }
func (n NestedInteger) IsList() bool              { return true }
func (n NestedInteger) GetInteger() int           { return 0 }
func (n NestedInteger) GetList() []*NestedInteger { return []*NestedInteger{} }
func (n NestedInteger) SetInteger(v int)          {}
func (n NestedInteger) Add(elem NestedInteger)    {}

// 1. either int
// 2. list containing ints

// DepthSum ...
func DepthSum(nestedList []*NestedInteger) int {
	return dfsSum(nestedList, 1)
}

// T -> O(max(depth) * max(num in list)) -> O(v + e) -> O(n) where n is the total number of integers.
// S -> O(depth) -> O(v) -> O(n)
func dfsSum(nestedList []*NestedInteger, level int) int {
	if len(nestedList) == 0 {
		return 0
	}

	if nestedList[0].IsInteger() {
		return (level * nestedList[0].GetInteger()) + dfsSum(nestedList[1:], level)
	}

	return dfsSum(nestedList[0].GetList(), level+1) + dfsSum(nestedList[1:], level)
}
