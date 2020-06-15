package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestGatewayNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"gtw1"},
		{"gtw2"},
		{"gtw3"},
		{"gtw4"},
	}

	s := gtw{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestGatewayShape(t *testing.T) {
	want := `label=<<b>GTW</b>>,shape="doublecircle",style="filled"`
	g := dot.NewGraph(dot.Directed)

	sketcher := gtw{}
	sketcher.sketch(g, Component{})

	if got := flatten(g.String()); !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
