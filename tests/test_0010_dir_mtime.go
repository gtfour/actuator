package main
import ("os" ; "fmt")

func main() {

fi, _:=os.Stat("/tmp/test2/test5")
mtime:=fi.ModTime()

fmt.Println(mtime)





}
