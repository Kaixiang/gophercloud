package main

import (
	"flag"
	"fmt"
	"github.com/rackspace/gophercloud"
)

var quiet = flag.Bool("quiet", false, "Quiet mode for acceptance testing.  $? non-zero on error though.")
var rgn = flag.String("r", "", "Datacenter region to interrogate.  Leave blank for provider-default region.")

func main() {
	flag.Parse()

	withIdentity(false, func(auth gophercloud.AccessProvider) {
		withServerApi(auth, func(servers gophercloud.CloudServersProvider) {
			images, err := servers.ListImages()
			if err != nil {
				panic(err)
			}

			if !*quiet {
				fmt.Println("ID,Name,MinRam,MinDisk")
				for _, image := range images {
					fmt.Printf("%s,\"%s\",%d,%d\n", image.Id, image.Name, image.MinRam, image.MinDisk)
				}
			}
		})
	})
}
