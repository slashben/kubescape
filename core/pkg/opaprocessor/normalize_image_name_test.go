package opaprocessor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_normalize_name(t *testing.T) {

	tests := []struct {
		name string
		img  string
		want string
	}{
		{
			name: "Normalize image name",
			img:  "nginx",
			want: "docker.io/library/nginx",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, has_signature(tt.img), tt.name)
		})
	}
}
