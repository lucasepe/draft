package edge

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
)

// Attribute is a function that apply a property to an edge.
type Attribute func(*dot.Edge)

// Label is the text attached to components.
func Label(label string) Attribute {
	return func(el *dot.Edge) {
		el.Attr("label", label)
	}
}

// FontName specify the font used for text.
func FontName(name string) Attribute {
	return func(el *dot.Edge) {
		el.Attr("fontname", name)
	}
}

// FontSize specify the font size, in points, used for text.
func FontSize(size float32) Attribute {
	return func(el *dot.Edge) {
		fs := fmt.Sprintf("%.2f", size)
		el.Attr("fontsize", fs)
	}
}

// Dir sets the ege direction, values: both, forward, back, none.
func Dir(dir string) Attribute {
	return func(el *dot.Edge) {
		if strings.TrimSpace(dir) != "" {
			el.Attr("dir", dir)
		}
	}
}

// Dashed set the edge line dashed.
func Dashed(dashed bool) Attribute {
	return func(el *dot.Edge) {
		if dashed {
			el.Attr("style", "dashed")
		}
	}
}

// Color set the color for an edge line.
func Color(color string) Attribute {
	return func(el *dot.Edge) {
		if strings.TrimSpace(color) != "" {
			el.Attr("color", color)
		} else {
			el.Attr("color", "#708090ff")
		}
	}
}

// Highlight makes the line thicker.
func Highlight(enable bool) Attribute {
	return func(el *dot.Edge) {
		if enable {
			el.Attr("penwidth", "1.2")
			el.Attr("arrowsize", "0.9")
		} else {
			el.Attr("penwidth", "0.6")
			el.Attr("arrowsize", "0.6")
		}
	}
}

// New add to dot.Graph a new connection line between two components.
func New(g *dot.Graph, fromNodeID, toNodeID string, attrs ...Attribute) error {
	n1, ok := g.FindNodeById(fromNodeID)
	if !ok {
		return fmt.Errorf("node with id=%s not found", fromNodeID)
	}

	n2, ok := g.FindNodeById(toNodeID)
	if !ok {
		return fmt.Errorf("node with id=%s not found", toNodeID)
	}

	el := g.Edge(n1, n2)

	FontName("Fira Mono")(&el)
	FontSize(8)(&el)
	Highlight(false)(&el)

	for _, opt := range attrs {
		opt(&el)
	}

	return nil
}
