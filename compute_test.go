package draft

import (
	"regexp"
	"strings"
	"testing"

	"github.com/emicklei/dot"
)

func TestServiceFigure(t *testing.T) {
	tests := []struct {
		kind string
		find string
		want bool
	}{
		{kindService, `shape="doubleoctagon"`, true},
	}

	gfx := dot.NewGraph(dot.Directed)

	cfg := Config{}

	for _, tt := range tests {
		com := Component{Kind: tt.kind}

		t.Run(tt.find, func(t *testing.T) {
			figService(cfg, com)(gfx)
			x := flatten(gfx.String())
			if got := verify(x, tt.find); got != tt.want {
				t.Errorf("got [%v] want [%v] : %s", got, tt.want, x)
			}
		})
	}
}

func TestFunctionFigure(t *testing.T) {
	tests := []struct {
		kind string
		find string
		want bool
	}{
		{kindFunction, `shape="signature"`, true},
	}

	gfx := dot.NewGraph(dot.Directed)

	cfg := Config{}

	for _, tt := range tests {
		com := Component{Kind: tt.kind}

		t.Run(tt.find, func(t *testing.T) {
			figFunction(cfg, com)(gfx)
			x := flatten(gfx.String())
			if got := verify(x, tt.find); got != tt.want {
				t.Errorf("got [%v] want [%v] : %s", got, tt.want, x)
			}
		})
	}
}

func TestContainersManagerFigure(t *testing.T) {
	tests := []struct {
		kind string
		find string
		want bool
	}{
		{kindContainersManager, `shape="square"`, true},
	}

	gfx := dot.NewGraph(dot.Directed)

	cfg := Config{}

	for _, tt := range tests {
		com := Component{Kind: tt.kind}

		t.Run(tt.find, func(t *testing.T) {
			figContainersManager(cfg, com)(gfx)

			x := flatten(gfx.String())
			if got := verify(x, tt.find); got != tt.want {
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
	r := regexp.MustCompile(`n2\[(.*)\]`)
	if m := r.FindAllStringSubmatch(src, -1); len(m) > 0 {
		return strings.Contains(m[0][1], kv)
	}
	return false
}
