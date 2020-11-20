const fs = require('fs');

function getJSFiles() {}
// Implement the function getJSFiles, it should take as arguments:
// a directory (string)
// a callback function
// the callback function must be invoked with an error and an array of js files only

getJSFiles('./directory', function (error, jsFiles) {
  console.log(`The jsFiles are: `, jsFiles);
});
