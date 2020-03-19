// Code generated by client-gen. DO NOT EDIT.

package v1beta2

import (
	v1beta2 "github.com/lyft/flinkk8soperator/pkg/apis/app/v1beta2"
	"github.com/lyft/flinkk8soperator/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type FlinkV1beta2Interface interface {
	RESTClient() rest.Interface
	FlinkApplicationsGetter
}

// FlinkV1beta2Client is used to interact with features provided by the flink.k8s.io group.
type FlinkV1beta2Client struct {
	restClient rest.Interface
}

func (c *FlinkV1beta2Client) FlinkApplications(namespace string) FlinkApplicationInterface {
	return newFlinkApplications(c, namespace)
}

// NewForConfig creates a new FlinkV1beta2Client for the given config.
func NewForConfig(c *rest.Config) (*FlinkV1beta2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &FlinkV1beta2Client{client}, nil
}

// NewForConfigOrDie creates a new FlinkV1beta2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *FlinkV1beta2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new FlinkV1beta2Client for the given RESTClient.
func New(c rest.Interface) *FlinkV1beta2Client {
	return &FlinkV1beta2Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta2.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FlinkV1beta2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}