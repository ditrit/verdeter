# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html)

## [0.3.0] - 2022/10/10

### Fixed 

- fixed a bug in the command initialization that affected multicommand apps: the initialisation function ran multiples times instead of running one time on the root command. Now the initialisation is ran in the pre run step only on the root command.
- fixed configuration search to be posix compliant.

## [0.2.0] - 2022/10/06

### Added 

- Add Uint type for config keys. It prevents the user from creating a validator to check that the value is not <0 .
- Add a validator named AuthorizedValues that check if the value of a config key is contained in a defined array of authorized values.

## [0.1.0] - 2022/09/06

### Added

- Setup project (ci and sonar).
- Setup e2e test solution (cucumber + docker).
- This project was imported from github.com/ditrit/gandalf/verdeter.

[0.3.0]: https://github.com/ditrit/verdeter/blob/v0.3.0/changelog.md
[0.2.0]: https://github.com/ditrit/verdeter/blob/v0.2.0/changelog.md
[0.1.0]: https://github.com/ditrit/verdeter/blob/0.1.0/changelog.md