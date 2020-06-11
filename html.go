package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type html struct {
	seq int16
}

func (rcv *html) nextID() string {
	rcv.seq++
	return fmt.Sprintf("htm%d", rcv.seq)
}

func (rcv *html) sketch(graph *dot.Graph, comp Component, bottomTop bool) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	cl := cluster.New(graph, id, cluster.BottomTop(bottomTop), cluster.Label(comp.Impl))

	node.New(cl, id,
		node.Label(comp.Label, true),
		node.FontColor(comp.FontColor),
		node.FillColor("", "transparent"),
		node.FontSize(7),
		node.Shape("plain"),
	)
}
