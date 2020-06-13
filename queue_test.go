package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestQueueComponentNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"qs1"},
		{"qs2"},
		{"qs3"},
		{"qs4"},
	}

	s := queue{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestQueueComponent(t *testing.T) {
	want := `label=<<table border="0" cellspacing="0" cellpadding="8"><tr><td border="0" colspan="3">&nbsp;</td></tr><tr><td border="1" style="dashed" bgcolor="#bdb76bff">msg N</td><td border="1" style="dashed" bgcolor="#bdb76bff">...</td><td border="1" style="dashed" bgcolor="#bdb76bff">msg 1</td></tr><tr><td border="0" colspan="3"><font point-size="7">&nbsp;</font></td></tr></table>>,shape="plain"`
	g := dot.NewGraph(dot.Directed)

	sketcher := queue{}
	sketcher.sketch(g, Component{})

	if got := flatten(g.String()); !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
