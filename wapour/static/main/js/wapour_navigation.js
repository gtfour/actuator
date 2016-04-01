wapourApp.directive('navItemLeft', ['websocketService',function (websocketService) {
    return {
        restrict: 'A',
        link:     function(scope, elem, attrs) {
            elem.bind('click', function() {
            dashboard_group_id = attrs.dashboardGroupId;
            dashboard_id       = attrs.dashboardId ; 
            var data               = {};
            var message            = {"datatype":"message_switch_dashboard"};
            var selected_dashboard = {"dashboardgroupid":dashboard_group_id, "dashboardid":dashboard_id};
            message["data"]        = selected_dashboard
            //alert("DGID:"+dashboard_group_id+" DID:"+dashboard_id);
            websocketService.sendRequest(message)
        });
        }
    }
}]);

