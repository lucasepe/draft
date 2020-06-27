package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

func init() {
	register(kindBlockStore, figBlockStore)
	register(kindFileStore, figFileStore)
	register(kindObjectStore, figObjectStore)
}

// figBlockStore render a Component of `kindBlockStore`.
func figBlockStore(ctx Config, com Component) func(gfx *dot.Graph) bool {
	return func(gfx *dot.Graph) bool {
		if !strings.EqualFold(com.Kind, kindBlockStore) {
			return false
		}

		if fc := strings.TrimSpace(com.FontColor); len(fc) == 0 {
			com.FontColor = "#000000ff"
		}

		if fc := strings.TrimSpace(com.FillColor); len(fc) == 0 {
			com.FillColor = "#606f5cff"
		}

		caption := strings.TrimSpace(com.Label)
		if len(caption) == 0 {
			caption = "Block Store"
		}

		label := strings.Replace(`<table border="0" cellspacing="4" cellpadding="4">
	<tr>
	  <td border="1" bgcolor="{{BGCOLOR}}" width="30"></td>
	  <td border="1" bgcolor="{{BGCOLOR}}"></td>
	</tr>
	<tr>
	  <td border="1" bgcolor="{{BGCOLOR}}" width="30"></td>
	  <td border="1" bgcolor="{{BGCOLOR}}"></td>
	</tr>
	<tr><td border="0" colspan="2"><font point-size="7">%s</font></td></tr>
	</table>`, "{{BGCOLOR}}", com.FillColor, -1)

		label = fmt.Sprintf(label, strings.ReplaceAll(caption, "\n", "<br/>"))

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		node.New(cl, com.ID,
			node.Label(label, true),
			node.FontColor(com.FontColor),
			node.FontSize(7),
			node.FillColor("transparent"),
			// ^^^ hack to set a transparent background
			// color since we will use the HTML table.
			node.Shape("plain"),
		)

		return true
	}
}

// figFileStore render a Component of `kindFileStore`.
func figFileStore(ctx Config, com Component) func(gfx *dot.Graph) bool {
	return func(gfx *dot.Graph) bool {
		if !strings.EqualFold(com.Kind, kindFileStore) {
			return false
		}

		if fc := strings.TrimSpace(com.FontColor); len(fc) == 0 {
			com.FontColor = "#000000ff"
		}

		if fc := strings.TrimSpace(com.FillColor); len(fc) == 0 {
			com.FillColor = "#dadd29ff"
		}

		if lab := strings.TrimSpace(com.Label); len(lab) == 0 {
			com.Label = "Files\nStore"
		}

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		el := node.New(cl, com.ID,
			node.Label(com.Label, false),
			node.FontColor(com.FontColor),
			node.FillColor(com.FillColor),
			node.Shape("folder"),
		)
		el.Attr("width", "0.7")

		return true
	}
}

// figObjectStore render a Component of `kindObjectStore`.
func figObjectStore(ctx Config, com Component) func(gfx *dot.Graph) bool {
	return func(gfx *dot.Graph) bool {
		if !strings.EqualFold(com.Kind, kindObjectStore) {
			return false
		}

		if fc := strings.TrimSpace(com.FontColor); len(fc) == 0 {
			com.FontColor = "#000000ff"
		}

		if fc := strings.TrimSpace(com.FillColor); len(fc) == 0 {
			com.FillColor = "#f5f5dcff"
		}

		if lab := strings.TrimSpace(com.Label); len(lab) == 0 {
			com.Label = "Object\nStore"
		}

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		el := node.New(cl, com.ID,
			node.Label(com.Label, false),
			node.FontSize(7),
			node.FontColor(com.FontColor),
			node.FillColor(com.FillColor),
			node.Shape("invtrapezium"),
		)
		el.Attr("width", "0.6")

		return true
	}
}
