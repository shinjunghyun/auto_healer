package baram_helper

type Rect struct {
	X      uint16
	Y      uint16
	Width  uint16
	Height uint16
}

var (
	BARAM_HP_BOX_RECT    = Rect{X: 859, Y: 647, Width: 140, Height: 16}
	BARAM_MP_BOX_RECT    = Rect{X: 859, Y: 667, Width: 140, Height: 17}
	BARAM_HP_MP_BOX_RECT = Rect{X: 859, Y: 647, Width: 140, Height: 37}
)
