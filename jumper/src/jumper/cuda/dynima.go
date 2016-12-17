package cuda

type Dynima struct {

    filters []Filter

}

func( d *Dynima ) AppendFilter (f *Filter)(error){
    return nil
}

func( d *Dynima ) RunFilters   ()(r *Result,err error){
    return r,err
}
