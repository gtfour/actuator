package main

import ( "fmt" ; "os" )

func main() {

    file, _ := os.Open("/tmp/test/test1/hello.txt")

    //fmt.Println(path)

    defer file.Close()


    file_info , _ := file.Stat()



    file_mode :=  file_info.Mode()
    perm      :=  file_mode.Perm()

    fmt.Printf("%s",perm)


}
