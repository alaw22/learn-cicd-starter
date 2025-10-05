package auth

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAuth(t *testing.T) {

	tests := map[string]struct {
		input http.Header
		want  string
	}{
		"no-auth": {
			input: http.Header{},
			want:  "",
		},
		"wrong-key": {
			input: http.Header{
				"Authorization": {"MyKey aljdlkjakljdf;ajdfl;jasd"},
			},
			want: "",
		},
		"no-key": {
			input: http.Header{
				"Authorization": {"ApiKey"},
			},
			want: "",
		},
		"normal": {
			input: http.Header{
				"Authorization": {"ApiKey asghu2uhgbf8238yrbwawjh!kjadu"},
			},
			want: "asghu2uhgbf8238yrbwawjh!kjadu",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := GetAPIKey(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}

		})
	}

}
