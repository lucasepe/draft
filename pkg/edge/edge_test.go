package edge

import (
	"strings"
	"testing"

	"github.com/emicklei/dot"
)

func TestDefaultAttributes(t *testing.T) {
	di := dot.NewGraph(dot.Directed)
	di.Node("A")
	di.Node("B")

	New(di, "A", "B")

	want := `digraph  {n1[label="A"];n2[label="B"];n1->n2[arrowsize="0.6",fontname="Fira Mono",fontsize="8.00",penwidth="0.6"];}`
	if got := flatten(di.String()); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestHighlightWithLabel(t *testing.T) {
	di := dot.NewGraph(dot.Directed)
	di.Node("A")
	di.Node("B")

	New(di, "A", "B", Highlight(true), Label("Go!"))

	want := `digraph  {n1[label="A"];n2[label="B"];n1->n2[arrowsize="0.9",fontname="Fira Mono",fontsize="8.00",label="Go!",penwidth="1.2"];}`
	if got := flatten(di.String()); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

// remove tabs and newlines and spaces
func flatten(s string) string {
	return strings.Replace((strings.Replace(s, "\n", "", -1)), "\t", "", -1)
}
