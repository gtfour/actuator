wapourApp.service('settingsService', ['$q','$rootScope', function($q, $rootScope) {
    var Service              = {};
    var Settings             = {};
    Settings["ws_url"]       = "" ;
    Settings["get_data_url"] = "" ;
    Settings["session_id"]   = "" ;
    function setSettings( settings ) {
        for ( var key in settings ){
            if (key == "ws_url") {
                Settings["ws_url"]       = settings["ws_url"];
            }
            if (key == "get_data_url") {
                Settings["get_data_url"] = settings["get_data_url"];

            }
            if (key == "session_id") {
                Settings["session_id"] = settings["session_id"];

            }
        }
        console.log("settingsService:"+JSON.stringify(Settings))
    }
    function getSettings() {
        return Settings;
    }
    Service.setSettings = setSettings;
    Service.getSettings = getSettings;
    return Service;
}]);

wapourApp.factory('dashboardDataService',['websocketService','settingsService', function(websocketService, settingsService, $http) {
    var service    = {};
    var dashboards = [];
    function SelectDashboard(dashboard_id, dashboard_group_id) {
        settings               = settingsService.getSettings();
        session_id             = settings["session_id"];
        var data               = {};
        var message            = {"datatype":"message_switch_dashboard"};
        var selected_dashboard = {"dashboardgroupid":dashboard_group_id, "dashboardid":dashboard_id};
        message["data"]        = selected_dashboard ;
        message["session_id"]  = session_id ;
        websocketService.sendRequest(message)
    }

    function getHttp(url){
        $http({method:"GET", url:url}).then(function(result){
            return result.data ;
        });
    }

    function GetDashboardData(dashboard_id, dashboard_group_id) {
        var settings     = settingsService.getSettings();
        var get_data_url = settings["get_data_url"];
    }

    service.SelectDashboard = SelectDashboard ; 
    return service ; 
}]);
