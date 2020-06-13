package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type storage struct {
	seq int16
}

func (rcv *storage) nextID() string {
	rcv.seq++
	return fmt.Sprintf("st%d", rcv.seq)
}

func (rcv *storage) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	cl := cluster.New(graph, id, cluster.BottomTop(comp.BottomTop()), cluster.Label(comp.Impl))

	el := node.New(cl, id,
		node.Label(comp.Label, false),
		node.Rounded(comp.Rounded),
		node.FontColor(comp.FontColor, "#000000ff"),
		node.FillColor(comp.FillColor, "#f0e77fff"),
		node.FontSize(8),
		node.Shape("folder"),
	)
	el.Attr("height", "0.4")
}
