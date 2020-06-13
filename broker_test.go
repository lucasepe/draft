package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestBrokerComponentNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"br1"},
		{"br2"},
		{"br3"},
		{"br4"},
	}

	s := broker{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestBrokerComponent(t *testing.T) {
	want := `label="Message Broker",shape="cds",style="filled"`
	g := dot.NewGraph(dot.Directed)

	sketcher := broker{}
	sketcher.sketch(g, Component{})

	if got := flatten(g.String()); !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
