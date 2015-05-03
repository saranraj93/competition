var app = angular.module('competition', ['ui.bootstrap']);

app.controller("MainCtrl", function($scope){
    $scope.name = "World";
});

app.controller('DropdownCtrl', function($scope){
    $scope.items = [
        "item",
        "item2",
        "item3"
    ];
})
