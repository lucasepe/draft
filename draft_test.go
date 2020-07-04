package draft

import (
	"fmt"
	"strings"
	"testing"

	"github.com/emicklei/dot"
)

func TestIDAutoGen(t *testing.T) {
	tests := []struct {
		kind string
		want string
	}{
		{kindCDN, "cdn1"},
		{kindCDN, "cdn2"},
		{kindCDN, "cdn3"},
		{kindService, "ser1"},
		{kindService, "ser2"},
		{kindCDN, "cdn4"},
		{kindService, "ser3"},
	}

	gen := idAutoGen()

	for _, tt := range tests {
		com := Component{Kind: tt.kind}

		t.Run(tt.kind, func(t *testing.T) {

			gen(&com)
			if got := com.ID; got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestSketchComponents(t *testing.T) {
	gfx := dot.NewGraph(dot.Directed)

	items := []Component{
		{Kind: kindGateway},
		{Kind: kindFunction},
	}

	if err := sketchComponents(gfx, Config{}, items); err != nil {
		t.Error(err)
	}

	if n, ok := gfx.FindNodeById("gtw1"); ok {
		got := fmt.Sprintf("%v", n.Value("label"))
		want := "default/gtw.png"

		if !strings.Contains(got, want) {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	}

	if n, ok := gfx.FindNodeById("fun1"); ok {
		got := fmt.Sprintf("%v", n.Value("label"))
		want := "default/fun.png"
		if !strings.Contains(got, want) {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	}
}
