var fs = require('fs')

exports.run = function (goPackageName, variable, destinationFileName, sourceFileName) {
  process.stdout.write('\n* Embedding in go: ' + sourceFileName + ' -> ' + destinationFileName +
        " \n ** package: '" + goPackageName + "', " + "variable: '" + variable + "'\n")
  var content = fs.readFileSync(sourceFileName, {encoding: 'utf8'})

  content = content.replace('`', '`+ "`" +`') // escape backticks

  fs.writeFileSync(destinationFileName,
    'package ' + goPackageName + '\n\n' +
        'var ' + variable + ' = []byte(`' +
        content + '`)',
    {
      encoding: 'utf8',
      flag: 'w'
    }
  )
}
