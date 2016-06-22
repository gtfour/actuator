package dusk

import "gopkg.in/mgo.v2/bson"
import "wengine/core/utah"
import "wengine/core/common"
import "wengine/core/dashboard"



//mongodb_user
func (d *MongoDb)GetUsers()([]string) {
    return d.Users
}

//mongodb_user
func (d *MongoDb)GetGroups()([]string) {
    return d.Users
}

//mongodb_user
func (d *MongoDb)CreateUser(user *utah.User)(user_id string, err error) {
    c           := d.Session.DB(d.dbname).C(d.users_c_name)
    user.Id,err = common.GenId()
    if err!=nil {
        return "",err
    }
    err         = c.Insert(user)
    return user.Id,err
}


//mongodb_user
func (d *MongoDb)GetUserById(id string)(user utah.User,err error) {
    c      := d.Session.DB(d.dbname).C(d.users_c_name)
    err    =  c.Find(bson.M{"id": id}).One(&user)
    if err!=nil {
        return user,err
    }
    return user,nil
}
//mongodb_user
func (d *MongoDb)GetAllUsers()(users []utah.User,err error) {
    c      := d.Session.DB(d.dbname).C(d.users_c_name)
    err    =  c.Find(nil).All(&users)
    if err!=nil {
        return users,err
    }
    return users,nil
}

//mongodb_user
func (d *MongoDb)GetUser(user_prop map[string]interface{})(user utah.User,err error) {
    c      := d.Session.DB(d.dbname).C(d.users_c_name)
    err    =  c.Find(bson.M(user_prop)).One(&user)
    if err!=nil {
        return user,err
    }
    return user,nil
}
//mongodb_user
func (d *MongoDb)RemoveUsersById(ids ...string)(err error){

    c      := d.Session.DB(d.dbname).C(d.users_c_name)
    for id := range ids {
        user_id:=ids[id]
        err    =  c.Remove(bson.M{"id": user_id})
        if err != nil { continue }
    }
    return nil
}

//mongodb_user
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

//mongodb_user
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

//mongodb_user
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
