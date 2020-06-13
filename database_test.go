package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestDatabaseComponentNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"db1"},
		{"db2"},
		{"db3"},
		{"db4"},
	}

	s := database{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestDatabaseComponent(t *testing.T) {
	want := `shape="cylinder",style="filled"`
	g := dot.NewGraph(dot.Directed)

	sketcher := database{}
	sketcher.sketch(g, Component{})

	if got := flatten(g.String()); !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
