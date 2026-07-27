package utility

import (
	"github.com/fatih/color"
)

func PrintCMD(text string, textColor string) {
    switch textColor {
    case "Red":
        color.Red(text)
    case "Green":
        color.Green(text)
    case "Yellow":
        color.Yellow(text)
    case "Blue":
        color.Blue(text)
    case "Orange":
        color.RGB(193, 95, 60).Print(text)
    case "White":
        color.RGB(177, 173, 161).Print(text)  
    }  
}