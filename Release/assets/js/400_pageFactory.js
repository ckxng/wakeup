// Copyright (c) 2015 Cameron King. All rights reserved.
// License: BSD 2-clause.
// Website: https://github.com/ckxng/wakeup

var pageFactory = angular.module('pageFactory', []);

pageFactory.factory('Page', function() {
  var title = '';
  var hh = '00';
  var mm = '00';
  var ss = '00';
  var tt = 'XM';
  
  return {
    title: function() { return title; },
    setTitle: function(x) { title = x },
	hh: function() { return hh; },
	mm: function() { return mm; },
	ss: function() { return ss; },
	tt: function() { return tt; },
	setTime: function(x) {
	  hh = x.format("hh");
	  mm = x.format("MM");
	  ss = x.format("ss");
	  tt = x.format("TT");
	},
  };
});