package chase

import "client/majesta"

type chaser struct {
    //
    wp       WorkerPool
    messages chan majesta.CompNotes
    //
}

func(c *chaser)Orient(messages chan majesta.CompNotes)(){
    //
    c.messages = messages
    // 
}


func(c *chaser)Follow(path string)(err error){
    //
    err = Listen(path, c.messages, c.wp)
    return err
    //
}

func NewChaser(mchan_size int)(*chaser){
    var c chaser
    c.wp       = WPCreate()
    if mchan_size > 0 {
        c.messages = make(chan majesta.CompNotes, mchan_size)
    }
    return &c
}
