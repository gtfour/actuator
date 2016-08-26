package dusk

//import "gopkg.in/mgo.v2/bson"

type Query struct {
    Type  int
    Table string
    Body  map[string]interface{}
}


func(q *Query)Run()(err error){
    return err
}

func AristoQuery(query map[string]interface{})(err error){
    return err
}
