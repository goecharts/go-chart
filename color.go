package charts

import "fmt"

type Color struct {
	H int
	E int
	X int
}

var (
	red = Color{
		H: 255,
		E: 0,
		X: 0,
	}
	orange = Color{
		H: 255,
		E: 165,
		X: 0,
	}
	yellow = Color{
		H: 255,
		E: 255,
		X: 0,
	}
	green = Color{
		H: 0,
		E: 128,
		X: 0,
	}

	cyan = Color{
		H: 0,
		E: 255,
		X: 255,
	}
	blue = Color{
		H: 0,
		E: 0,
		X: 255,
	}
	purple = Color{
		H: 160,
		E: 32,
		X: 240,
	}
	black = Color{
		H: 0,
		E: 0,
		X: 0,
	}
	white = Color{
		H: 255,
		E: 255,
		X: 255,
	}
)

var (
	RGB  = "rgb(%d,%d,%d)"
	RGBA = "rgba(%d,%d,%d,%f)"
)

func toColor(c Color) string {
	return fmt.Sprintf(RGB, c.H, c.E, c.X)
}

func toColorA(c Color, alpha float32) string {
	return fmt.Sprintf(RGBA, c.H, c.E, c.X, alpha)
}
func Red() string {
	return toColor(red)
}

func RedAlpha(alpha float32) string {
	return toColorA(red, alpha)
}

func Orange() string {
	return toColor(orange)
}

func OrangeAlpha(alpha float32) string {
	return toColorA(orange, alpha)
}

func Yellow() string {
	return toColor(yellow)
}
func YellowAlpha(alpha float32) string {
	return toColorA(yellow, alpha)
}

func Green() string {
	return toColor(green)
}

func GreenAlpha(alpha float32) string {
	return toColorA(green, alpha)
}

func Cyan() string {
	return toColor(cyan)
}

func CyanAlpha(alpha float32) string {
	return toColorA(cyan, alpha)
}

func Blue() string {
	return toColor(blue)
}

func BlueAlpha(alpha float32) string {
	return toColorA(blue, alpha)
}

func Purple() string {
	return toColor(purple)
}
func PurpleAlpha(alpha float32) string {
	return toColorA(purple, alpha)
}

func Black() string {
	return toColor(black)
}

func BlackAlpha(alpha float32) string {
	return toColorA(black, alpha)
}

func White() string {
	return toColor(white)
}

func WhiteAplha(alpha float32) string {
	return toColorA(white, alpha)
}
