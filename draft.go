package draft

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/edge"
	"github.com/lucasepe/draft/pkg/graph"
)

// Connection is a link between two components.
type Connection struct {
	Origin  string `yaml:"origin"`
	Targets []struct {
		ID            string  `yaml:"id"`
		Label         string  `yaml:"label,omitempty"`
		LabelDistance float32 `yaml:"labeldistance,omitempty"`
		LabelAngle    float32 `yaml:"labelangle,omitempty"`
		MinLen        float32 `yaml:"minlen,omitempty"`
		Num           int     `yaml:"num,omitempty"`
		Color         string  `yaml:"color,omitempty"`
		Dashed        bool    `yaml:"dashed,omitempty"`
		Dir           string  `yaml:"dir,omitempty"`
		Highlight     bool    `yaml:"highlight,omitempty"`
	} `yaml:"targets"`
}

// Component is a basic architecture unit.
type Component struct {
	ID        string `yaml:"id,omitempty"`
	Kind      string `yaml:"kind"`
	Label     string `yaml:"label,omitempty"`
	Outline   string `yaml:"outline,omitempty"`
	Impl      string `yaml:"impl,omitempty"`
	Provider  string `yaml:"provider,omitempty"`
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
	fixKind := validateKind()
	fixProvider := validateProvider()

	for _, it := range items {
		genID(&it)

		fixKind(&it)
		fixProvider(&it)
		setImpl(&it)

		if cfg.verbose {
			bin, _ := json.Marshal(it)
			fmt.Fprintf(os.Stderr, "  • component: %s\n", string(bin))
		}

		parent := gfx
		if box := strings.TrimSpace(it.Outline); len(box) > 0 {
			parent = cluster.New(gfx, it.Outline,
				cluster.PenColor("#78959b"),
				cluster.FontSize(10),
				cluster.FontColor("#63625b"))
			parent.Attr("style", "dashed,rounded")
		}

		render(cfg, it, parent)
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
				edge.Label(x.Num, x.Label),
				edge.Dir(x.Dir),
				edge.Color(x.Color),
				edge.Dashed(x.Dashed),
				edge.LabelDistance(x.LabelDistance),
				edge.LabelAngle(x.LabelAngle),
				edge.MinLen(x.MinLen),
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
