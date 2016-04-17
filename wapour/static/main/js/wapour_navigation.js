wapourApp.directive('navItemLeft', ['websocketService','settingsService',function (websocketService, settingsService) {
    return {
        restrict: 'AE',
        link:     function(scope, elem, attrs) {
            elem.bind('click', function() {
            dashboard_group_id = attrs.dashboardGroupId;
            dashboard_id       = attrs.dashboardId ; 
            settings           = settingsService.getSettings();
            session_id         = settings["session_id"];
            var data               = {};
            var message            = {"datatype":"message_switch_dashboard"};
            var selected_dashboard = {"dashboardgroupid":dashboard_group_id, "dashboardid":dashboard_id};
            message["data"]        = selected_dashboard ; 
            message["session_id"]  = session_id ; 
            //alert("DGID:"+dashboard_group_id+" DID:"+dashboard_id);
            websocketService.sendRequest(message)
        });
        }
    }
}]);

