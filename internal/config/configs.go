package config

type Config struct {
	Porxy struct {
		TargetUrl string
	}
	Auth struct {
		KeyTabPath         string
		KerberosConfigPath string
		ServiceName        string
		Realm              string
		DisablePAFXFAST    bool
		SPNHostsMapping    map[string]string
	}
}
