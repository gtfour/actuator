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

