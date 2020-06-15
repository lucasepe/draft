package node

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
)

// Attribute is a function that apply a property to a node.
type Attribute func(*dot.Node)

// Label is the node caption. If 'htm' is true the
// caption is treated as HTML code.
func Label(label string, htm bool) Attribute {
	return func(el *dot.Node) {
		if htm {
			el.Attr("label", dot.HTML(label))
		} else {
			el.Attr("label", label)
		}
	}
}

// Shape sets the shape of a node.
func Shape(shape string) Attribute {
	return func(el *dot.Node) {
		el.Attr("shape", shape)
	}
}

// Rounded sets the shape with round corners.
func Rounded(rounded bool) Attribute {
	return func(el *dot.Node) {
		if rounded {
			el.Attr("style", "rounded,filled")
		} else {
			el.Attr("style", "filled")
		}
	}
}

// FillColor sets the node fill color.
func FillColor(color string) Attribute {
	return func(el *dot.Node) {
		if strings.TrimSpace(color) != "" {
			el.Attr("fillcolor", color)
		}
	}
}

// FontColor specify the text color.
func FontColor(color string) Attribute {
	return func(el *dot.Node) {
		if strings.TrimSpace(color) != "" {
			el.Attr("fontcolor", color)
		}
	}
}

// FontName specify the font used for text.
func FontName(name string) Attribute {
	return func(el *dot.Node) {
		el.Attr("fontname", name)
	}
}

// FontSize specify the font size, in points, used for text.
func FontSize(size float32) Attribute {
	return func(el *dot.Node) {
		fs := fmt.Sprintf("%.2f", size)
		el.Attr("fontsize", fs)
	}
}

// New create a new node with the specified attributes.
func New(cluster *dot.Graph, id string, attrs ...Attribute) *dot.Node {
	el := cluster.Node(id)

	// default attributes
	Rounded(false)(&el)
	FontName("Fira Mono")(&el)
	FontSize(9)(&el)

	for _, opt := range attrs {
		opt(&el)
	}

	return &el
}
