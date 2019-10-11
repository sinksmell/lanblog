/* eslint-disable */

var parser = new UAParser()

var uaResult = parser.getResult()

if (uaResult.device.tags === 'mobile') {
  window.location.href = './is-mobile.html'
}
