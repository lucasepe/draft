package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestGatewayComponentNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"gt1"},
		{"gt2"},
		{"gt3"},
		{"gt4"},
	}

	s := gateway{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestGatewayComponent(t *testing.T) {
	want := `label="GW",shape="doublecircle",style="filled"`
	g := dot.NewGraph(dot.Directed)

	sketcher := gateway{}
	sketcher.sketch(g, Component{})

	if got := flatten(g.String()); !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
