package render

import (
	"fmt"
	"ztalloc/ztalloc"

	"github.com/goccy/go-graphviz/cgraph"
)

type (
	TreeRenderer interface {
		RenderNode(ztalloc.Node,*cgraph.Node)
		RenderEdge(ztalloc.TransformName, *cgraph.Edge)
	}

	ZTreeRenderer struct {}
	CTreeRenderer struct {}
)


func (r ZTreeRenderer) RenderNode(dNode ztalloc.Node, gNode *cgraph.Node) {
	label := fmt.Sprintf("%d(%d)", dNode.Value(), dNode.Value()%54)
	gNode.SetLabel(label)
	gNode.SetColorScheme("oranges9")
	gNode.SetStyle(cgraph.FilledNodeStyle)
	gNode.SetShape(cgraph.OvalShape)

	var fillc, fontc string
	switch dNode.Value() % 54 {
	case 4:
		fillc = "2"
		fontc = "black"
	case 16:
		fillc = "3"
		fontc = "black"
	case 22:
		fillc = "4"
		fontc = "black"
	case 34:
		fillc = "5"
		fontc = "black"
	case 40:
		fillc = "6"
		fontc = "black"
	case 52:
		fillc = "7"
		fontc = "black"
	default: //not expected
		gNode.SetShape(cgraph.DiamondShape)
		gNode.SetColorScheme("")
		fillc = "firebrick"
		fontc = "white"
	}

	gNode.SetFillColor(fillc)
	gNode.SetFontColor(fontc)
}

func (r ZTreeRenderer) RenderEdge(t ztalloc.TransformName, gedge *cgraph.Edge) {
	gedge.SetLabel(t)
	gedge.SetColorScheme("piyg7")
	// gedge.Set(cgraph.VeeArrow)

	var fontc string
	switch t {
	case ztalloc.Z22:
		fontc = "1"
	case ztalloc.Z2222:
		fontc = "2"
	case ztalloc.Z32:
		fontc = "3"
	case ztalloc.Z322:
		fontc = "4"
	case ztalloc.Z3222:
		fontc = "5"
	default: //not expected
		gedge.SetColorScheme("")
		fontc = "firebrick"
	}
	gedge.SetFontColor(fontc)
	gedge.SetColor(fontc)
}


func (r CTreeRenderer) RenderNode(dNode ztalloc.Node, gNode *cgraph.Node) {
	label := fmt.Sprintf("%d(%d)", dNode.Value(), dNode.Value()%54)
	gNode.SetLabel(label)
	gNode.SetColorScheme("oranges9")
	gNode.SetStyle(cgraph.FilledNodeStyle)
	gNode.SetShape(cgraph.OvalShape)

	var fillc, fontc string
	switch dNode.Value() % 54 {
	case 4:
		fillc = "2"
		fontc = "black"
	case 16:
		fillc = "3"
		fontc = "black"
	case 22:
		fillc = "4"
		fontc = "black"
	case 34:
		fillc = "5"
		fontc = "black"
	case 40:
		fillc = "6"
		fontc = "black"
	case 52:
		fillc = "7"
		fontc = "black"
	default: //not expected
		gNode.SetShape(cgraph.DiamondShape)
		gNode.SetColorScheme("")
		fillc = "firebrick"
		fontc = "white"
	}

	gNode.SetFillColor(fillc)
	gNode.SetFontColor(fontc)
}

func (r CTreeRenderer) RenderEdge(t ztalloc.TransformName, gedge *cgraph.Edge) {
	gedge.SetLabel(t)
	gedge.SetColorScheme("piyg7")
	// gedge.Set(cgraph.VeeArrow)

	var fontc string
	switch t {
	case ztalloc.Z22:
		fontc = "1"
	case ztalloc.Z2222:
		fontc = "2"
	case ztalloc.Z32:
		fontc = "3"
	case ztalloc.Z322:
		fontc = "4"
	case ztalloc.Z3222:
		fontc = "5"
	default: //not expected
		gedge.SetColorScheme("")
		fontc = "firebrick"
	}
	gedge.SetFontColor(fontc)
	gedge.SetColor(fontc)
}
