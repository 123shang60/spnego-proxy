package proxy

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/123shang60/spnego-proxy/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/jcmturner/gokrb5/v8/spnego"
	"github.com/sirupsen/logrus"
)

var (
	spn          string
	loadIconOnce sync.Once
)

func getSpn() {
	host := strings.SplitN(strings.SplitN(config.C.Porxy.TargetUrl, "://", 2)[1], ":", 2)[0]

	if h, ok := config.C.Auth.SPNHostsMapping[host]; ok && h != "" {
		host = h
	}

	spn = fmt.Sprintf("%s/%s", config.C.Auth.ServiceName, host)

	logrus.Debugf("生成认证使用的 SPN : %s", spn)
}

func DoSpnego(c *gin.Context) {
	logrus.Debugf("收到请求- Method: %s, Path: %s", c.Request.Method, c.Request.URL.Path)
	loadIconOnce.Do(getSpn)

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
	url := fmt.Sprintf("%s/%s", config.C.Porxy.TargetUrl, strings.TrimPrefix(c.Request.URL.Path, "/"))
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

	r.Header = c.Request.Header.Clone()
	r.Header.Del("Authorization")

	logrus.Debugf("本次请求的 SPN: %s", spn)
	spnegoCl := spnego.NewClient(krb5Client, client, spn)
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
	logrus.Debugf("请求返回- Path: %s, StatusCode: %d, Body: %s", c.Request.URL.Path, resp.StatusCode, string(data))
	c.Writer.WriteHeader(resp.StatusCode)
	c.Writer.Write(data)
}

func HandleError(c *gin.Context, err error) {
	errMsg := Resperr{Msg: err.Error()}

	c.JSON(http.StatusInternalServerError, errMsg)
}
