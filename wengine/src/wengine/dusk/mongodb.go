package dusk
import "gopkg.in/mgo.v2"
//import "gopkg.in/mgo.v2/bson"

type MongoDb struct {

    Session *mgo.Session
    Users  []string
    Groups []string


}
func (d *Mongodb)GetUsers()([]string) {
    return d.Users
}

func (d *Mongodb)GetUsers()([]string) {
    return d.Users
}

func (d *Mongodb)GetUsers()([]string) {
    return d.Users
}
func (d *Mongodb)GetUsers()([]string) {
    return d.Users
}



func Connect( username string, password string, host string, dbname string ) ( session *mgo.Session, err error  ) {

    return mgo.Dial("mongodb://"+username+":"+password+"@"+host+"/"+dbname)


}



