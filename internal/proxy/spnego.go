package proxy

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/123shang60/spnego-proxy/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/jcmturner/gokrb5/v8/spnego"
	"github.com/sirupsen/logrus"
)

func DoSpnego(c *gin.Context) {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       600 * time.Second,
		TLSHandshakeTimeout:   600 * time.Second,
		ExpectContinueTimeout: 600 * time.Second,
		MaxIdleConnsPerHost:   100,
	}
	client := &http.Client{Transport: transport}
	url := fmt.Sprintf("%s/%s", config.C.Porxy.TargetUrl, c.Request.URL.Path)
	if c.Request.URL.RawQuery != "" {
		url += "?" + c.Request.URL.RawQuery
	}

	r, err := http.NewRequest(c.Request.Method, url, c.Request.Body)
	if err != nil {
		logrus.Error("创建请求失败！", err)
		HandleError(c, err)
		return
	}
	defer r.Body.Close()

	host := strings.SplitN(strings.SplitN(config.C.Porxy.TargetUrl, "://", 2)[0], ":", 2)[0]

	if h, ok := config.C.Auth.SPNHostsMapping[host]; ok && h != "" {
		host = h
	}

	logrus.Debugf("认证使用的 SPN : %s", host)

	spnegoCl := spnego.NewClient(krb5Client, client, host)
	resp, err := spnegoCl.Do(r)
	if err != nil {
		logrus.Error("spnego 认证失败！", err)
		HandleError(c, err)
		return
	}
	defer resp.Body.Close()
	for k, v := range resp.Header {
		for _, vv := range v {
			c.Writer.Header().Add(k, vv)
		}
	}

	data, _ := io.ReadAll(resp.Body)
	c.Writer.WriteHeader(resp.StatusCode)
	c.Writer.Write(data)
}

func HandleError(c *gin.Context, err error) {
	errMsg := Resperr{Msg: err.Error()}

	c.JSON(http.StatusInternalServerError, errMsg)
}
