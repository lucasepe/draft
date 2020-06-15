package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type doc struct {
	seq int16
}

func (rcv *doc) nextID() string {
	rcv.seq++
	return fmt.Sprintf("doc%d", rcv.seq)
}

func (rcv *doc) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	cl := cluster.New(graph, id, cluster.BottomTop(comp.BottomTop()), cluster.Label(comp.Impl))
	guessImpl(&comp, cl)

	el := node.New(cl, id,
		node.Label(comp.Label, false),
		node.Rounded(comp.Rounded),
		node.FontColor("#000000ff"),
		node.FillColor("#d1c8d4ff"),
		node.Shape("note"),
	)
	el.Attr("height", "0.5")
}
