/**
 * The app module, as both AngularJS as well as RequireJS module.
 * Splitting an app in several Angular modules serves no real purpose in Angular 1.0/1.1.
 * (Hopefully this will change in the near future.)
 * Splitting it into several RequireJS modules allows async loading. We cannot take full advantage
 * of RequireJS and lazy-load stuff because the angular modules have their own dependency system.
 */
define(["angular", "ui-bootstrap", "ui-bootstrap-tpls", "home", "user", "dashboard", "saveti", "live", "security", "stanari"], function(angular) {
  "use strict";

  // We must already declare most dependencies here (except for common), or the submodules' routes
  // will not be resolved
  return angular.module("app", ["ngResource", "savet4.home", "savet4.user", "savet4.dashboard","savet4.saveti", "savet4.live", "savet4.security", "savet4.stanari"]);
});
