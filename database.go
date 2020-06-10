package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type database struct {
	seq int16
}

func (rcv *database) nextID() string {
	rcv.seq++
	return fmt.Sprintf("db%d", rcv.seq)
}

func (rcv *database) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	cl := cluster.New(graph, id, cluster.Label(comp.Impl))

	el := node.New(cl, id,
		node.Label(comp.Label, false),
		node.Rounded(comp.Rounded),
		node.FontColor(comp.FontColor),
		node.FillColor(comp.FillColor, "#f5f5dcff"),
		node.Shape("cylinder"),
	)
	el.Attr("height", "0.5")
	el.Attr("fontsize", "6")
}
