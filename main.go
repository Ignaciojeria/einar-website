package main

import (
	_ "einar-website-api/app/shared/configuration"
	"einar-website-api/app/shared/constants"
	_ "embed"
	"log"
	"os"

	ioc "github.com/Ignaciojeria/einar-ioc/v2"
	_ "einar-website-api/app/shared/infrastructure/serverwrapper"
	_ "einar-website-api/app/shared/infrastructure/healthcheck"
)

//go:embed .version
var version string

func main() {
	os.Setenv(constants.Version, version)
	if err := ioc.LoadDependencies(); err != nil {
		log.Fatal(err)
	}
}
