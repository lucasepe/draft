package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestCacheNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"mem1"},
		{"mem2"},
		{"mem3"},
		{"mem4"},
	}

	s := mem{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestCacheShape(t *testing.T) {
	want := `label=<<table border="0" cellspacing="1" cellpadding="4"><tr><td border="0" colspan="2">&nbsp;</td></tr><tr>  <td bgcolor="#f04d30ff"><font color="#f5f5f5" point-size="7">key</font></td>  <td bgcolor="#f04d30ff"><font color="#f5f5f5" point-size="7">val</font></td></tr><tr>  <td bgcolor="#f04d30ff" style="dashed"></td>  <td bgcolor="#f04d30ff" style="dashed"></td></tr><tr>  <td bgcolor="#f04d30ff" style="dashed"></td>  <td bgcolor="#f04d30ff" style="dashed"></td></tr><tr>  <td colspan="2">&nbsp;</td>    </tr></table>>`
	g := dot.NewGraph(dot.Directed)

	sketcher := mem{}
	sketcher.sketch(g, Component{})

	if got := flatten(g.String()); !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
