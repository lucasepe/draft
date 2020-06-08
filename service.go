package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type service struct {
	seq int16
}

func (rcv *service) nextID() string {
	rcv.seq++
	return fmt.Sprintf("ms%d", rcv.seq)
}

func (rcv *service) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	cl := cluster.New(graph, id, cluster.Label(comp.Impl))

	el := node.New(cl, id,
		node.Label(comp.Label),
		node.Rounded(comp.Rounded),
		node.FontColor(comp.FontColor),
		node.FillColor(comp.FillColor, "#f5f5dcff"),
		node.Shape("octagon"),
	)
	el.Attr("height", "0.5")
}
