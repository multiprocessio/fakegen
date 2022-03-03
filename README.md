# fakegen: Single binary CLI for generating a random schema of M columns to populate N rows of data

This program generates a random schema of M columns and then generates
N rows of that schema. So all value types within a column across all
rows will be consistent. For example, if a value is an int in one
row's column, it will be an int in the same column across all other
row's.

## Install

Binaries for amd64 (x86_64) are provided for each release.

### macOS, Linux, WSL

On macOS, Linux, and WSL you can run the following:

```bash
$ curl -LO "https://github.com/multiprocessio/fakegen/releases/download/0.1.0/fakegen-$(uname -s | awk '{ print tolower($0) }')-x64-0.1.0.zip"
$ unzip fakegen-*-0.1.0.zip
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
