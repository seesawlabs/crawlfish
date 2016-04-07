'use strict';

/**
 * @ngdoc filter
 * @name crawlfishApp.filter:secondsToDateTime
 * @function
 * @description
 * # secondsToDateTime
 * Filter in the crawlfishApp.
 */
angular
    .module('crawlfishApp')
    .filter('secondsToDateTime', secondsToDateTime);

function secondsToDateTime() {
    return function(seconds) {
        return new Date(1970, 0, 1).setSeconds(seconds);
    };
}
