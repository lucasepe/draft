package draft

//go:generate statik -p statik -src=./assets

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/node"

	// init the embedded file system
	_ "github.com/lucasepe/draft/statik"
	"github.com/rakyll/statik/fs"
)

// render a Component using the cloud provider icon.
func render(ctx Config, com Component, gfx *dot.Graph) *dot.Node {

	img := iconPath(ctx, com)

	if fc := strings.TrimSpace(com.FontColor); len(fc) == 0 {
		com.FontColor = "#000000ff"
	}

	if imp := strings.TrimSpace(com.Impl); len(imp) == 0 {
		com.Impl = "&nbsp;"
	}

	var sb strings.Builder
	sb.WriteString(`<table border="0" cellborder="0">`)
	if ctx.showImpl {
		fmt.Fprintf(&sb, `<tr><td><font point-size="8">%s</font></td></tr>`, com.Impl)
	}

	sb.WriteString("<tr>")
	fmt.Fprintf(&sb, `<td fixedsize="true" width="50" height="50"><img src="%s" /></td>`, img)
	sb.WriteString("</tr>")

	label := "&nbsp;"
	if s := strings.TrimSpace(com.Label); len(s) > 0 {
		label = s
	}
	fmt.Fprintf(&sb, `<tr><td><font point-size="7">%s</font></td></tr>`, label)
	sb.WriteString("</table>")

	return node.New(gfx, com.ID,
		node.Label(sb.String(), true),
		node.FillColor("transparent"),
		node.Shape("plain"),
	)
}

// iconPath resolve the component icon path by it's provider attribute.
// If the icon does not exists returns a default image.
func iconPath(ctx Config, com Component) string {
	if prov := strings.TrimSpace(com.Provider); len(prov) == 0 {
		com.Provider = "default"
	} else {
		com.Provider = strings.ToLower(com.Provider)
	}

	fn := fmt.Sprintf("%s.png", com.Kind)

	dst := filepath.Join(ctx.iconsPath, com.Provider, fn)
	if fileExists(dst) {
		return dst
	}

	src := filepath.Join(ctx.iconsPath, "default", fn)

	if ctx.verbose {
		fmt.Fprintf(os.Stderr, "  ! file '%s' not found\n", dst)
	}

	if !fileExists(src) {
		if err := copyFileTo(fmt.Sprintf("/icons/%s", fn), src); err != nil {
			if os.IsNotExist(err) {
				src = filepath.Join(ctx.iconsPath, "default", "def.png")
				copyFileTo("/icons/def.png", src)
			}
		}
		fmt.Fprintf(os.Stderr, "    • using default image: '%s'\n", src)
	}

	return src
}

// copyFileTo extract the embedded icon to the user file system.
func copyFileTo(fromFn, toFn string) error {
	sfs, err := fs.New()
	if err != nil {
		return err
	}

	from, err := sfs.Open(fromFn)
	if err != nil {
		return err
	}
	defer from.Close()

	to, err := os.OpenFile(toFn, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	return err
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
