package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestBlockStorageNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"bst1"},
		{"bst2"},
		{"bst3"},
		{"bst4"},
	}

	s := bst{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestBlockStorageShape(t *testing.T) {
	want := `label=<<table border="0" cellspacing="4" cellpadding="4"><tr>  <td border="1" bgcolor="#606f5cff" width="30"></td>  <td border="1" bgcolor="#606f5cff"></td></tr><tr>  <td border="1" bgcolor="#606f5cff" width="30"></td>  <td border="1" bgcolor="#606f5cff"></td></tr></table>>`
	g := dot.NewGraph(dot.Directed)

	sketcher := bst{}
	sketcher.sketch(g, Component{})

	if got := flatten(g.String()); !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
