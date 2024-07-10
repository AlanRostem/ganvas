package ganvas

import "github.com/go-gl/gl/v4.6-core/gl"

func Clear(color Color) {
	gl.ClearColor(
		float32(color[0])/255,
		float32(color[1])/255,
		float32(color[2])/255,
		float32(color[3])/255,
	)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}
