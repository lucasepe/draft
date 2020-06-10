package draft

import (
	"fmt"
	"io"
	"strings"

	"github.com/emicklei/dot"
	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/edge"
	"github.com/lucasepe/draft/pkg/graph"
	"gopkg.in/yaml.v2"
)

const (
	kindHTML     = "html"
	kindClient   = "client"
	kindGateway  = "gateway"
	kindService  = "service"
	kindQueue    = "queue"
	kindBroker   = "broker"
	kindStorage  = "storage"
	kindDatabase = "database"
	kindFunction = "function"
	kindBalancer = "balancer"
	kindCDN      = "cdn"
	kindDNS      = "dns"
)

// Connection is a link between two components.
type Connection struct {
	Origin struct {
		ComponentID string `yaml:"componentId"`
	} `yaml:"origin"`
	Targets []struct {
		ComponentID string `yaml:"componentId"`
		Label       string `yaml:"label,omitempty"`
		Color       string `yaml:"color,omitempty"`
		Dashed      bool   `yaml:"dashed,omitempty"`
		Dir         string `yaml:"dir,omitempty"`
		Highlight   bool   `yaml:"highlight,omitempty"`
	} `yaml:"targets"`
}

// Component is a basic architecture unit.
type Component struct {
	ID        string `yaml:"id,omitempty"`
	Kind      string `yaml:"kind"`
	Label     string `yaml:"label,omitempty"`
	Impl      string `yaml:"impl,omitempty"`
	Provider  string `yaml:"provider,omitempty"`
	FillColor string `yaml:"fillColor,omitempty"`
	FontColor string `yaml:"fontColor,omitempty"`
	Rounded   bool   `yaml:"rounded,omitempty"`
}

// Draft represents a whole diagram.
type Draft struct {
	Title           string       `yaml:"title,omitempty"`
	BackgroundColor string       `yaml:"backgroundColor,omitempty"`
	Components      []Component  `yaml:"components"`
	Connections     []Connection `yaml:"connections,omitempty"`

	sketchers map[string]interface {
		sketch(*dot.Graph, Component)
	}
}

// NewDraft returns a new decoded Draft struct
func NewDraft(r io.Reader) (*Draft, error) {
	res := &Draft{
		sketchers: map[string]interface {
			sketch(*dot.Graph, Component)
		}{
			kindHTML:     &html{},
			kindClient:   &client{},
			kindGateway:  &gateway{},
			kindService:  &service{},
			kindBroker:   &broker{},
			kindQueue:    &queue{},
			kindFunction: &function{},
			kindStorage:  &storage{},
			kindDatabase: &database{},
			kindBalancer: &balancer{},
			kindCDN:      &cdn{},
			kindDNS:      &dns{},
		},
	}

	// Init new YAML decode
	d := yaml.NewDecoder(r)

	// Start YAML decoding from file
	if err := d.Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}

// Sketch generates the GraphViz definition for this architecture diagram.
func (ark *Draft) Sketch() (string, error) {
	g := graph.New(graph.BackgroundColor(ark.BackgroundColor), graph.Label(ark.Title))

	if err := sketchComponents(g, ark); err != nil {
		return "", err
	}

	if err := sketchConnections(g, ark); err != nil {
		return "", err
	}

	return g.String(), nil
}

func sketchComponents(graph *dot.Graph, draft *Draft) error {
	for _, el := range draft.Components {
		sketcher, ok := draft.sketchers[el.Kind]
		if !ok {
			return fmt.Errorf("render not found for component of kind '%s'", el.Kind)
		}

		parent := graph
		if strings.TrimSpace(el.Provider) != "" {
			parent = cluster.New(graph, el.Provider,
				cluster.PenColor("#d9cc31"),
				cluster.FontName("Fira Mono"),
				cluster.FontSize(10),
				cluster.FontColor("#63625b"))
		}

		sketcher.sketch(parent, el)
	}

	return nil
}

func sketchConnections(graph *dot.Graph, draft *Draft) error {
	for _, el := range draft.Connections {
		var from = el.Origin.ComponentID

		for _, x := range el.Targets {
			err := edge.New(graph, from, x.ComponentID,
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
