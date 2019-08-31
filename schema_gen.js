const {
    mergeSchemas
} = require('graphql-tools');
const { promisify } = require('util');
const { print, isSpecifiedScalarType, isSpecifiedDirective } = require('graphql');

function printSchemaWithDirectives(schema) {
  const str = Object
    .keys(schema.getTypeMap())
    .filter(k => !k.match(/^__/))
    .reduce((accum, name) => {
      const type = schema.getType(name);
      return !isSpecifiedScalarType(type)
        ? accum += `${print(type.astNode)}\n`
        : accum;
    }, '');

  return schema
    .getDirectives()
    .reduce((accum, d) => {
      return !isSpecifiedDirective(d)
        ? accum += `${print(d.astNode)}\n`
        : accum;
    }, str + `${print(schema.astNode)}\n`);
}

const {readFile: readFileRegular, writeFile: writeFileRegular } = require('fs');
const writeFile = promisify(writeFileRegular);
const readFile = promisify(readFileRegular);
const glob = promisify(require('glob'));
const outfileName = "merged.graphql";
const blacklist = [ outfileName];

const main = async () => {
    const files = (await glob("**/*.graphql")).filter(v => !blacklist.some(b => b == v));

    const schemas = await Promise.all(files.map(async filename => 
        (await readFile(filename)).toString()
    ));

    console.log(`loaded ${schemas.length} schemas: \n${files.join("\n").replace(/^/gm, "\t")}`)

    const merged = mergeSchemas({ schemas: schemas, mergeDirectives: true })
    debugger;
    await writeFile(outfileName, printSchemaWithDirectives(merged));
}

main().catch(err => { throw new Error(err) });