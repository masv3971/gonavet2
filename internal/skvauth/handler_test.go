package skvauth

import (
	"testing"
)

func TestUse(t *testing.T) {
	tts := []struct {
		name string
		have string
		want string
	}{
		{
			name: "OK",
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			// Test function
			//	got, err := funcName
			//	assert.NoError(t, err)
			//	assert.Equal(t, tt.want, got)
		})
	}
}
