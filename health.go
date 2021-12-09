package consuldiscovery

// Health is a set of functions for the health of services
type Health interface {
	HealthByNode(nodeName string) ([]HealthCheck, error)
	HealthByService(serviceName string) (HealthNodes, error)
	HealthByState(state string) ([]HealthCheck, error)
}

// [{"Node":{"Node":"skailhq.local","Address":"192.168.50.1"},
//  "Service":{"ID":"simple_service","Service":"simple_service","Tags":["tag1","tag2"],"Port":6666},
//  "Checks":[{"Node":"skailhq.local","CheckID":"serfHealth","Name":"Serf Health Status","Status":"passing","Notes":"","Output":"","ServiceID":"","ServiceName":""}]}]

// HealthNodes summarizes the health checks for all Nodes for a single Service
type HealthNodes []HealthForNode

// HealthForNode summarizes the health checks for a single Nodes for a single Service
type HealthForNode struct {
	Node    HealthNode
	Service HealthService
	Checks  []HealthCheck
}

// HealthNode indicates a server/node being described by HealthNodes
type HealthNode struct {
	Node    string
	Address string

	ID              string
	Datacenter      string
	TaggedAddresses TaggedAddressesNode
}

// HealthService indicates a service being described by HealthNodes
type HealthService struct {
	ServiceID              string                    `json:"ID"`
	ServiceName            string                    `json:"Service"`
	ServiceTags            []string                  `json:"Tags"`
	ServicePort            uint64                    `json:"Port"`
	ServiceAddress         string                    `json:"Address"`
	ServiceTaggedAddresses map[string]AddressService `json:"TaggedAddresses"`
	ServiceMeta            map[string]string         `json:"Meta"`
}

// "TaggedAddressesNode":{"lan":{"Address":"192.168.0.55","Port":8000},"wan":{"Address":"198.18.0.23","Port":80}}

type TaggedAddressesNode struct {
	Lan     string `json:"lan"`
	LanIpv4 string `json:"lan_ipv4"`
	Wan     string `json:"wan"`
	WanIpv4 string `json:"wan_ipv4"`
}

type AddressService struct {
	Address string
	Port    int
}

// HealthCheck contains a current health check result
type HealthCheck struct {
	Node        string
	CheckID     string
	Name        string
	Status      string
	Notes       string
	Output      string
	ServiceID   string
	ServiceName string
	ServiceTags []string
	Type        string
	Namespace   string
	Partition   string
}

// HealthByNode returns the health checks for a specific node
func (c *Client) HealthByNode(nodeName string) (result []HealthCheck, err error) {
	err = c.doGET("health/node/"+nodeName, &result)
	return
}

// HealthByService returns a list of advertised service names and their tags
func (c *Client) HealthByService(serviceName string) (result HealthNodes, err error) {
	err = c.doGET("health/service/"+serviceName, &result)
	return
}

// HealthByState returns the health checks with a specific state
func (c *Client) HealthByState(state string) (checks []HealthCheck, err error) {
	err = c.doGET("health/state/"+state, &checks)
	return
}
