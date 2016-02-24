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
    users_c_name            string
    actions_c_name          string
    dashboards_c_name       string
    dashboard_groups_c_name string
    tokens_c_name           string
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

func (d *MongoDb)CreateDashboardGroup(dgroup *dashboard.DashboardGroup)(dashboardgroup_id string, err error){
    c           := d.Session.DB(d.dbname).C(d.dashboard_groups_c_name)
    dgroup.Id,err = common.GenId()
    if err!=nil {
        return "",err
    }
    err         = c.Insert(dgroup)
    return dgroup.Id,err
}

func(d *MongoDb)RemoveDashboardGroup(dashboardgroup_id string)(err error){

    c      := d.Session.DB(d.dbname).C(d.dashboard_groups_c_name)
    err    =  c.Remove(bson.M{"id": dashboardgroup_id})
    return err



}

func (d *MongoDb)AddDashboardToGroup(dashboardgroup_id string,dashboard_id string) (err error) {
    dboard      := dashboard.Dashboard{}
    dgroup      := dashboard.DashboardGroup{}
    cdashboards := d.Session.DB(d.dbname).C(d.dashboards_c_name)
    cdgroups    := d.Session.DB(d.dbname).C(d.dashboard_groups_c_name)
    err         =  cdashboards.Find(bson.M{"id": dashboard_id}).One(&dboard)
    if err!=nil {return DashboardDoesNotExist()}
    err         =  cdgroups.Find(bson.M{"id": dashboardgroup_id}).One(&dgroup)
    if err!=nil {return DashboardGroupDoesNotExist()}
    err         =  cdgroups.Update(bson.M{"id": dashboardgroup_id},bson.M{"$push":bson.M{"list": dashboard_id}})
    return err
}

func (d *MongoDb)RemoveDashboardFromGroup(dashboardgroup_id string,dashboard_id string) (err error) {
    dboard      := dashboard.Dashboard{}
    dgroup      := dashboard.DashboardGroup{}
    cdashboards := d.Session.DB(d.dbname).C(d.dashboards_c_name)
    cdgroups    := d.Session.DB(d.dbname).C(d.dashboard_groups_c_name)
    err         =  cdashboards.Find(bson.M{"id": dashboard_id}).One(&dboard)
    if err!=nil {return DashboardDoesNotExist()}
    err         =  cdgroups.Find(bson.M{"id": dashboardgroup_id}).One(&dgroup)
    if err!=nil {return DashboardGroupDoesNotExist()}
    err         =  cdgroups.Update(bson.M{"id": dashboardgroup_id},bson.M{"$pull":bson.M{"list": dashboard_id}})
    return err
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

func (d *MongoDb)GetAllUsers()(users []utah.User,err error) {
    c      := d.Session.DB(d.dbname).C(d.users_c_name)
    err    =  c.Find(nil).All(&users)
    if err!=nil {
        return users,err
    }
    return users,nil
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

func(d *MongoDb)TokenExists(user_id string,token_id string ) (bool) {
    token  := utah.Token{}
    c      := d.Session.DB(d.dbname).C(d.tokens_c_name)
    err    :=  c.Find(bson.M{"userid": user_id, "id":token_id}).One(&token)
    if err!=nil { return false } else { return true }
}

func (d *MongoDb)DashboardExists(dashboard_id string)(bool){
    dashboard  := dashboard.Dashboard{}
    c          := d.Session.DB(d.dbname).C(d.dashboards_c_name)
    err        := c.Find(bson.M{"id": dashboard_id}).One(&dashboard)
    if err!=nil { return false } else { return true }
}

func (d *MongoDb)DashboardGroupExists(dgroup_id string)(bool){
    dgroup     := dashboard.DashboardGroup{}
    c          := d.Session.DB(d.dbname).C(d.dashboard_groups_c_name)
    err        := c.Find(bson.M{"id": dgroup_id}).One(&dgroup)
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
    if d.DashboardExists(dashboard_id) == false { return DashboardDoesNotExist()  }
    c      := d.Session.DB(d.dbname).C(d.users_c_name)
    err    :=  c.Update(bson.M{"id": user_id},bson.M{"$push":bson.M{"dashboards": dashboard_id}})
    if err != nil { return err }
    return nil
}
func(d *MongoDb) AttachDashboardGroupToUser (user_id,dgroup_id string)(error) {
    //user   := utah.User{}
    if d.DashboardGroupExists(dgroup_id) == false { return DashboardGroupDoesNotExist()  }
    c      := d.Session.DB(d.dbname).C(d.users_c_name)
    err    :=  c.Update(bson.M{"id": user_id},bson.M{"$push":bson.M{"dashboardgroups":dgroup_id}})
    if err != nil { return err }
    return nil
}


func(d *MongoDb) GetMyDashboardList (user_id,token_id string)(dashboard_list dashboard.DashboardList,err error) {
    //user   := utah.User{}
    dashboard_list.List = make([]dashboard.Dashboard, 0)
    if d.TokenExists(user_id,token_id) == false { return dashboard_list,TokenDoesNotExist() }
    user:=utah.User{}
    cu      := d.Session.DB(d.dbname).C(d.users_c_name)
    err     =  cu.Find(bson.M{"id": user_id}).One(&user)
    if err!= nil { return dashboard_list,err }
    cd      := d.Session.DB(d.dbname).C(d.dashboards_c_name)
    for d := range user.Dashboards {
        d_id:=user.Dashboards[d]
        dashboard:=dashboard.Dashboard{}
        err     =  cd.Find(bson.M{"id": d_id}).One(&dashboard)
        if err == nil { dashboard_list.List = append(dashboard_list.List, dashboard)  }
    }
    return dashboard_list,nil
}

func(d *MongoDb) GetMyDashboardGroupList (user_id,token_id string)(dgroup_list dashboard.DashboardGroupList, err error) {
    //user   := utah.User{}
    dgroup_list.List = make([]dashboard.DashboardGroup, 0)
    if d.TokenExists(user_id,token_id) == false { return dgroup_list, TokenDoesNotExist() }
    user:=utah.User{}
    cu      := d.Session.DB(d.dbname).C(d.users_c_name)
    err     =  cu.Find(bson.M{"id": user_id}).One(&user)
    if err!= nil { return dgroup_list,err }
    cd      := d.Session.DB(d.dbname).C(d.dashboard_groups_c_name)
    for d := range  user.DashboardGroups{
        dg_id := user.DashboardGroups[d]
        dgroup := dashboard.DashboardGroup{}
        err     =  cd.Find(bson.M{"id": dg_id}).One(&dgroup)
        if err == nil { dgroup_list.List = append(dgroup_list.List, dgroup) }
    }
    return dgroup_list, nil
}


func (d *MongoDb)Close()() {
    d.Session.Close()
}

func (d *MongoDb)Connect() ( err error ) {
    d.Session,err = mgo.Dial("mongodb://"+d.username+":"+d.password+"@"+d.host+"/"+d.dbname)
    d.Session.SetMode(mgo.Monotonic, true)
    d.users_c_name            ="dashboard_users"
    d.tokens_c_name           = "user_tokens"
    d.dashboards_c_name       = "dashboards"
    d.dashboard_groups_c_name = "dashboard_groups"
    return err
}
