package config

type Config struct {
	Porxy struct {
		TargetUrl string
	}
	Auth struct {
		KeyTabPath         string
		KerberosConfigPath string
		ServiceName        string
		UserName           string
		Realm              string
		DisablePAFXFAST    bool
		SPNHostsMapping    map[string]string
	}
	Log struct {
		Level string
	}
	Server struct {
		Port             int32
		EnablePprof      bool
		EnablePrometheus bool
	}
}

var C = new(Config)
