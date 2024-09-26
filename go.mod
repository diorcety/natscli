module github.com/nats-io/natscli

go 1.21.0

toolchain go1.21.13

require (
	github.com/AlecAivazis/survey/v2 v2.3.7
	github.com/HdrHistogram/hdrhistogram-go v1.1.2
	github.com/choria-io/fisk v0.6.2
	github.com/dustin/go-humanize v1.0.1
	github.com/emicklei/dot v1.6.2
	github.com/expr-lang/expr v1.16.9
	github.com/fatih/color v1.17.0
	github.com/ghodss/yaml v1.0.0
	github.com/google/go-cmp v0.6.0
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510
	github.com/gosuri/uiprogress v0.0.1
	github.com/guptarohit/asciigraph v0.7.1
	github.com/jedib0t/go-pretty/v6 v6.5.9
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51
	github.com/klauspost/compress v1.17.9
	github.com/mattn/go-isatty v0.0.20
	github.com/nats-io/jsm.go v0.1.2
	github.com/nats-io/jwt/v2 v2.5.8
	github.com/nats-io/nats-server/v2 v2.10.18
	github.com/nats-io/nats.go v1.36.0
	github.com/nats-io/nkeys v0.4.7
	github.com/nats-io/nuid v1.0.1
	github.com/prometheus/client_golang v1.19.1
	github.com/prometheus/common v0.55.0
	github.com/santhosh-tekuri/jsonschema/v5 v5.3.1
	github.com/tylertreat/hdrhistogram-writer v0.0.0-20210816161836-2e440612a39f
	golang.org/x/crypto v0.27.0
	golang.org/x/exp v0.0.0-20240719175910-8a7402abbf56
	golang.org/x/term v0.24.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/gosuri/uilive v0.0.4 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/minio/highwayhash v1.0.3 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	golang.org/x/time v0.6.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/nats-io/jsm.go => github.com/diorcety/jsm.go v0.1.3-0.20240926144904-80e3c7361679

replace github.com/nats-io/nats-server/v2 => github.com/diorcety/nats-server/v2 v2.0.0-20240926143622-cf9f2e51b94c
