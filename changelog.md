# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html)

## [0.3.2] - 2023/01/20

### Added

- Added a normalizer package.
- Added `LowerString`, `UpperString` normalizers.
- Added an URL validator.
- Now all validation errors are displayed if there are any. Before verdeter quit on the first error.
- Now all natives feature of cobra are available through the VerdeterConfig type and the build function.

### Fixed

- Fixed a bug: the SetDefault method now panic if called on an undeclared config key.

### Changed

- The signature of AuthorizedValues was altered to make it more understandable.

## [0.3.1] - 2022/11/09

### Fixed

- Fixed an initialization bug. The initialization function did not run when a subcommand was called. This caused the command to fail it's validation step.

## [0.3.0] - 2022/10/31

### Added

- A method named `Lookup` on the VerdeterCommand type. It allow to search in both local and global config keys. If no config key is found, it return nil.

### Changed

- The tasks in the CI are now ran in parallel. 
- The VerdeterCommand method `SetValidator` is now named `AddValidator`. Although the argument list did not change, the behavior did. Now the method will add a validator to the ConfigKey validators list.

### Fixed

- fixed a bug regarding the validation cascade in multicommand app. The validation function of a root Command  return an error on valid input when a subcommand was called.

## [0.2.1] - 2022/10/11

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

[0.3.2]: https://github.com/ditrit/verdeter/blob/v0.3.2/changelog.md
[0.3.1]: https://github.com/ditrit/verdeter/blob/v0.3.1/changelog.md
[0.3.0]: https://github.com/ditrit/verdeter/blob/v0.3.0/changelog.md
[0.2.1]: https://github.com/ditrit/verdeter/blob/v0.2.1/changelog.md
[0.2.0]: https://github.com/ditrit/verdeter/blob/v0.2.0/changelog.md
[0.1.0]: https://github.com/ditrit/verdeter/blob/0.1.0/changelog.md