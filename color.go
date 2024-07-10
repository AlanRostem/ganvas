package ganvas

type Color [4]uint8

func Red() Color {
	return Color{255, 0, 0, 0}
}

func Green() Color {
	return Color{0, 255, 0, 0}
}

func Blue() Color {
	return Color{0, 0, 255, 0}
}
