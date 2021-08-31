package suffix_tree_test

import (
	"testing"

	st "github.com/maladroitthief/tool-chest/v2/go/data_structures/suffix_tree"
)

func TestNewSuffixTree(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Basic Suffix Tree Test",
			args: args{
				s: "glass",
			},
		},
		{
			name: "Complex word test",
			args: args{
				s: "mississippi",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st.NewSuffixTree(tt.args.s)
		})
	}
}
