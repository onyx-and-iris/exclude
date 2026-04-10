![Windows](https://img.shields.io/badge/Windows-0078D6?style=for-the-badge&logo=windows&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)
![macOS](https://img.shields.io/badge/mac%20os-000000?style=for-the-badge&logo=macos&logoColor=F0F0F0)

# exclude

[Exclude files][exclude] from local git tracking

## Install

```console
go install github.com/onyx-and-iris/exclude@latest
```

## Configuration

*flags*

-   --path/-p: Path the exclude file resides in (default is .git/info/)

*environment variables*

```bash
#!/usr/bin/env bash

export EXCLUDE_PATH="./cmd/testdata/"
```

## Commands

### Add

usage: *exclude add [pattern-1] [pattern-2]...*

```console
exclude add "*.log", "temp/"
```

### Del

usage: *exclude del [pattern]*

```console
exclude del "*.log"
```

### List

```console
exclude list
```

### Reset

```console
exclude reset
```

## Special Thanks

-   [spf13](https://github.com/spf13) for the [cobra](https://github.com/spf13/cobra) and [viper](https://github.com/spf13/viper) packages.
-   The developers at [charmbracelet](https://github.com/charmbracelet) for the [fang](https://github.com/charmbracelet/fang) package.


[exclude]: https://docs.github.com/en/get-started/git-basics/ignoring-files#excluding-local-files-without-creating-a-gitignore-file