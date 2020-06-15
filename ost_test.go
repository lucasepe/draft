package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestObjectStoreNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"ost1"},
		{"ost2"},
		{"ost3"},
		{"ost4"},
	}

	s := ost{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestObjectStoreShape(t *testing.T) {
	g := dot.NewGraph(dot.Directed)

	sketcher := ost{}
	sketcher.sketch(g, Component{})

	if got, want := flatten(g.String()), `shape="invtrapezium",style="filled"`; !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
