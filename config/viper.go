package config

import(
	"io"
	"strings"

	"github.com/spf13/viper"
)

// SetupViper will setup viper for reading env vars and accepts an optional
// io.Reader to read from a config file.
func SetupViper(cfgfile io.Reader) *viper.Viper {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".","_"))
	if cfgfile != nil {
		v.ReadConfig(cfgfile)
	}
	return v
}
