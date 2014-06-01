// Generated by CoffeeScript 1.6.3
/*
 * Stanari controllers.
*/


(function() {
  define(["angular"], function() {
    "use strict";
    var StanariCtrl;
    StanariCtrl = function($scope, $location, $routeParams, $resource) {
      var Savet, Stanari, savetId;
      Savet = $resource('/savet/:id', {
        id: '@id'
      });
      Stanari = $resource("/stanari/:savetId", {
        savetId: "@savetId"
      });
      savetId = $routeParams.savetId;
      $scope.savet = Savet.get({
        id: savetId
      });
      return $scope.stanari = Stanari.query({
        savetId: $routeParams.savetId
      });
    };
    StanariCtrl.$inject = ["$scope", "$location", "$routeParams", "$resource"];
    return {
      StanariCtrl: StanariCtrl
    };
  });

}).call(this);