package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestCSComponentNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"cos1"},
		{"cos2"},
		{"cos3"},
		{"cos4"},
	}

	s := containerService{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestCSComponent(t *testing.T) {
	want := `shape="component",style="filled"`
	g := dot.NewGraph(dot.Directed)

	sketcher := containerService{}
	sketcher.sketch(g, Component{})

	if got := flatten(g.String()); !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
