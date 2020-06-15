package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type mem struct {
	seq int16
}

func (rcv *mem) nextID() string {
	rcv.seq++
	return fmt.Sprintf("mem%d", rcv.seq)
}

func (rcv *mem) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	caption := strings.TrimSpace(comp.Label)
	if len(caption) == 0 {
		caption = "&nbsp;"
	}

	fillColor := comp.FillColor
	if strings.TrimSpace(comp.FillColor) == "" {
		fillColor = "#f04d30ff"
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
	</table>`, "{{BGCOLOR}}", fillColor, -1)

	label = fmt.Sprintf(label, caption)

	cl := cluster.New(graph, id, cluster.BottomTop(comp.BottomTop()), cluster.Label(comp.Impl))
	guessImpl(&comp, cl)

	el := node.New(cl, id,
		node.Label(label, true),
		node.FontColor("#000000ff"),
		node.FillColor("transparent"),
		// ^^^ hack to set a transparent background
		// color since we will use the HTML table.
		node.Shape("plain"),
	)
	el.Attr("color", "#f5f5f5ff")
}
