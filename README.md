# bc-wallet-common-lib-redis

## Description

Library for manage redis config and connections

Library contains:
* common redis config struct
* connection manager

## Usage example

Examples of create connection and write database communication code

### Config and connection

```go
package main

import (
	"context"
	
	commonEnvConfig "gitlab.heronodes.io/bc-platform/bc-wallet-common-lib-config/pkg/envconfig"
	commonRedis "gitlab.heronodes.io/bc-platform/bc-wallet-common-lib-postgres/pkg/redis"
	commonVault "gitlab.heronodes.io/bc-platform/bc-wallet-common-lib-vault/pkg/vault"
	commonVaultTokenClient "gitlab.heronodes.io/bc-platform/bc-wallet-common-lib-vault/pkg/vault/client/token"
)

type VaultWrappedConfig struct {
	*commonVault.BaseConfig
	*commonVaultTokenClient.AuthConfig
}

func main() {
	ctx := context.Background()
	// vault config and client prepare 
	vaultCfg := &VaultWrappedConfig{
		BaseConfig: &commonVault.BaseConfig{},
		AuthConfig: &commonVaultTokenClient.AuthConfig{},
	}
	vaultClientSrv, err := commonVaultTokenClient.NewClient(ctx, vaultCfg)
	if err != nil {
		panic(err)
	}
	// vault service-component creation prepare 
	vaultSrc, err := commonVault.NewService(ctx, vaultCfg, vaultClientSrv)
	if err != nil {
		panic(err)
	}

	_, err = vaultSrc.Login(ctx)
	if err != nil {
		panic(err)
	}

	// REDIS CONFIG PREPARE
	redisConfigSrc := commonRedis.RedisConfig{}
	rdCfgPreparerSrv := commonEnvConfig.NewConfigManager()
	err = rdCfgPreparerSrv.PrepareTo(redisConfigSrc).With(vaultSrc).Do(ctx)
	if err != nil {
		panic(err)
	}

	// REDIS CONNECTION
	rdConn := commonRedis.NewConnection(ctx, redisConfigSrc, loggerSvc)
	_, err = rdConn.Connect()
	if err != nil {
		panic(err)
	}
	
	rdClient := rdConn.GetClient()
	redisCMD := rdClient.Get(ctx, "some_key")
	rawData, err := redisCMD.Result()
	if err != nil {
		panic(err)
	}
}
```

## Contributors

* Author and maintainer - [@gudron (Alex V Kotelnikov)](https://github.com/gudron)

**bc-wallet-common-lib-redis** has a proprietary license.

Switched to proprietary license from MIT - [CHANGELOG.MD - v0.0.5](./CHANGELOG.md)