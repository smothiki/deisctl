package cmd

import (
	"fmt"
	"github.com/deis/deisctl/systemd"
	"github.com/deis/deisctl/utils"
)

const (
	DownloadDir = "/home/core/deis/systemd/"
	numServices = 11
)

func getService() map[string][]string {
	service := map[string][]string{
		"builder": {
			"deis-builder.service",
			"deis-builder-data.service",
		},
		"cache": {
			"deis-cache.service",
		},
		"controller": {
			"deis-controller.service",
		},
		"database": {
			"deis-database.service",
			"deis-database-data.service",
		},
		"logger": {
			"deis-logger.service",
			"deis-logger-data.service",
		},
		"registry": {
			"deis-registry.service",
			"deis-registry-data.service",
		},
		"router": {
			"deis-router.service",
		},
	}
	return service
}

func Install(service []string) {
	fmt.Println("starting systemd units")
	files, _ := utils.ListFiles(DownloadDir + "*.service")
	if len(files) != numServices {
		fmt.Println("not enough files")
		return
	}
	conn, _ := systemd.NewSystemdUnitManager()
	conn.Enable(files)
	Services := getService()
	if len(service) == 0 {
		for _, serv := range Services {
			for _, sub := range serv {
				conn.Start(sub)
			}
		}
	} else {
		Start(service, false)
	}
}

func Uninstall(service []string) {
	fmt.Println("starting systemd units")
	conn, _ := systemd.NewSystemdUnitManager()
	Services := getService()
	if len(service) == 0 {
		for _, serv := range Services {
			for _, sub := range serv {
				conn.Stop(sub)
			}
		}
	} else {
		Stop(service, false)
	}

	conn.Disable(service)
}

func Start(service []string, data bool) {
	Services := getService()
	conn, _ := systemd.NewSystemdUnitManager()
	for _, name := range service {
		for _, DeisService := range Services[name] {
			conn.Start(DeisService)
			if !data {
				break
			}
		}
	}
}

func Stop(service []string, data bool) {
	Services := getService()
	conn, _ := systemd.NewSystemdUnitManager()
	for _, name := range service {
		for _, DeisService := range Services[name] {
			conn.Stop(DeisService)
			if !data {
				break
			}
		}
	}
}
