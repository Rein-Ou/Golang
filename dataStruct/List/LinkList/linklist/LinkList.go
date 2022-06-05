package linklist

type ElemType int
type LNode struct {
	data ElemType
	p    *LNode
}
type LinkList LNode
