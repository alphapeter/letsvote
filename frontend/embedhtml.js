var embed = require('./embedingo.js')

var goPackageName = 'gui'
var property = 'Html'
var destinationFileName = '../server/gui/Html.go'
var sourceFileName = './dist/index.html'

embed.run(goPackageName, property, destinationFileName, sourceFileName)
