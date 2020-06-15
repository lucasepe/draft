package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestPubSubNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"msg1"},
		{"msg2"},
		{"msg3"},
		{"msg4"},
	}

	s := msg{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestPubSubShape(t *testing.T) {
	want := `label=<<b>Pub / Sub</b>>,shape="cds",style="filled"`
	g := dot.NewGraph(dot.Directed)

	sketcher := msg{}
	sketcher.sketch(g, Component{})

	if got := flatten(g.String()); !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
