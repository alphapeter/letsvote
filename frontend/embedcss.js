var embed = require('./embedingo.js')

var goPackageName = 'webui'
var property = 'Css'
var destinationFileName = '../server/webui/CssBundle.go'
var sourceFileName = './dist/static/css/app.css'

embed.run(goPackageName, property, destinationFileName, sourceFileName)
