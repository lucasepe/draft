package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type waf struct {
	seq int16
}

func (rcv *waf) nextID() string {
	rcv.seq++
	return fmt.Sprintf("waf%d", rcv.seq)
}

func (rcv *waf) sketch(graph *dot.Graph, comp Component, bottomTop bool) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	fontColor := "#fafafaff"
	if strings.TrimSpace(comp.FontColor) != "" {
		fontColor = comp.FontColor
	}

	cl := cluster.New(graph, id, cluster.BottomTop(bottomTop), cluster.Label(comp.Impl))

	el := node.New(cl, id,
		node.Label("FW", false),
		node.FontColor(fontColor),
		node.FillColor(comp.FillColor, "#f3190b"),
		node.Shape("invhouse"),
	)
	el.Attr("width", "0.3")
}
