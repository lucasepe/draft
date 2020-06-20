package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestCDNFigure(t *testing.T) {
	tests := []struct {
		kind string
		find string
		want bool
	}{
		{kindCDN, `shape="Mcircle"`, true},
	}

	gfx := dot.NewGraph(dot.Directed)

	cfg := Config{}

	for _, tt := range tests {
		com := Component{Kind: tt.kind}

		t.Run(tt.find, func(t *testing.T) {
			figCDN(cfg, com)(gfx)

			x := flatten(gfx.String())
			if got := verify(x, tt.find); got != tt.want {
				t.Errorf("got [%v] want [%v] : %s", got, tt.want, x)
			}
		})
	}
}

func TestDNSFigure(t *testing.T) {
	tests := []struct {
		kind string
		find string
		want bool
	}{
		{kindDNS, `shape="Msquare"`, true},
	}

	gfx := dot.NewGraph(dot.Directed)

	cfg := Config{}

	for _, tt := range tests {
		com := Component{Kind: tt.kind}

		t.Run(tt.find, func(t *testing.T) {
			figDNS(cfg, com)(gfx)

			x := flatten(gfx.String())
			if got := verify(x, tt.find); got != tt.want {
				t.Errorf("got [%v] want [%v] : %s", got, tt.want, x)
			}
		})
	}
}

func TestLBAFigure(t *testing.T) {
	tests := []struct {
		kind string
		find string
		want bool
	}{
		{kindLBA, `shape="Mdiamond"`, true},
	}

	gfx := dot.NewGraph(dot.Directed)

	cfg := Config{}

	for _, tt := range tests {
		com := Component{Kind: tt.kind}

		t.Run(tt.find, func(t *testing.T) {
			figLBA(cfg, com)(gfx)

			x := flatten(gfx.String())
			if got := verify(x, tt.find); got != tt.want {
				t.Errorf("got [%v] want [%v] : %s", got, tt.want, x)
			}
		})
	}
}

func TestGatewayFigure(t *testing.T) {
	tests := []struct {
		kind string
		find string
		want bool
	}{
		{kindGateway, `shape="doublecircle"`, true},
	}

	gfx := dot.NewGraph(dot.Directed)

	cfg := Config{}

	for _, tt := range tests {
		com := Component{Kind: tt.kind}

		t.Run(tt.find, func(t *testing.T) {
			figGateway(cfg, com)(gfx)

			x := flatten(gfx.String())
			if got := verify(x, tt.find); got != tt.want {
				t.Errorf("got [%v] want [%v] : %s", got, tt.want, x)
			}
		})
	}
}
