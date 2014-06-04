/**
 * User service, exposes user model to the rest of the app.
 */
define([ "angular", "./controllers", "ui-bootstrap", "ui-bootstrap-tpls", "common", "security" ], function(angular, controllers) {
	"use strict";

	var mod = angular.module("user.services",
			[ "savet4.common", "ui.bootstrap", "ui.bootstrap.tpls", "security.retryQueue" ]);
	mod.factory("userService", [
			"$http",
			"$q",
			function($http, $q) {
				var user, token;
				return {
					loginUser : function(credentials) {
						return $http.post("/login",
								credentials).then(
								function(response) {
									// return promise so we can chain easily
									token = response.data.token;
									// in a real app we could use the token to
									// fetch the
									// user data
									return $http.get("/currentUser",{"token": token}).then(
													function(response) {
														user = response.data; // Extract
														// user
														// data
														// from
														// user() request
														// user.email =
														// credentials.email;
														return user;
													});
								}, function(response) {
									// Login failure case!
									return undefined;
								});
					},
					logout : function() {
						// Logout on server in a real app
						return $http.post("/logout").then(function(response) {
									user = undefined;
								});

					},
					getUser : function() {
						return user;
					},
					refreshUser : function() {
						var u = $http.get("/currentUser", {doNotRetry : true} ).then(function(response) {
									user = response.data; // Extract user
									return user;
									// data from
								});
					}
				};
			} ]);
	/**
	 * Add this object to a route definition to only allow resolving the route
	 * if the user is logged in. This also adds the contents of the objects as a
	 * dependency of the controller.
	 */
	mod.constant("userResolve", {
		user : [ "$q", "userService", function($q, userService) {
			var deferred = $q.defer();
			var user = userService.getUser();
			if (user) {
				deferred.resolve(user);
			} else {
				deferred.reject();
			}
			return deferred.promise;
		} ]
	});

	mod.factory("securityService", [ "$http", "$q", "$modal", "$log", "securityRetryQueue", 
			function($http, $q, $modal, $log, queue	) {
				
				var modalLogin = null;
				
				// Register a handler for when an item is added to the retry queue
				  queue.onItemAddedCallbacks.push(function(retryItem) {
				    if ( queue.hasMore() ) {
				      service.showLogin();
				    }
				  });
				  
				var service = {
					
					showLogin : function() {

							    modalLogin = $modal.open({
								templateUrl : '/html/templates/user/modalLogin.html',
								controller : controllers.ModalLoginCtrl,
								size : "sm"
							});

							modalLogin.result.then(function(user) {
								$log.info('Login done for : ' + user);
								queue.retryAll();
							}, function() {
								$log.info('Login canceled: ' + new Date());
								queue.cancelAll();
							});
						}
				
					  
				}
				
				return service;
			} ]);

	/**
	 * If the current route does not resolve, go back to the start page.
	 */
	var handleRouteError = function($rootScope, $location) {
		$rootScope.$on("$routeChangeError", function(e, next, current) {
			$location.path("/");
		});
	};
	handleRouteError.$inject = [ "$rootScope", "$location" ];
	mod.run(handleRouteError);
	return mod;
});
