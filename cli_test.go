package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestClientComponentNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"cli1"},
		{"cli2"},
		{"cli3"},
		{"cli4"},
	}

	s := cli{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestClientComponent(t *testing.T) {
	want := `shape="underline",style="filled"`
	g := dot.NewGraph(dot.Directed)

	sketcher := cli{}
	sketcher.sketch(g, Component{})

	if got := flatten(g.String()); !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
