package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type kub struct {
	seq int16
}

func (rcv *kub) nextID() string {
	rcv.seq++
	return fmt.Sprintf("kub%d", rcv.seq)
}

func (rcv *kub) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	cl := cluster.New(graph, id, cluster.BottomTop(comp.BottomTop()), cluster.Label(comp.Impl))
	guessImpl(&comp, cl)

	node.New(cl, id,
		node.Label("<b>K8s<br/>Engine</b>", true),
		node.FontColor("#fafafaff"),
		node.FillColor("#64a365"),
		node.Shape("square"),
	)
}
