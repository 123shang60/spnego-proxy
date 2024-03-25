package server

import (
	"github.com/123shang60/spnego-proxy/internal/common"
	"github.com/123shang60/spnego-proxy/internal/config"
	"github.com/spf13/cobra"
)

var c = new(config.Config)

var Server = &cobra.Command{
	Use:   "server",
	Short: "Run the spnego-proxy server",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		common.SetLogger()

		return nil
	},
	Run: Run,
}

func init() {
	Server.Flags().StringVar(&c.Porxy.TargetUrl, "target-url", "", "The target URL to proxy requests to")
	Server.Flags().StringVar(&c.Auth.KeyTabPath, "keytab-path", "", "The path to the keytab file")
	Server.Flags().StringVar(&c.Auth.KerberosConfigPath, "kerberos-config-path", "/etc/krb5.conf", "The path to the kerberos config file")
	Server.Flags().StringVar(&c.Auth.ServiceName, "servicename", "HTTP", "The service name")
	Server.Flags().StringVar(&c.Auth.Realm, "realm", "", "The realm")
	Server.Flags().BoolVar(&c.Auth.DisablePAFXFAST, "disable-pafx-fast", false, "Disable the use of PA-FX-FAST")
	Server.Flags().StringToStringVar(&c.Auth.SPNHostsMapping, "spn-hosts-mapping", nil, "A mapping of SPNs to hosts")
}

func Run(_ *cobra.Command, _ []string) {

}
