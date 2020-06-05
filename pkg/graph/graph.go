package graph

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
)

type Attribute func(*dot.Graph)

func Label(label string) Attribute {
	return func(el *dot.Graph) {
		if strings.TrimSpace(label) != "" {
			el.Attr("label", label)
		}
	}
}

func FontName(name string) Attribute {
	return func(el *dot.Graph) {
		el.Attr("fontname", name)
	}
}

func FontSize(size float32) Attribute {
	return func(el *dot.Graph) {
		fs := fmt.Sprintf("%.2f", size)
		el.Attr("fontsize", fs)
	}
}

func LeftToRight() Attribute {
	return func(el *dot.Graph) {
		el.Attr("rankdir", "LR")
	}
}

func TopToBottom() Attribute {
	return func(el *dot.Graph) {
		el.Attr("rankdir", "TB")
	}
}

func RankSep(size float32) Attribute {
	return func(el *dot.Graph) {
		fs := fmt.Sprintf("%.2f", size)
		el.Attr("ranksep", fs)
	}
}

func BackgroundColor(color string) Attribute {
	return func(el *dot.Graph) {
		if strings.TrimSpace(color) != "" {
			el.Attr("bgcolor", color)
		} else {
			el.Attr("bgcolor", "transparent")
		}
	}
}

func New(attrs ...Attribute) *dot.Graph {
	el := dot.NewGraph(dot.Directed)

	FontName("Fira Mono Bold")(el)
	FontSize(13)(el)
	LeftToRight()(el)
	RankSep(1.1)(el)

	for _, opt := range attrs {
		opt(el)
	}

	return el
}
