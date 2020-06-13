package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestLoadBalancerComponentNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"lb1"},
		{"lb2"},
		{"lb3"},
		{"lb4"},
	}

	s := balancer{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestLoadBalancerComponent(t *testing.T) {
	want := `label="LB",shape="Mdiamond",style="filled"`
	g := dot.NewGraph(dot.Directed)

	sketcher := balancer{}
	sketcher.sketch(g, Component{})

	if got := flatten(g.String()); !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
