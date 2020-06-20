package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestPubSubFigure(t *testing.T) {
	tests := []struct {
		kind string
		find string
		want bool
	}{
		{kindPubSub, `shape="cds"`, true},
	}

	gfx := dot.NewGraph(dot.Directed)

	cfg := Config{}

	for _, tt := range tests {
		com := Component{Kind: tt.kind}

		t.Run(tt.find, func(t *testing.T) {
			figPubSub(cfg, com)(gfx)

			x := flatten(gfx.String())
			if got := verify(x, tt.find); got != tt.want {
				t.Errorf("got [%v] want [%v] : %s", got, tt.want, x)
			}
		})
	}
}

func TestQueueFigure(t *testing.T) {
	tests := []struct {
		kind string
		find string
		want bool
	}{
		{kindQueue, `label=<<table border="0" cellspacing="0" cellpadding="8"><tr><td border="0" colspan="3">&nbsp;</td></tr><tr><td border="1" bgcolor="#bdb76bff"><font point-size="6">msg N</font></td><td border="1" bgcolor="#bdb76bff"><font point-size="6">...</font></td><td border="1" bgcolor="#bdb76bff"><font point-size="6">msg 1</font></td></tr><tr><td border="0" colspan="3"><font point-size="7">&nbsp;</font></td></tr></table>>`, true},
	}

	gfx := dot.NewGraph(dot.Directed)

	cfg := Config{}

	for _, tt := range tests {
		com := Component{Kind: tt.kind}

		t.Run(tt.find, func(t *testing.T) {
			figQueue(cfg, com)(gfx)

			x := flatten(gfx.String())
			if got := verify(x, tt.find); got != tt.want {
				t.Errorf("got [%v] want [%v] : %s", got, tt.want, x)
			}
		})
	}
}
