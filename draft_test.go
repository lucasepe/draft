package draft

import (
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
		{Kind: kindNoSQL},
	}

	if err := sketchComponents(gfx, Config{}, items); err != nil {
		t.Error(err)
	}

	want := `digraph  {
		subgraph cluster_s5 {
			fontname="Fira Mono Bold";fontsize="9.00";label="";labelloc="t";pencolor="transparent";
			n6[fillcolor="#d1c8d4ff",fontcolor="#000000ff",fontname="Fira Mono",fontsize="9.00",height="0.5",label="NoSQL",shape="note",style="filled"];
		}
		subgraph cluster_s3 {
			fontname="Fira Mono Bold";fontsize="9.00";label="";labelloc="t";pencolor="transparent";
			n4[fillcolor="#abd9e9ff",fontcolor="#000000ff",fontname="Fira Mono",fontsize="6",height="0.5",label="",shape="signature",style="filled"];
		}
		subgraph cluster_s1 {
			fontname="Fira Mono Bold";fontsize="9.00";label="";labelloc="t";pencolor="transparent";
			n2[color="#ff7f00ff",fillcolor="#ff7f00ff",fontcolor="#f5f5f5ff",fontname="Fira Mono",fontsize="9.00",label=<<b>GTW</b>>,shape="doublecircle",style="filled",width="0.1"];
		}
	}`
	if got := flatten(gfx.String()); got != flatten(want) {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
