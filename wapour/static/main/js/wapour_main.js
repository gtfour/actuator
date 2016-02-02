wapourApp.factory('websocketService', ['$q','$rootScope', function($q, $rootScope) {
    // 2 200 000
    var Service           = {};
    var callbacks         = {};
    var currentCallbackId = 0;
    var ws = new WebSocket("ws://10.10.111.143:8090/entry");
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
    var request            = {};
    request['author']      = "user1";
    request['message']     = "Hello all!";
    var send_test_message  = function(){websocketService.sendRequest(request)};
    var data               = websocketService.wsReady(ws_connect_retries, send_test_message);

    $scope.dashboard_select = function(dashboard_name) {
        alert("Selected dashboard:"+dashboard_name); 
    }
    $scope.dashboards_list = function() {


    }


}]);

