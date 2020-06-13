package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestStorageComponentNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"st1"},
		{"st2"},
		{"st3"},
		{"st4"},
	}

	s := storage{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestStorageComponent(t *testing.T) {
	g := dot.NewGraph(dot.Directed)

	sketcher := storage{}
	sketcher.sketch(g, Component{})

	if got, want := flatten(g.String()), `shape="folder",style="filled"`; !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
