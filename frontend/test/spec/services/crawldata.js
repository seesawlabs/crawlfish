'use strict';

describe('Service: crawlData', function () {

  // load the service's module
  beforeEach(module('crawlfishApp'));

  // instantiate service
  var crawlData;
  beforeEach(inject(function (_crawlData_) {
    crawlData = _crawlData_;
  }));

  it('should do something', function () {
    expect(!!crawlData).toBe(true);
  });

});
