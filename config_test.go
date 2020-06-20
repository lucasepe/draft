package draft

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProviderOpt(t *testing.T) {
	tests := []struct {
		val  string
		want string
	}{
		{"aws", "aws"},
		{"gcp", "gcp"},
		{"cloudflare", ""},
		{"azure", "azure"},
	}

	cfg := NewConfig()

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			Provider(tt.val)(&cfg)
			if got := cfg.provider; got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestBottomTopOpt(t *testing.T) {
	cfg := NewConfig()

	BottomTop(true)(&cfg)
	if got := cfg.bottomTop; got != true {
		t.Errorf("got [%v] want [true]", got)
	}
}

func TestVerboseOpt(t *testing.T) {
	cfg := NewConfig()

	Verbose(true)(&cfg)
	if got := cfg.verbose; got != true {
		t.Errorf("got [%v] want [true]", got)
	}
}

func TestLoadFromHTTPUri(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `title: Upload file to S3 using Lambda for pre-signed URL
backgroundColor: '#ffffff'
components:
  -
    kind: cli
    label: "Web App"
    impl: SPA
  -
    kind: gtw
  -
    kind: fun
    label: Get Pre-Signed URL
  -
    kind: ost
    label: "*.jpg\n*.png"`)
	}))
	defer ts.Close()

	cfg := NewConfig(URI(ts.URL))
	prj, err := Load(cfg)
	if err != nil {
		t.Error(err)
	}

	want := "Upload file to S3 using Lambda for pre-signed URL"
	if got := prj.Title; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}

	if got := len(prj.Components); got != 4 {
		t.Errorf("got [%v] want [3]", got)
	}

	want = "cli"
	if got := prj.Components[0].Kind; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}

	want = "Get Pre-Signed URL"
	if got := prj.Components[2].Label; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
