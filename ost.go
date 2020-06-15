package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type ost struct {
	seq int16
}

func (rcv *ost) nextID() string {
	rcv.seq++
	return fmt.Sprintf("ost%d", rcv.seq)
}

func (rcv *ost) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	fontColor := "#000000ff"
	if fc := strings.TrimSpace(comp.FontColor); fc != "" {
		fontColor = fc
	}

	fillColor := "#f5f5dcff"
	if fc := strings.TrimSpace(comp.FillColor); len(fc) > 0 {
		fillColor = fc
	}

	cl := cluster.New(graph, id, cluster.BottomTop(comp.BottomTop()), cluster.Label(comp.Impl))
	guessImpl(&comp, cl)

	el := node.New(cl, id,
		node.Label(comp.Label, false),
		node.Rounded(comp.Rounded),
		node.FontColor(fontColor),
		node.FillColor(fillColor),
		node.Shape("invtrapezium"),
	)
	el.Attr("width", "0.6")
}
