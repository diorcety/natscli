// Copyright 2020 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"os"

	"github.com/nats-io/natscli/cli"
	"gopkg.in/alecthomas/kingpin.v2"
)

var version = "development"

func main() {
	help := `NATS Utility

NATS Server and JetStream administration.

See 'nats cheat' for a quick cheatsheet of commands
	`

	ncli := kingpin.New("nats", help)
	ncli.Author("NATS Authors <info@nats.io>")
	ncli.UsageWriter(os.Stdout)
	ncli.Version(version)
	ncli.HelpFlag.Short('h')

	opts := &cli.Options{}

	ncli.Flag("server", "NATS server urls").Short('s').Envar("NATS_URL").PlaceHolder("NATS_URL").StringVar(&opts.Servers)
	ncli.Flag("user", "Username or Token").Envar("NATS_USER").PlaceHolder("NATS_USER").StringVar(&opts.Username)
	ncli.Flag("password", "Password").Envar("NATS_PASSWORD").PlaceHolder("NATS_PASSWORD").StringVar(&opts.Password)
	ncli.Flag("creds", "User credentials").Envar("NATS_CREDS").PlaceHolder("NATS_CREDS").StringVar(&opts.Creds)
	ncli.Flag("nkey", "User NKEY").Envar("NATS_NKEY").PlaceHolder("NATS_NKEY").StringVar(&opts.Nkey)
	ncli.Flag("tlscert", "TLS public certificate").Envar("NATS_CERT").PlaceHolder("NATS_CERT").ExistingFileVar(&opts.TlsCert)
	ncli.Flag("tlskey", "TLS private key").Envar("NATS_KEY").PlaceHolder("NATS_KEY").ExistingFileVar(&opts.TlsKey)
	ncli.Flag("tlsca", "TLS certificate authority chain").Envar("NATS_CA").PlaceHolder("NATS_CA").ExistingFileVar(&opts.TlsCA)
	ncli.Flag("timeout", "Time to wait on responses from NATS").Default("5s").Envar("NATS_TIMEOUT").PlaceHolder("NATS_TIMEOUT").DurationVar(&opts.Timeout)
	ncli.Flag("js-api-prefix", "Subject prefix for access to JetStream API").PlaceHolder("PREFIX").StringVar(&opts.JsApiPrefix)
	ncli.Flag("js-event-prefix", "Subject prefix for access to JetStream Advisories").PlaceHolder("PREFIX").StringVar(&opts.JsEventPrefix)
	ncli.Flag("js-domain", "JetStream domain to access").PlaceHolder("PREFIX").PlaceHolder("DOMAIN").StringVar(&opts.JsDomain)
	ncli.Flag("domain", "JetStream domain to access").PlaceHolder("PREFIX").PlaceHolder("DOMAIN").Hidden().StringVar(&opts.JsDomain)
	ncli.Flag("context", "Configuration context").Envar("NATS_CONTEXT").StringVar(&opts.CfgCtx)
	ncli.Flag("trace", "Trace API interactions").BoolVar(&opts.Trace)

	cli.Configure(ncli, opts, true, nil)

	log.SetFlags(log.Ltime)

	kingpin.MustParse(ncli.Parse(os.Args[1:]))
}
