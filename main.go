package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "encoding/json"
  "log"
)



var defaultTaskList []byte = []byte(
`{
  "header": "default",
  "tasks": [
    {
      "task": "Пример",
      "status": false
    },
    {
      "task": "Пример2",
      "status": true
    }
  ]
}`)


type taskStruct struct {
    Check  bool `json:"status"`
    Task   string `json:"task"`
}

type notepud struct{
    Header string `json:"header"`
    Task []taskStruct `json:"tasks"`
}

type windowsStruct struct {
    width  uint
    height uint
    x      uint
    y      uint
}

func loadDefaultTaskList() notepud {
    var a notepud
    if err:=json.Unmarshal(defaultTaskList, &a);err!=nil{
        log.Fatalf("Ошибка файла данных: %s",err);
    }
    return a
}

func drawTaskList(taskList []taskStruct,window windowsStruct){
    col := int(window.x)
    w := int(window.width)
    
    for i, task := range taskList{
        row := i+2
        fmt.Printf("\u001B[%d;%dH",row,col)
        fmt.Print(strings.Repeat(" ",w))
        fmt.Printf("\u001B[%d;%dH",row,col)
        var checkStr string
        if task.Check {
            checkStr = "☑"
        }else{
            checkStr = "☐"
        }
        fmt.Printf("  %s  %s",checkStr,task.Task)
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

    oldState, err := makeRaw(0)
    if err != nil {
            panic(err)
    }
    defer restore(0, oldState)
    
    width, height := getTerminalSize()
    //~ WindowsHeader,WindowsTask := drawWindows(width,height)
    _,windowTask := drawWindows(width,height)
    
    nodepudList := loadDefaultTaskList();
    taskList:= nodepudList.Task
    drawTaskList(taskList,windowTask)
    
    fmt.Print("\u001B[?1002h");

L:  for {
        task := &taskStruct{}
        fmt.Printf("\u001B[%d;2H\u001B[0KВведите имя задачи: ",height-1)
        //~ buf, _ := reader.ReadString('\n')
        switch signal:= getSignal(); signal{
            case "exit": break L
            case "none": break
            default:
                task.Check = false
                task.Task = signal
                taskList = append(taskList,*task);
                drawTaskList(taskList,windowTask);
        }
    }

    fmt.Printf("\u001B[%d;%dH",height,1)
    fmt.Print("\u001B[?1002l");
}
