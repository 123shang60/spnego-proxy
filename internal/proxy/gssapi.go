package proxy

import (
	"log"

	"github.com/jcmturner/gokrb5/v8/client"
	"github.com/jcmturner/gokrb5/v8/config"
	"github.com/jcmturner/gokrb5/v8/keytab"
	"github.com/sirupsen/logrus"
)

var krb5Client *client.Client

func InitKrb5Cli() {
	cfg, err := config.Load("/etc/security/xh_krb5.conf")
	if err != nil {
		logrus.Error("加载 krb5.conf 失败！")
		logrus.Fatal(err)
	}

	ktFile, err := keytab.Load("/root/demo.keytab")
	if err != nil {
		logrus.Error("加载 /root/user.keytab 失败！")
		logrus.Fatal(err)
	}

	l := log.New(logrus.StandardLogger().Out, "GOKRB5 Client: ", log.Ldate|log.Ltime|log.Lshortfile)

	krb5Client = client.NewWithKeytab("noah", "HADOOP.COM", ktFile, cfg, client.Logger(l))

	err = krb5Client.Login()
	if err != nil {
		logrus.Error("kerberos 认证失败！")
		logrus.Fatal(err)
	}
}
