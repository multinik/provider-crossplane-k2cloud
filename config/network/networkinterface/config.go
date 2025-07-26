package networkinterface

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_network_interface", func(r *config.Resource) {
		r.Kind       = "NetworkInterface" // безопасное имя
		r.ShortGroup = "network"
	})
}