package draft

import (
	"regexp"
	"strings"
	"testing"

	"github.com/emicklei/dot"
)

func TestWAFComponentNextID(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"waf1"},
		{"waf2"},
		{"waf3"},
		{"waf4"},
	}

	s := waf{}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := s.nextID(); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestWAFComponent(t *testing.T) {
	g := dot.NewGraph(dot.Directed)

	sketcher := waf{}
	sketcher.sketch(g, Component{})

	if got, want := flatten(g.String()), `shape="invhouse",style="filled"`; !verify(got, want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

// remove tabs and newlines and spaces
func flatten(s string) string {
	return strings.Replace((strings.Replace(s, "\n", "", -1)), "\t", "", -1)
}

func verify(src, kv string) bool {
	r := regexp.MustCompile(`n2\[(.*)\]`)
	if m := r.FindAllStringSubmatch(src, -1); len(m) > 0 {
		return strings.Contains(m[0][1], kv)
	}
	return false
}
