package draft

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/node"
)

// figIcon render a Component using the cloud provider icon.
func figIcon(ctx Config, com Component) func(gfx *dot.Graph) bool {

	return func(gfx *dot.Graph) bool {
		if len(ctx.provider) == 0 {
			return false
		}

		img := filepath.Join(ctx.iconsPath, ctx.provider, fmt.Sprintf("%s.png", com.Kind))
		if !fileExists(img) {
			if ctx.verbose {
				fmt.Fprintf(os.Stderr, "  ! file '%s' not found\n", img)
			}
			return false
		}

		if fc := strings.TrimSpace(com.FontColor); len(fc) == 0 {
			com.FontColor = "#000000ff"
		}

		caption := "&nbsp;"
		if val := strings.TrimSpace(com.Label); len(val) > 0 {
			caption = strings.ReplaceAll(val, "\n", "<br/>")
		}

		var sb strings.Builder
		sb.WriteString(`<table border="0" cellborder="0">`)
		sb.WriteString("<tr>")
		fmt.Fprintf(&sb, `<td fixedsize="true" width="50" height="50"><img src="%s" /></td>`, img)
		sb.WriteString("</tr>")
		if len(caption) > 0 {
			fmt.Fprintf(&sb, `<tr><td><font point-size="7">%s</font></td></tr>`, caption)
		}
		sb.WriteString("</table>")

		cl := cluster.New(gfx, com.ID, cluster.BottomTop(ctx.bottomTop), cluster.Label(com.Impl))

		node.New(cl, com.ID,
			node.Label(sb.String(), true),
			node.FillColor("transparent"),
			node.Shape("plain"),
		)

		return true
	}
}
