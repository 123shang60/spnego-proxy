package proxy

import (
	"io"
	"net/http"

	"github.com/jcmturner/gokrb5/v8/spnego"
	"github.com/sirupsen/logrus"
)

func DoSpnego() {
	r, _ := http.NewRequest("GET", "http://localhost:9200/_cat/indices", nil)
	defer r.Body.Close()
	spnegoCl := spnego.NewClient(krb5Client, nil, "")
	resp, err := spnegoCl.Do(r)
	if err != nil {
		logrus.Error("spnego 认证失败！")
		logrus.Fatal(err)
	}

	data, _ := io.ReadAll(resp.Body)

	logrus.Info(string(data))
}
