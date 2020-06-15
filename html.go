package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type html struct {
	seq int16
}

func (rcv *html) nextID() string {
	rcv.seq++
	return fmt.Sprintf("htm%d", rcv.seq)
}

func (rcv *html) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	fontColor := "#000000ff"
	if fc := strings.TrimSpace(comp.FontColor); fc != "" {
		fontColor = fc
	}

	cl := cluster.New(graph, id, cluster.BottomTop(comp.BottomTop()), cluster.Label(comp.Impl))

	node.New(cl, id,
		node.Label(comp.Label, true),
		node.FontColor(fontColor),
		node.FillColor("transparent"),
		node.FontSize(7),
		node.Shape("plain"),
	)
}
