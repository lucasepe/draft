package draft

import (
	"testing"

	"github.com/lucasepe/draft/pkg/cluster"
	"github.com/lucasepe/draft/pkg/graph"
)

func TestCloudImpl(t *testing.T) {
	tests := []struct {
		provider string
		key      string
		want     string
	}{
		{"aws", "bst", "Elastic Block\nStore (EBS)"},
		{"aws", "lba", "Elastic\nLoad Balancer"},
		{"aws", "ost", "Simple Storage\nService (S3)"},

		{"gcp", "kub", "Google Kubernetes\nEngine"},
		{"gcp", "mem", "Cloud Memorystore"},
		{"gcp", "ost", "Cloud Storage"},

		{"azure", "dns", "Azure DNS"},
		{"azure", "mem", "Redis Caches"},
		{"azure", "waf", "Azure Firewall"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := getCloudImpl(tt.provider, tt.key); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestGuessImpl(t *testing.T) {
	tests := []struct {
		comp Component
		want string
	}{
		{
			Component{provider: "aws", Kind: "bst"},
			"Elastic Block\nStore (EBS)",
		},
		{
			Component{provider: "aws", Kind: "lba"},
			"Elastic\nLoad Balancer",
		},
		{
			Component{provider: "aws", Kind: "ost"},
			"Simple Storage\nService (S3)",
		},

		{
			Component{provider: "gcp", Kind: "kub"},
			"Google Kubernetes\nEngine",
		},
		{
			Component{provider: "gcp", Kind: "mem"},
			"Cloud Memorystore",
		},
		{
			Component{provider: "gcp", Kind: "ost"},
			"Cloud Storage",
		},

		{
			Component{provider: "azure", Kind: "dns"},
			"Azure DNS",
		},
		{
			Component{provider: "azure", Kind: "mem"},
			"Redis Caches",
		},
		{
			Component{provider: "azure", Kind: "waf"},
			"Azure Firewall",
		},
	}

	gr := graph.New()
	cl := cluster.New(gr, "DUMMY")

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			guessImpl(&tt.comp, cl)
			if got := cl.AttributesMap.Value("label"); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}
