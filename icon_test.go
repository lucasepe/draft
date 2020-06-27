package draft

import (
	"os"
	"testing"

	"github.com/emicklei/dot"
)

func TestIconFigure(t *testing.T) {
	tests := []struct {
		provider string
		kind     string
		find     string
		want     bool
	}{
		{"aws", kindFirewall, `src="icons/aws/waf.png"`, true},
		{"aws", kindFunction, `src="icons/aws/fun.png"`, true},
		{"google", kindRDB, `src="icons/google/rdb.png"`, true},
		{"google", kindCache, `src="icons/google/mem.png"`, true},
		{"azure", kindDNS, `src="icons/azure/dns.png"`, true},
		{"azure", kindCDN, `src="icons/azure/cdn.png"`, true},
	}

	os.Setenv("DRAFT_ICONS_PATH", "./icons")

	gfx := dot.NewGraph(dot.Directed)

	for _, tt := range tests {
		cfg := Config{provider: tt.provider}
		com := Component{Kind: tt.kind}

		t.Run(tt.find, func(t *testing.T) {
			figIcon(cfg, com)(gfx)

			x := flatten(gfx.String())
			if got := verify(x, tt.find); got != tt.want {
				t.Errorf("got [%v] want [%v] : %s", got, tt.want, x)
			}
		})
	}
}
