package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/choria-io/go-security/puppetsec"
)

func puppetEnroll() error {
	if puppetSSlDir == "" {
		return errors.New("the Puppet scheme requires its ssl directory set on the CLI, see --help")
	}

	cfg := puppetsec.Config{
		Identity:   identity,
		SSLDir:     puppetSSlDir,
		DisableSRV: true,
	}

	re := regexp.MustCompile(`^((([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9]))\:(\d+)$`)

	if re.MatchString(puppetCA) {
		parts := strings.Split(puppetCA, ":")
		cfg.PuppetCAHost = parts[0]

		p, err := strconv.Atoi(parts[1])
		if err != nil {
			return err
		}

		cfg.PuppetCAPort = p
	}

	prov, err := puppetsec.New(puppetsec.WithConfig(&cfg), puppetsec.WithLog(log.WithField("provider", "puppet")))
	if err != nil {
		return err
	}

	return prov.Enroll(ctx, wait, func(try int) { fmt.Printf("Attempting to download certificate for %s, try %d.\n", identity, try) })
}
