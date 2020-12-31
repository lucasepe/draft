package draft

//go:generate statik -f -p statik -src=./assets

import (
	"testing"
)

func TestReadCsvFile(t *testing.T) {
	dat, err := readCsvFile("/default.csv")
	if err != nil {
		t.Fatal(err)
	}

	if len(dat) == 0 {
		t.Error("len(dat) should not be 0")
	}
}

func TestCloudImpl(t *testing.T) {
	tests := []struct {
		provider string
		key      string
		want     string
	}{
		{"aws", "bst", "EBS"},
		{"aws", "lba", "ELB"},
		{"aws", "ost", "S3"},

		{"google", "kub", `Kubernetes\nEngine`},
		{"google", "mem", "Memorystore"},
		{"google", "ost", "Storage"},

		{"azure", "dns", "DNS"},
		{"azure", "mem", "Redis Caches"},
		{"azure", "waf", "Firewall"},

		{"default", "dns", "DNS"},
		{"default", "mem", "Cache"},
		{"default", "waf", "Firewall"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := getCloudImpl(tt.provider, tt.key); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestSetImpl(t *testing.T) {
	tests := []struct {
		provider string
		kind     string
		want     string
	}{
		{"aws", "bst", "EBS"},
		{"aws", "lba", "ELB"},
		{"aws", "ost", "S3"},

		{"google", "kub", `Kubernetes\nEngine`},
		{"google", "mem", "Memorystore"},
		{"google", "ost", "Storage"},

		{"azure", "dns", "DNS"},
		{"azure", "mem", "Redis Caches"},
		{"azure", "waf", "Firewall"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			com := Component{Provider: tt.provider, Kind: tt.kind}
			setImpl(&com)

			if got := com.Impl; got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}
