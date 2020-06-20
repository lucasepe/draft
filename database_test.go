package draft

import (
	"testing"

	"github.com/emicklei/dot"
)

func TestRDBFigure(t *testing.T) {
	tests := []struct {
		kind string
		find string
		want bool
	}{
		{kindRDB, `shape="tab"`, true},
	}

	gfx := dot.NewGraph(dot.Directed)

	cfg := Config{}

	for _, tt := range tests {
		com := Component{Kind: tt.kind}

		t.Run(tt.find, func(t *testing.T) {
			figRDB(cfg, com)(gfx)

			x := flatten(gfx.String())
			if got := verify(x, tt.find); got != tt.want {
				t.Errorf("got [%v] want [%v] : %s", got, tt.want, x)
			}
		})
	}
}

func TestNoSQLFigure(t *testing.T) {
	tests := []struct {
		kind string
		find string
		want bool
	}{
		{kindNoSQL, `shape="note"`, true},
	}

	gfx := dot.NewGraph(dot.Directed)

	cfg := Config{}

	for _, tt := range tests {
		com := Component{Kind: tt.kind}

		t.Run(tt.find, func(t *testing.T) {
			figNoSQL(cfg, com)(gfx)

			x := flatten(gfx.String())
			if got := verify(x, tt.find); got != tt.want {
				t.Errorf("got [%v] want [%v] : %s", got, tt.want, x)
			}
		})
	}
}

func TestCacheFigure(t *testing.T) {
	tests := []struct {
		kind string
		find string
		want bool
	}{
		{kindCache, `label=<<table border="0" cellspacing="1" cellpadding="4"><tr><td border="0" colspan="2">&nbsp;</td></tr><tr>  <td bgcolor="#f04d30ff"><font color="#f5f5f5" point-size="7">key</font></td>  <td bgcolor="#f04d30ff"><font color="#f5f5f5" point-size="7">val</font></td></tr><tr>  <td bgcolor="#f04d30ff" style="dashed"></td>  <td bgcolor="#f04d30ff" style="dashed"></td></tr><tr>  <td bgcolor="#f04d30ff" style="dashed"></td>  <td bgcolor="#f04d30ff" style="dashed"></td></tr><tr>  <td colspan="2">Cache</td>    </tr></table>>`, true},
	}

	gfx := dot.NewGraph(dot.Directed)

	cfg := Config{}

	for _, tt := range tests {
		com := Component{Kind: tt.kind}

		t.Run(tt.find, func(t *testing.T) {
			figCache(cfg, com)(gfx)

			x := flatten(gfx.String())
			if got := verify(x, tt.find); got != tt.want {
				t.Errorf("got [%v] want [%v] : %s", got, tt.want, x)
			}
		})
	}
}
