tnvdapp.controller('menubarCntl', function($scope, $http) {
	
	var ret = $http({
		method: 'GET',
		url: '/admin/api/mainMenu'
	})
	ret.success(function(response, status, headers, config) {
		$scope.menuBar = response;
	})
});

tnvdapp.controller('serverStates', function($scope, $http) {
	var ret = $http({
		method: "GET",
		url: '/admin/api/sysStatus'
	})
	ret.success(function(response, status, headers, config) {
		$scope.sysStatus = response;
	})
})

tnvdapp.controller('dashboard', function($scope, $http){
	
})

