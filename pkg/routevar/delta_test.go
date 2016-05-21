package routevar

import (
	"reflect"
	"strings"
	"testing"

	"sourcegraph.com/sourcegraph/sourcegraph/go-sourcegraph/sourcegraph"

	"github.com/kr/pretty"
)

func TestDeltas(t *testing.T) {
	tests := []struct {
		delta         Delta
		wantRouteVars map[string]string
	}{
		{
			delta: Delta{
				Base: RepoRev{RepoSpec: sourcegraph.RepoSpec{URI: "samerepo"}, Rev: "base-rev"},
				Head: RepoRev{RepoSpec: sourcegraph.RepoSpec{URI: "samerepo"}, Rev: "head-rev"},
			},
			wantRouteVars: map[string]string{
				"Repo":         "samerepo",
				"Rev":          "@base-rev",
				"DeltaHeadRev": "@head-rev",
			},
		},
	}
	for _, test := range tests {
		vars := DeltaRouteVars(test.delta)
		if !reflect.DeepEqual(vars, test.wantRouteVars) {
			t.Errorf("got route vars != want\n\n%s", strings.Join(pretty.Diff(vars, test.wantRouteVars), "\n"))
		}

		delta := ToDelta(vars)
		if !reflect.DeepEqual(delta, test.delta) {
			t.Errorf("got delta != original delta\n\n%s", strings.Join(pretty.Diff(delta, test.delta), "\n"))
		}
	}
}
