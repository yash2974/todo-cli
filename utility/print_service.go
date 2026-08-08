package utility

import (
	"github.com/fatih/color"
)

func PrintCMD(text string, textColor string) {
    switch textColor {
    case "Red":
        color.New(color.FgRed).Print(text)
    case "Green":
        color.New(color.FgGreen).Print(text)
    case "Yellow":
        color.New(color.FgYellow).Print(text)
    case "Blue":
        color.New(color.FgBlue).Print(text)
    case "Orange":
        color.RGB(193, 95, 60).Print(text)
    case "White":
        color.RGB(177, 173, 161).Print(text)
    }
}