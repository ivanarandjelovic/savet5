(function(requirejs) {
	"use strict";

	// -- DEV RequireJS config --
	requirejs.config({
		// Packages = top-level folders; loads a contained file named "main.js"
		packages : [ "common", "home", "user", "dashboard", "saveti", "live",
				"security", "stanari" ]

		,shim: {
	        angular: {
	            exports: "angular"
	        }
	        ,"angular-resource": {
	            deps: [ "angular" ]
	            /*,exports: "ngResource"*/
	        }
	        
	        ,"angular-cookies":     ["angular"]
	        ,"angular-route":       ["angular"]
	        ,"ui-bootstrap":   ["angular"]
	        ,"ui-bootstrap-tpls": ["angular"]
	        ,"bootstrap":      ["jquery"]
	    }

	});

	requirejs.onError = function(err) {
		console.log(err);
	};

	// Load the app. This is kept minimal so it doesn't need much updating.
	require([ "angular", "angular-cookies", "angular-route",
			"angular-resource", "jquery", "bootstrap", "ui-bootstrap",
			"ui-bootstrap-tpls", "./app" ],
			function(angular) {
				angular.bootstrap(document, [ "app" ]);
			});
})(requirejs);
