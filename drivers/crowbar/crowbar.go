package crowbar  // http://github.com/opencrowbar/core

import (
	"github.com/docker/machine/drivers"
	"github.com/docker/machine/provider"
	"github.com/docker/machine/state"
)

type CrowbarDriver struct {
	NodeRaw		string
	URL 		string	// default 192.168.124.10
	Username	string  // default "crowbar"
	Password	string  // default "crowbar"
	SourcePool	string  // default "system"
	TargetPool	string  // default "docker-machines"
	ReadyState	string  // default "docker-ready"
}

func init() {
	drivers.Register("crowbar", &drivers.RegisteredDriver{
		New:            NewDriver,
		GetCreateFlags: GetCreateFlags,
	})
}

func NewDriver(url: string, username: string, password: string, source: string, target: string, state: string) (drivers.Driver, error) {
    if url == "" {
    	url = "http://192.168.124.10"
  	}
  	if username == "" {
  		username = "crowbar"
  	}
  	if password == "" {
  		password = "crowbar"
  	}
  	if source == "" {
  		source = "system"
  	}
  	if target == "" {
  		target = "docker-machines"
  	}
  	if state == "" {
  		state = "docker-machine-ready"
  	}
	return &CrowbarDriver{NodeState: state.None, URL: url, Username: username, Password: password, SourcePool: source, TargetPool: target, ReadyState: state}, nil
}

func (d *CrowbarDriver) DriverName() string {
	return "crowbar"
}

func (d *CrowbarDriver) AuthorizePort(ports []*drivers.Port) error {
	return nil
}

func (d *CrowbarDriver) DeauthorizePort(ports []*drivers.Port) error {
	return nil
}

func (d *CrowbarDriver) SetConfigFromFlags(flags drivers.DriverOptions) error {
	return nil
}

func (d *CrowbarDriver) GetURL() (string, error) {
	ip, err := d.NodeAddress()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("tcp://%s:2376", ip), nil
}

func (d *CrowbarDriver) GetIP() (string, error) {
	if d.NodeAddress == "" {
		return "", fmt.Errorf("IP address is not set")
	}
	return d.NodeAddress, nil
}

func (d *CrowbarDriver) GetMachineName() string {
	return d.NodeName
}

func (d *CrowbarDriver) GetProviderType() provider.ProviderType {
	return provider.Remote
}

func (d *CrowbarDriver) GetSSHHostname() (string, error) {
	return d.GetIP()
}

func (d *CrowbarDriver) GetSSHKeyPath() string {
	return filepath.Join(d.storePath, "id_rsa")
}

func (d *CrowbarDriver) GetSSHPort() (int, error) {
	return 22, nil
}

func (d *CrowbarDriver) GetSSHUsername() string {
	return "root"
}

func (d *CrowbarDriver) GetState() (state.State, error) {

	return d.NodeState, nil
}

func (d *CrowbarDriver) PreCreateCheck() error {
	// does Pool have enough machines?
	return nil
}

func (d *CrowbarDriver) Create() error {
	// get first machine from pool
	// move machine into target pool
	// assign machine to ready state
	// commit machine
	return nil
}

func (d *CrowbarDriver) Remove() error {
	return nil
}

func (d *CrowbarDriver) Start() error {
	// machine power on
	d.NodeState = state.Running
	return nil
}

func (d *CrowbarDriver) Stop() error {
	// machine power off
	d.NodeState = state.Stopped
	return nil
}

func (d *CrowbarDriver) Restart() error {
	return nil
}

func (d *CrowbarDriver) Kill() error {
	return nil
}

func (d *CrowbarDriver) Upgrade() error {
	return nil
}

func (d *CrowbarDriver) StartDocker() error {
	return nil
}

func (d *CrowbarDriver) StopDocker() error {
	return nil
}

func (d *CrowbarDriver) GetDockerConfigDir() string {
	return ""
}

func (d *CrowbarDriver) GetSSHCommand(args ...string) (*exec.Cmd, error) {
	return &exec.Cmd{}, nil
}
