package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type bst struct {
	seq int16
}

func (rcv *bst) nextID() string {
	rcv.seq++
	return fmt.Sprintf("bst%d", rcv.seq)
}

func (rcv *bst) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	fillColor := comp.FillColor
	if strings.TrimSpace(comp.FillColor) == "" {
		fillColor = "#606f5cff"
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
	</table>`, "{{BGCOLOR}}", fillColor, -1)

	cl := cluster.New(graph, id, cluster.BottomTop(comp.BottomTop()), cluster.Label(comp.Impl))
	guessImpl(&comp, cl)

	node.New(cl, id,
		node.Label(label, true),
		node.FontColor("#000000ff"),
		node.FontSize(7),
		node.FillColor("transparent"),
		// ^^^ hack to set a transparent background
		// color since we will use the HTML table.
		node.Shape("plain"),
	)
}
