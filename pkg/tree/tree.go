package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) Value() *int {
	if n == nil {
		return nil
	}

	return &n.Val
}

func (n *Node) SliceLen() int {
	if n == nil {
		return 0

	}
	l := 0
	n.sliceLen(n, 0, &l)
	return l + 1
}

func (n *Node) sliceLen(t *Node, pos int, len *int) {
	if t == nil {
		return
	}

	lChildPos := pos*2 + 1
	rChildPos := pos*2 + 2

	if t.Left != nil {
		if lChildPos > *len {
			*len = lChildPos
		}
	}

	if t.Right != nil {
		if rChildPos > *len {
			*len = rChildPos
		}
	}

	n.sliceLen(t.Left, lChildPos, len)
	n.sliceLen(t.Right, rChildPos, len)
}

func (t *Node) ToSlice() []*int {
	res := make([]*int, t.SliceLen())
	res[0] = &t.Val

	walkNodes(t, 0, &res)

	return res
}

func walkNodes(n *Node, pos int, arr *[]*int) {
	if n == nil || n.Left == nil && n.Right == nil {
		return
	}

	lChildPos := pos*2 + 1
	rChildPos := pos*2 + 2

	if len(*arr) > lChildPos {
		(*arr)[lChildPos] = n.Left.Value()
	}
	if len(*arr) > rChildPos {
		(*arr)[rChildPos] = n.Right.Value()
	}

	walkNodes(n.Left, lChildPos, arr)
	walkNodes(n.Right, rChildPos, arr)
}

func FromSlice(s []*int) Node {
	if len(s) == 0 || s[0] == nil {
		return Node{}
	}

	root := Node{
		Val: *s[0],
	}

	fillTree(&root, s, 0)

	return root
}

func fillTree(n *Node, s []*int, pos int) {
	lChildPos := pos*2 + 1
	rChildPos := pos*2 + 2

	if lChildPos < len(s) && s[lChildPos] != nil {
		n.Left = &Node{
			Val: *s[lChildPos],
		}
		fillTree(n.Left, s, lChildPos)
	}

	if rChildPos < len(s) && s[rChildPos] != nil {
		n.Right = &Node{
			Val: *s[rChildPos],
		}
		fillTree(n.Right, s, rChildPos)
	}
}

func ToPointersSlice(s []int) []*int {
	res := make([]*int, len(s))

	for i := range s {
		res[i] = &s[i]
	}

	return res
}
