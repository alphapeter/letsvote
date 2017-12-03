var embed = require('./embedingo.js')

var goPackageName = 'webui'
var property = 'Html'
var destinationFileName = '../server/webui/Html.go'
var sourceFileName = './dist/index.html'

embed.run(goPackageName, property, destinationFileName, sourceFileName)
