package main

import (
	_ "einar-website-api/app/shared/configuration"
	"einar-website-api/app/shared/constants"
	_ "embed"
	"log"
	"os"

	_ "einar-website-api/app/shared/infrastructure/healthcheck"
	_ "einar-website-api/app/shared/infrastructure/serverwrapper"

	ioc "github.com/Ignaciojeria/einar-ioc/v2"
)

//go:embed .version
var version string

func main() {
	os.Setenv(constants.Version, version)
	os.Setenv("PROJECT_NAME", "einar-website-api")
	os.Setenv("ENVIRONMENT", "production")
	if err := ioc.LoadDependencies(); err != nil {
		log.Fatal(err)
	}
}
