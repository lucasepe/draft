package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type cdn struct {
	seq int16
}

func (rcv *cdn) nextID() string {
	rcv.seq++
	return fmt.Sprintf("cn%d", rcv.seq)
}

func (rcv *cdn) sketch(graph *dot.Graph, comp Component, bottomTop bool) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	cl := cluster.New(graph, id, cluster.BottomTop(bottomTop), cluster.Label(comp.Impl))

	el := node.New(cl, id,
		node.Label("CDN", false),
		node.FontColor(comp.FontColor),
		node.FillColor(comp.FillColor, "#47df9aff"),
		node.Shape("Mcircle"),
	)
	el.Attr("height", "0.5")
}
