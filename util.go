package ganvas

func convertColor(c Color) [4]float32 {
	return [4]float32{
		float32(c[0]) / 255,
		float32(c[1]) / 255,
		float32(c[2]) / 255,
		float32(c[3]) / 255,
	}
}
