package ganvas

type DrawCallback func()

type Options struct {
	windowWidth  int
	windowHeight int
	windowTitle  string
	drawCallback DrawCallback
}

func NewOptions() *Options {
	return &Options{}
}

func (g *Options) WindowTitle(title string) *Options {
	g.windowTitle = title
	return g
}

func (g *Options) WindowSize(width, height int) *Options {
	g.windowWidth = width
	g.windowHeight = height
	return g
}

func (g *Options) DrawCallback(callback DrawCallback) *Options {
	g.drawCallback = callback
	return g
}
