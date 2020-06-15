package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestDNSComponentNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"dn1"},
		{"dn2"},
		{"dn3"},
		{"dn4"},
	}

	s := dns{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestDNSComponent(t *testing.T) {
	want := `label=<<b>DNS</b>>,shape="Msquare",style="filled"`
	g := dot.NewGraph(dot.Directed)

	sketcher := dns{}
	sketcher.sketch(g, Component{})

	if got := flatten(g.String()); !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
