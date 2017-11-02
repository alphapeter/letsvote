var embed = require('./embedingo.js')

var goPackageName = 'gui'
var property = 'Css'
var destinationFileName = '../server/gui/CssBundle.go'
var sourceFileName = './dist/static/css/app.css'

embed.run(goPackageName, property, destinationFileName, sourceFileName)
