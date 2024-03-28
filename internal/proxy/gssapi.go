package proxy

import (
	"log"

	serverConfig "github.com/123shang60/spnego-proxy/internal/config"
	"github.com/jcmturner/gokrb5/v8/client"
	"github.com/jcmturner/gokrb5/v8/config"
	"github.com/jcmturner/gokrb5/v8/keytab"
	"github.com/sirupsen/logrus"
)

var krb5Client *client.Client

func InitKrb5Cli() error {
	cfg, err := config.Load(serverConfig.C.Auth.KerberosConfigPath)
	if err != nil {
		logrus.Error("加载 krb5.conf 失败！", err)
		return err
	}

	ktFile, err := keytab.Load(serverConfig.C.Auth.KeyTabPath)
	if err != nil {
		logrus.Error("加载 keytab 文件失败！", err)
		return err
	}

	var l *log.Logger

	if logrus.GetLevel() == logrus.DebugLevel {
		l = log.New(logrus.StandardLogger().Out, "GOKRB5 Client: ", log.Ldate|log.Ltime|log.Lshortfile)
	}

	krb5Client = client.NewWithKeytab(serverConfig.C.Auth.UserName, serverConfig.C.Auth.Realm, ktFile, cfg, client.DisablePAFXFAST(serverConfig.C.Auth.DisablePAFXFAST), client.Logger(l))

	err = krb5Client.Login()
	if err != nil {
		logrus.Error("kerberos 认证失败！", err)
		return err
	}

	logrus.Info("kerberos 认证成功！client 缓存已建立！")

	return nil
}
