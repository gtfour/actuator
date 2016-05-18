package main
import "wengine/dusk"
import "wengine/core/dashboard"
//import "gopkg.in/mgo.v2/bson"
import .  "wengine/core/utah"
//import "fmt"

func main() {

    d                     := dusk.OpenDatabase("mongo","wengine","OpenStack123","127.0.0.1","wengine")
    user_root             := &User{Name:"root", Password:"OpenStack123"}
    user_riah             := &User{Name:"riah", Password:"OpenStack123"}
    user_kelly            := &User{Name:"kelly", Password:"OpenStack123"}
    user_root_id,_        := d.CreateUser(user_root)
    user_riah_id,_        := d.CreateUser(user_riah)
    user_kelly_id,_        := d.CreateUser(user_kelly)

    new_dashboard1       :=&dashboard.Dashboard{Id:"groups_dashboard",              Title:"Groups",     Url:"groups-dashboard", Icon:"fa fa-circle-o"}
    new_dashboard2       :=&dashboard.Dashboard{Id:"networking_hosts_dashboard",    Title:"Hosts",      Url:"networking-hosts-dashboard", Icon:"fa fa-circle-o"}
    new_dashboard3       :=&dashboard.Dashboard{Id:"networking_nsswitch_dashboard", Title:"Nsswitch",   Url:"networking-nsswitch-dashboard", Icon:"fa fa-circle-o"}
    new_dashboard4       :=&dashboard.Dashboard{Id:"hardware_partitions_dashboard", Title:"Partitions", Url:"hardware-partitions-dashboard", Icon:"fa fa-circle-o"}
    new_dashboard5       :=&dashboard.Dashboard{Id:"hardware_dmidecode_dashboard",  Title:"Dmidecode",  Url:"hardware-dmidecode-dashboard", Icon:"fa fa-circle-o"}
    new_dashboard6       :=&dashboard.Dashboard{Id:"users_dashboard",               Title:"Users",      Url:"users-dashboard", Icon:"fa fa-circle-o"}

    d1_id,_    :=d.CreateDashboard(new_dashboard1)
    d2_id,_    :=d.CreateDashboard(new_dashboard2)
    d3_id,_    :=d.CreateDashboard(new_dashboard3)
    d4_id,_    :=d.CreateDashboard(new_dashboard4)
    d5_id,_    :=d.CreateDashboard(new_dashboard5)
    d6_id,_    :=d.CreateDashboard(new_dashboard6)


    d.AttachDashboardToUser(user_root_id, d1_id)
    d.AttachDashboardToUser(user_root_id, d2_id)
    d.AttachDashboardToUser(user_root_id, d3_id)
    d.AttachDashboardToUser(user_root_id, d4_id)
    d.AttachDashboardToUser(user_root_id, d5_id)
    d.AttachDashboardToUser(user_root_id, d6_id)

    d.AttachDashboardToUser(user_riah_id, d5_id)
    d.AttachDashboardToUser(user_riah_id, d6_id)

    d.AttachDashboardToUser(user_kelly_id, d2_id)
    d.AttachDashboardToUser(user_kelly_id, d3_id)


    // d.AttachDashboardToUser("AF35CEFC-1AEA-A399-7448-C2EF4B80E77F","8835CEFC-1AEA-A399-2222-C2EF4B80E77F")
    // user:=&User{Name:"Anna", Password:"SecretPassword123"}
    // d.CreateUser(user)
    //existing_user,err:=d.GetUserById("60F8FEE2-A6B9-45CF-24CA-B2795002C779")
    //fmt.Printf("--\n%v\n%v\n--",existing_user,err)
    //query:=make(map[string]interface{})
    //query["name"] = "Mike"
    //query["secondname"] = "Livshieshch"
    //existing_user,err:=d.GetUser(query)
    //fmt.Printf("==\n%v\n==\n%v\n==bson==\n%v",existing_user,err,bson.M(query))
    //d.RemoveUsersById("a","b","159E2D96-0AFF-3EBC-D01C-C2E3F3AD16A9")
    //token,err:=d.CreateToken("C5952D91-9AA5-4EEB-A21A-F138445103D5")
    //fmt.Printf("token exists %v",d.TokenExists("AF35CEFC-1AEA-A399-7448-C2EF4B80E77F", "8D52B9F2-2E19-427F-4E72-04AF9BF91571"))
    //fmt.Printf("New token: %s Err: %v",token,err)
    dgroup_usermanagement := &dashboard.DashboardGroup{Icon:"fa-child"           ,Title:"User Management",Name:"user_management",List:[]string{d1_id,d6_id}}
    dgroup_networking     := &dashboard.DashboardGroup{Icon:"fa-share-alt-square",Title:"Networking"     ,Name:"networking"     ,List:[]string{d2_id,d3_id}}
    dgroup_hardware       := &dashboard.DashboardGroup{Icon:"fa-tv"              ,Title:"Hardware"       ,Name:"hardware"       ,List:[]string{d4_id,d5_id}}

    dg1_id,_:=d.CreateDashboardGroup(dgroup_usermanagement)
    dg2_id,_:=d.CreateDashboardGroup(dgroup_networking)
    dg3_id,_:=d.CreateDashboardGroup(dgroup_hardware)

    d.AttachDashboardGroupToUser(user_root_id,dg1_id)
    d.AttachDashboardGroupToUser(user_root_id,dg2_id)
    d.AttachDashboardGroupToUser(user_root_id,dg3_id)

    d.AttachDashboardGroupToUser(user_riah_id, dg3_id)
    d.AttachDashboardGroupToUser(user_kelly_id, dg2_id)
}
