package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type dns struct {
	seq       int16
	bottomTop bool
}

func (rcv *dns) nextID() string {
	rcv.seq++
	return fmt.Sprintf("dn%d", rcv.seq)
}

func (rcv *dns) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	cl := cluster.New(graph, id, cluster.BottomTop(rcv.bottomTop), cluster.Label(comp.Impl))
	guessImpl(&comp, cl)

	el := node.New(cl, id,
		node.Label("<b>DNS</b>", true),
		node.FontColor("#f5f5f5ff"),
		node.FillColor("#854eadff"),
		node.Shape("Msquare"),
	)
	el.Attr("height", "0.3")
	el.Attr("color", "#f5f5f5ff")
}
