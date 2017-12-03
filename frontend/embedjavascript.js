var embed = require('./embedingo.js')

var goPackageName = 'webui'
var property = 'Javascript'
var destinationFileName = '../server/webui/JavascriptBundle.go'
var sourceFileName = 'dist/static/js/app.js'

embed.run(goPackageName, property, destinationFileName, sourceFileName)
