package draft

import (
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

func init() {
	register(kindFirewall, figFirewall)
}

// figFirewall render a Component of `kindFirewall`.
func figFirewall(ctx Config, com Component) func(gfx *dot.Graph) bool {
	return func(gfx *dot.Graph) bool {
		if !strings.EqualFold(com.Kind, kindFirewall) {
			return false
		}

		if fc := strings.TrimSpace(com.FontColor); len(fc) == 0 {
			com.FontColor = "#fafafaff"
		}

		if fc := strings.TrimSpace(com.FillColor); len(fc) == 0 {
			com.FillColor = "#f3190b"
		}

		if lab := strings.TrimSpace(com.Label); len(lab) == 0 {
			com.Label = "WAF"
		}

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		el := node.New(cl, com.ID,
			node.Label(com.Label, false),
			node.FontColor(com.FontColor),
			node.FillColor(com.FillColor),
			node.Shape("invhouse"),
		)
		el.Attr("width", "0.3")

		return true
	}
}
