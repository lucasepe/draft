package draft

import (
	"testing"
)

func TestValidateProvider(t *testing.T) {
	tests := []struct {
		provider string
		want     string
	}{
		{"aws", "aws"},
		{"google", "google"},
		{"alibaba", "default"},
		{"digitalocean", "default"},
		{"azure", "azure"},
	}

	validate := validateProvider()

	for _, tt := range tests {
		t.Run(tt.provider, func(t *testing.T) {
			com := Component{Provider: tt.provider}
			validate(&com)

			if com.Provider != tt.want {
				t.Errorf("got [%v] want [%v]", com.Provider, tt.want)
			}
		})
	}
}
