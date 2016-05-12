wapourApp.service('settingsService', ['$q','$rootScope', function($q, $rootScope) {
    var Service              = {};
    var Settings             = {};
    Settings["ws_url"]       = "" ;
    Settings["get_data_url"] = "" ;
    Settings["session_id"]   = "" ;
    Settings["app_data_url"] = "" ;
    Settings["websocket"]    = false ;
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
            if (key == "app_data_url") {
                Settings["app_data_url"] = settings["app_data_url"];
            }
            if (key == "websocket") {
                if ( settings["websocket"] == 'true' ) {
                    Settings["websocket"] = true;
                } 
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

wapourApp.factory('dashboardDataService',['websocketService','settingsService','$http', function(websocketService, settingsService, $http) {
    var service    = {};
    var dashboards = [];
    var callbacks  = [];
    var callbacks_url = [];
    function SelectDashboardById(dashboard_id, dashboard_group_id) {
        settings               = settingsService.getSettings();
        if ( settings.websocket == true ) {
            var session_id         = settings["session_id"];
            var data               = {};
            var message            = {"datatype":"message_switch_dashboard"};
            var selected_dashboard = {"dashboardgroupid":dashboard_group_id, "dashboardid":dashboard_id};

            message["data"]        = selected_dashboard ;
            message["session_id"]  = session_id ;
            websocketService.sendRequest(message)
        }
        Notify(dashboard_id, dashboard_group_id);
    }

    function SelectDashboardByUrl(dashboard_url) {
        settings               = settingsService.getSettings();
        session_id             = settings["session_id"];
        var data               = {};
        var message            = {"datatype":"message_switch_dashboard"};
        var selected_dashboard = {"dashboardgroupid":dashboard_group_id, "dashboardid":dashboard_id};
        message["data"]        = selected_dashboard ;
        message["session_id"]  = session_id ;
        websocketService.sendRequest(message)
        NotifyByUrl(dashboard_url);
    }

    function GetHttp(url){
        return $http({method:"GET", url:url}).then(function(result){
            return result.data ;
        });
    }

    function GetDashboardData(dashboard_id, dashboard_group_id) {
        var settings     = settingsService.getSettings();
        var get_data_url = settings["get_data_url"];
        var dashboard_url = get_data_url+dashboard_group_id+"/"+dashboard_id
        var data_promise = GetHttp(dashboard_url);
        data_promise.then(function(result) {  
           console.log(result);
           return result;
        });
    }
    function GetDashboardDataByUrl(dashboard_url) {
        var settings     = settingsService.getSettings();
        var get_data_url = settings["get_data_url"];
        var full_dashboard_url = get_data_url+dashboard_url
        var data_promise = GetHttp(full_dashboard_url);
        data_promise.then(function(result) {
           console.log(result);
           return result;
        });
    }
    function AddCallback(callback) {
        callbacks.push(callback);
    }
    function AddUrlCallback(callback) {
        callbacks_url.push(callback);
    }
    function Notify(dashboard_id, dashboard_group_id){
        for (var key in callbacks){
            var callback = callbacks[key];
            callback(dashboard_id, dashboard_group_id);
        }
    }
    function NotifyByUrl(dashboard_url){
        for (var key in callbacks_url){
            var callback = callbacks_url[key];
            callback(dashboard_url);
        }
    }

    service.SelectDashboardById   = SelectDashboardById ; 
    service.GetDashboardData      = GetDashboardData ;  
    service.GetDashboardDataByUrl = GetDashboardDataByUrl; 
    service.AddCallback           = AddCallback;
    service.AddUrlCallback        = AddUrlCallback;
    service.GetHttp               = GetHttp ;
    
    return service ; 
}]);
