package slicetools

import (
	"testing"
)

func TestChunk(t *testing.T) {
	tCases := []struct {
		name           string
		inputSize      int
		inputSlice     []interface{}
		expectedChunks [][]interface{}
	}{
		{
			name:           "",
			inputSize:      3,
			inputSlice:     []any{nil, nil, nil, nil, nil},
			expectedChunks: [][]any{{nil, nil, nil}, {nil, nil}},
		},
		{
			name:           "",
			inputSize:      2,
			inputSlice:     []any{nil, nil, nil, nil, nil, nil, nil},
			expectedChunks: [][]any{{nil, nil}, {nil, nil}, {nil, nil}, {nil}},
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := Chunk(tc.inputSize, tc.inputSlice)
			if len(actual) != len(tc.expectedChunks) {
				tt.Errorf("unexpected actual chunks size (expected=%d, actual=%d)", len(tc.expectedChunks), len(actual))
			}
			for i := 0; i < len(actual); i++ {
				if len(actual[i]) != len(tc.expectedChunks[i]) {
					tt.Errorf("unexpected actual inner chunk size (expected=%d, actual=%d)", len(tc.expectedChunks[i]), len(actual[i]))
				}
			}
		})
	}
}

func TestChunkFor(t *testing.T) {
	tCases := []struct {
		name           string
		inputSize      int
		inputSlice     []interface{}
		expectedChunks [][]interface{}
	}{
		{
			name:           "",
			inputSize:      3,
			inputSlice:     []any{nil, nil, nil, nil, nil},
			expectedChunks: [][]any{{nil, nil}, {nil, nil}, {nil}},
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := ChunkFor(tc.inputSlice, tc.inputSize)
			if len(actual) != len(tc.expectedChunks) {
				tt.Errorf("unexpected actual chunks size (expected=%d, actual=%d)", len(tc.expectedChunks), len(actual))
			}
			for i := 0; i < len(actual); i++ {
				if len(actual[i]) != len(tc.expectedChunks[i]) {
					tt.Errorf("unexpected actual inner chunk size (expected=%d, actual=%d)", len(tc.expectedChunks[i]), len(actual[i]))
				}
			}
		})
	}
}
