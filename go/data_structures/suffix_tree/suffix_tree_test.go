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
				s: "abbc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st.NewSuffixTree(tt.args.s)
		})
	}
}

func Test_suffixTree_LongestRepeatedSubstring(t *testing.T) {
	type fields struct {
		inputString string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Basic substring search test",
			fields: fields{
				inputString: "banana",
			},
			want:    "ana",
			wantErr: false,
		},
		{
			name: "No duplicate substring search test",
			fields: fields{
				inputString: "abc",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := st.NewSuffixTree(tt.fields.inputString)
			got, err := s.LongestRepeatedSubstring()
			if (err != nil) != tt.wantErr {
				t.Errorf("suffixTree.LongestRepeatedSubstring() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("suffixTree.LongestRepeatedSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
