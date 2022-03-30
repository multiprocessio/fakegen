# fakegen: Single binary CLI for generating a random schema of M columns to populate N rows of JSON, CSV, Excel, etc.

This program generates a random schema of M columns and then generates
N rows of that schema. So all value types within a column across all
rows will be consistent. For example, if a value is an int in one
row's column, it will be an int in the same column across all other
row's.

It generates JSON by default but can generate other formats like CSV,
TSV, Excel, etc.

## Install

Binaries for amd64 (x86_64) are provided for each release.

### macOS, Linux, WSL

On macOS, Linux, and WSL you can run the following:

```bash
$ curl -LO "https://github.com/multiprocessio/fakegen/releases/download/0.2.0/fakegen-$(uname -s | awk '{ print tolower($0) }')-x64-0.2.0.zip"
$ unzip fakegen-*-0.2.0.zip
$ sudo mv fakegen /usr/local/bin/fakegen
```

Or install manually from the [releases
page](https://github.com/multiprocessio/fakegen/releases), unzip and add
`fakegen` to your `$PATH`.

### Windows, not WSL

Download the [latest Windows
release](https://github.com/multiprocessio/fakegen/releases), unzip it,
and add `fakegen` to your `$PATH`.

### Manual, and other Go platforms

If you are on another platform or architecture or want to grab the
latest release, you can do so with Go 1.17+:

```
$ go install github.com/multiprocessio/fakegen@latest
```

`fakegen` will likely work on other platforms that Go is ported to such as
AARCH64 and OpenBSD, but tests and builds are only run against x86_64
Windows/Linux/macOS.

## Usage

Pass the number of rows and columns you want and `fakegen` will give
you a JSON array of objects with that many rows and unique columns.

```
$ fakegen --rows 2 --cols 5 | jq .
[
  {
    "enecate": 1322845113,
    "irruptions": "et tempore suscipit dignissimos odit ut accusantium dolores cumque est dignissimos ut dolorem saepe quia laborum doloribus quisquam sapiente illo omnis dolor consequatur incidunt quisquam vero tempore quae eos doloribus temporibus et eligendi aspernatur molestias sed pariatur qui officiis voluptate quis tempore laboriosam fugit in recusandae explicabo nemo ut neque est quia aliquam ex animi reprehenderit sint neque eaque quibusdam eius ducimus consequatur nostrum ut facilis id quam non rerum architecto. dolor reiciendis autem reprehenderit nostrum assumenda tempore et ex est vero error sequi ut magni quis molestias nemo voluptatum omnis nesciunt.",
    "phototelescopic": "facere non iusto a pariatur vero qui magnam nostrum quibusdam magnam a omnis adipisci molestias dolores commodi at consequatur architecto doloribus tempora qui inventore quia officiis illo nemo et eos doloremque maxime omnis fuga qui quibusdam. et sit molestiae iste dolor totam facere debitis quae ullam nam sed amet at ipsam culpa repellendus expedita sit maiores quaerat odio exercitationem qui et itaque voluptas dolores nesciunt quia mollitia nesciunt laudantium fuga in nulla doloribus omnis et odio necessitatibus soluta asperiores est velit nobis nihil nulla ea et necessitatibus aut eius pariatur enim inventore qui nobis corrupti nam non ullam et esse exercitationem totam qui.",
    "restrictively": false,
    "upwafted": "2020-10-09T21:03:23.109945329Z"
  },
  {
    "enecate": 1771337749,
    "irruptions": "voluptas quis commodi qui commodi soluta ut debitis ipsam reprehenderit odio quaerat animi temporibus praesentium repellendus quae sapiente alias id assumenda dolorum rem aut numquam repellendus sed asperiores nulla ut accusamus consectetur incidunt a accusamus qui blanditiis ut maxime velit et inventore vel aliquid sit autem ex quo quae rerum aspernatur ullam ut aut sed dolor eos ut dolorem id quia aperiam libero magnam perferendis cupiditate qui ex corrupti numquam dicta id laboriosam corporis illum asperiores enim soluta animi debitis deleniti totam corporis corporis dolores a qui.",
    "phototelescopic": "non suscipit illo omnis explicabo omnis omnis quia eligendi quidem suscipit tenetur odit dicta incidunt asperiores nisi vel sit porro voluptatem commodi error autem exercitationem dicta quas totam necessitatibus neque et et et quia consectetur facere suscipit repellat dolor aliquam culpa harum aspernatur dolorem nihil dolorem dolorum ex in culpa molestiae nihil odio et doloremque repellendus blanditiis et quae et similique nam culpa ratione fugit et dolorum dolorum unde qui qui veniam occaecati sit nemo asperiores ipsa excepturi soluta odit dolores excepturi occaecati sit. aut harum vel vitae dicta quam quibusdam magni quam qui architecto odit excepturi officiis eum rerum aliquam est molestias similique assumenda sunt autem velit molestiae tempora dolor et esse quisquam consequuntur ducimus deserunt consequatur earum doloribus ratione eius repellendus quidem omnis quaerat deserunt officia qui possimus officia dicta sit qui neque sunt blanditiis illo veritatis consequatur eaque praesentium quibusdam ratione rem dolores magni odio quisquam tempora. consequatur laudantium itaque omnis temporibus mollitia dolores quisquam ab vero inventore et dolorem ea ut quia laudantium neque odit veniam voluptatem vero et delectus rerum quaerat architecto ab vitae tempore error omnis dolor et doloremque dolor.",
    "restrictively": true,
    "upwafted": "1979-01-09T05:20:25.7650053Z"
  }
]
```

### Formats

You can change the output format by passing the `-f` or `--format` flag.

```
$ fakegen -r 10 -c 2 -f csv
seavy,spangle-baby
2070-08-27T17:07:35Z,-28123.3
2063-10-30T15:06:37Z,-82834.6
2066-04-29T00:10:38Z,-74438.5
2006-08-31T11:45:50Z,-85888.4
1981-11-20T20:49:30Z,-78208.3
2065-07-02T12:31:05Z,13387.5
2003-12-29T22:37:19Z,-62472.3
1981-08-29T14:41:20Z,-55740.1
2062-07-09T08:56:47Z,54202.2
2010-05-30T08:28:33Z,-73394.6
```

Here are the supported format strings:

| Format | Description |
|--------|--------------|
| `json`   | Array of JSON objects             |
| `jsonl` | JSON objects separated by newlines |
| `csv` | |
| `tsv` | |
| `xlsx` | Excel file with one sheet, "Sheet1". |

In the future other formats like OpenOffice ODS, Parquet, Apache ORC,
etc. would be great to have.

### Output file

You can specify a file to write to with the `-o` or `--out` flag. Some
formats like `xlsx` require this flag.

If you specify this with a wellknown extension you can omit the format
flag.

```
$ fakegen -r 10 -c 2 -o data.csv
$ cat data.csv
courtyard,pagods
2008-01-06T18:00:42Z,false
2015-04-08T21:26:41Z,true
2055-04-21T08:48:07Z,false
2023-02-17T06:37:25Z,false
<nil>,true
2008-05-18T06:57:03Z,false
1987-11-18T21:15:48Z,true
2014-04-12T08:38:32Z,false
2012-11-06T01:40:37Z,false
1992-11-26T03:51:10Z,false
```

### Null frequency

By default nulls will show up in 10% of generated cells. You can
modify this by setting the `-n` or `--null-frequency` flag with a
decimal value.

To disable nulls, set the flag to `0`.

```
$ fakegen -r 10 -c 2 -n 0
[{"unctious":"1977-10-31T22:05:05.68606544Z","misagent":"2019-02-05T08:10:49.013647805Z"},
{"unctious":"2052-11-28T15:10:40.998426932Z","misagent":"1995-10-05T14:02:26.732748512Z"},
{"unctious":"1982-10-14T04:41:34.326758028Z","misagent":"2070-03-20T05:50:11.749294271Z"},
{"unctious":"1984-06-04T00:09:05.594979649Z","misagent":"2047-08-31T23:08:52.655138927Z"},
{"unctious":"1979-01-10T05:38:35.725041374Z","misagent":"2043-06-14T19:22:48.02132443Z"},
{"unctious":"2012-03-05T06:43:01.640412792Z","misagent":"2031-10-02T20:41:54.617712604Z"},
{"unctious":"2029-02-23T05:48:40.869202594Z","misagent":"1992-08-18T18:07:09.712263831Z"},
{"misagent":"2064-12-09T01:31:53.965240833Z","unctious":"1999-10-18T04:57:53.869159811Z"},
{"unctious":"1992-08-18T00:58:12.024110889Z","misagent":"2024-09-07T05:58:36.481215844Z"},
{"unctious":"2064-04-05T16:46:31.701345883Z","misagent":"2066-09-11T14:15:54.357142854Z"}]
```

To get nulls 50% of the time, set the flag to `.5`.

```
$ fakegen -r 10 -c 2 -n .5
[{"retraded":-58640.1,"high-compression":null},
{"high-compression":1519695743,"retraded":null},
{"high-compression":-1466727411,"retraded":-59521.6},
{"high-compression":null,"retraded":-6683.1},
{"high-compression":null,"retraded":null},
{"high-compression":null,"retraded":null},
{"high-compression":null,"retraded":null},
{"high-compression":721870540,"retraded":-18646.4},
{"high-compression":-922344240,"retraded":19933.6},
{"high-compression":-1471776625,"retraded":null}]
```
