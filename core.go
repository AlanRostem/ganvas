package ganvas

import (
	"runtime"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func initGlfw() {
	runtime.LockOSThread()

	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
}

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	return prog
}

func finalize() {
	glfw.Terminate()
}

func Start(opt *Options) {
	initGlfw()
	prg := initOpenGL()

	window, err := glfw.CreateWindow(opt.windowWidth, opt.windowHeight, opt.windowTitle, nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	_graphicsStates.init()

	for !window.ShouldClose() {
		// Do OpenGL stuff.
		gl.UseProgram(prg)

		window.SwapBuffers()
		glfw.PollEvents()

		opt.drawCallback()
	}

	finalize()
}
