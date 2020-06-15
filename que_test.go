package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestQueueComponentNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"que1"},
		{"que2"},
		{"que3"},
		{"que4"},
	}

	s := que{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestQueueComponent(t *testing.T) {
	want := `label=<<table border="0" cellspacing="0" cellpadding="8"><tr><td border="0" colspan="3">&nbsp;</td></tr><tr><td border="1" bgcolor="#bdb76bff"><font point-size="6">msg N</font></td><td border="1" bgcolor="#bdb76bff"><font point-size="6">...</font></td><td border="1" bgcolor="#bdb76bff"><font point-size="6">msg 1</font></td></tr><tr><td border="0" colspan="3"><font point-size="7">&nbsp;</font></td></tr></table>>`
	g := dot.NewGraph(dot.Directed)

	sketcher := que{}
	sketcher.sketch(g, Component{})

	if got := flatten(g.String()); !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
