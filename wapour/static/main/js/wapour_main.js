wapourApp.controller('mainController', ['$scope', function($scope) {
    var sock = new SockJS("127.0.0.1:8090/entry")
    $scope.sendMessage = function() {
        sock.send($scope.messageText);
        $scope.message = 'Text' ; 
    };
    sock.onmessage = function(e) {
        $scope.messages.push(e.data)
        alert("Message"+e.data);
        $scope.apply();
    };
}]);
