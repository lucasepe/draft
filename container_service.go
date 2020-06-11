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

func (rcv *containerService) sketch(graph *dot.Graph, comp Component, bottomTop bool) {
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
		node.Label("Container\nService", false),
		node.FontColor(fontColor),
		node.FillColor(comp.FillColor, "#64a365"),
		node.Shape("component"),
	)
	el.Attr("height", "0.3")
}
