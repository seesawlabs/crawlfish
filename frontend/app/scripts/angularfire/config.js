angular.module('firebase.config', [])
  .constant('FBURL', 'https://crawlfish-dev.firebaseio.com')
  .constant('SIMPLE_LOGIN_PROVIDERS', ['password','google'])

  .constant('loginRedirectPath', '/login');
