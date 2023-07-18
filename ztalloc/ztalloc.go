package core

import(

)

type(
	Node interface {
		Parent() *Node
		LeftChild() *Node
		RightChild() *Node
	}
)