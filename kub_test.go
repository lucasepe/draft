package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestContainerManagerNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"kub1"},
		{"kub2"},
		{"kub3"},
		{"kub4"},
	}

	s := kub{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestContainerManagerShape(t *testing.T) {
	want := `shape="square",style="filled"`
	g := dot.NewGraph(dot.Directed)

	sketcher := kub{}
	sketcher.sketch(g, Component{})

	if got := flatten(g.String()); !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
