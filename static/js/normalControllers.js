tnvdapp.controller('menubarCntl', function($scope, $http) {
	
	var ret = $http({
		method: 'GET',
		url: '/api/mainMenu'
	})
	ret.success(function(response, status, headers, config) {
		$scope.menuBar = response;
	})
});

tnvdapp.controller('mainPage', function($scope, $http) {
	
})

