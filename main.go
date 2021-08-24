package main


import  "ztalloc/pkg/core" 

func main(){
	v := func(a core.Node) { println(a.ToString())}
	traverse := core.DfoGen(core.PREORDER,v)

	a := core.NodeFromInt(16)
	traverse(a, 10)

}
