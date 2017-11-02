var fs = require('fs')

exports.run = function(package, variable, destinationFileName, sourceFileName){
    process.stdout.write("\n* Embedding in go: " + sourceFileName + " -> " + destinationFileName
        + " \n ** package: '" + package + "', " + "variable: '" + variable +  "'\n");
    var content = fs.readFileSync(sourceFileName,{encoding: 'utf8'});

    fs.writeFileSync(destinationFileName,
        "package " + package + "\n\n" +
        "var " + variable + " = []byte(`"
        + content + "`)",
        {
            encoding: "utf8",
            flag: "w"
        }
    );
};





