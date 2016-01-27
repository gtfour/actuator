package dusk
import "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"
import "wengine/core/utah"
import "wengine/core/common"
import "wengine/core/dashboard"

type MongoDb struct {

    Session  *mgo.Session
    Users    []string
    Groups   []string
    username string
    password string
    host     string
    dbname   string
    CollectionsNames
}
type CollectionsNames struct {
    users_c_name      string
    actions_c_name    string
    dashboards_c_name string
    tokens_c_name     string
}
func (d *MongoDb)GetUsers()([]string) {
    return d.Users
}

func (d *MongoDb)GetGroups()([]string) {
    return d.Users
}

func (d *MongoDb)CreateUser(user *utah.User)(user_id string, err error) {
    c           := d.Session.DB(d.dbname).C(d.users_c_name)
    user.Id,err = common.GenId()
    if err!=nil {
        return "",err
    }
    err         = c.Insert(user)
    return user.Id,err
}

func (d *MongoDb)CreateDashboard(dashboard *dashboard.Dashboard)(dashboard_id string, err error) {
    c           := d.Session.DB(d.dbname).C(d.dashboards_c_name)
    dashboard.Id,err = common.GenId()
    if err!=nil {
        return "",err
    }
    err         = c.Insert(dashboard)
    return dashboard.Id,err
}

func (d *MongoDb)CreateToken(userid string)(token_id string,err error){
    _,err         = d.GetUserById(userid)
    if (err!=nil) {return "",err}
    token_id,err = common.GenId()
    if (err!=nil) {return "",err}
    token       := utah.Token{UserId:userid,Id:token_id}
    c           := d.Session.DB(d.dbname).C(d.tokens_c_name)
    err         = c.Insert(token)
    return token_id,err

}



func (d *MongoDb)GetUserById(id string)(user utah.User,err error) {
    c      := d.Session.DB(d.dbname).C(d.users_c_name)
    err    =  c.Find(bson.M{"id": id}).One(&user)
    if err!=nil {
        return user,err
    }
    return user,nil
}

func (d *MongoDb)GetUser(user_prop map[string]interface{})(user utah.User,err error) {
    c      := d.Session.DB(d.dbname).C(d.users_c_name)
    err    =  c.Find(bson.M(user_prop)).One(&user)
    if err!=nil {
        return user,err
    }
    return user,nil
}

func (d *MongoDb)RemoveUsersById(ids ...string)(err error){

    c      := d.Session.DB(d.dbname).C(d.users_c_name)
    for id := range ids {
        user_id:=ids[id]
        err    =  c.Remove(bson.M{"id": user_id})
        if err != nil { continue }
    }
    return nil
}

func(d *MongoDb)TokenIsExist(user_id string,token_id string ) (bool) {
    token  := utah.Token{}
    c      := d.Session.DB(d.dbname).C(d.tokens_c_name)
    err    :=  c.Find(bson.M{"userid": user_id, "id":token_id}).One(&token)
    if err!=nil { return false } else { return true }
}
func(d *MongoDb)RemoveToken(token_id string ,user_id string)(error) {

    c      := d.Session.DB(d.dbname).C(d.tokens_c_name)
    err    := c.Remove(bson.M{"id": token_id,"userid":user_id })
    return err

}

func(d *MongoDb)GetAnUserToken(user_id string) (string,error) {
    token  := utah.Token{}
    c      := d.Session.DB(d.dbname).C(d.tokens_c_name)
    err    :=  c.Find(bson.M{"userid": user_id}).One(&token)
    if err!=nil { return "",err } else { return token.Id,nil  }
}

func(d *MongoDb)UserPasswordIsCorrect(username,password string)(string,string,bool) {
    user   := utah.User{}
    c      := d.Session.DB(d.dbname).C(d.users_c_name)
    err    :=  c.Find(bson.M{"name": username, "password":password}).One(&user)
    if err!=nil { return "","",false } else {
        token_id,err:= d.GetAnUserToken(user.Id)
        if err!= nil {
            new_token_id,err := d.CreateToken(user.Id)
            if err!=nil {
                return "","",false
            } else {
                return user.Id,new_token_id,true
            }

        } else {
            return user.Id, token_id, true

        }
        return "","",false
    }
}

func(d *MongoDb) AttachDashboardToUser (user_id,dashboard_id string)(error) {
    //user   := utah.User{}
    c      := d.Session.DB(d.dbname).C(d.users_c_name)
    err    :=  c.Update(bson.M{"id": user_id},bson.M{"$push":bson.M{"dashboardids": dashboard_id}})
    if err != nil { return err }
    return nil

}

func (d *MongoDb)Close()() {
    d.Session.Close()
}

func (d *MongoDb)Connect() ( err error ) {
    d.Session,err = mgo.Dial("mongodb://"+d.username+":"+d.password+"@"+d.host+"/"+d.dbname)
    d.Session.SetMode(mgo.Monotonic, true)
    d.users_c_name      ="dashboard_users"
    d.tokens_c_name     = "user_tokens"
    d.dashboards_c_name = "dashboards"
    return err
}
