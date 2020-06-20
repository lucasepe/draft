package graph

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
)

// Attribute is a function that apply a property to a Graph.
type Attribute func(*dot.Graph)

// Label is the Graph title.
func Label(label string) Attribute {
	return func(el *dot.Graph) {
		if strings.TrimSpace(label) != "" {
			el.Attr("label", label)
		}
	}
}

// FontName specify the font used for the Graph title.
func FontName(name string) Attribute {
	return func(el *dot.Graph) {
		el.Attr("fontname", name)
	}
}

// FontSize specify the font size, in points, used for title.
func FontSize(size float32) Attribute {
	return func(el *dot.Graph) {
		fs := fmt.Sprintf("%.2f", size)
		el.Attr("fontsize", fs)
	}
}

// LeftToRight sets direction of Graph layout from left to right.
func LeftToRight() Attribute {
	return func(el *dot.Graph) {
		el.Attr("rankdir", "LR")
	}
}

// BottomTop sets direction of Graph layout from bottom to top.
func BottomTop(enable bool) Attribute {
	return func(el *dot.Graph) {
		if enable {
			el.Attr("rankdir", "BT")
		}
	}
}

// RankSep gives the desired rank separation, in inches.
// This is the minimum vertical distance between the bottom
// of the nodes in one rank and the tops of nodes in the next.
func RankSep(size float32) Attribute {
	return func(el *dot.Graph) {
		fs := fmt.Sprintf("%.2f equally", size)
		el.Attr("ranksep", fs)
	}
}

// NodeSep specifies the minimum space between two
// adjacent nodes in the same rank, in inches.
func NodeSep(size float32) Attribute {
	return func(el *dot.Graph) {
		fs := fmt.Sprintf("%.2f", size)
		el.Attr("nodesep", fs)
	}
}

// Ortho controls how edges are represented, if true lines are orthogonal.
func Ortho(enable bool) Attribute {
	return func(el *dot.Graph) {
		if enable {
			el.Attr("splines", "ortho")
		} else {
			el.Attr("splines", "line")
		}
	}
}

// BackgroundColor sets the Graph background color.
func BackgroundColor(color string) Attribute {
	return func(el *dot.Graph) {
		if strings.TrimSpace(color) != "" {
			el.Attr("bgcolor", color)
		} else {
			el.Attr("bgcolor", "transparent")
		}
	}
}

// New create a new Graph with the specified attributes.
func New(attrs ...Attribute) *dot.Graph {
	el := dot.NewGraph(dot.Directed)
	el.Attr("newrank", "true")
	el.Attr("labelloc", "t")

	FontName("Fira Mono Bold")(el)
	FontSize(13)(el)
	LeftToRight()(el)
	RankSep(1.1)(el)
	NodeSep(0.4)(el)

	for _, opt := range attrs {
		opt(el)
	}

	return el
}
