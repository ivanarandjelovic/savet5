/**
 * Dashboard routes.
 */
define(["angular", "./controllers", "common"], function(angular, controllers) {
  "use strict";

  var mod = angular.module("dashboard.routes", ["savet4.common"]);
  mod.config(["$routeProvider", "userResolve", function($routeProvider, userResolve) {
    $routeProvider
      .when("/dashboard",  {templateUrl: "/html/templates/dashboard/dashboard.html",  controller:controllers.DashboardCtrl, resolve:userResolve});
      //.when("/admin/dashboard",  {templateUrl: "/html/templates/dashboard/admin.html",  controller:controllers.AdminDashboardCtrl})
  }]);
  return mod;
});
