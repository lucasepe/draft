package cluster

import (
	"regexp"
	"strings"
	"testing"

	"github.com/emicklei/dot"
)

func TestDefaultAttributes(t *testing.T) {
	want := `digraph  {subgraph cluster_s1 {fontname="Fira Mono Bold";fontsize="9.00";label="DUMMY";pencolor="transparent";}}`

	di := dot.NewGraph(dot.Directed)
	New(di, "DUMMY")

	if got := flatten(di.String()); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestFontColor(t *testing.T) {
	di := dot.NewGraph(dot.Directed)
	New(di, "DUMMY", FontColor("#fafafaff"))

	want := `fontcolor="#fafafaff";fontname="Fira Mono Bold";fontsize="9.00";label="DUMMY"`
	got := flatten(di.String())
	if !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestFontName(t *testing.T) {
	di := dot.NewGraph(dot.Directed)
	New(di, "DUMMY", FontName("Inconsolata"))

	want := `fontname="Inconsolata";fontsize="9.00";label="DUMMY"`
	got := flatten(di.String())
	if !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestFontSize(t *testing.T) {
	di := dot.NewGraph(dot.Directed)
	New(di, "DUMMY", FontSize(14))

	want := `fontsize="14.00";label="DUMMY"`
	got := flatten(di.String())
	if !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestPenColor(t *testing.T) {
	di := dot.NewGraph(dot.Directed)
	New(di, "DUMMY", PenColor("red"))

	want := `fontname="Fira Mono Bold";fontsize="9.00";label="DUMMY";pencolor="red"`
	got := flatten(di.String())
	if !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestLabel(t *testing.T) {
	di := dot.NewGraph(dot.Directed)
	New(di, "DUMMY", Label("AaA"))

	want := `fontname="Fira Mono Bold";fontsize="9.00";label="AaA";pencolor="transparent"`
	got := flatten(di.String())
	if !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestBottomTop(t *testing.T) {
	di := dot.NewGraph(dot.Directed)
	New(di, "DUMMY", BottomTop(true))

	want := `fontname="Fira Mono Bold";fontsize="9.00";label="DUMMY";labelloc="b"`
	got := flatten(di.String())
	if !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

// remove tabs and newlines and spaces
func flatten(s string) string {
	return strings.Replace((strings.Replace(s, "\n", "", -1)), "\t", "", -1)
}

func verify(src, kv string) bool {
	r := regexp.MustCompile(`\{(.*)\}`)
	if m := r.FindAllStringSubmatch(src, -1); len(m) > 0 {
		return strings.Contains(m[0][1], kv)
	}
	return false
}
