package integration

type PlanAWS struct {
	Etcd                         []AWSNodeDeets
	Master                       []AWSNodeDeets
	Worker                       []AWSNodeDeets
	MasterNodeFQDN               string
	MasterNodeShortName          string
	SSHUser                      string
	SSHKeyFile                   string
	HomeDirectory                string
	AllowPackageInstallation     bool
	AutoConfiguredDockerRegistry bool
}

type AWSNodeDeets struct {
	Instanceid string
	Publicip   string
	Privateip  string
	Hostname   string
}

const planAWSOverlay = `cluster:
  name: kubernetes
  admin_password: abbazabba
  allow_package_installation: {{.AllowPackageInstallation}}
  networking:
    type: overlay
    pod_cidr_block: 172.16.0.0/16
    service_cidr_block: 172.17.0.0/16
    policy_enabled: false
    update_hosts_files: false
  certificates:
    expiry: 17520h
    location_city: Troy
    location_state: New York
    location_country: US
  ssh:
    user: {{.SSHUser}}
    ssh_key: {{.SSHKeyFile}}
    ssh_port: 22
  docker_registry:
    setup_internal: {{.AutoConfiguredDockerRegistry}}
etcd:
  expected_count: {{len .Etcd}}
  nodes:{{range .Etcd}}
  - host: {{.Hostname}}
    ip: {{.Publicip}}
    internalip: {{.Privateip}}{{end}}
master:
  expected_count: {{len .Master}}
  nodes:{{range .Master}}
  - host: {{.Hostname}}
    ip: {{.Publicip}}
    internalip: {{.Privateip}}{{end}}
  load_balanced_fqdn: {{.MasterNodeFQDN}}
  load_balanced_short_name: {{.MasterNodeShortName}}
worker:
  expected_count: {{len .Worker}}
  nodes:{{range .Worker}}
  - host: {{.Hostname}}
    ip: {{.Publicip}}
    internalip: {{.Privateip}}{{end}}
`
