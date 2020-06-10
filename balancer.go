package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type balancer struct {
	seq int16
}

func (rcv *balancer) nextID() string {
	rcv.seq++
	return fmt.Sprintf("lb%d", rcv.seq)
}

func (rcv *balancer) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	cl := cluster.New(graph, id, cluster.Label(comp.Impl))

	el := node.New(cl, id,
		node.Label("LB", false),
		node.Rounded(comp.Rounded),
		node.FontColor(comp.FontColor),
		node.FillColor(comp.FillColor, "#1a5276ff"),
		node.Shape("Mdiamond"),
	)
	el.Attr("height", "0.3")
}
