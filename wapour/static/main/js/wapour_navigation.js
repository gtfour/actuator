wapourApp.directive('navItemLeft', [function () {
    return {
        restrict: 'A',
        link:     function(scope, elem, attrs) {
            elem.bind('click', function() {
            dashboard_group_id = attrs.dashboardGroupId;
            dashboard_id       = attrs.dashboardId ; 
            alert("DGID:"+dashboard_group_id+" DID:"+dashboard_id);
        });
        }
    }
}]);

