package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

func init() {
	register(kindRDB, figRDB)
	register(kindNoSQL, figNoSQL)
	register(kindCache, figCache)
}

// figRDB render a Component of `kindRDB`.
func figRDB(ctx Config, com Component) func(gfx *dot.Graph) bool {
	return func(gfx *dot.Graph) bool {
		if !strings.EqualFold(com.Kind, kindRDB) {
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
			node.Shape("tab"),
		)
		el.Attr("height", "0.5")

		return true
	}
}

// figNoSQL render a Component of `kindNoSQL`.
func figNoSQL(ctx Config, com Component) func(gfx *dot.Graph) bool {
	return func(gfx *dot.Graph) bool {
		if !strings.EqualFold(com.Kind, kindNoSQL) {
			return false
		}

		if fc := strings.TrimSpace(com.FontColor); len(fc) == 0 {
			com.FontColor = "#000000ff"
		}

		if fc := strings.TrimSpace(com.FillColor); len(fc) == 0 {
			com.FillColor = "#d1c8d4ff"
		}

		if lab := strings.TrimSpace(com.Label); len(lab) == 0 {
			com.Label = "NoSQL"
		}

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		el := node.New(cl, com.ID,
			node.Label(com.Label, false),
			node.FontColor(com.FontColor),
			node.FillColor(com.FillColor),
			node.Shape("note"),
		)
		el.Attr("height", "0.5")

		return true
	}
}

// figCache render a Component of `kindCache`.
func figCache(ctx Config, com Component) func(gfx *dot.Graph) bool {
	return func(gfx *dot.Graph) bool {
		if !strings.EqualFold(com.Kind, kindCache) {
			return false
		}

		if fc := strings.TrimSpace(com.FontColor); len(fc) == 0 {
			com.FontColor = "#000000ff"
		}

		if fc := strings.TrimSpace(com.FillColor); len(fc) == 0 {
			com.FillColor = "#f04d30ff"
		}

		caption := strings.TrimSpace(com.Label)
		if len(caption) == 0 {
			caption = "Cache"
		}

		label := strings.Replace(`<table border="0" cellspacing="1" cellpadding="4">
	<tr><td border="0" colspan="2">&nbsp;</td></tr>
	<tr>
	  <td bgcolor="{{BGCOLOR}}"><font color="#f5f5f5" point-size="7">key</font></td>
	  <td bgcolor="{{BGCOLOR}}"><font color="#f5f5f5" point-size="7">val</font></td>
	</tr>
	<tr>
	  <td bgcolor="{{BGCOLOR}}" style="dashed"></td>
	  <td bgcolor="{{BGCOLOR}}" style="dashed"></td>
	</tr>
	<tr>
	  <td bgcolor="{{BGCOLOR}}" style="dashed"></td>
	  <td bgcolor="{{BGCOLOR}}" style="dashed"></td>
	</tr>
	<tr>
	  <td colspan="2">%s</td>
    </tr>
	</table>`, "{{BGCOLOR}}", com.FillColor, -1)

		label = fmt.Sprintf(label, caption)

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		el := node.New(cl, com.ID,
			node.Label(label, true),
			node.FontColor(com.FontColor),
			node.FillColor("transparent"),
			// ^^^ hack to set a transparent background
			// color since we will use the HTML table.
			node.Shape("plain"),
		)
		el.Attr("color", "#f5f5f5ff")

		return true
	}
}
