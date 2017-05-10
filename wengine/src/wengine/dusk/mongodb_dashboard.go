package dusk

import "gopkg.in/mgo.v2/bson"
//import "wengine/core/common"
import "jumper/common/gen"
import "wengine/core/dashboard"


//mongodb_dashboard
func (d *MongoDb)CreateDashboard(dashboard *dashboard.Dashboard)(dashboard_id string, err error) {
    c           := d.Session.DB(d.dbname).C(d.dashboards_c_name)
    dashboard.Id,err = gen.GenId()
    if err!=nil {
        return "",err
    }
    err         = c.Insert(dashboard)
    return dashboard.Id,err
}

//mongodb_dashboard
func (d *MongoDb)CreateDashboardGroup(dgroup *dashboard.DashboardGroup)(dashboardgroup_id string, err error){
    c           := d.Session.DB(d.dbname).C(d.dashboard_groups_c_name)
    dgroup.Id,err = gen.GenId()
    if err!=nil {
        return "",err
    }
    err         = c.Insert(dgroup)
    return dgroup.Id,err
}

//mongodb_dashboard
func(d *MongoDb)RemoveDashboardGroup(dashboardgroup_id string)(err error){

    c      := d.Session.DB(d.dbname).C(d.dashboard_groups_c_name)
    err    =  c.Remove(bson.M{"id": dashboardgroup_id})
    return err



}

//mongodb_dashboard
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
//mongodb_dashboard
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


//mongodb_dashboard
func (d *MongoDb)DashboardExists(dashboard_id string)(bool){
    dashboard  := dashboard.Dashboard{}
    c          := d.Session.DB(d.dbname).C(d.dashboards_c_name)
    err        := c.Find(bson.M{"id": dashboard_id}).One(&dashboard)
    if err!=nil { return false } else { return true }
}
//mongodb_dashboard
func (d *MongoDb)DashboardGroupExists(dgroup_id string)(bool){
    dgroup     := dashboard.DashboardGroup{}
    c          := d.Session.DB(d.dbname).C(d.dashboard_groups_c_name)
    err        := c.Find(bson.M{"id": dgroup_id}).One(&dgroup)
    if err!=nil { return false } else { return true }
}


//mongodb_dashboard
func(d *MongoDb) AttachDashboardToUser (user_id,dashboard_id string)(error) {
    //user   := utah.User{}
    if d.DashboardExists(dashboard_id) == false { return DashboardDoesNotExist()  }
    c      := d.Session.DB(d.dbname).C(d.users_c_name)
    err    :=  c.Update(bson.M{"id": user_id},bson.M{"$push":bson.M{"dashboards": dashboard_id}})
    if err != nil { return err }
    return nil
}
//mongodb_dashboard
func(d *MongoDb) AttachDashboardGroupToUser (user_id,dgroup_id string)(error) {
    //user   := utah.User{}
    if d.DashboardGroupExists(dgroup_id) == false { return DashboardGroupDoesNotExist()  }
    c      := d.Session.DB(d.dbname).C(d.users_c_name)
    err    :=  c.Update(bson.M{"id": user_id},bson.M{"$push":bson.M{"dashboardgroups":dgroup_id}})
    if err != nil { return err }
    return nil
}




