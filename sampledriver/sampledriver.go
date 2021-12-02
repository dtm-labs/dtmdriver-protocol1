package sample

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/yedf/dtmdriver"
	"google.golang.org/grpc/resolver"
)

// 这里的sampleDriver，演示了一个的driver应当如何编写

type sampleDriver struct {
}

func (d *sampleDriver) GetName() string {
	return "dtm-sample-driver"
}

func (d *sampleDriver) RegisterGrpcService(url string, endpoint string) error {
	// 如果使用etcd，polaris等注册/发现组件的话，那么在这里，将endpoint注册到相应的url中
	// 这里的sample仅作为演示用，没有实际注册
	return nil
}

func (d *sampleDriver) RegisterGrpcResolver() {
	resolver.Register(&sampleBuilder{})
}

func (d *sampleDriver) ParseServerMethod(uri string) (server string, method string, err error) {
	if !strings.Contains(uri, "//") { // 处理无scheme的情况，如果您没有直连，可以不处理
		sep := strings.IndexByte(uri, '/')
		if sep == -1 {
			return "", "", fmt.Errorf("bad url: '%s'. no '/' found", uri)
		}
		return uri[:sep], uri[sep:], nil

	}
	u, err := url.Parse(uri)
	if err != nil {
		return "", "", err
	}
	return u.Scheme + "://" + u.Host, u.Path, nil
}

func init() {
	dtmdriver.Register(&sampleDriver{})
}
