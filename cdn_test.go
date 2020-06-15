package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestCDNCNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"cdn1"},
		{"cdn2"},
		{"cdn3"},
		{"cdn4"},
	}

	s := cdn{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestCDNShape(t *testing.T) {
	want := `shape="Mcircle",style="filled"`
	g := dot.NewGraph(dot.Directed)

	sketcher := cdn{}
	sketcher.sketch(g, Component{})

	if got := flatten(g.String()); !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
