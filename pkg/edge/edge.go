package edge

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/emicklei/dot"
)

// Attribute is a function that apply a property to an edge.
type Attribute func(*dot.Edge)

// Label is the edge caption. If 'htm' is true the
// caption is treated as HTML code.
func Label(num int, text string) Attribute {
	return func(el *dot.Edge) {
		var render = false
		var sb strings.Builder
		sb.WriteString(`<table border="0">`)
		if num > 0 {
			render = true
			sb.WriteString(`<tr><td><font color="#1f6c7c" point-size="9"><b>`)
			sb.WriteString(strconv.Itoa(num))
			sb.WriteString("</b></font></td></tr>")
		}

		if lab := strings.TrimSpace(text); len(lab) > 0 {
			render = true
			sb.WriteString("<tr><td>")
			sb.WriteString(lab)
			sb.WriteString("</td></tr>")
		}
		sb.WriteString("</table>")

		if render {
			el.Attr("taillabel", dot.HTML(sb.String()))
		}
	}
}

// LabelDistance adjusts the distance that the
// headlabel(taillabel) is from the head(tail) node.
func LabelDistance(dist float32) Attribute {
	return func(el *dot.Edge) {

		el.Attr("labeldistance", fmt.Sprintf("%.2f", dist))
	}
}

// LabelAngle along with labeldistance, determine where
// the headlabel (taillabel) are placed with respect
// to the head (tail) in polar coordinates.
// The origin in the coordinate system is the point
// where the edge touches the node.
// The ray of 0 degrees goes from the origin back along
// the edge, parallel to the edge at the origin.
// The angle, in degrees, specifies the rotation from
// the 0 degree ray, with positive angles moving counterclockwise
// and negative angles moving clockwise.
func LabelAngle(angle float32) Attribute {
	return func(el *dot.Edge) {
		el.Attr("labelangle", fmt.Sprintf("%.2f", angle))
	}
}

// MinLen sets the minimum edge length (rank difference between head and tail).
func MinLen(len float32) Attribute {
	return func(el *dot.Edge) {
		if len <= 0 {
			return
		}
		el.Attr("minlen", fmt.Sprintf("%.2f", len))
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
