package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestFileStoreNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"fst1"},
		{"fst2"},
		{"fst3"},
		{"fst4"},
	}

	s := fst{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestFileStoreShape(t *testing.T) {
	g := dot.NewGraph(dot.Directed)

	sketcher := fst{}
	sketcher.sketch(g, Component{})

	if got, want := flatten(g.String()), `shape="folder",style="filled"`; !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
