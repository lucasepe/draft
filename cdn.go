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
	return fmt.Sprintf("cdn%d", rcv.seq)
}

func (rcv *cdn) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	cl := cluster.New(graph, id, cluster.BottomTop(comp.BottomTop()), cluster.Label(comp.Impl))
	guessImpl(&comp, cl)

	el := node.New(cl, id,
		node.Label("<b>CDN</b>", true),
		node.FontColor("#000000ff"),
		node.FillColor("#47df9aff"),
		node.Shape("Mcircle"),
	)
	el.Attr("height", "0.5")
	el.Attr("color", "#f5f5f5ff")
}
