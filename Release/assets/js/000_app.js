// Copyright (c) 2015 Cameron King. All rights reserved.
// License: BSD 2-clause.
// Website: https://github.com/ckxng/wakeup

var wakeupWebApp = angular.module('wakeupWebApp', [
	'ngRoute',
	'wakeupWebAppControllers',
	'pageFactory'
]);

wakeupWebApp.config(
    ['$routeProvider', '$locationProvider',
    function($routeProvider, $locationProvider) {
        $locationProvider.hashPrefix('!');
        $routeProvider.
            when('/clock', {
                templateUrl: 'parts/clock.html',
                controller: 'ClockCtrl'
            }).
            when('/wakeup/rise', {
                templateUrl: 'parts/wakeup-rise.html',
                controller: 'WakeupRiseCtrl'
            }).
            when('/wakeup/dress', {
                templateUrl: 'parts/wakeup-dress.html',
                controller: 'WakeupDressCtrl'
            }).
            when('/wakeup/breakfast', {
                templateUrl: 'parts/wakeup-breakfast.html',
                controller: 'WakeupBreakfastCtrl'
            }).
            when('/wakeup/teeth', {
                templateUrl: 'parts/wakeup-teeth.html',
                controller: 'WakeupTeethCtrl'
            }).
            when('/wakeup/backpack', {
                templateUrl: 'parts/wakeup-backpack.html',
                controller: 'WakeupBackpackCtrl'
            }).
            when('/wakeup/leave', {
                templateUrl: 'parts/wakeup-leave.html',
                controller: 'WakeupLeaveCtrl'
            }).
            otherwise({
                redirectTo: '/clock'
            });
    }
]);
