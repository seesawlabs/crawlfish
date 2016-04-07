'use strict';

/**
 * @ngdoc service
 * @name crawlfishApp.crawlData
 * @description
 * # crawlData
 * Factory in the crawlfishApp.
 */
angular.module('crawlfishApp')
    .factory('crawlData', crawlData);

crawlData.$inject = ['Ref', '$firebaseArray'];


function crawlData(Ref, $firebaseArray) {
    var crawlArray = $firebaseArray(Ref);

    var loadedCrawlArray = crawlArray.$loaded()
        .then(function (crawls) {
            return crawls;
        })
        .catch(function (error) {
            console.log('Error:', error);
        });

    return {
        loadedCrawlArray: loadedCrawlArray
    };
}