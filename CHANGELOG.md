# Changelog

All notable changes to this project will be documented here.

To see more information about this file, see README.md section CHANGELOG file

This project is inspired by [Keep a Changelog](https://keepachangelog.com/en/1.1.0/) and does not use semantic versioning.

## [Unreleased]
## [Current]
### Changed
- Rewrote flag package almost completly
- Added flags type on which Add(to add flag), Parse(to parse cli into \*flag type objects) and Status(to get flags status: parsed/read) are called
- Added flag type, returned by flags.Add(FlagDescription) method and on wich method Get()(returning string value, flag status(set/not_set;assigned_with_value/not_assigned;read/wan't_read) and error) is called.
- Now flag may be specified only once, otherwise parsing fails.
- Now flag may have aliases, description(for future error support), default value, require or prohibit value setting
## [2026-07-25_15-21-35Z__9cf356a7e030016800a11b73f7cc7556e708c432]
### Changed
- Now handler returns cmd.Cmd(interface with methods Exec and Update) and error instead of just error.
### Added
- CHANGELOG.md

