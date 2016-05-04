wapourApp.directive('navigationItem', ['dashboardDataService',function (dashboardDataService) {
    return {
        restrict: 'AE',
        link:     function(scope, elem, attrs) {
            elem.bind('click', function() {
                dashboard_group_id = attrs.dashboardGroupId;
                dashboard_id       = attrs.dashboardId ; 
                dashboardDataService.SelectDashboard(dashboard_id, dashboard_group_id);
        });
        }
    }
}]);

wapourApp.directive('navigationMenu', ['dashboardDataService',function (dashboardDataService) {
    return {
        restrict: 'AE',
        link:     function(scope, elem, attrs) {
            elem.bind('click', function() {
                dashboard_group_id = attrs.dashboardGroupId;
                dashboard_id       = attrs.dashboardId ;
                dashboardDataService.SelectDashboard(dashboard_id, dashboard_group_id);
        });
        }
    }
}]);

