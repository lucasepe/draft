package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type fun struct {
	seq int16
}

func (rcv *fun) nextID() string {
	rcv.seq++
	return fmt.Sprintf("fun%d", rcv.seq)
}

func (rcv *fun) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	label := id
	if strings.TrimSpace(comp.Label) != "" {
		label = comp.Label
	}

	fillColor := "#abd9e9ff"
	if fc := strings.TrimSpace(comp.FillColor); len(fc) > 0 {
		fillColor = fc
	}

	cl := cluster.New(graph, id, cluster.BottomTop(comp.BottomTop()), cluster.Label(comp.Impl))
	guessImpl(&comp, cl)

	el := node.New(cl, id,
		node.Label(label, false),
		node.Rounded(comp.Rounded),
		node.FontColor("#000000ff"),
		node.FillColor(fillColor),
		node.Shape("signature"),
	)
	el.Attr("fontsize", "6")
	el.Attr("height", "0.5")
}
