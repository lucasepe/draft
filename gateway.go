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

func (rcv *gateway) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	label := comp.Label
	if strings.TrimSpace(comp.Label) == "" {
		label = "API Gateway"
	}

	cl := cluster.New(graph, id, cluster.Label(comp.Provider))

	el := node.New(cl, id,
		node.Label(label),
		node.Rounded(comp.Rounded),
		node.FontColor(comp.FontColor),
		node.FillColor(comp.FillColor, "#ff7f00ff"),
		node.FontSize(7),
		node.Shape("point"),
	)
	el.Attr("width", "0.3")
}
