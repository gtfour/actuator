var wapourApp = angular.module('wapourApp', ['ui.router']);
wapourApp.config(function($stateProvider, $urlRouterProvider ) {
    $urlRouterProvider.otherwise('/index');
    $stateProvider
        .state('index', {
            url:'',
            templateUrl: 'dashboard/overview',
            controller: function($scope, $templateCache ) {
                $templateCache.remove('dashboard/overview')
            }
        })
        .state('actions', {
             url:'/dashboard/actions',
             templateUrl: 'dashboard/actions',
             controller: function($scope, $templateCache ) {
                 $templateCache.remove('dashboard/actions')
             }
            
        })
        .state('files', {
            url:'/dashboard/files',
            templateUrl: 'dashboard/files',
            controller: function($scope, $templateCache ) {
                 $templateCache.remove('dashboard/files')
            }
        })
        .state('hosts', {
             url:'/dashboard/hosts',
             templateUrl: 'dashboard/hosts',
             controller: function($scope, $templateCache ) {
                 $templateCache.remove('dashboard/hosts')
             }

        })
        .state('websocket', {
             url:'/wspage',
             templateUrl: 'wspage',
             controller: function($scope, $templateCache ) {
                 $templateCache.remove('wspage')
             }

        });


});
