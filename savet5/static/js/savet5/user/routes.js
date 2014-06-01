/**
 * Configure routes of user module.
 */
define(["angular", "./controllers", "common"], function(angular, controllers) {
  var mod = angular.module("user.routes", ["user.services", "savet4.common"]);
  mod.config(["$routeProvider", function($routeProvider) {
    $routeProvider
      .when("/login", {templateUrl:"/html/templates/user/login.html", controller:controllers.LoginCtrl});
      //.when("/users", {templateUrl:"/html/templates/user/users.html", controller:controllers.UserCtrl})
      //.when("/users/:id", {templateUrl:"/html/templates/user/editUser.html", controller:controllers.UserCtrl});
  }]);
  return mod;
});
