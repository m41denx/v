package structs

import octoconfig "github.com/nothinux/octo-proxy/pkg/config"

type NodeConfig struct {
	HostIP   string `json:"host_ip"`
	DockerIP string `json:"docker_ip"`
	CaCert   string `json:"ca_cert"`
	Cert     string `json:"cert"`
	CertKey  string `json:"cert_key"`

	OctoConfig    octoconfig.Config `json:"-"`
	SidecarConfig string            `json:"sidecar_config"` // base64(yaml)
}
