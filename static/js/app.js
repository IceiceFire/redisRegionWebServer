// app.js
// create angular app
var validationApp = angular.module('validationApp', [ 'ui.bootstrap' ]);
 
// create angular controller
validationApp.controller('mainController', function($scope, $modal, $http, $log) {

    $scope.expire = '-1';
    $scope.starting = '0';

    // function to submit the form after all validation has occurred            
    $scope.submitForm = function(isValid) {

        // list
        $scope.items = [ '         Region：' + $scope.region, 
                         '  Redis Address：' + $scope.redisip, 
                         '     Key Prefix：' + $scope.keyprefix, 
                         '    Expire Time：' + $scope.expire, 
                         'Starting Status：' + $scope.starting ];

        var modalInstance = $modal.open({

            templateUrl : 'myModelContent.html',
            controller : 'ModalInstanceCtrl', // specify controller for modal
            size : '',
            resolve : {
                items : function() {
                    return $scope.items;
                },
                isValid : function() {
                    return isValid
                }
            }
        });
        // modal return result
        modalInstance.result.then(function(selectedItem) {

            if (selectedItem == '1') {

                $http.post('/region/', {Region: $scope.region, Redisip: $scope.redisip, Keyprefix: $scope.keyprefix, Expire: $scope.expire, Starting: $scope.starting}).
                  success(function() {

                        // submitted stat
                        $scope.submitted = false;
                        // Entry force information init
                        $scope.region = '';
                        $scope.redisip = '';
                        $scope.keyprefix = '';
                        $scope.expire = '-1';
                        $scope.starting = '0';
                    })
            }
        }, function() {
            $log.info('Modal dismissed at: ' + new Date())
        });
        // submitted stat
        $scope.submitted = true;
    };
 
})

// modal controller
.controller('ModalInstanceCtrl', function($scope, $modalInstance, isValid, items) {
    
    $scope.items = items;

    // ok click
    $scope.ok = function() {
        // check to make sure the form is completely valid
        if (isValid) {
            $modalInstance.close('1');
        }
    };
    // cancel click
    $scope.cancel = function() {
        $modalInstance.dismiss('cancel');
    }
});