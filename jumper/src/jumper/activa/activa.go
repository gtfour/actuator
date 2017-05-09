package activa

type MotionsPool struct {


}

func (pool *MotionsPool)Handle()(){


}


func Handle( motions chan *Motion )() {
    for {
        select {
            case motion:=<-motions:
                _ = motion
                // cross.WriteMotion(motion)
                // fmt.Printf("<<New motion:\n%v\n>>", motion)

            }
        }
}

