package draft

import (
	"testing"
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

		{"google", "kub", "Google Kubernetes\nEngine"},
		{"google", "mem", "Cloud Memorystore"},
		{"google", "ost", "Cloud Storage"},

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
		prov string
		comp Component
		want string
	}{
		{
			"aws", Component{Kind: "bst"},
			"Elastic Block\nStore (EBS)",
		},
		{
			"aws", Component{Kind: "lba"},
			"Elastic\nLoad Balancer",
		},
		{
			"aws", Component{Kind: "ost"},
			"Simple Storage\nService (S3)",
		},

		{
			"google", Component{Kind: "kub"},
			"Google Kubernetes\nEngine",
		},
		{
			"google", Component{Kind: "mem"},
			"Cloud Memorystore",
		},
		{
			"google", Component{Kind: "ost"},
			"Cloud Storage",
		},

		{
			"azure", Component{Kind: "dns"},
			"Azure DNS",
		},
		{
			"azure", Component{Kind: "mem"},
			"Redis Caches",
		},
		{
			"azure", Component{Kind: "waf"},
			"Azure Firewall",
		},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			guessImplByProvider(tt.prov)(&tt.comp)
			if got := tt.comp.Impl; got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}
