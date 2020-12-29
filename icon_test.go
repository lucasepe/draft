package draft

//go:generate statik -f -p statik -src=./assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/emicklei/dot"
)

func TestCopyFileTo(t *testing.T) {
	dir, err := ioutil.TempDir("", "icons")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir) // clean up

	if err := os.Mkdir(filepath.Join(dir, "default"), os.ModePerm); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		src string
		dst string
	}{
		{"/icons/waf.png", fmt.Sprintf("%s/default/waf.png", dir)},
		{"/icons/fun.png", fmt.Sprintf("%s/default/fun.png", dir)},
		{"/icons/kub.png", fmt.Sprintf("%s/default/kub.png", dir)},
	}

	for _, tt := range tests {
		t.Run(tt.src, func(t *testing.T) {
			if err := copyFileTo(tt.src, tt.dst); err != nil {
				t.Fatal(err)
			}

			if !fileExists(tt.dst) {
				t.Errorf("file [%v] does not exists", tt.dst)
			}
		})
	}
}

func TestIconPath(t *testing.T) {
	dir, err := ioutil.TempDir("", "icons")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir) // clean up

	tests := []struct {
		provider string
		kind     string
		want     string
	}{
		{"aws", kindFirewall, fmt.Sprintf("%s/default/waf.png", dir)},
		{"aws", kindFunction, fmt.Sprintf("%s/default/fun.png", dir)},
		{"google", kindRDB, fmt.Sprintf("%s/default/rdb.png", dir)},
		{"google", kindCache, fmt.Sprintf("%s/default/mem.png", dir)},
		{"azure", kindDNS, fmt.Sprintf("%s/default/dns.png", dir)},
		{"azure", kindCDN, fmt.Sprintf("%s/default/cdn.png", dir)},
	}

	cfg := NewConfig(IconsPath(dir))

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			got := iconPath(cfg, Component{Kind: tt.kind})
			if got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestIconFigure(t *testing.T) {
	dir, err := ioutil.TempDir("", "icons")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir) // clean up

	tests := []struct {
		provider string
		kind     string
		find     string
		want     bool
	}{
		{"aws", kindFirewall, fmt.Sprintf(`src="%s/default/waf.png"`, dir), true},
		{"aws", kindFunction, fmt.Sprintf(`src="%s/default/fun.png"`, dir), true},
		{"google", kindRDB, fmt.Sprintf(`src="%s/default/rdb.png"`, dir), true},
		{"google", kindCache, fmt.Sprintf(`src="%s/default/mem.png"`, dir), true},
		{"azure", kindDNS, fmt.Sprintf(`src="%s/default/dns.png"`, dir), true},
		{"azure", kindCDN, fmt.Sprintf(`src="%s/default/cdn.png"`, dir), true},
	}

	cfg := NewConfig(IconsPath(dir))

	gfx := dot.NewGraph(dot.Directed)

	for _, tt := range tests {
		com := Component{Kind: tt.kind}

		t.Run(tt.find, func(t *testing.T) {
			n := render(cfg, com, gfx)
			x := fmt.Sprintf("%v", n.AttributesMap.Value("label"))
			if got := strings.Contains(x, tt.find); got != tt.want {
				t.Errorf("got [%v] want [%v] : %s", got, tt.want, x)
			}
		})
	}
}

// remove tabs and newlines and spaces
func flatten(s string) string {
	return strings.Replace((strings.Replace(s, "\n", "", -1)), "\t", "", -1)
}

func verify(src, kv string) bool {
	r := regexp.MustCompile(`n1\[(.*)\]`)
	if m := r.FindAllStringSubmatch(src, -1); len(m) > 0 {
		return strings.Contains(m[0][1], kv)
	}
	return false
}
