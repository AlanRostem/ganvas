package ganvas_test

import (
	"testing"

	"github.com/AlanRostem/ganvas"
)

func TestGanvas(t *testing.T) {
	opt := ganvas.NewOptions()

	opt.WindowTitle("Test Ganvas")
	opt.WindowSize(640, 480)
	opt.DrawCallback(func() {
		ganvas.Clear(ganvas.Green())
	})

	ganvas.Start(opt)
}
