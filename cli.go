package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type cli struct {
	seq int16
}

func (rcv *cli) nextID() string {
	rcv.seq++
	return fmt.Sprintf("cli%d", rcv.seq)
}

func (rcv *cli) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	fontColor := "#000000ff"
	if fc := strings.TrimSpace(comp.FontColor); len(fc) > 0 {
		fontColor = fc
	}

	cl := cluster.New(graph, id, cluster.BottomTop(comp.BottomTop()), cluster.Label(comp.Impl))

	el := node.New(cl, id,
		node.Label(comp.Label, false),
		node.Rounded(comp.Rounded),
		node.FontColor(fontColor),
		node.FillColor("transparent"),
		node.Shape("underline"),
	)
	el.Attr("fontsize", "8")
	el.Attr("height", "0.4")
}
