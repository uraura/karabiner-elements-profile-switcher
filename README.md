# karabiner-elements-profile-switcher



## Description
automatically switch (with hammerspoon) karabiner-elements profile when connect/disconnect external keyboard.

## Usage
```
karabiner-elements-profile-switcher use ${Profile_Name}
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/uraura/karabiner-elements-profile-switcher
```

1. install hammerspoon.
2. copy switcher.lua to `$HOME/.hammerspoon` and change command paths, msgs, etc...
3. add following to 'init.lua'

       require 'switcher'

## Contribution

1. Fork ([https://github.com/uraura/karabiner-profile-switcher/fork](https://github.com/uraura/karabiner-profile-switcher/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[uraura](https://github.com/uraura)
