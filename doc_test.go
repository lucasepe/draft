package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestNoSQLNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"doc1"},
		{"doc2"},
		{"doc3"},
		{"doc4"},
	}

	s := doc{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestNoSQLShape(t *testing.T) {
	want := `shape="note",style="filled"`
	g := dot.NewGraph(dot.Directed)

	sketcher := doc{}
	sketcher.sketch(g, Component{})

	if got := flatten(g.String()); !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
