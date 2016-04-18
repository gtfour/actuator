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

/*wapourApp.directive('wapourAppSettings', ['settingsService',function (settingsService) {
    var directive = {};
    directive.restrict = 'AE';
    directive.link = function( scope, elements, attrs ) {
        var ws_url               = attrs.wsUrl ;
        var get_data_url         = attrs.getDataUrl ;
        var settings             = {};
        settings["ws_url"]       = ws_url;
        settings["get_data_url"] = get_data_url;
        console.log("wapourAppSettings:"+JSON.stringify(settings));
        settingsService.setSettings(settings)
    }
    return directive ; 
}]);*/

wapourApp.factory('getDataService', ['$q','$rootScope','settingsService', function(settingsService, $q, $rootScope) {

    var Service = {};
    function getDashboardData( dashboard_group_id, dashboard_id, session_id ){

    }
}]);


wapourApp.service('websocketService', ['$q','$rootScope','settingsService', function($q, $rootScope, settingsService) {
    // 2 200 000
    var Service           = {};
    var SubscribersDashboards = [] ;
    var SubscribersDataboxes  = [] ; 
    var callbacks         = {};
    var currentCallbackId = 0;
    console.log("<<websocketService>>");


    //var settings = settingsService.getSettings();
    //console.log("websocketService settings "+ JSON.stringify(settings));

    var ws = (function () { return; })(); // setting ws variable as undefined

    function createWsConnection(ws_url) {
        ws = new WebSocket(ws_url);
        ws.onopen = function (){
            var settings    = settingsService.getSettings();
            var session_id  = settings["session_id"];
            var state       = {"state":"open"}; 
            var message  = { "datatype":"message_ws_state","session_id":session_id, "data":state };
            ws.send(JSON.stringify(message));
            console.log("Socket has been opened!");
        };
        /*ws.onclose = function (){
            var settings    = settingsService.getSettings();
            var session_id  = settings["session_id"];
            var state       = {"state":"close"};
            var message  = { "datatype":"message_ws_state","session_id":session_id, "data":state };
            ws.send(JSON.stringify(message));
            console.log("Socket has been closed!");
        };*/

        ws.onmessage = function(message){
            listener(JSON.parse(message.data));
        };
    };
    function closeWsConnection() {
        ws.onclose = function () {}; // disable onclose handler first
        var settings    = settingsService.getSettings();
        var session_id  = settings["session_id"];
        var state       = {"state":"close"};
        var message  = { "datatype":"message_ws_state","session_id":session_id, "data":state };
        ws.send(JSON.stringify(message));
        ws.close();
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
    Service.createWsConnection = createWsConnection ; 
    Service.closeWsConnection  = closeWsConnection  ; 


    return Service;
}]);

wapourApp.controller('initController',['settingsService','websocketService','$scope', function initController(settingsService,websocketService, $scope) {
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

