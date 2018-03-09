var embed = require('./embedingo.js')

{
  let goPackageName = 'webui'
  let property = 'Html'
  let destinationFileName = '../server/webui/Html.go'
  let sourceFileName = './dist/index.html'
  embed.run(goPackageName, property, destinationFileName, sourceFileName)
}

{
  let goPackageName = 'webui'
  let property = 'AdminHtml'
  let destinationFileName = '../server/webui/AdminHtml.go'
  let sourceFileName = './dist/admin.html'
  embed.run(goPackageName, property, destinationFileName, sourceFileName)
}
