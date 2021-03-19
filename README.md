# github-latest

A CLI tool that fetches latest commit, tag or release information from Github repository.

You should set up your Github token in $HOME/.github-latest.yaml before use it. Like:

GITHUB_TOKEN: [your Github token]

## Usage

Fetch latest commit on default branch

```sh
github-latest commit spf13/viper
```

Fetch latest commit on specified branch

```sh
github-latest commit spf13/viper -b master
```

Fetch latest tag

```sh
github-latest tag spf13/viper
```

Fetch latest release

```sh
github-latest release spf13/viper
```

## Copyright

Copyright © 2021 Qian Hum

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
