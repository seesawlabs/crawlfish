'use strict';

describe('Filter: secondsToDateTime', function () {

  // load the filter's module
  beforeEach(module('crawlfishApp'));

  // initialize a new instance of the filter before each test
  var secondsToDateTime;
  beforeEach(inject(function ($filter) {
    secondsToDateTime = $filter('secondsToDateTime');
  }));

  it('should return the input prefixed with "secondsToDateTime filter:"', function () {
    var text = 'angularjs';
    expect(secondsToDateTime(text)).toBe('secondsToDateTime filter: ' + text);
  });

});
