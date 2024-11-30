// go
package rpn_test

import (
	"testing"
	"github.com/Nagib_227/rpn/pkg/rpn"
)

func TestNewWorld(t *testing.T) {
	type test struct {
		height  int
		width   int
		wantErr bool
	}

	tests := []test{
		{height: 0, width: 4, wantErr: true},
		{height: -1, width: 4, wantErr: true},
		{height: 5, width: 0, wantErr: true},
		{height: 5, width: 6, wantErr: false},
	}

	for _, tt := range tests {
		height := tt.height
		width := tt.width
		world, err := NewWorld(height, width)
		if err != nil {
			if tt.wantErr {
				continue
			}
			t.Errorf("Unexpected error: %s", err)
		}

		if world.Height != height {
			t.Errorf("expected height: %d, actual height: %d", height, world.Height)
		}
		if world.Width != width {
			t.Errorf("expected width: %d, actual width: %d", width, world.Width)
		}

		if len(world.Cells) != height {
			t.Errorf("expected height: %d, actual number of rows: %d", height, len(world.Cells))
		}

		for i, row := range world.Cells {
			if len(row) != width {
				t.Errorf("expected width: %d, actual row's %d len: %d", width, i, world.Width)
			}
		}
	}
}