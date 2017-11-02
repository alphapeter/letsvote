var embed = require('./embedingo.js');

var package = "gui";
var property = "Html";
var destinationFileName = "../server/gui/Html.go";
var sourceFileName = "./dist/index.html";

embed.run(package, property, destinationFileName, sourceFileName);
