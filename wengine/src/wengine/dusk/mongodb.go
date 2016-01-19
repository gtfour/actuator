package dusk
import "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"
import "wengine/core/utah"
import "wengine/core/common"

type MongoDb struct {

    Session  *mgo.Session
    Users    []string
    Groups   []string
    username string
    password string
    host     string
    dbname   string

}
func (d *MongoDb)GetUsers()([]string) {
    return d.Users
}

func (d *MongoDb)GetGroups()([]string) {
    return d.Users
}

func (d *MongoDb)CreateUser(user *utah.User)(err error) {
    c           := d.Session.DB(d.dbname).C("dashboard_users")
    user.Id,err = common.GenId()
    if err!=nil {
        return err
    }
    err         = c.Insert(user)
    return err
}
func (d *MongoDb)GetUser(id string)(user utah.User,err error) {
    c      := d.Session.DB(d.dbname).C("dashboard_users")
    err    =  c.Find(bson.M{"id": id}).One(&user)
    if err!=nil {
        return user,err
    }
    return user,nil
}

func (d *MongoDb)Close()() {
    d.Session.Close()
}

func (d *MongoDb)Connect() ( err error ) {
    d.Session,err = mgo.Dial("mongodb://"+d.username+":"+d.password+"@"+d.host+"/"+d.dbname)
    d.Session.SetMode(mgo.Monotonic, true)
    return err
}
