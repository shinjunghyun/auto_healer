package baram_helper

type Point struct {
	X uint16
	Y uint16
}

type PixelPointColor struct {
	Point
	color int32
}

var (
	BaramNumberPixlesMap map[int8][]PixelPointColor = map[int8][]PixelPointColor{
		0: {
			{Point: Point{X: 10, Y: 10}, color: 0xFFFFFF},
			{Point: Point{X: 11, Y: 12}, color: 0xFFFFFF},
		},

		1: {
			{Point: Point{X: 10, Y: 10}, color: 0xFFFFFF},
			{Point: Point{X: 11, Y: 12}, color: 0xFFFFFF},
		},

		2: {
			{Point: Point{X: 10, Y: 10}, color: 0xFFFFFF},
			{Point: Point{X: 11, Y: 12}, color: 0xFFFFFF},
		},

		3: {
			{Point: Point{X: 10, Y: 10}, color: 0xFFFFFF},
			{Point: Point{X: 11, Y: 12}, color: 0xFFFFFF},
		},

		4: {
			{Point: Point{X: 10, Y: 10}, color: 0xFFFFFF},
			{Point: Point{X: 11, Y: 12}, color: 0xFFFFFF},
		},

		5: {
			{Point: Point{X: 10, Y: 10}, color: 0xFFFFFF},
			{Point: Point{X: 11, Y: 12}, color: 0xFFFFFF},
		},

		6: {
			{Point: Point{X: 10, Y: 10}, color: 0xFFFFFF},
			{Point: Point{X: 11, Y: 12}, color: 0xFFFFFF},
		},

		7: {
			{Point: Point{X: 10, Y: 10}, color: 0xFFFFFF},
			{Point: Point{X: 11, Y: 12}, color: 0xFFFFFF},
		},

		8: {
			{Point: Point{X: 10, Y: 10}, color: 0xFFFFFF},
			{Point: Point{X: 11, Y: 12}, color: 0xFFFFFF},
		},

		9: {
			{Point: Point{X: 10, Y: 10}, color: 0xFFFFFF},
			{Point: Point{X: 11, Y: 12}, color: 0xFFFFFF},
		},
	}
)
