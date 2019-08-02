package main

import "fmt"
import "unicode/utf8"

func center(s string, w int) string {
	return fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w + utf8.RuneCountInString(s))/2, s))
}

func setcolor(s string, color string) string{
    RESET := "\u001B[0m";
    var colorDict = map[string]string{
        "BLACK":"\u001B[30m",
        "RED" : "\u001B[31m",
        "GREEN" : "\u001B[32m",
        "YELLOW" : "\u001B[33m",
        "BLUE" : "\u001B[34m",
        "PURPLE" : "\u001B[35m",
        "CYAN" : "\u001B[36m",
        "WHITE" : "\u001B[37m",
    }
    return fmt.Sprintf("%s%s%s",colorDict[color],s,RESET);
}
