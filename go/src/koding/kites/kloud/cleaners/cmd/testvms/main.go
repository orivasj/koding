package main

import (
	"fmt"
	"koding/kites/kloud/cleaners/lookup"
	"time"

	"github.com/koding/multiconfig"
	"github.com/mitchellh/goamz/aws"
)

type Config struct {
	Terminate bool

	// AWS Access and Secret Key
	AccessKey string `required:"true"`
	SecretKey string `required:"true"`
}

func main() {
	conf := new(Config)

	// Load the config, it's reads environment variables or from flags
	multiconfig.New().MustLoad(conf)

	auth := aws.Auth{
		AccessKey: conf.AccessKey,
		SecretKey: conf.SecretKey,
	}

	l := lookup.New(auth)

	fmt.Printf("Searching for instances tagged with [sandbox, dev] and older than 1 day ...\n")

	instances := l.FetchInstances()
	instances.OlderThan(time.Hour * 24)
	instances.WithTag("koding-env", "sandbox", "dev")

	fmt.Println(instances)

	fmt.Printf("All regions total: %+v\n", instances.Total())

	// if conf.Terminate {
	// 	l.TerminateAll()
	// 	fmt.Printf("Terminated '%d' instances\n", total)
	// } else if total > 0 {
	// 	fmt.Printf("To delete all VMs run the command again with the flag -terminate\n")
	// }
}
