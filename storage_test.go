package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestBlockStoreFigure(t *testing.T) {
	tests := []struct {
		kind string
		find string
		want bool
	}{
		{kindBlockStore, `label=<<table border="0" cellspacing="4" cellpadding="4"><tr>  <td border="1" bgcolor="#606f5cff" width="30"></td>  <td border="1" bgcolor="#606f5cff"></td></tr><tr>  <td border="1" bgcolor="#606f5cff" width="30"></td>  <td border="1" bgcolor="#606f5cff"></td></tr></table>>`, true},
	}

	gfx := dot.NewGraph(dot.Directed)

	cfg := Config{}

	for _, tt := range tests {
		com := Component{Kind: tt.kind}

		t.Run(tt.find, func(t *testing.T) {
			figBlockStore(cfg, com)(gfx)

			x := flatten(gfx.String())
			if got := verify(x, tt.find); got != tt.want {
				t.Errorf("got [%v] want [%v] : %s", got, tt.want, x)
			}
		})
	}
}

func TestFileStoreFigure(t *testing.T) {
	tests := []struct {
		kind string
		find string
		want bool
	}{
		{kindFileStore, `shape="folder"`, true},
	}

	gfx := dot.NewGraph(dot.Directed)

	cfg := Config{}

	for _, tt := range tests {
		com := Component{Kind: tt.kind}

		t.Run(tt.find, func(t *testing.T) {
			figFileStore(cfg, com)(gfx)

			x := flatten(gfx.String())
			if got := verify(x, tt.find); got != tt.want {
				t.Errorf("got [%v] want [%v] : %s", got, tt.want, x)
			}
		})
	}
}

func TestObjectStoreFigure(t *testing.T) {
	tests := []struct {
		kind string
		find string
		want bool
	}{
		{kindObjectStore, `shape="invtrapezium"`, true},
	}

	gfx := dot.NewGraph(dot.Directed)

	cfg := Config{}

	for _, tt := range tests {
		com := Component{Kind: tt.kind}

		t.Run(tt.find, func(t *testing.T) {
			figObjectStore(cfg, com)(gfx)

			x := flatten(gfx.String())
			if got := verify(x, tt.find); got != tt.want {
				t.Errorf("got [%v] want [%v] : %s", got, tt.want, x)
			}
		})
	}
}
