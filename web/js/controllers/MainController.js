/**
 * Created by sionix on 16/05/2016.
 */

app.controller('MainController', ['$scope', '$http', function ($scope, $http) {
    $scope.title = 'Clippings';

    $http.get('http://localhost:8080/api/clippings').success(function (data) {
        $scope.clippings = data;
    });
}]);
