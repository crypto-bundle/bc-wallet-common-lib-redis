# Change Log

## [v0.0.1] - 31.04.2022 20:42 MSK
### Added
* Connection config
* Connection wrapper with client option preparation flow

## [v0.0.2] - 17.03.2023 16:35 MSK
### Changed
* Lib-redis moved to another repository - https://github.com/crypto-bundle/bc-wallet-common-lib-redis

## [v0.0.3] - 07.04.2023 18:49 MSK
### Changed
* Changed redis client config
  * Added supporting of _secret_ tags for lib-config
  * Small refactoring for using config interface implementation

## [v0.0.4] - 09.05.2023
### Added
* Added Dragonfly helm-chart for local development. Chart cloned from [official Dragonfly repository](https://github.com/dragonflydb/dragonfly/tree/main/contrib/charts/dragonfly)
### Changed
* Changed content of license file - MIT license
* Changed go-namespace

## [v0.0.5] - 14.04.2024
### Added
* Added support of healthcheck flow, which required by [lib-healthcheck](https://github.com/crypto-bundle/bc-wallet-common-lib-healthcheck)