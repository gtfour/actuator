var wapourApp = angular.module('wapourApp', ['ui.router']);
wapourApp.config(function($stateProvider, $urlRouterProvider) {
    $urlRouterProvider.otherwise('/index');
    $stateProvider
        .state('index', {
            url:'/index',
            templateUrl: 'index'
        })
        .state('files', {
        })
        .state('actions', {
        });


});
