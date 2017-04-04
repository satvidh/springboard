package main

import (
	"context"
	"flag"
	"fmt"

	b64 "encoding/base64"
	"github.com/google/subcommands"
	vaultapi "github.com/hashicorp/vault/api"
)
type setCmd struct {
	secrets string
	vault string
	path string
	key string
	value string
}

func (*setCmd) Name() string     { return "set" }
func (*setCmd) Synopsis() string { return "Encrypt a value and set it to the specified key" }
func (*setCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *setCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.secrets, "secrets", "", "path to secrets file")
	f.StringVar(&p.vault, "vault", "", "vault host address")
	f.StringVar(&p.path, "path", "", "vault path")
	f.StringVar(&p.key, "key", "", "secret key")
	f.StringVar(&p.value, "value", "", "plain text secret")

}

func (p *setCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
    enc := b64.StdEncoding.EncodeToString([]byte(p.value))


	cfg := vaultapi.DefaultConfig()
	cfg.Address = p.vault
	client, err := vaultapi.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	client.SetToken("horde")
	c := client.Logical()

	v, err := c.Write("transit/encrypt/cub",
	    map[string]interface{}{
			"plaintext": enc,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", v.Data["ciphertext"])

	return subcommands.ExitSuccess
}
