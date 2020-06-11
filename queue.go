package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type queue struct {
	seq int16
}

func (rcv *queue) nextID() string {
	rcv.seq++
	return fmt.Sprintf("qs%d", rcv.seq)
}

func (rcv *queue) sketch(graph *dot.Graph, comp Component, bottomTop bool) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	cl := cluster.New(graph, id, cluster.BottomTop(bottomTop), cluster.Label(comp.Impl))

	el := node.New(cl, id,
		node.Rounded(comp.Rounded),
		node.FontColor(comp.FontColor),
		node.FontSize(7),
		node.FillColor("", "transparent"),
		// ^^^ hack to set a transparent background
		// color since we will use the HTML table.
		node.Shape("plain"),
	)

	caption := strings.TrimSpace(comp.Label)
	if len(caption) == 0 {
		caption = "&nbsp;"
	}

	fillColor := comp.FillColor
	if strings.TrimSpace(comp.FillColor) == "" {
		fillColor = "#bdb76bff"
	}

	label := fmt.Sprintf(`<table border="0" cellspacing="0" cellpadding="8">
	<tr><td border="0" colspan="3">&nbsp;</td></tr>
	<tr><td border="1" style="dashed" bgcolor="%s">msg N</td>
		<td border="1" style="dashed" bgcolor="%s">...</td>
		<td border="1" style="dashed" bgcolor="%s">msg 1</td></tr>
	<tr><td border="0" colspan="3"><font point-size="7">%s</font></td></tr>
	</table>`, fillColor, fillColor, fillColor, caption)

	el.Attr("label", dot.HTML(label))
}
