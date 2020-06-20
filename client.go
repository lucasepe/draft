package draft

import (
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

func init() {
	register(kindClient, figClient)
}

// figClient render a Component of `kindClient`.
func figClient(ctx Config, com Component) func(gfx *dot.Graph) bool {
	return func(gfx *dot.Graph) bool {
		if !strings.EqualFold(com.Kind, kindClient) {
			return false
		}

		if fc := strings.TrimSpace(com.FontColor); len(fc) == 0 {
			com.FontColor = "#000000ff"
		}

		if fc := strings.TrimSpace(com.FillColor); len(fc) == 0 {
			com.FillColor = "transparent"
		}

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		el := node.New(cl, com.ID,
			node.Label(com.Label, false),
			node.FontColor(com.FontColor),
			node.FillColor(com.FillColor),
			node.Shape("underline"),
		)
		el.Attr("fontsize", "8")
		el.Attr("height", "0.4")

		return true
	}
}
