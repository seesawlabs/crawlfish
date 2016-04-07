'use strict';

/**
 * @ngdoc function
 * @name crawlfishApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the crawlfishApp
 */
angular.module('crawlfishApp')
    .controller('MainCtrl', function ($scope, $timeout, $log, crawlHistory, $mdConstant, $http) {

        var vm = this;

        // Use common key codes found in $mdConstant.KEY_CODE...
        vm.keys = [$mdConstant.KEY_CODE.ENTER];
        vm.tags = [];
        vm.url = '';
        vm.customKeys = [$mdConstant.KEY_CODE.ENTER];
        vm.toggleCard = toggleCard;
        vm.startCrawl = startCrawl;
        vm.searchTerm = '';
    
        var BASE_URL = 'https://crawlfish-dev.herokuapp.com';

        function toggleCard(crawl) {
            crawl.clicked = !crawl.clicked;
            crawl.found = crawl.words_found.length;
            crawl.max = crawl.words.length;
        }
    
        function startCrawl() {
            var tags = '';
            for (var tag in vm.tags) {
                
                tags += vm.tags[tag];
            
                if (parseInt(tag) !== vm.tags.length - 1) {
                    tags += ' \n ';
                }
                
            }
            var payload = JSON.stringify({
                url: vm.url,
                words: tags
            });
            console.log(payload);


            post('/api/v1/crawl', payload)
                .then(postComplete)
                .catch(postError);

        }

        function postComplete(resp) {
            console.log('success', resp);

            // Clear Crawl data
            vm.tags = [];
            vm.url = ''
        }

        function postError(err) {
            console.log('error', err);
        }

        function post(endpoint, data) {
            return $http.post(BASE_URL + endpoint, data);
//            return $http({
//                url: BASE_URL + endpoint,
//                method: 'POST',
//                headers: {
//                    'Content-Type': 'application/json; charset=utf-8'
//                },
//                data: data
//            });
        }


        $scope.crawlHistory = crawlHistory[0];
        console.log($scope.crawlHistory);

    });
