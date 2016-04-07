wapourApp.factory('settingsService', ['$q','$rootScope', function($q, $rootScope) {
    var Service              = {};
    var Settings             = {};
    Settings["ws_url"]       = "" ;
    Settings["get_data_url"] = "" ;
    function setSettings( settings ) {
        for ( var key in settings ){
            if (key == "ws_url") {
                Settings["ws_url"]       = settings["ws_url"]
            }
            if (key == "get_data_url") {
                Settings["get_data_url"] = settings["get_data_url"]

            }
        }
    }
    function getSettings() {
        return Settings;
    }
    Service.setSettings = setSettings;
    Service.getSettings = getSettings;
    return Service;
}]);

wapourApp.directive('wapourAppSettings', ['settingsService',function (websocketService) {
    var directive = {};
    directive.restrict = 'AE';
    directive.link = function( scope, elements, attrs ) {
        var ws_url               = attrs.wsUrl ;
        var get_data_url         = attrs.getDataUrl ;
        var settings             = {};
        settings["ws_url"]       = ws_url;
        settings["get_data_url"] = get_data_url;
        settingsService.setSettings(settings)
    }
    return directive ; 
}]);

wapourApp.factory('getDataService', ['$q','$rootScope', function($q, $rootScope) {

    var Service = {} ;
    function getDashboardData( dashboard_group_id, dashboard_id, session_id ) {

    }
}]);


wapourApp.factory('websocketService', ['$q','$rootScope','wapourAppSettings', function($q, $rootScope) {
    // 2 200 000
    var Service           = {};
    var SubscribersDashboards = [] ;
    var SubscribersDataboxes  = [] ; 
    var callbacks         = {};
    var currentCallbackId = 0;

    settings = wapourAppSettings.getSettings()

    var ws = new WebSocket(settings["ws_url"]);
    ws.onopen = function (){
        console.log("Socket has been opened!");
        //var test = {}; 
        //test['author']="user";
        //test['body']  ="logged_in";
        //ws.send(JSON.stringify(test));
    };
    ws.onmessage = function(message){
        listener(JSON.parse(message.data));
    };
    function sendRequest(request){
        var defer = $q.defer();
        var callbackId = getCallbackId();
        callbacks[callbackId] = {
            time: new Date(),
            cb: defer
        };
        request.callback_id = callbackId ;
        console.log("Sending request", request);
        ws.send(JSON.stringify(request));
        return defer.promise;
    }
    function listener(data){
        var messageObj = data ;
        console.log("Received data from websocket: ",messageObj);
        if(callbacks.hasOwnProperty(messageObj.callback_id)) {
            console.log(callbacks[messageObj.callback_id]);
            $rootScope.$apply(callbacks[messageObj.callback_id].cb.resolve(messageObj.data));
            delete callbacks[messageObj.callbackID];
        }
    }
    // This creates a new callback ID for a request
    function getCallbackId() {
      currentCallbackId += 1;
      if(currentCallbackId > 10000) {
        currentCallbackId = 0;
      }
      return currentCallbackId;
    }

    function waitForSocketConnection(retries,callback){
        var count = retries ; 
        setTimeout(
            function () {
                if (ws.readyState === 1) {
                    console.log("Connection is made");
                    if(callback != null){
                        data = callback();
                        return JSON.parse(data)
                    }

            } else {
                if (count == 0) { return null }
                count = count - 1;
                console.log("wait for connection...");
                waitForSocketConnection(count,null);
            }

        }, 5); // wait 5 milisecond for the connection...
    }

    // Define a "getter" for getting customer data
    Service.getCustomers = function() {
      var request = {
        type: "get_customers"
      }
      // Storing in a variable for clarity on what sendRequest returns
      var promise = sendRequest(request); 
      return promise;
    }
    Service.sendRequest = function(request) {
      var promise = sendRequest(request);
      return promise;
    }
    Service.wsReady = function(callback) {
      var promise = waitForSocketConnection(callback);
      return promise;
    }


    return Service;
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

