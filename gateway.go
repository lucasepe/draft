package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type gateway struct {
	seq int16
}

func (rcv *gateway) nextID() string {
	rcv.seq++
	return fmt.Sprintf("gt%d", rcv.seq)
}

func (rcv *gateway) sketch(graph *dot.Graph, comp Component, bottomTop bool) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	cl := cluster.New(graph, id, cluster.BottomTop(bottomTop), cluster.Label(comp.Impl))

	el := node.New(cl, id,
		node.Label("GW", false),
		node.Rounded(comp.Rounded),
		node.FontColor(comp.FontColor),
		node.FillColor(comp.FillColor, "#ff7f00ff"),
		node.Shape("doublecircle"),
	)
	el.Attr("width", "0.2")
}
