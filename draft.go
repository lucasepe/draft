package draft

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/edge"
	"github.com/lucasepe/draft/pkg/graph"
)

const (
	kindClient            = "cli"
	kindGateway           = "gtw"
	kindService           = "ser"
	kindQueue             = "que"
	kindPubSub            = "msg"
	kindObjectStore       = "ost"
	kindRDB               = "rdb"
	kindNoSQL             = "doc"
	kindFunction          = "fun"
	kindLBA               = "lba"
	kindCDN               = "cdn"
	kindDNS               = "dns"
	kindFirewall          = "waf"
	kindContainersManager = "kub"
	kindBlockStore        = "bst"
	kindCache             = "mem"
	kindFileStore         = "fst"
)

// Connection is a link between two components.
type Connection struct {
	Origin  string `yaml:"origin"`
	Targets []struct {
		ID        string `yaml:"id"`
		Label     string `yaml:"label,omitempty"`
		Color     string `yaml:"color,omitempty"`
		Dashed    bool   `yaml:"dashed,omitempty"`
		Dir       string `yaml:"dir,omitempty"`
		Highlight bool   `yaml:"highlight,omitempty"`
	} `yaml:"targets"`
}

// Component is a basic architecture unit.
type Component struct {
	ID        string `yaml:"id,omitempty"`
	Kind      string `yaml:"kind"`
	Label     string `yaml:"label,omitempty"`
	Impl      string `yaml:"impl,omitempty"`
	Outline   string `yaml:"outline,omitempty"`
	FillColor string `yaml:"fillColor,omitempty"`
	FontColor string `yaml:"fontColor,omitempty"`
}

// Design represents a whole diagram.
type Design struct {
	Title           string       `yaml:"title,omitempty"`
	BackgroundColor string       `yaml:"backgroundColor,omitempty"`
	Components      []Component  `yaml:"components"`
	Connections     []Connection `yaml:"connections,omitempty"`
	Ranks           []Rank       `yaml:"ranks,omitempty"`
}

// Rank define the nodes laying on the same level.
type Rank struct {
	Name       string   `yaml:"name"`
	Components []string `yaml:"components"`
}

// Sketch generates the GraphViz definition for this architecture diagram.
func Sketch(cfg Config) (string, error) {
	prj, err := Load(cfg)
	if err != nil {
		return "", err
	}

	if cfg.verbose {
		fmt.Fprintf(os.Stderr, "elaborating draft architecture definition: %s\n", cfg.uri)
	}

	gfx := graph.New(graph.BackgroundColor(prj.BackgroundColor),
		//graph.Ortho(ark.ortho),
		graph.BottomTop(cfg.bottomTop),
		graph.Label(prj.Title))

	if err := sketchComponents(gfx, cfg, prj.Components); err != nil {
		return "", err
	}

	if err := sketchConnections(gfx, cfg, prj.Connections); err != nil {
		return "", err
	}

	sketchSameRanks(gfx, cfg, prj.Ranks)

	return gfx.String(), nil
}

// fig is a function that render a Component according the Config.
type fig func(ctx Config, com Component) func(dia *dot.Graph) bool

var (
	// figRegistry keeps tracks of all 'figure' functions.
	figRegistry = make(map[string]fig)
)

// register a 'figure' function to the global registry.
func register(kind string, fig fig) {
	if _, ok := figRegistry[kind]; !ok {
		figRegistry[kind] = fig
	}
}

// idAutoGen auto generate a component id.
func idAutoGen() func(comp *Component) {
	counters := make(map[string]int16)

	return func(comp *Component) {
		if strings.TrimSpace(comp.ID) == "" {
			key := comp.Kind
			counters[key]++
			comp.ID = fmt.Sprintf("%s%d", key, counters[key])
		}
	}
}

// sketchComponents draws all components.
func sketchComponents(gfx *dot.Graph, cfg Config, items []Component) error {
	genID := idAutoGen()
	getImpl := guessImplByProvider(cfg.provider)

	for _, it := range items {
		genID(&it)
		getImpl(&it)

		if cfg.verbose {
			bin, _ := json.Marshal(it)
			fmt.Fprintf(os.Stderr, "  • component: %s\n", string(bin))
		}

		if ok := figIcon(cfg, it)(gfx); !ok {
			if fig, found := figRegistry[it.Kind]; found {
				fig(cfg, it)(gfx)
			} else {
				return fmt.Errorf("sketcher not found for component of kind <%s>", it.Kind)
			}
		}
	}

	return nil
}

// sketchConnections draws connections
func sketchConnections(gfx *dot.Graph, ctx Config, items []Connection) error {
	for _, el := range items {

		for _, x := range el.Targets {
			if ctx.verbose {
				fmt.Fprintf(os.Stderr, "  • connection: <%s> to <%s>\n", el.Origin, x.ID)
			}

			err := edge.New(gfx, el.Origin, x.ID,
				edge.Label(x.Label),
				edge.Dir(x.Dir),
				edge.Color(x.Color),
				edge.Dashed(x.Dashed),
				edge.Highlight(x.Highlight))

			if err != nil {
				return err
			}
		}
	}

	return nil
}

// sketchSameRanks groups component belonging to the same rank.
func sketchSameRanks(gfx *dot.Graph, ctx Config, items []Rank) {
	for _, grp := range items {
		if name := strings.TrimSpace(grp.Name); len(name) > 0 {
			for _, el := range grp.Components {
				if n, ok := gfx.FindNodeById(el); ok {
					gfx.AddToSameRank(name, n)
				}
			}
		}
	}
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
