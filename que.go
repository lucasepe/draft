package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type que struct {
	seq int16
}

func (rcv *que) nextID() string {
	rcv.seq++
	return fmt.Sprintf("que%d", rcv.seq)
}

func (rcv *que) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	fillColor := comp.FillColor
	if strings.TrimSpace(comp.FillColor) == "" {
		fillColor = "#bdb76bff"
	}

	caption := strings.TrimSpace(comp.Label)
	if len(caption) == 0 {
		caption = "&nbsp;"
	}

	label := strings.Replace(`<table border="0" cellspacing="0" cellpadding="8">
	<tr><td border="0" colspan="3">&nbsp;</td></tr>
	<tr><td border="1" bgcolor="{{BGCOLOR}}"><font point-size="6">msg N</font></td>
		<td border="1" bgcolor="{{BGCOLOR}}"><font point-size="6">...</font></td>
		<td border="1" bgcolor="{{BGCOLOR}}"><font point-size="6">msg 1</font></td></tr>
	<tr><td border="0" colspan="3"><font point-size="7">%s</font></td></tr>
	</table>`, "{{BGCOLOR}}", fillColor, -1)

	label = fmt.Sprintf(label, caption)

	cl := cluster.New(graph, id, cluster.BottomTop(comp.BottomTop()), cluster.Label(comp.Impl))
	guessImpl(&comp, cl)

	node.New(cl, id,
		node.Label(label, true),
		node.FontColor("#000000ff"),
		node.FontSize(8),
		node.FillColor("transparent"),
		// ^^^ hack to set a transparent background
		// color since we will use the HTML table.
		node.Shape("plain"),
	)
}
