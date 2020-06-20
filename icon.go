package draft

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

// figIcon render a Component using the cloud provider icon.
func figIcon(ctx Config, com Component) func(gfx *dot.Graph) bool {
	iconsPath := os.Getenv("DRAFT_ICONS_PATH")

	return func(gfx *dot.Graph) bool {
		if len(ctx.provider) == 0 {
			return false
		}

		img := filepath.Join(iconsPath, ctx.provider, fmt.Sprintf("%s.png", com.Kind))
		if !fileExists(img) {
			if ctx.verbose {
				fmt.Fprintf(os.Stderr, "  ! file '%s' not found\n", img)
			}
			return false
		}

		label := fmt.Sprintf(`<table border="0" cellborder="0">
		<tr>
		  <td fixedsize="true" width="50" height="50"><img src="%s" /></td>
		</tr>
		</table>`, img)

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		node.New(cl, com.ID,
			node.Label(label, true),
			node.FillColor("transparent"),
			node.Shape("plain"),
		)

		return true
	}
}
