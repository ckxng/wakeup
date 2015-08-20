/*  Copyright (C) 2009 Elijah Rutschman

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details, available at
    <http://www.gnu.org/licenses/>.
/*

/*
a typical cron entry has either wildcards (*) or an integer:

 .---------------- minute (0 - 59) 
 |  .------------- hour (0 - 23)
 |  |  .---------- day of month (1 - 31)
 |  |  |  .------- month (1 - 12)
 |  |  |  |  .---- day of week (0 - 6) (Sunday=0)
 |  |  |  |  |
 *  *  *  *  *

*/

var Cron = {
 "jobs" : [],
 "process" : function() {
  var now = new Date();
  for (var i=0; i<Cron.jobs.length; i++) {
   if ( Cron.jobs[i].minute == "*" || parseInt(Cron.jobs[i].minute) == now.getMinutes() )
    if ( Cron.jobs[i].hour == "*" || parseInt(Cron.jobs[i].hour) == now.getHours() )
     if ( Cron.jobs[i].date == "*" || parseInt(Cron.jobs[i].date) == now.getDate() )
      if ( Cron.jobs[i].month == "*" || (parseInt(Cron.jobs[i].month) - 1) == now.getMonth() )
       if ( Cron.jobs[i].day == "*" || parseInt(Cron.jobs[i].day) == now.getDay() )
        Cron.jobs[i].run();
  }
  now = null;
 },
 "id" : 0,
 "start" : function() {
  Cron.stop();
  Cron.id = setInterval("Cron.process()",60000);
 },
 "stop" : function() {
  clearInterval(Cron.id);
 },
 "Job" : function(cronstring, fun) {
  var _Job = this;
  var items = cronstring.match(/^([0-9]+|\*{1})[ \n\t\b]+([0-9]+|\*{1})[ \n\t\b]+([0-9]+|\*{1})[ \n\t\b]+([0-9]+|\*{1})[ \n\t\b]+([0-9]+|\*{1})[ \n\t\b]*$/);
  _Job.minute = items[1];
  _Job.hour = items[2];
  _Job.date = items[3];
  _Job.month = items[4];
  _Job.day = items[5];
  _Job.run = fun;
  Cron.jobs.push(_Job);
  _Job = null;
  items = null;
 }
}