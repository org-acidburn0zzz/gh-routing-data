const fs = require("fs")
const path = require("path")
const child_process = require("child_process")

const patterns = ["*_bike2", "*_mtb", "*_racingbike"]
const bucketName = "gh-data"

function unlinkIfExists(file) {
 try {
   fs.unlinkSync(file)
 }
 catch (ignore) {
 }
}

function spawn(command, args, data) {
  process.stderr.write(`${command} ${args.join(" ")}\n`)
  return new Promise((resolve, reject) => {
    const child = child_process.spawn(command, args, data == null ? {stdio: ["ignore", process.stdout, process.stderr]} : {stdio: ["pipe", process.stdout, process.stderr]})
    child.on("error", reject)
    child.on("close", code => {
      if (code === 0) {
        resolve()
      }
      else {
        reject(new Error(`${command} exited with code ${code}`))
      }
    })

    if (data != null) {
      child.stdin.end(data)
    }
  })
}

// 7za very slow - compression ratio doesn't worth spent time (2m52.212s vs 11m16.306s)
async function main() {
  const remoteDir = new Date().toISOString().substr(0, 10)

  process.chdir(process.env.MAP_DIR)

  const resultName = process.argv[2]
  const dir = `${resultName}.osm-gh`
  const fileStat = fs.statSync(path.join(dir, "shortcuts_fastest_bike2"))
  const archiveFiles = []
  if (fileStat.size < 600000000) {
    const outFile = `${resultName}.osm-gh.zip`
    archiveFiles.push(outFile)
    if (process.env.SKIP_ZIP == null) {
      unlinkIfExists(outFile)
      await spawn("zip", ["-r", "-9", outFile, dir])
    }
  }
  else {
    // if > 700 MB, compress as 2 zip files
    const outFile1 = `${resultName}-part1.osm-gh.zip`
    const outFile2 = `${resultName}-part2.osm-gh.zip`
    archiveFiles.push(outFile1)
    archiveFiles.push(outFile2)

    if (process.env.SKIP_ZIP == null) {
      unlinkIfExists(outFile1)
      unlinkIfExists(outFile2)
      await Promise.all([
        spawn("zip", ["-r", "-9", outFile1, dir, "-i"].concat(patterns)),
        spawn("zip", ["-r", "-9", outFile2, dir, "-x"].concat(patterns)),
      ])
    }
  }

  if (process.env.SKIP_UPLOAD == null) {
    const promises = []
    for (const file of archiveFiles) {
      promises.push(spawn("aws", ["s3", "cp", file, `s3://${bucketName}/${remoteDir}/${file}`, "--profile", "sw", "--acl", "public-read"]))
    }
    await Promise.all(promises)

    let data = "<locusActions>"
    for (const file of archiveFiles) {
      const remotePath = `${remoteDir}/${file}`
      data += `<download>
  <source><![CDATA[http://gh-data.s3.nl-ams.scw.cloud/${remotePath}]]></source>
  <dest><![CDATA[/mapsVector/${resultName}]]></dest>
  <after>extract|deleteSource</after>
</download>
`
    }
    data += "</locusActions>"

    await spawn("aws", ["s3", "cp", "-", `s3://${bucketName}/${remoteDir}/${resultName}.locus.xml`, "--profile", "sw", "--acl", "public-read"], data)
  }
}

main()
  .catch(e => {
    console.error(e.stack)
    process.exit(1)
  })