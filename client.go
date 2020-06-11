package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type client struct {
	seq int16
}

func (rcv *client) nextID() string {
	rcv.seq++
	return fmt.Sprintf("cl%d", rcv.seq)
}

func (rcv *client) sketch(graph *dot.Graph, comp Component, bottomTop bool) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	cl := cluster.New(graph, id, cluster.BottomTop(bottomTop), cluster.Label(comp.Impl))

	el := node.New(cl, id,
		node.Label(comp.Label, false),
		node.Rounded(comp.Rounded),
		node.FontColor(comp.FontColor),
		node.FillColor(comp.FillColor, "#90ee90ff"),
		node.Shape("underline"),
	)
	el.Attr("fontsize", "8")
	el.Attr("height", "0.3")
}
