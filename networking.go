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

		if fc := strings.TrimSpace(com.FontColor); len(fc) == 0 {
			com.FontColor = "#f5f5f5ff"
		}

		if fc := strings.TrimSpace(com.FillColor); len(fc) == 0 {
			com.FillColor = "#47df9aff"
		}

		if lab := strings.TrimSpace(com.Label); len(lab) == 0 {
			com.Label = "CDN"
		}

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		el := node.New(cl, com.ID,
			node.Label(com.Label, false),
			node.FontColor(com.FontColor),
			node.FillColor(com.FillColor),
			node.Shape("Mcircle"),
		)
		el.Attr("height", "0.5")
		el.Attr("color", com.FontColor)

		return true
	}
}

// figDNS render a Component of `kindCDN`.
func figDNS(ctx Config, com Component) func(gfx *dot.Graph) bool {
	return func(gfx *dot.Graph) bool {
		if !strings.EqualFold(com.Kind, kindDNS) {
			return false
		}

		if fc := strings.TrimSpace(com.FontColor); len(fc) == 0 {
			com.FontColor = "#f5f5f5ff"
		}

		if fc := strings.TrimSpace(com.FillColor); len(fc) == 0 {
			com.FillColor = "#854eadff"
		}

		if lab := strings.TrimSpace(com.Label); len(lab) == 0 {
			com.Label = "DNS"
		}

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		el := node.New(cl, com.ID,
			node.Label(com.Label, false),
			node.FontColor(com.FontColor),
			node.FillColor(com.FillColor),
			node.Shape("Msquare"),
		)
		el.Attr("height", "0.3")
		el.Attr("color", com.FontColor)

		return true
	}
}

// figLBA render a Component of `kindLBA`.
func figLBA(ctx Config, com Component) func(gfx *dot.Graph) bool {
	return func(gfx *dot.Graph) bool {
		if !strings.EqualFold(com.Kind, kindLBA) {
			return false
		}

		if fc := strings.TrimSpace(com.FontColor); len(fc) == 0 {
			com.FontColor = "#f5f5f5ff"
		}

		if fc := strings.TrimSpace(com.FillColor); len(fc) == 0 {
			com.FillColor = "#1a5276ff"
		}

		if lab := strings.TrimSpace(com.Label); len(lab) == 0 {
			com.Label = "LB"
		}

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		el := node.New(cl, com.ID,
			node.Label(com.Label, false),
			node.FontColor(com.FontColor),
			node.FillColor(com.FillColor),
			node.Shape("Mdiamond"),
		)
		el.Attr("width", "0.3")
		el.Attr("height", "0.3")
		el.Attr("color", com.FontColor)

		return true
	}
}

// figGateway render a Component of `kindGateway`.
func figGateway(ctx Config, com Component) func(gfx *dot.Graph) bool {
	return func(gfx *dot.Graph) bool {
		if !strings.EqualFold(com.Kind, kindGateway) {
			return false
		}

		if fc := strings.TrimSpace(com.FontColor); len(fc) == 0 {
			com.FontColor = "#f5f5f5ff"
		}

		if fc := strings.TrimSpace(com.FillColor); len(fc) == 0 {
			com.FillColor = "#ff7f00ff"
		}

		if lab := strings.TrimSpace(com.Label); len(lab) == 0 {
			com.Label = "AGW"
		}

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		el := node.New(cl, com.ID,
			node.Label(com.Label, false),
			node.FontColor(com.FontColor),
			node.FillColor(com.FillColor),
			node.Shape("doublecircle"),
		)
		el.Attr("width", "0.1")
		el.Attr("color", com.FontColor)

		return true
	}
}
