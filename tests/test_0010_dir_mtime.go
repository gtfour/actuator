package main
import ("os" ; "fmt")

func main() {

fi, _:=os.Stat("/etc/apt")
mtime:=fi.ModTime()

fmt.Println(mtime)





}
