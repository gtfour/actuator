wapourApp.directive('navigationItem', ['dashboardDataService',function (dashboardDataService) {
    return {
        restrict: 'AE',
        link:     function(scope, elem, attrs) {
            elem.bind('click', function() {
                dashboard_group_id = attrs.dashboardGroupId;
                dashboard_id       = attrs.dashboardId ;
                dashboardDataService.SelectDashboardById(dashboard_id, dashboard_group_id);
        });
        }
    }
}]);


wapourApp.directive('navigationSubitem', ['dashboardDataService',function (dashboardDataService) {
    return {
        restrict: 'AE',
        link:     function(scope, elem, attrs) {
            elem.bind('click', function() {
                dashboard_group_id = attrs.dashboardGroupId;
                dashboard_id       = attrs.dashboardId ; 
                dashboardDataService.SelectDashboardById(dashboard_id, dashboard_group_id);
        });
        }
    }
}]);

wapourApp.directive('subitemMenu', function () {
    return {
        restrict: 'AE',
        scope: {
            subitems:'=',
            parentid:'='

        },
        link:     function(scope, elem, attrs) {
        },
        template: '<li navigation-subitem ng-repeat="subitem in subitems"   dashboard-group-id="parentid" dashboard-id="{{subitem.id}}"><a ui-sref="subitem.url"><i class="fa fa-circle-o"></i>{{subitem.name}}</a></li>'
    }
});

wapourApp.directive('navigationMenu',function () {
    return {
        restrict: 'AE',
        scope: {
            navigation_items:'=navigationItems'
        },
        link:     function(scope, elem, attrs) {
        },
        template: '<li class="header">MAIN NAVIGATION</li>'+
                  '<li class="treeview" ng-repeat="item in navigation_items">'+
                  '<a href="#"><i class="fa {{ item.icon }}"></i> <span>{{ item.name }}</span> <i class="fa fa-angle-left pull-right"></i>'+
                  '</a>'+
                  '<ul class="treeview-menu" subitem-menu subitems=item.subitems parentid="{{ item.id }}"></ul>'+
                  '</li>'
    }
});

