package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("ifconfig","-a")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
        lines:= strings.Split(out.String(),"\n")
        for i:=range lines {

            fmt.Printf("\n%s",lines[i])


        }
	//fmt.Printf("in all caps: %q\n", out.String())
}
