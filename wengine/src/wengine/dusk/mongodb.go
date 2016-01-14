package dusk
import "gopkg.in/mgo.v2"
//import "gopkg.in/mgo.v2/bson"

type MongoDb struct {



}

func CreateMongoDBConn( username string, password string, host string, dbname string ) ( session *mgo.Session, err error  ) {

    return mgo.Dial("mongodb://"+username+":"+password+"@"+host+"/"+dbname)


}



