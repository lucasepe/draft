package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type msg struct {
	seq int16
}

func (rcv *msg) nextID() string {
	rcv.seq++
	return fmt.Sprintf("msg%d", rcv.seq)
}

func (rcv *msg) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	cl := cluster.New(graph, id, cluster.BottomTop(comp.BottomTop()), cluster.Label(comp.Impl))
	guessImpl(&comp, cl)

	el := node.New(cl, id,
		node.Label("<b>Pub / Sub</b>", true),
		node.FontColor("#000000ff"),
		node.FillColor("#e0eeeeff"),
		node.Shape("cds"),
	)
	el.Attr("height", "0.6")
	el.Attr("width", "1.3")
}
