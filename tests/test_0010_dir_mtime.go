package main
import ("os" ; "fmt")

func main() {

fi, _:=os.Stat("/var/test")
mtime:=fi.ModTime()

fmt.Println(mtime)





}
