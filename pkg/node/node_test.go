package node

import (
	"strings"
	"testing"

	"github.com/emicklei/dot"
)

func TestDefaultAttributes(t *testing.T) {
	di := dot.NewGraph(dot.Directed)
	New(di, "NODE_1")

	want := `digraph  {n1[fontname="Fira Mono",fontsize="9.00",label="NODE_1",style="filled"];}`
	if got := flatten(di.String()); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestFontColorFallback(t *testing.T) {
	di := dot.NewGraph(dot.Directed)
	New(di, "NODE_1", FontColor("", "#000000ff"))

	want := `digraph  {n1[fontcolor="#000000ff",fontname="Fira Mono",fontsize="9.00",label="NODE_1",style="filled"];}`
	if got := flatten(di.String()); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestFontColor(t *testing.T) {
	di := dot.NewGraph(dot.Directed)
	New(di, "NODE_1", FontColor("#fafafaff", "#000000"))

	want := `digraph  {n1[fontcolor="#fafafaff",fontname="Fira Mono",fontsize="9.00",label="NODE_1",style="filled"];}`
	if got := flatten(di.String()); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestFillColorFallback(t *testing.T) {
	di := dot.NewGraph(dot.Directed)
	New(di, "NODE_1", FillColor("", "#ff0000ff"))

	want := `digraph  {n1[fillcolor="#ff0000ff",fontname="Fira Mono",fontsize="9.00",label="NODE_1",style="filled"];}`
	if got := flatten(di.String()); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestShape(t *testing.T) {
	di := dot.NewGraph(dot.Directed)
	New(di, "NODE_1", Shape("box"))

	want := `digraph  {n1[fontname="Fira Mono",fontsize="9.00",label="NODE_1",shape="box",style="filled"];}`
	if got := flatten(di.String()); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestPlainLabel(t *testing.T) {
	di := dot.NewGraph(dot.Directed)
	New(di, "NODE_1", Label("Service", false))

	want := `digraph  {n1[fontname="Fira Mono",fontsize="9.00",label="Service",style="filled"];}`
	if got := flatten(di.String()); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestHTMLLabel(t *testing.T) {
	di := dot.NewGraph(dot.Directed)
	New(di, "NODE_1", Label("<b>Service<b>", true))

	want := `digraph  {n1[fontname="Fira Mono",fontsize="9.00",label=<<b>Service<b>>,style="filled"];}`
	if got := flatten(di.String()); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

// remove tabs and newlines and spaces
func flatten(s string) string {
	return strings.Replace((strings.Replace(s, "\n", "", -1)), "\t", "", -1)
}
