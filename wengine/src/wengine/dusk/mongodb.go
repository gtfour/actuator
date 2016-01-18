package dusk
import "gopkg.in/mgo.v2"
import "wengine/core/utah"
//import "gopkg.in/mgo.v2/bson"

type MongoDb struct {

    session  *mgo.Session
    Users    []string
    Groups   []string
    username string
    password string
    host     string
    dbname   string

}
func (d MongoDb)GetUsers()([]string) {
    return d.Users
}

func (d MongoDb)GetGroups()([]string) {
    return d.Users
}

func (d MongoDb)CreateUser(user *utah.User)(err error) {
    c      :=d.session.DB(d.dbname).C("users")
    err = c.Insert(user)
    return err
}
func (d MongoDb)GetUser()([]string) {
    return d.Users
}

func (d MongoDb)Close()() {
    d.session.Close()
}

func (d MongoDb)Connect() ( err error  ) {
    d.session,err = mgo.Dial("mongodb://"+d.username+":"+d.password+"@"+d.host+"/"+d.dbname)
    d.session.SetMode(mgo.Monotonic, true)
    return err
}
