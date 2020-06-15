package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestFunctionNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"fun1"},
		{"fun2"},
		{"fun3"},
		{"fun4"},
	}

	s := fun{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestFunctionShape(t *testing.T) {
	want := `shape="signature",style="filled"`
	g := dot.NewGraph(dot.Directed)

	sketcher := fun{}
	sketcher.sketch(g, Component{})

	if got := flatten(g.String()); !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
