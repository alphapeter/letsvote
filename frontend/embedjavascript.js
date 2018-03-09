var embed = require('./embedingo.js')

{
  let goPackageName = 'webui'
  let property = 'Javascript'
  let destinationFileName = '../server/webui/JavascriptBundle.go'
  let sourceFileName = 'dist/static/js/app.js'
  embed.run(goPackageName, property, destinationFileName, sourceFileName)
}

{
  let goPackageName = 'webui'
  let property = 'AdminJavascript'
  let destinationFileName = '../server/webui/AdminJavascriptBundle.go'
  let sourceFileName = 'dist/static/js/admin.js'
  embed.run(goPackageName, property, destinationFileName, sourceFileName)
}
