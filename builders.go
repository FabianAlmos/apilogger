package logger

import (
	"bufio"
	"fmt"
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

func write(msg string) {
	writer := bufio.NewWriter(logFile)
	writer.WriteString(msg + "\n")
	writer.Flush()
}

func msgBuilder(msg string, codes ...escCode) string {
	resetOption := optionBuilder(reset)
	style := optionBuilder(codes...)
	builtMsg := style + msg + resetOption
	switch {
	case Config.Out&TERMINAL == TERMINAL && Config.Out != ALL:
		{
			fmt.Println(builtMsg)
			return msg
		}
	case Config.Out&FILE == FILE && Config.Out != ALL:
		{
			write(msg)
			return msg
		}
	case Config.Out&ALL == ALL:
		{
			fmt.Println(builtMsg)
			write(msg)
			return msg
		}
	default:
		return msg
	}
}

func msgBuilderRGB(msg string, code escCode, rgb RGBCode) string {
	resetOption := optionBuilder(reset)
	style := optionBuilder(code, dim, rgb.R, rgb.G, rgb.B)
	builtMsg := style + msg + resetOption
	switch {
	case Config.Out&TERMINAL == TERMINAL:
		{
			fmt.Println(builtMsg)
			return msg
		}
	case Config.Out&FILE == FILE:
		{
			write(msg)
			return msg
		}
	case Config.Out&ALL == ALL:
		{
			fmt.Println(builtMsg)
			write(msg)
			return msg
		}
	default:
		return msg
	}
}
