package src

import (
	"math"
	"strconv"
	"strings"
)

type HSL struct {
	H, S, L float64
}
type RGB struct {
	R, G, B int
}

func GetColor(s string) string {
	Colors := make(map[string]string)
	Colors["RED"] = "\033[31;1m"
	Colors["GREEN"] = "\033[32;1m"
	Colors["YELLOW"] = "\033[33;1m"
	Colors["ORANGE"] = "\033[33;1m"
	Colors["BLUE"] = "\033[34;1m"
	Colors["PURPLE"] = "\033[35;1m"
	Colors["CYAN"] = "\033[36;1m"
	Colors["GREY"] = "\033[37;1m"
	s = strings.TrimPrefix(s, "--color=")
	if strings.Contains(s, "rgb") {
		co := strings.ReplaceAll(strings.TrimPrefix(s, "rgb("), " ", "")
		co = strings.ReplaceAll(co, ")", "")
		c := strings.Split(co, ",")
		return "\033[38;2;" + c[0] + ";" + c[1] + ";" + c[2] + "m"
	}
	if strings.Contains(s, "hsl") {
		co := strings.ReplaceAll(strings.TrimPrefix(s, "hsl("), " ", "")
		co = strings.ReplaceAll(co, ")", "")
		c := strings.Split(co, ",")
		vals := []float64{}
		for _, value := range c {
			value = strings.ReplaceAll(value, "%", "")
			val, _ := strconv.ParseFloat(value, 32)
			vals = append(vals, val)
		}
		h, s, l := vals[0], vals[1], vals[2]
		if strings.Contains(c[0], "%") {
			h = vals[0] / 100.0
		}
		if strings.Contains(c[1], "%") {
			s = vals[1] / 100.0
		}
		if strings.Contains(c[2], "%") {
			l = vals[2] / 100.0
		}
		rgb := hslToRGB(HSL{h, s, l})
		return "\033[38;2;" + strconv.Itoa(rgb.R) + ";" + strconv.Itoa(rgb.G) + ";" + strconv.Itoa(rgb.B) + "m"
	}
	return Colors[strings.ToUpper(s)]
}
func hslToRGB(hsl HSL) RGB {
	s := hsl.S
	l := hsl.L
	c := (1 - math.Abs(2*l-float64(1))) * s
	x := c * (float64(1) - math.Abs(math.Mod(float64(hsl.H)/60, 2)-float64(1)))
	m := l - c/2
	var r, g, b float64
	switch {
	case hsl.H < 60:
		r = c
		g = x
		b = 0
	case hsl.H < 120 && hsl.H >= 60:
		r = x
		g = c
		b = 0
	case hsl.H < 180 && hsl.H >= 120:
		r = 0
		g = c
		b = x
	case hsl.H < 240 && hsl.H >= 180:
		r = 0
		g = x
		b = c
	case hsl.H < 300 && hsl.H >= 240:
		r = x
		g = 0
		b = c
	case hsl.H < 360 && hsl.H >= 300:
		r = c
		g = 0
		b = x
	}
	R := int(math.Round((r + m) * 255))
	G := int(math.Round((g + m) * 255))
	B := int(math.Round((b + m) * 255))
	return RGB{R, G, B}
}
