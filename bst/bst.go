package bst

type BST struct {
	value int

	left  *BST
	right *BST
}

func New(value int) *BST {
	return &BST{
		value: value,
	}
}

func (t *BST) Insert(value int) *BST {
	if value < t.value {
		if t.left == nil {
			t.left = &BST{
				value: value,
			}
			return t.left
		}
		return t.left.Insert(value)
	}
	if t.right == nil {
		t.right = &BST{
			value: value,
		}
		return t.right
	}
	return t.right.Insert(value)
}
