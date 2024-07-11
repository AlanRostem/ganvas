package ganvas

import "github.com/go-gl/gl/v4.6-core/gl"

func Clear(color Color) {
	c := convertColor(color)
	gl.ClearColor(c[0], c[1], c[2], c[3])
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func Triangle(v0, v1, v2 Vec2) {

}
