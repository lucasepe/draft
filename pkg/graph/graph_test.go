package graph

import (
	"strings"
	"testing"
)

func TestDefaultAttributes(t *testing.T) {
	g := New()

	want := `digraph  {fontname="Fira Mono Bold";fontsize="13.00";labelloc="t";newrank="true";nodesep="0.80";rankdir="LR";ranksep="1.10 equally";}`
	if got := flatten(g.String()); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestBottomTopLayout(t *testing.T) {
	g := New(BottomTop(true))

	want := `digraph  {fontname="Fira Mono Bold";fontsize="13.00";labelloc="t";newrank="true";nodesep="0.80";rankdir="BT";ranksep="1.10 equally";}`
	if got := flatten(g.String()); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

// remove tabs and newlines and spaces
func flatten(s string) string {
	return strings.Replace((strings.Replace(s, "\n", "", -1)), "\t", "", -1)
}
