package logger

import "fmt"

type escCode int8

const (
	reset      escCode = 0
	dim        escCode = 2
	slowblink  escCode = 5
	foreground escCode = 38
	background escCode = 48
)

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
