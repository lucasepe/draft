package draft

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

type icon struct {
	seq int16
}

func (rcv *icon) nextID() string {
	rcv.seq++
	return fmt.Sprintf("ico%d", rcv.seq)
}

func (rcv *icon) sketch(graph *dot.Graph, comp Component) {
	if strings.TrimSpace(comp.Label) == "" {
		return
	}

	id := comp.ID
	if strings.TrimSpace(comp.ID) == "" {
		id = rcv.nextID()
	}

	iconsPath := os.Getenv("DRAFT_ICONS_PATH")
	img := path.Join(iconsPath, comp.Label)

	label := fmt.Sprintf(`<table border="0" cellborder="0">
	<tr>
	  <td><img src="%s"/></td>
	</tr>
	</table>`, img)

	cl := cluster.New(graph, id, cluster.BottomTop(comp.BottomTop()), cluster.Label(comp.Impl))

	node.New(cl, id,
		node.Label(label, true),
		node.FillColor("transparent"),
		node.Shape("plain"),
	)
}
