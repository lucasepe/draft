package edge

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
)

type Attribute func(*dot.Edge)

func Label(label string) Attribute {
	return func(el *dot.Edge) {
		el.Attr("label", label)
	}
}

func FontName(name string) Attribute {
	return func(el *dot.Edge) {
		el.Attr("fontname", name)
	}
}

func FontSize(size float32) Attribute {
	return func(el *dot.Edge) {
		fs := fmt.Sprintf("%.2f", size)
		el.Attr("fontsize", fs)
	}
}

func Dir(dir string) Attribute {
	return func(el *dot.Edge) {
		if strings.TrimSpace(dir) != "" {
			el.Attr("dir", dir)
		}
	}
}

func Dashed(dashed bool) Attribute {
	return func(el *dot.Edge) {
		if dashed {
			el.Attr("style", "dashed")
		}
	}
}

func Color(color string) Attribute {
	return func(el *dot.Edge) {
		if strings.TrimSpace(color) != "" {
			el.Attr("color", color)
		} else {
			el.Attr("color", "#708090ff")
		}
	}
}

func PenWidth(size float32) Attribute {
	return func(el *dot.Edge) {
		pw := fmt.Sprintf("%.2f", size)
		el.Attr("penwidth", pw)
	}
}

func ArrowSize(size float32) Attribute {
	return func(el *dot.Edge) {
		pw := fmt.Sprintf("%.2f", size)
		el.Attr("arrowsize", pw)
	}
}

func Highlight(ok bool) Attribute {
	return func(el *dot.Edge) {
		if ok {
			el.Attr("penwidth", "1.2")
			el.Attr("arrowsize", "0.9")
		} else {
			el.Attr("penwidth", "0.6")
			el.Attr("arrowsize", "0.6")
		}
	}
}

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
