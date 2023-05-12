package logger

import (
	"fmt"
)

type escCode uint8

// option codes
const (
	reset      escCode = 0
	dim        escCode = 2
	slowblink  escCode = 5
	foreground escCode = 38
	background escCode = 48
)

// colors
const (
	RED    escCode = 9
	YELLOW escCode = 11
	AQUA   escCode = 14
)

type RGBCode struct {
	r, g, b escCode
}

func optionBuilder(codes ...escCode) string {
	res := "\033["
	for index, code := range codes {
		if index < len(codes)-1 {
			res += fmt.Sprintf("%d", code) + ";"
		} else {
			res += fmt.Sprintf("%d", code)
		}
	}
	res += "m"
	return res
}

func msgBuilder(msg string, codes ...escCode) string {
	resetOption := optionBuilder(reset)
	style := optionBuilder(codes...)
	return style + msg + resetOption
}

func msgBuilderRGB(msg string, code escCode, rgb RGBCode) string {
	resetOption := optionBuilder(reset)
	style := optionBuilder(code, dim, rgb.r, rgb.g, rgb.b)
	return style + msg + resetOption
}
