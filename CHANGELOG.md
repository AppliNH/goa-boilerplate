## 0.2.1

- Fix Makefile docker build command
- Improve CI workflow
  - Use go 1.24 for golangci-lint-action as it is not compatible with go 1.25
  - Remove unnecessary `go mod download` for docker build

## 0.2.0

- Add config validation, default values and automated registration of config keys.
- Add sloghttp middleware to exclude health checks from logging.

## 0.1.2

- Fix typo in readme.

## 0.1.1

- Update readme to skip .git folder when using copier.

## 0.1.0

- Initial release with goa design generation and basic project structure.