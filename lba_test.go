package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestLoadBalancerNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"lba1"},
		{"lba2"},
		{"lba3"},
		{"lba4"},
	}

	s := lba{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestLoadBalancerShape(t *testing.T) {
	want := `label=<<b>LB</b>>,shape="Mdiamond",style="filled"`
	g := dot.NewGraph(dot.Directed)

	sketcher := lba{}
	sketcher.sketch(g, Component{})

	if got := flatten(g.String()); !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
