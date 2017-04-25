package main

import "fmt"
import "errors"

var targetWasNotConfigured         =  errors.New("cuda:Target was not configured")

type TargetList        []*Target
type TargetListRigaSLR []Target

type Target struct {
    //
    //
    //  #  Get:       ()(lineAsArray [][]string, err error)
    //  #  GetType:   ()(typ int)
    //  #  Gather:    ()(error)
    //  #  PushPart:  ([][]string)(error)
    //
    //  #  Target could  store content of line, file or also just decribe an directory 
    //  #  correction: section could not be determined as section on this level of processing
    //
    //
    selfIndex                 int         //  // self uniq   number 
    parentIndex               int         //  // uniq parent target number
    typ                       int
    path                      string
    pathShort                 string
    //lineAsArray             [][]string
    lines                     []string
    configured                bool
    gatherFailed              bool
    //
    diving                    bool  // gathering nested directories. seems that i can't implement this feauture yet here
    nestedTargets             []*Target
    //
    //
    isLogFile                 bool
    isDirectoryWithLogFiles   bool
    offset                    int64 // for log files 
    //
    //
}

func(tl *TargetList)Append(t *Target)(err error){
    if t.configured {
        (*tl) = append((*tl), t)
        return nil
    } else {
        return targetWasNotConfigured
    }
}

func(tl *TargetListRigaSLR)Append(t *Target)(err error){
    if t.configured {
        var target Target
        target = *t
        (*tl) = append((*tl), target)
        return nil
    } else {
        return targetWasNotConfigured
    }
}


func main(){

    var targets      TargetList
    var targetsSLR   TargetListRigaSLR

    var target       Target

    target.lines      = []string{"a","b"}
    target.configured = true

    targets.Append(&target)
    targets.Append(&target)
    fmt.Printf("\nslr len %d\n",len(targetsSLR))
    targetsSLR.Append(&target)
    fmt.Printf("\nslr len %d\n",len(targetsSLR))
    targetsSLR.Append(&target)
    fmt.Printf("\nslr len %d\n",len(targetsSLR))

    fmt.Printf("test#1:\nTargets:\n")
    for i:= range targets {
        target:=targets[i]
        fmt.Printf("%v\n", target)

    }
    fmt.Printf("TargetsSLR : %v\n"         , targetsSLR )

    target.lines = []string{"y","z"}
    fmt.Printf("test#2:\nTargets:\n")
    for i:= range targets {
        target:=targets[i]
        fmt.Printf("%v\n", target)

    }
    fmt.Printf("TargetsSLR : %v\n"         , targetsSLR )

}
