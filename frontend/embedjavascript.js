var embed = require('./embedingo.js')

var goPackageName = 'gui'
var property = 'Javascript'
var destinationFileName = '../server/gui/JavascriptBundle.go'
var sourceFileName = 'dist/static/js/app.js'

embed.run(goPackageName, property, destinationFileName, sourceFileName)
