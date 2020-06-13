package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestFunctionComponentNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"fn1"},
		{"fn2"},
		{"fn3"},
		{"fn4"},
	}

	s := function{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestFunctionComponent(t *testing.T) {
	want := `shape="signature",style="filled"`
	g := dot.NewGraph(dot.Directed)

	sketcher := function{}
	sketcher.sketch(g, Component{})

	if got := flatten(g.String()); !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
