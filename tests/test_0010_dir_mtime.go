package main
import ("os" ; "fmt")

func main() {

fi, _:=os.Stat("/tmp")
mtime:=fi.ModTime()

fmt.Println(mtime)





}
