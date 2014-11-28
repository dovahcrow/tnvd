tnvdapp.config(function($routeProvider) {
  $routeProvider.
      when('/serverStates', {templateUrl: '/static/tpl/admin/serverStates.html', controller: "serverStates"}).
      when('/dashboard', {templateUrl: '/static/tpl/admin/dashboard.html', controller: "dashboard"}).
      otherwise({redirectTo: '/dashboard'});

});
