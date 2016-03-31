var testApp = angular.module('testApp',[]);

testApp.factory('dataService', function(){
    return {messages:["Hello all!","Buy!","Nice to see you!"],
            data:function get_table_data(table_id){} ,};
});

testApp.controller('ResponseController', function ResponseController($scope, dataService) {
    $scope.answers = dataService.messages ; 
    $scope.word = "Habrahabr";
}); 

testApp.directive('habraHabr', function () {

    return function(scope,element,attrs) {
        scope.$watch(attrs.habraHabr, function(value) {
            element.text(value+attrs.habr);
        });
        scope.$watch(attrs.habraHabr, function(value) {
            //element.text(value+attrs.habr);
            alert(value);
        });
        }
});

testApp.directive('helloWorld', function () {
    return {
        restrict: 'AE',
        replace:  'true',
        template: '<h3>Hello world!</h3>'
    };
});




