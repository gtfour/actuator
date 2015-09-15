package main
import ("os" ; "fmt" ; "time" ; "syscall")

func main() {

fi, _:=os.Stat("/tmp/test")
//mtime:=fi.ModTime()

stat := fi.Sys().(*syscall.Stat_t)
atime := time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
//ctime := time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))

fmt.Println(atime)





}
