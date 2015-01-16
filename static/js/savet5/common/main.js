/**
 * Common functionality.
 */
define(["angular", "./services/helper", "./filters", "./directives/example"],
    function(angular) {
  "use strict";

  return angular.module("savet4.common", ["common.helper", "common.filters",
    "common.directives.example"]);
});
