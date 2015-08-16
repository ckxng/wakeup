// Copyright (c) 2015 Cameron King. All rights reserved.
// License: BSD 2-clause.
// Website: https://github.com/ckxng/wakeup

var wakeupWebAppControllers = angular.module('wakeupWebAppControllers', []);

wakeupWebAppControllers.filter('stringify', function(){ 
    return function(input) {
        return JSON.stringify(input);
    };
});

var wakeupWebAppControllers_helper_updateTime = function(Page) {
	Page.setTime(new Date())
};

wakeupWebAppControllers.controller('MainCtrl', [
    '$scope', '$interval', 'Page',
    function ($scope, $interval, Page) {
        $scope.Page = Page;
		Page.setTitle("Wakeup App")
		
		$interval(function() {
			Page.setTime(new Date());
		}, 1000);
    }
]);

wakeupWebAppControllers.controller('ClockCtrl', [
    '$scope', 'Page',
    function ($scope, Page) {
        Page.setTitle("Clock");
    }
]);

wakeupWebAppControllers.controller('WakeupRiseCtrl', [
    '$scope', 'Page',
    function ($scope, Page) {
        Page.setTitle("Rise and shine!");
    }
	
	// play music
	
	// start timer then go back to clock
	
]);

wakeupWebAppControllers.controller('WakeupDressCtrl', [
    '$scope', 'Page',
    function ($scope, Page) {
        Page.setTitle("Get dressed!");
    }
	
	// play music
	
	// start timer then go back to clock
	
]);

wakeupWebAppControllers.controller('WakeupBreakfastCtrl', [
    '$scope', 'Page',
    function ($scope, Page) {
        Page.setTitle("Eat breakfast!");
    }
	
	// play music
	
	// start timer then go back to clock
	
]);

wakeupWebAppControllers.controller('WakeupTeethCtrl', [
    '$scope', 'Page',
    function ($scope, Page) {
        Page.setTitle("Brush your teeth!");
    }
	
	// play music
	
	// start timer then go back to clock
	
]);

wakeupWebAppControllers.controller('WakeupBackpackCtrl', [
    '$scope', 'Page',
    function ($scope, Page) {
        Page.setTitle("Put on your backpack!");
    }
	
	// play music
	
	// start timer then go back to clock
	
]);

wakeupWebAppControllers.controller('WakeupLeaveCtrl', [
    '$scope', 'Page',
    function ($scope, Page) {
        Page.setTitle("Time to go!");
    }
	
	// play music
	
	// start timer then go back to clock
	
]);
