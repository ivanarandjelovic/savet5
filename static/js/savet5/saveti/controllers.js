// Generated by CoffeeScript 1.6.3
/*
 * Saveti controllers.
*/


(function() {
  define(["angular"], function() {
    "use strict";
    var SavetiCtrl;
    SavetiCtrl = function($scope, $location, $routeParams, $resource) {
      var Savet, savetId;
      Savet = $resource('/savet/:id', {
        id: '@id'
      });
      savetId = $routeParams.id;
      if (savetId) {
        $scope.savet = Savet.get({
          id: savetId
        });
      } else {
        $scope.saveti = Savet.query();
      }
      $scope.open = function(id) {
        return $location.path("/stanari/" + id);
      };
      $scope.showCreate = function() {
        return $location.path("/saveti/create");
      };
      $scope.create = function(savet) {
        var savetNew;
        savetNew = new Savet(savet);
        return savetNew.$save(null, function(savedSavet) {
          $scope.saveti = Savet.query();
          return $location.path("/saveti");
        });
      };
    };
    SavetiCtrl.$inject = ["$scope", "$location", "$routeParams", "$resource"];
    return {
      SavetiCtrl: SavetiCtrl
    };
  });

}).call(this);