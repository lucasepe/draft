package draft

import (
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

func init() {
	register(kindService, figService)
	register(kindFunction, figFunction)
	register(kindContainersManager, figContainersManager)
}

// figService render a Component of `kindService`.
func figService(ctx Config, com Component) func(gfx *dot.Graph) bool {
	return func(gfx *dot.Graph) bool {
		if !strings.EqualFold(com.Kind, kindService) {
			return false
		}

		if fc := strings.TrimSpace(com.FontColor); len(fc) == 0 {
			com.FontColor = "#000000ff"
		}

		if fc := strings.TrimSpace(com.FillColor); len(fc) == 0 {
			com.FillColor = "#f5f5dcff"
		}

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		el := node.New(cl, com.ID,
			node.Label(com.Label, false),
			node.FontColor(com.FontColor),
			node.FillColor(com.FillColor),
			node.Shape("doubleoctagon"),
		)
		el.Attr("width", "0.7")
		el.Attr("color", com.FillColor)

		return true
	}
}

// figFunction render a Component of `kindFunction`.
func figFunction(ctx Config, com Component) func(gfx *dot.Graph) bool {
	return func(gfx *dot.Graph) bool {
		if !strings.EqualFold(com.Kind, kindFunction) {
			return false
		}

		if fc := strings.TrimSpace(com.FontColor); len(fc) == 0 {
			com.FontColor = "#000000ff"
		}

		if fc := strings.TrimSpace(com.FillColor); len(fc) == 0 {
			com.FillColor = "#abd9e9ff"
		}

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		el := node.New(cl, com.ID,
			node.Label(com.Label, false),
			node.FontColor(com.FontColor),
			node.FillColor(com.FillColor),
			node.Shape("signature"),
		)
		el.Attr("fontsize", "6")
		el.Attr("height", "0.5")

		return true
	}
}

// figContainersManager render a Component of `kindContainersManager`.
func figContainersManager(ctx Config, com Component) func(gfx *dot.Graph) bool {
	return func(gfx *dot.Graph) bool {
		if !strings.EqualFold(com.Kind, kindContainersManager) {
			return false
		}

		if lab := strings.TrimSpace(com.Label); len(lab) == 0 {
			com.Label = "Containers\nEngine"
		}

		if fc := strings.TrimSpace(com.FontColor); len(fc) == 0 {
			com.FontColor = "#fafafaff"
		}

		if fc := strings.TrimSpace(com.FillColor); len(fc) == 0 {
			com.FillColor = "#64a365"
		}

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		node.New(cl, com.ID,
			node.Label(com.Label, false),
			node.FontSize(7),
			node.FontColor(com.FontColor),
			node.FillColor(com.FillColor),
			node.Shape("square"),
		)

		return true
	}
}
