package console

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)


func colorOut(message,color string){
	fmt.Fprintln(os.Stdout,ansi.Color(message,color))
}
func Success(msg string){
	colorOut(msg,"green")
}
func Warning(msg string){
	colorOut(msg,"yellow")
}
func Error(msg string){
	colorOut(msg,"red")
}
func Exit(msg string){
	Error(msg)
	os.Exit(1)
}

func ExitIf(err error){
	if err!=nil{
		Exit(err.Error())
	}
}