package node

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
)

type Attribute func(*dot.Node)

func Label(label string, htm bool) Attribute {
	return func(el *dot.Node) {
		if htm {
			el.Attr("label", dot.HTML(label))
		} else {
			el.Attr("label", label)
		}
	}
}

func Shape(shape string) Attribute {
	return func(el *dot.Node) {
		el.Attr("shape", shape)
	}
}

func Rounded(rounded bool) Attribute {
	return func(el *dot.Node) {
		if rounded {
			el.Attr("style", "rounded,filled")
		} else {
			el.Attr("style", "filled")
		}
	}
}

func FillColor(color, fallback string) Attribute {
	return func(el *dot.Node) {
		if strings.TrimSpace(color) != "" {
			el.Attr("fillcolor", color)
		} else {
			el.Attr("fillcolor", fallback)
		}
	}
}

func FontColor(color string) Attribute {
	return func(el *dot.Node) {
		if strings.TrimSpace(color) != "" {
			el.Attr("fontcolor", color)
		} else {
			el.Attr("fontcolor", "#000000ff")
		}
	}
}

func FontName(name string) Attribute {
	return func(el *dot.Node) {
		el.Attr("fontname", name)
	}
}

func FontSize(size float32) Attribute {
	return func(el *dot.Node) {
		fs := fmt.Sprintf("%.2f", size)
		el.Attr("fontsize", fs)
	}
}

func New(cluster *dot.Graph, id string, attrs ...Attribute) *dot.Node {
	el := cluster.Node(id)
	el.Attr("style", "filled")

	// default attributes
	FontName("Fira Mono")(&el)
	FontSize(9)(&el)

	for _, opt := range attrs {
		opt(&el)
	}

	return &el
}
