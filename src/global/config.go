package global

var (
	Conf Config
)

type Config struct {
	Service struct {
		LogFile string `flag:"logfile" cfg:"logfile" toml:"logfile"`
	}
	Http struct {
		Port      int `flag:"port" cfg:"port" toml:"port"`
		PprofPort int `flag:"pprofport" cfg:"pprofport" toml:"pprofport"`
	}
}
