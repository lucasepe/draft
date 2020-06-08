package cluster

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
)

type Attribute func(*dot.Graph)

func Label(label string) Attribute {
	return func(el *dot.Graph) {
		el.Attr("label", label)
		if strings.TrimSpace(label) != "" {
			el.Attr("pencolor", "#f5deb3")
			el.Attr("style", "dashed")
		}
	}
}

func PenColor(color string) Attribute {
	return func(el *dot.Graph) {
		if strings.TrimSpace(color) != "" {
			el.Attr("pencolor", color)
		}
	}
}

func FontColor(color string) Attribute {
	return func(el *dot.Graph) {
		if strings.TrimSpace(color) != "" {
			el.Attr("fontcolor", color)
		} else {
			el.Attr("fontcolor", "#000000ff")
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

func New(parent *dot.Graph, id string, attrs ...Attribute) *dot.Graph {
	cluster := parent.Subgraph(id, dot.ClusterOption{})

	// default attributes
	FontName("Fira Mono Bold")(cluster)
	FontSize(9)(cluster)
	FontColor("#000000")
	PenColor("transparent")(cluster)

	for _, opt := range attrs {
		opt(cluster)
	}
	return cluster
}
