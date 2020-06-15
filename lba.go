package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type lba struct {
	seq int16
}

func (rcv *lba) nextID() string {
	rcv.seq++
	return fmt.Sprintf("lba%d", rcv.seq)
}

func (rcv *lba) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	cl := cluster.New(graph, id, cluster.BottomTop(comp.BottomTop()), cluster.Label(comp.Impl))
	guessImpl(&comp, cl)

	el := node.New(cl, id,
		node.Label("<b>LB</b>", true),
		node.Rounded(comp.Rounded),
		node.FontColor("#f5f5f5ff"),
		node.FillColor("#1a5276ff"),
		node.Shape("Mdiamond"),
	)
	el.Attr("width", "0.3")
	el.Attr("height", "0.3")
	el.Attr("color", "#f5f5f5ff")
}
