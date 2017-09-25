wapourApp.directive('wapourDataBox', ['websocketService',function (websocketService) {
    /*return {
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
    } */
    var directive = {
        restrict: 'E',
        //scope: { 
        //    boxdata:'='
        //},
        //scope:{},
        //link: function( scope, element, attrs ) {
            //console.log("wapourDataBox data="+scope.boxdata);
            //scope.$watch("boxdata", function(value) {
            //    console.log("wapourDataBox data="+scope.boxdata);
            //}, true);
            //scope.elemtype = attrs.boxdata ;  
            
        //}
        //template: '<div><div {{elemtype}}>'+
        //          '</div></div>',


    };
    return directive ; 

}]);

wapourApp.directive('wapourDashboard', ['websocketService',function (websocketService) {
    var directive = {};
    directive.restrict = 'AE';
    directive.link = function( scope, elements, attrs ) {
        dashboard_group_id = attrs.dashboardGroupId;
        dashboard_id       = attrs.dashboardId ;
    }
    return directive ; 
}]);

wapourApp.directive('wapourTable', ['websocketService',function (websocketService) {
    /*return {
        restrict: 'A',
        link:     function(scope, elem, attrs) {
            elem.bind('click', function() {
            dashboard_group_id = attrs.dashboardGroupId;
            dashboard_id       = attrs.dashboardId ;
            var data               = {};
            var message            = {"datatype":"message_switch_dashboard"};
            var selected_dashboard = {"dashboardgroupid":dashboard_group_id, "dashboardid":dashboard_id};
            mes
sage["data"]        = selected_dashboard
            //alert("DGID:"+dashboard_group_id+" DID:"+dashboard_id);
            websocketService.sendRequest(message)
        });
        }
    }*/
    var directive = {};
    directive.restrict = 'C';
    directive.template = '<my-table></my-table>';
    directive.link = function( scope, elements, attrs ) {
        console.log("my-table");
    }
    return directive;
}]);

wapourApp.directive('wapourChart', ['websocketService',function (websocketService) {
    var directive = {};
    directive.restrict = 'C';
    directive.link = function( scope, elements, attrs ) {
        console.log("my-chart");
    };
    directive.template = '<my-chart></my-chart>';
    return directive;

}]);

wapourApp.directive('wapourTableRow', ['websocketService',function (websocketService) {

    var directive = {}; 
    directive.restrict = 'E';
    directive.template = "<tr></tr>"
    return directive ; 
}]);

wapourApp.directive('wapourInfoBoxArray', ['websocketService',function (websocketService) {
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

/*
        <div class="col-md-3 col-sm-6 col-xs-12">
          <div class="info-box">
            <span class="info-box-more">Likes</span>
            <span class="info-box-number">41,410</span>
            <span class="info-box-more" style="font-size:8px;">4E4BFE80-6FCC-B5E9-1AE6-DA906A88BD3B</span>
            <a href="www.google.com"  class="info-box-more" style="font-size:8px;">More Info</a>
          </div>
        </div>
*/
