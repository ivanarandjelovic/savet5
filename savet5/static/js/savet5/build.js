/*
 * A custom build profile that is passed to the optimizer via requireJsShim in build.sbt.
 * Play does this via settings it as the mainConfigFile:
 * http://requirejs.org/docs/optimization.html#mainConfigFile
 */
requirejs.config({
  packages: ["common", "home", "user", "dashboard", "saveti", "live", "security", "stanari"],
  paths: {
    // Make the optimizer ignore CDN assets
    "_" : "empty:",
    "jquery": "empty:",
    "bootstrap": "empty:",
    "angular": "empty:",
    "angular-cookies": "empty:",
    "angular-route": "empty:",
    "angular-resource": "empty",
    "ui-bootstrap": "empty"

  }
});
