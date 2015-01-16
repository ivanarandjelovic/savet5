/**
 * Home routes.
 */
define(["angular", "./controllers", "common"], function(angular, controllers) {
  "use strict";

  var mod = angular.module("home.routes", ["savet4.common"]);
  mod.config(["$routeProvider", function($routeProvider) {
    $routeProvider
      .when("/",  {templateUrl: "/html/templates/home/home.html", controller:controllers.HomeCtrl})
      .otherwise( {templateUrl: "/html/templates/home/notFound.html"});
  }]);
  return mod;
});
