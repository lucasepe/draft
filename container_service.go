package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type containerService struct {
	seq int16
}

func (rcv *containerService) nextID() string {
	rcv.seq++
	return fmt.Sprintf("cos%d", rcv.seq)
}

func (rcv *containerService) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	cl := cluster.New(graph, id, cluster.BottomTop(comp.BottomTop()), cluster.Label(comp.Impl))

	el := node.New(cl, id,
		node.Label("Container\nService", false),
		node.FontColor(comp.FontColor, "#fafafaff"),
		node.FillColor(comp.FillColor, "#64a365"),
		node.Shape("component"),
	)
	el.Attr("height", "0.3")
}
