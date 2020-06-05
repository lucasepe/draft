package draft

import (
	"fmt"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type broker struct {
	seq int16
}

func (rcv *broker) nextID() string {
	rcv.seq++
	return fmt.Sprintf("br%d", rcv.seq)
}

func (rcv *broker) sketch(graph *dot.Graph, comp Component) {
	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	label := comp.Label
	if strings.TrimSpace(comp.Label) == "" {
		label = "Message Broker"
	}

	cl := cluster.New(graph, id, cluster.Label(comp.Provider))

	el := node.New(cl, id,
		node.Label(label),
		node.Rounded(comp.Rounded),
		node.FontColor(comp.FontColor),
		node.FillColor(comp.FillColor, "#e0eeeeff"),
		node.Shape("cds"),
	)
	el.Attr("height", "0.8")
}

/** Alternative

label=<table border="0" cellspacing="0">
                    <tr><td border="1" align="center"><font point-size="8"><b>&nbsp;topic 1&nbsp;</b></font></td></tr>
                    <tr><td border="1" align="center"><font point-size="8"><b>&nbsp;topic 2&nbsp;</b></font></td></tr>
                    <tr><td border="1" align="center"><font point-size="8"><b>&nbsp;...&nbsp;</b></font></td></tr>
                    <tr><td border="1" align="center"><font point-size="8"><b>&nbsp;topic N&nbsp;</b></font></td></tr>
				</table> >
shape="plain"

**/
