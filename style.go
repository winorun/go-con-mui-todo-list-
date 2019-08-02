package main

import "fmt"
import "strings"

func showTitle(w int, title string){
    fmt.Printf("┏%s┓\n",strings.Repeat("━",w));
    fmt.Printf("┃%s┃\n",center(title,w))
    fmt.Printf("┗%s┛\n\n",strings.Repeat("━",w));
}

func showHeader2(w int, header string){
    fmt.Printf("%s\n\n",setcolor(center(header,w),"BLUE"))
}

func showHeader1(w int, header string){
    fmt.Printf("%s\n\n",setcolor(strings.ToUpper(center(header,w)),"BLUE"))
}

func showItalicText(text string){
    RESET := "\u001B[0m";
    fmt.Printf("\u001B[3m%s%s\n",text,RESET)
}

func showBoldText(text string){
    RESET := "\u001B[0m";
    fmt.Printf("\u001B[1m%s%s\n",text,RESET)
}

func showTablet(){
    
}
