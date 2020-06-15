package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type gtw struct {
	seq int16
}

func (rcv *gtw) nextID() string {
	rcv.seq++
	return fmt.Sprintf("gtw%d", rcv.seq)
}

func (rcv *gtw) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	cl := cluster.New(graph, id, cluster.BottomTop(comp.BottomTop()), cluster.Label(comp.Impl))
	guessImpl(&comp, cl)

	el := node.New(cl, id,
		node.Label("<b>GTW</b>", true),
		node.Rounded(comp.Rounded),
		node.FontColor("#f5f5f5ff"),
		node.FillColor("#ff7f00ff"),
		node.Shape("doublecircle"),
	)
	el.Attr("width", "0.1")
	el.Attr("color", "#ff7f00ff")
}
