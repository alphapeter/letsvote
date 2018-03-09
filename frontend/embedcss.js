var embed = require('./embedingo.js')

{
  let goPackageName = 'webui'
  let property = 'Css'
  let destinationFileName = '../server/webui/CssBundle.go'
  let sourceFileName = './dist/static/css/app.css'

  embed.run(goPackageName, property, destinationFileName, sourceFileName)
}

{
  let goPackageName = 'webui'
  let property = 'AdminCss'
  let destinationFileName = '../server/webui/AdminCssBundle.go'
  let sourceFileName = './dist/static/css/admin.css'

  embed.run(goPackageName, property, destinationFileName, sourceFileName)
}
