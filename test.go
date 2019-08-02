package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

//~ import "golang.org/x/crypto/ssh/terminal"

type taskStruct struct {
    check  bool
    task   string
}

type windowsStruct struct {
    width  uint
    height uint
    x      uint
    y      uint
}

//~ import "github.com/gbin/goncurses"
//~ import "log"

func drawTaskList(taskList []taskStruct,window windowsStruct){
    col := int(window.x)
    w := int(window.width)
    
    for i, task := range taskList{
        row := i+2
        fmt.Printf("\u001B[%d;%dH",row,col)
        fmt.Print(strings.Repeat(" ",w))
        fmt.Printf("\u001B[%d;%dH",row,col)
        var checkStr string
        if task.check {
            checkStr = "☑"
        }else{
            checkStr = "☐"
        }
        fmt.Printf("  %s  %s",checkStr,task.task)
    }
}

func drawWindows(w uint, h uint) (windowsStruct,windowsStruct){
    fmt.Print("\u001B[2J\u001B[1;1H")
    wWindowHeader:=w/3
    wWindowTask:=w-wWindowHeader
    repeatWidthWindowsHeader := int(wWindowHeader-1)
    repeatWidthWindowsTask:=int(wWindowTask-2)
    heightWindows := h-4
    fmt.Printf("┏%s┳%s┓\n",strings.Repeat("━",repeatWidthWindowsHeader),strings.Repeat("━",repeatWidthWindowsTask));
    for i:= uint(0); i< heightWindows;i++ {
        fmt.Printf("┃%s┃%s┃\n",strings.Repeat(" ",repeatWidthWindowsHeader),strings.Repeat(" ",repeatWidthWindowsTask))
    }
    fmt.Printf("┗%s┻%s┛",strings.Repeat("━",repeatWidthWindowsHeader),strings.Repeat("━",repeatWidthWindowsTask));
    windowHeader := windowsStruct{x:2, y:2, width: uint(repeatWidthWindowsHeader), height: heightWindows}
    windowTask := windowsStruct{x:(2+wWindowHeader), y:2, width: uint(repeatWidthWindowsTask), height: heightWindows}
    return windowHeader,windowTask
}

func getSignal() string {
    return getSignalMain()
}

func getSignalMain() string {
    reader := bufio.NewReader(os.Stdin)
    
    var buf []byte=make([]byte, 10)
    
    size, _ := reader.Read(buf)
    switch size{
        case 1: 
        switch buf[0]{
            case 'q': return "exit"
        }
        //~ fmt.Printf("%c",buf[0]);
        case 3: break
        case 6: break
    }
    return "none"
}

func main() {
    //~ fmt.Println(getTerminalSize());
    //~ showTitle(80,"Привет, Мир");
    //~ showHeader1(80,"Заголовок первого уровня");
    //~ showHeader2(80,"Заголовок второго уровня");
    //~ showToDoItem(true,"True item");
    //~ showToDoItem(false,"False item");
    //~ fmt.Printf("\n");showItalicText("Italic");fmt.Printf("\n");
    //~ fmt.Printf("\n");showBoldText("BoldText");fmt.Printf("\n");
    
    //~ var name string
    //~ var age int
    
    oldState, err := makeRaw(0)
    if err != nil {
            panic(err)
    }
    defer restore(0, oldState)
    
    task := &taskStruct{}
    
    width, height := getTerminalSize()
    //~ WindowsHeader,WindowsTask := drawWindows(width,height)
    _,windowTask := drawWindows(width,height)
    drawWindows(width,height)
    
    var taskList []taskStruct
    
    fmt.Print("\u001B[?1002h");

L:  for {
        fmt.Printf("\u001B[%d;2H\u001B[0KВведите имя задачи: ",height-1)
        //~ buf, _ := reader.ReadString('\n')
        switch signal:= getSignal(); signal{
            case "exit": break L
            case "none": break
            default:
                task.check = false
                task.task = signal
                taskList = append(taskList,*task);
                drawTaskList(taskList,windowTask);
        }
    }

    fmt.Printf("\u001B[%d;%dH",height,1)
    fmt.Print("\u001B[?1002l");
    
    //~ fmt.Print("\u001B[0F\u001B[2KВведите пол:")
    //~ fmt.Print("\u001B[0F\u001B[2KВведите возраст:")
    
    //~ src, err := goncurses.Init()
    //~ _, err := goncurses.Init()
    //~ if err != nil {
        //~ log.Fatal("init:", err)
    //~ }
    //~ defer goncurses.End()

    //~ goncurses.Echo(false);
    //~ goncurses.CBreak(true);
    //~ goncurses.MouseMask(goncurses.M_ALL,nil);

    //~ src.Keypad(true);

    //~ //fmt.Printf("\u001Bc %X flag - %d x - %d y - %d \n",name[0:3], name[4], name[5]-32,name[6]-32);
  //~ L:
    //~ for ;;{
        //~ fmt.Print("\u001B[0F\u001B[2KВведите возраст:")
        //~ switch char := src.GetChar(); char{
            //~ case goncurses.KEY_TAB:
                //~ break L
            //~ case goncurses.KEY_RETURN:
                //~ break L
            //~ case goncurses.KEY_MOUSE:
                //~ mouse:=goncurses.GetMouse()
                //~ src.Printf("X - %d Y- %d st -%d \n",mouse.X,mouse.Y,mouse.State)
            //~ default:
                //~ src.Printf("%X ",char)
                //~ src.Printf("%X ",goncurses.KEY_MOUSE)
            //~ }
    //~ }
    
}
