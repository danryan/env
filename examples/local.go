package main

import (
	"fmt"
	env "github.com/danryan/env"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Sdump()

type Config struct {
	Name    string `env:"key=NAME required"`
	Port    int    `env:"key=PORT default=9000"`
	Adapter string `env:"key=ADAPTER default=shell in=shell,slack,hipchat"`
	Enabled bool   //`env:"key=IS_ENABLED default=true"`
}

// HAL_NAME=hal
// HAL_PORT=
// HAL_ADAPTER=
func main() {
	// os.Setenv("IS_ENABLED", "asdf")
	c := &Config{}
	if err := env.Process(c); err != nil {
		fmt.Println("error:", err)
		return
		// panic(err)
	}
	fmt.Printf("name: %s, port: %d, adapter: %s, enabled: %v\n", c.Name, c.Port, c.Adapter, c.Enabled)
}
