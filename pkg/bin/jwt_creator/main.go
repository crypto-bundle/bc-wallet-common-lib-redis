package main

import (
	"flag"
	"log"

	jwtvalitator "github.com/cryptowize-tech/bc-wallet-common/pkg/jwt"

	_ "github.com/mailru/easyjson/gen"
	"go.uber.org/zap"
)

func main() {
	var (
		key, uuid, expiration string
	)

	flag.StringVar(&key, "key", "", "secret key")
	flag.StringVar(&uuid, "uuid", "", "merchant uuid")
	flag.StringVar(&expiration, "expiration", "", "expiration date with format: '2006-01-02'")
	flag.Parse()
	if key == "" {
		log.Fatalf("missing required -%v argument/flag\n", "key")
	}
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	jwtSrv := jwtvalitator.NewService(key, logger)
	token, err := jwtSrv.GenerateJWT(uuid, expiration)
	if err != nil {
		log.Fatalf("can not make JWT token. Error: %v ", err.Error())
	}
	log.Println("Token: ", token)
}
