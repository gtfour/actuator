wapourApp.controller('appController',['settingsService','websocketService','$scope', function initController(settingsService,websocketService, $scope) {
    $scope.initApp = function(settings) {
        settingsService.setSettings(settings);
        websocketService.createWsConnection(settings["ws_url"]);
    };
    $scope.$on("$destroy", function(){
        websocketService.closeWsConnection();
        console.log("Exit from initController . Closing ws-connetion");
    });
    /*$scope.$on('$routeChangeStart', function(){
        websocketService.closeWsConnection();
        alert("Catching routeChangeStart");
        console.log("Exit from initController  via routeChangeStart . Closing ws-connetion");
    });*/
    
}]);

wapourApp.controller('mainController', ['$scope','websocketService', function($scope, websocketService) {
    var ws_connect_retries = 10 ; 
    var data            = {};
    var message         = {"datatype":"message_chat"};
    
    data['author']      = "user1";
    data['message']     = "Hello all!";
    message["data"]     = data
    
    var send_test_message  = function(){websocketService.sendRequest(message)};
    var data               = websocketService.wsReady(ws_connect_retries, send_test_message);

    $scope.dashboard_select = function(dashboard_name) {
        alert("Selected dashboard:"+dashboard_name); 
    }
    $scope.dashboards_list = function() {

    }
}]);

wapourApp.controller('dashboardController', ['$scope','dashboardDataService', function($scope, dashboardDataService) {

    $scope.dashboard_data = {} ;

    var notifier = function(dashboard_id, dashboard_group_id) {
        $scope.dashboard_data = dashboardDataService.GetDashboardData(dashboard_id, dashboard_group_id); 
        console.log("-- dashboard_data --")
        console.log($scope.dashboard_data);
    }
    dashboardDataService.AddCallback(notifier);



}]);

