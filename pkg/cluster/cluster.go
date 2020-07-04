package cluster

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
)

// Attribute is a function that apply a property to a cluster.
type Attribute func(*dot.Graph)

// Label is the cluster caption.
func Label(label string) Attribute {
	return func(el *dot.Graph) {
		el.Attr("label", label)
	}
}

// PenColor set the color used to draw the
// bounding box around a cluster.
func PenColor(color string) Attribute {
	return func(el *dot.Graph) {
		if strings.TrimSpace(color) != "" {
			el.Attr("pencolor", color)
		}
	}
}

// FontColor specify the text color.
func FontColor(color string) Attribute {
	return func(el *dot.Graph) {
		if strings.TrimSpace(color) != "" {
			el.Attr("fontcolor", color)
		} else {
			el.Attr("fontcolor", "#000000ff")
		}
	}
}

// FontName specify the font used for text.
func FontName(name string) Attribute {
	return func(el *dot.Graph) {
		el.Attr("fontname", name)
	}
}

// FontSize specify the font size, in points, used for text.
func FontSize(size float32) Attribute {
	return func(el *dot.Graph) {
		fs := fmt.Sprintf("%.2f", size)
		el.Attr("fontsize", fs)
	}
}

// BottomTop sets the label location according to layout.
func BottomTop(bt bool) Attribute {
	return func(el *dot.Graph) {
		if bt {
			el.Attr("labelloc", "b")
		} else {
			el.Attr("labelloc", "t")
		}
	}
}

// New create a new cluster with the specified attributes.
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
