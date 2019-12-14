package main

import (
	"context"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	identity string
	verbose  bool
	scheme   string
	wait     time.Duration

	puppetSSlDir string
	puppetCA     string

	log *logrus.Logger

	ctx    context.Context
	cancel func()

	version = "0.0.1"
)

// Run runs the command
func Run() {
	app := kingpin.New("pki-enroll", "Enrolls with various PKI systems using the Choria framework")
	app.Author("R.I.Pienaar <rip@devco.net>")
	app.Version(version)

	app.Arg("identity", "Identity to enroll as").Required().StringVar(&identity)
	app.Flag("scheme", "Provider to enroll with, only support 'puppet'").Default("puppet").EnumVar(&scheme, "puppet")
	app.Flag("wait", "How long to wait for the certificate to be signed").Default("30m").DurationVar(&wait)

	app.Flag("puppet-ssldir", "The directory to write the Puppet compatible SSL structure").PlaceHolder("PATH").StringVar(&puppetSSlDir)
	app.Flag("puppet-ca", "PuppetCA in host:port format").Default("puppet:8140").StringVar(&puppetCA)

	app.Flag("verbose", "Verbose logging").Default("false").BoolVar(&verbose)

	kingpin.MustParse(app.Parse(os.Args[1:]))

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	log = logrus.New()
	log.Out = os.Stdout

	log.SetLevel(logrus.InfoLevel)

	if verbose {
		log.SetLevel(logrus.DebugLevel)
	}

	switch scheme {
	case "puppet":
		kingpin.FatalIfError(puppetEnroll(), "Enrolling with Puppet failed")
	default:
		kingpin.Fatalf("Invalid scheme '%s'", scheme)
	}
}
