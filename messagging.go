package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

func init() {
	register(kindPubSub, figPubSub)
	register(kindQueue, figQueue)
}

// figPubSub render a Component of `kindPubSub`.
func figPubSub(ctx Config, com Component) func(gfx *dot.Graph) bool {
	return func(gfx *dot.Graph) bool {
		if !strings.EqualFold(com.Kind, kindPubSub) {
			return false
		}

		if fc := strings.TrimSpace(com.FontColor); len(fc) == 0 {
			com.FontColor = "#000000ff"
		}

		if fc := strings.TrimSpace(com.FillColor); len(fc) == 0 {
			com.FillColor = "#e0eeeeff"
		}

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		el := node.New(cl, com.ID,
			node.Label("<b>Pub / Sub</b>", true),
			node.FontColor(com.FontColor),
			node.FillColor(com.FillColor),
			node.Shape("cds"),
		)
		el.Attr("height", "0.6")
		el.Attr("width", "1.3")

		return true
	}
}

// figQueue render a Component of `kindQueue`.
func figQueue(ctx Config, com Component) func(gfx *dot.Graph) bool {
	return func(gfx *dot.Graph) bool {
		if !strings.EqualFold(com.Kind, kindQueue) {
			return false
		}

		if fc := strings.TrimSpace(com.FontColor); len(fc) == 0 {
			com.FontColor = "#000000ff"
		}

		if fc := strings.TrimSpace(com.FillColor); len(fc) == 0 {
			com.FillColor = "#bdb76bff"
		}

		caption := strings.TrimSpace(com.Label)
		if len(caption) == 0 {
			caption = "&nbsp;"
		}

		label := strings.Replace(`<table border="0" cellspacing="0" cellpadding="8">
	<tr><td border="0" colspan="3">&nbsp;</td></tr>
	<tr><td border="1" bgcolor="{{BGCOLOR}}"><font point-size="6">msg N</font></td>
		<td border="1" bgcolor="{{BGCOLOR}}"><font point-size="6">...</font></td>
		<td border="1" bgcolor="{{BGCOLOR}}"><font point-size="6">msg 1</font></td></tr>
	<tr><td border="0" colspan="3"><font point-size="7">%s</font></td></tr>
	</table>`, "{{BGCOLOR}}", com.FillColor, -1)

		label = fmt.Sprintf(label, caption)

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		node.New(cl, com.ID,
			node.Label(label, true),
			node.FontColor(com.FontColor),
			node.FontSize(8),
			node.FillColor("transparent"),
			// ^^^ hack to set a transparent background
			// color since we will use the HTML table.
			node.Shape("plain"),
		)

		return true
	}
}
