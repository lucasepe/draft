package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestFirewallFigure(t *testing.T) {
	tests := []struct {
		kind string
		find string
		want bool
	}{
		{kindFirewall, `shape="invhouse"`, true},
	}

	gfx := dot.NewGraph(dot.Directed)

	cfg := Config{}

	for _, tt := range tests {
		com := Component{Kind: tt.kind}

		t.Run(tt.find, func(t *testing.T) {
			figFirewall(cfg, com)(gfx)

			x := flatten(gfx.String())
			if got := verify(x, tt.find); got != tt.want {
				t.Errorf("got [%v] want [%v] : %s", got, tt.want, x)
			}
		})
	}
}
