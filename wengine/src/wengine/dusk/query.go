package dusk

import "gopkg.in/mgo.v2/bson"

type Query struct {

    Type      int
    Table     string
    KeyBody   map[string]interface{}
    QueryBody map[string]interface{}

}


func(d *MongoDb)RunQuery(q Query)(result map[string]interface{}, err error){
    c      := d.Session.DB(d.dbname).C(q.Table)

    if q.Type == CREATE_NEW || q.Type == UPDATE || q.Type == EDIT   {


    } else if q.Type == GET  || q.Type == CHECK_EXIST {


    } else  if q.Type == REMOVE {


    } else {
        return nil,err


    }

    err    =  c.Find(bson.M(q.KeyBody)).One(&result)
    if err != nil {
        return result,err
    }
    return result, err
}
