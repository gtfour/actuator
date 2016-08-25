package aristo

type Query struct {
    Type  int
    Table string
    Body  map[string]interface{}
}


func(q *Query)Run()(err error){
    return err
}
