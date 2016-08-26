package dusk

//import "gopkg.in/mgo.v2/bson"

type Query struct {

    Type      int
    Table     string
    KeyBody   map[string]interface{}
    QueryBody map[string]interface{}

}


func(d *MongoDb)RunQuery(q *Query)(result map[string]interface{}, err error){
    return result, err
}

