var embed = require('./embedingo.js');

var package = "gui";
var property = "Javascript";
var destinationFileName = "../server/gui/JavascriptBundle.go";
var sourceFileName = "dist/static/js/app.js";

embed.run(package, property, destinationFileName, sourceFileName);
