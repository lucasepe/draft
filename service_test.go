package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestServiceComponentNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"ms1"},
		{"ms2"},
		{"ms3"},
		{"ms4"},
	}

	s := service{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestServiceComponent(t *testing.T) {
	g := dot.NewGraph(dot.Directed)

	sketcher := service{}
	sketcher.sketch(g, Component{})

	if got, want := flatten(g.String()), `shape="doubleoctagon",style="filled"`; !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
