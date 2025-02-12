
package main


import (
	"github.com/psychedelicnekopunch/httpclient-checker/src/infrastructure"
)


func main() {
	config := infrastructure.NewConfig()
	routing := infrastructure.NewRouting(config)
	routing.Start()
}
