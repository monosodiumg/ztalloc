package render

import (
	"fmt"
	"ztalloc/ztalloc"

	"github.com/goccy/go-graphviz/cgraph"
)

type ()

func RenderNode(z ztalloc.Node, gnode *cgraph.Node) {
	label := fmt.Sprintf("%d(%d)", z.Value(), z.Value()%54)
	gnode.SetLabel(label)
	gnode.SetColorScheme("oranges9")
	gnode.SetStyle(cgraph.FilledNodeStyle)
	gnode.SetShape(cgraph.OvalShape)

	var fillc, fontc string
	switch z.Value() % 54 {
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
		gnode.SetShape(cgraph.DiamondShape)
		gnode.SetColorScheme("")
		fillc = "firebrick"
		fontc = "white"
	}

	gnode.SetFillColor(fillc)
	gnode.SetFontColor(fontc)
}

func RenderEdge(t ztalloc.TransformName, gedge *cgraph.Edge) {
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
