package draft

import (
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

func init() {
	register(kindCDN, figCDN)
	register(kindDNS, figDNS)
	register(kindLBA, figLBA)
	register(kindGateway, figGateway)
}

// figCDN render a Component of `kindCDN`.
func figCDN(ctx Config, com Component) func(gfx *dot.Graph) bool {
	return func(gfx *dot.Graph) bool {
		if !strings.EqualFold(com.Kind, kindCDN) {
			return false
		}

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		el := node.New(cl, com.ID,
			node.Label("<b>CDN</b>", true),
			node.FontColor("#000000ff"),
			node.FillColor("#47df9aff"),
			node.Shape("Mcircle"),
		)
		el.Attr("height", "0.5")
		el.Attr("color", "#f5f5f5ff")

		return true
	}
}

// figDNS render a Component of `kindCDN`.
func figDNS(ctx Config, com Component) func(gfx *dot.Graph) bool {
	return func(gfx *dot.Graph) bool {
		if !strings.EqualFold(com.Kind, kindDNS) {
			return false
		}

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		el := node.New(cl, com.ID,
			node.Label("<b>DNS</b>", true),
			node.FontColor("#f5f5f5ff"),
			node.FillColor("#854eadff"),
			node.Shape("Msquare"),
		)
		el.Attr("height", "0.3")
		el.Attr("color", "#f5f5f5ff")

		return true
	}
}

// figLBA render a Component of `kindLBA`.
func figLBA(ctx Config, com Component) func(gfx *dot.Graph) bool {
	return func(gfx *dot.Graph) bool {
		if !strings.EqualFold(com.Kind, kindLBA) {
			return false
		}

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		el := node.New(cl, com.ID,
			node.Label("<b>LB</b>", true),
			node.FontColor("#f5f5f5ff"),
			node.FillColor("#1a5276ff"),
			node.Shape("Mdiamond"),
		)
		el.Attr("width", "0.3")
		el.Attr("height", "0.3")
		el.Attr("color", "#f5f5f5ff")

		return true
	}
}

// figGateway render a Component of `kindGateway`.
func figGateway(ctx Config, com Component) func(gfx *dot.Graph) bool {
	return func(gfx *dot.Graph) bool {
		if !strings.EqualFold(com.Kind, kindGateway) {
			return false
		}

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		el := node.New(cl, com.ID,
			node.Label("<b>GTW</b>", true),
			node.FontColor("#f5f5f5ff"),
			node.FillColor("#ff7f00ff"),
			node.Shape("doublecircle"),
		)
		el.Attr("width", "0.1")
		el.Attr("color", "#ff7f00ff")

		return true
	}
}
