package models

import (
	"errors"
	"fmt"
)

// Vendor represents a supported network vendor.
type Vendor string

const (
	VendorCiscoIOS     Vendor = "cisco_ios"
	VendorCiscoNXOS    Vendor = "cisco_nxos"
	VendorAristaEOS    Vendor = "arista_eos"
	VendorJuniperJunos Vendor = "juniper_junos"
)

// InterfaceType represents the type of a network interface.
type InterfaceType string

const (
	InterfaceTypeEthernet       InterfaceType = "ethernet"
	InterfaceTypeGigabit        InterfaceType = "gigabit"
	InterfaceTypeTenGigabit     InterfaceType = "tengigabit"
	InterfaceTypeFortyGigabit   InterfaceType = "fortygigabit"
	InterfaceTypeHundredGigabit InterfaceType = "hundredgigabit"
	InterfaceTypeLoopback       InterfaceType = "loopback"
	InterfaceTypeVLAN           InterfaceType = "vlan"
	InterfaceTypePortChannel    InterfaceType = "port_channel"
	InterfaceTypeMgmt           InterfaceType = "management"
)

// SwitchportMode represents the switchport mode of an interface.
type SwitchportMode string

const (
	SwitchportModeAccess           SwitchportMode = "access"
	SwitchportModeTrunk            SwitchportMode = "trunk"
	SwitchportModeDynamicAuto      SwitchportMode = "dynamic auto"
	SwitchportModeDynamicDesirable SwitchportMode = "dynamic desirable"
)

// STPMode represents a Spanning Tree Protocol mode.
type STPMode string

const (
	STPModePVST      STPMode = "pvst"
	STPModeRapidPVST STPMode = "rapid-pvst"
	STPModeMST       STPMode = "mst"
)

// ACLAction represents the action in an ACL entry.
type ACLAction string

const (
	ACLActionPermit ACLAction = "permit"
	ACLActionDeny   ACLAction = "deny"
)

// ACLProtocol represents the protocol in an ACL entry.
type ACLProtocol string

const (
	ACLProtocolIP   ACLProtocol = "ip"
	ACLProtocolTCP  ACLProtocol = "tcp"
	ACLProtocolUDP  ACLProtocol = "udp"
	ACLProtocolICMP ACLProtocol = "icmp"
)

// PolicyAction represents a permit or deny action in a policy.
type PolicyAction string

const (
	PolicyActionPermit PolicyAction = "permit"
	PolicyActionDeny   PolicyAction = "deny"
)

// Interface represents a network interface configuration.
type Interface struct {
	Name              string          `json:"name"`
	Type              InterfaceType   `json:"interface_type"`
	Description       string          `json:"description,omitempty"`
	IPAddress         *string         `json:"ip_address,omitempty"`
	SubnetMask        *string         `json:"subnet_mask,omitempty"`
	Enabled           *bool           `json:"enabled,omitempty"`
	Speed             *string         `json:"speed,omitempty"`
	Duplex            *string         `json:"duplex,omitempty"`
	MTU               int             `json:"mtu,omitempty"`
	SwitchportMode    *SwitchportMode `json:"switchport_mode,omitempty"`
	AccessVLAN        *int            `json:"access_vlan,omitempty"`
	VoiceVLAN         *int            `json:"voice_vlan,omitempty"`
	TrunkAllowedVLANs *string         `json:"trunk_allowed_vlans,omitempty"`
	TrunkNativeVLAN   *int            `json:"trunk_native_vlan,omitempty"`
	ChannelGroup      *int            `json:"channel_group,omitempty"`
	ChannelGroupMode  *string         `json:"channel_group_mode,omitempty"`
}

// NewInterface creates a new Interface with default values.
func NewInterface() *Interface {
	enabled := true
	return &Interface{
		Enabled: &enabled,
		MTU:     1500,
	}
}

// Validate checks the interface configuration for correctness.
func (i *Interface) Validate() error {
	if i.AccessVLAN != nil && (*i.AccessVLAN < 1 || *i.AccessVLAN > 4094) {
		return fmt.Errorf("access VLAN must be between 1 and 4094, got %d", *i.AccessVLAN)
	}
	if i.MTU < 64 || i.MTU > 9216 {
		return fmt.Errorf("MTU must be between 64 and 9216, got %d", i.MTU)
	}
	return nil
}

// VLAN represents a VLAN configuration.
type VLAN struct {
	ID          int    `json:"vlan_id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	State       string `json:"state,omitempty"`
}

// NewVLAN creates a new VLAN with default values.
func NewVLAN() *VLAN {
	return &VLAN{
		State: "active",
	}
}

// Validate checks the VLAN configuration for correctness.
func (v *VLAN) Validate() error {
	if v.ID < 1 || v.ID > 4094 {
		return fmt.Errorf("VLAN ID must be between 1 and 4094, got %d", v.ID)
	}
	if v.Name == "" {
		return fmt.Errorf("VLAN name cannot be empty")
	}
	return nil
}

// ACLEntry represents a single ACL entry/rule.
type ACLEntry struct {
	Sequence            int         `json:"sequence"`
	Action              ACLAction   `json:"action"`
	Protocol            ACLProtocol `json:"protocol"`
	Source              string      `json:"source"`
	SourceWildcard      string      `json:"source_wildcard,omitempty"`
	Destination         string      `json:"destination,omitempty"`
	DestinationWildcard string      `json:"destination_wildcard,omitempty"`
	SourcePort          *string     `json:"source_port,omitempty"`
	DestinationPort     *string     `json:"destination_port,omitempty"`
	Log                 bool        `json:"log,omitempty"`
	Remark              string      `json:"remark,omitempty"`
}

// NewACLEntry creates a new ACLEntry with default values.
func NewACLEntry() *ACLEntry {
	return &ACLEntry{
		SourceWildcard:      "0.0.0.0",
		Destination:         "any",
		DestinationWildcard: "0.0.0.0",
	}
}

// Validate checks the ACL entry for correctness.
func (e *ACLEntry) Validate() error {
	if e.Sequence <= 0 {
		return fmt.Errorf("ACL entry sequence must be a positive integer, got %d", e.Sequence)
	}
	return nil
}

// ACL represents an Access Control List.
type ACL struct {
	Name       string     `json:"name"`
	Entries    []ACLEntry `json:"entries,omitempty"`
	IsExtended *bool      `json:"is_extended,omitempty"`
}

// NewACL creates a new ACL with default values.
func NewACL() *ACL {
	isExtended := true
	return &ACL{
		IsExtended: &isExtended,
	}
}

// Validate checks the ACL and all its entries for correctness.
func (a *ACL) Validate() error {
	for _, entry := range a.Entries {
		if err := entry.Validate(); err != nil {
			return fmt.Errorf("invalid ACL entry (sequence %d): %w", entry.Sequence, err)
		}
	}
	return nil
}

// StaticRoute represents a static route configuration.
type StaticRoute struct {
	Destination   string `json:"destination"`
	Mask          string `json:"mask"`
	NextHop       string `json:"next_hop"`
	AdminDistance int    `json:"admin_distance,omitempty"`
	Name          string `json:"name,omitempty"`
	Permanent     bool   `json:"permanent,omitempty"`
}

// NewStaticRoute creates a new StaticRoute with default values.
func NewStaticRoute() *StaticRoute {
	return &StaticRoute{
		AdminDistance: 1,
	}
}

// Validate checks the static route configuration for correctness.
func (sr *StaticRoute) Validate() error {
	if sr.AdminDistance < 1 || sr.AdminDistance > 255 {
		return fmt.Errorf("admin distance must be 1-255, got %d", sr.AdminDistance)
	}
	return nil
}

// OSPFNetwork represents an OSPF network statement.
type OSPFNetwork struct {
	Network  string `json:"network"`
	Wildcard string `json:"wildcard"`
	Area     int    `json:"area"`
}

// OSPFConfig represents OSPF configuration.
type OSPFConfig struct {
	ProcessID                   int           `json:"process_id"`
	RouterID                    *string       `json:"router_id,omitempty"`
	Networks                    []OSPFNetwork `json:"networks,omitempty"`
	PassiveInterfaces           []string      `json:"passive_interfaces,omitempty"`
	DefaultInformationOriginate bool          `json:"default_information_originate,omitempty"`
	ReferenceBandwidth          int           `json:"reference_bandwidth,omitempty"`
}

// NewOSPFConfig creates a new OSPFConfig with default values.
func NewOSPFConfig() *OSPFConfig {
	return &OSPFConfig{
		ReferenceBandwidth: 100,
	}
}

// Validate checks the OSPF configuration for correctness.
func (o *OSPFConfig) Validate() error {
	if o.ProcessID < 0 || o.ProcessID > 65535 {
		return fmt.Errorf("OSPF process ID must be 0-65535, got %d", o.ProcessID)
	}
	return nil
}

// BGPNeighbor represents a BGP neighbor configuration.
type BGPNeighbor struct {
	IPAddress    string  `json:"ip_address"`
	RemoteAS     uint32  `json:"remote_as"`
	Description  string  `json:"description,omitempty"`
	Password     *string `json:"password,omitempty"`
	UpdateSource *string `json:"update_source,omitempty"`
	EBGPmultihop int     `json:"ebgp_multihop,omitempty"`
	RouteMapIn   *string `json:"route_map_in,omitempty"`
	RouteMapOut  *string `json:"route_map_out,omitempty"`
}

// Validate checks the BGP neighbor for correctness.
func (n *BGPNeighbor) Validate() error {
	if n.RemoteAS == 0 {
		return fmt.Errorf("BGP remote AS must be 1-4294967295, got %d", n.RemoteAS)
	}
	if n.IPAddress == "" {
		return fmt.Errorf("BGP neighbor IP address cannot be empty")
	}
	return nil
}

// BGPConfig represents BGP configuration.
type BGPConfig struct {
	LocalAS            uint32        `json:"local_as"`
	RouterID           *string       `json:"router_id,omitempty"`
	Neighbors          []BGPNeighbor `json:"neighbors,omitempty"`
	Networks           []string      `json:"networks,omitempty"`
	LogNeighborChanges *bool         `json:"log_neighbor_changes,omitempty"`
	Redistribute       []string      `json:"redistribute,omitempty"`
}

// NewBGPConfig creates a new BGPConfig with default values.
func NewBGPConfig() *BGPConfig {
	logChanges := true
	return &BGPConfig{
		LogNeighborChanges: &logChanges,
	}
}

// Validate checks the BGP configuration for correctness.
func (b *BGPConfig) Validate() error {
	if b.LocalAS == 0 {
		return fmt.Errorf("BGP AS must be 1-4294967295, got %d", b.LocalAS)
	}
	for _, neighbor := range b.Neighbors {
		if err := neighbor.Validate(); err != nil {
			return fmt.Errorf("invalid BGP neighbor (%s): %w", neighbor.IPAddress, err)
		}
	}
	return nil
}

// EIGRPNetwork represents an EIGRP network statement.
type EIGRPNetwork struct {
	Network  string  `json:"network"`
	Wildcard *string `json:"wildcard,omitempty"`
}

// EIGRPConfig represents EIGRP configuration.
type EIGRPConfig struct {
	ASNumber          int            `json:"as_number"`
	RouterID          *string        `json:"router_id,omitempty"`
	Networks          []EIGRPNetwork `json:"networks,omitempty"`
	PassiveInterfaces []string       `json:"passive_interfaces,omitempty"`
	AutoSummary       bool           `json:"auto_summary,omitempty"`
	Redistribute      []string       `json:"redistribute,omitempty"`
	NamedMode         bool           `json:"named_mode,omitempty"`
	Name              string         `json:"name,omitempty"`
}

// NewEIGRPConfig creates a new EIGRPConfig with default values.
func NewEIGRPConfig() *EIGRPConfig {
	return &EIGRPConfig{
		Name: "EIGRP_PROCESS",
	}
}

// Validate checks the EIGRP configuration for correctness.
func (e *EIGRPConfig) Validate() error {
	if e.ASNumber < 1 || e.ASNumber > 65535 {
		return fmt.Errorf("EIGRP AS must be 1-65535, got %d", e.ASNumber)
	}
	return nil
}

// PrefixListEntry represents a single prefix-list entry.
type PrefixListEntry struct {
	Sequence int          `json:"sequence"`
	Action   PolicyAction `json:"action"`
	Prefix   string       `json:"prefix"`
	GE       *int         `json:"ge,omitempty"`
	LE       *int         `json:"le,omitempty"`
}

// Validate checks the prefix-list entry for correctness.
func (e *PrefixListEntry) Validate() error {
	if e.Sequence <= 0 {
		return fmt.Errorf("prefix-list entry sequence must be a positive integer, got %d", e.Sequence)
	}
	if e.Prefix == "" {
		return fmt.Errorf("prefix-list entry prefix cannot be empty")
	}
	if e.GE != nil && e.LE != nil && *e.LE < *e.GE {
		return fmt.Errorf("less-than-or-equal value (%d) cannot be less than greater-than-or-equal value (%d)", *e.LE, *e.GE)
	}
	return nil
}

// PrefixList represents an IP prefix-list for route filtering.
type PrefixList struct {
	Name    string            `json:"name"`
	Entries []PrefixListEntry `json:"entries,omitempty"`
}

// Validate checks the PrefixList and all its entries for correctness.
func (pl *PrefixList) Validate() error {
	for _, entry := range pl.Entries {
		if err := entry.Validate(); err != nil {
			return fmt.Errorf("invalid PrefixList entry (sequence %d): %w", entry.Sequence, err)
		}
	}
	return nil
}

// RouteMapEntry represents a single route-map entry.
type RouteMapEntry struct {
	Sequence         int          `json:"sequence"`
	Action           PolicyAction `json:"action"`
	MatchPrefixList  *string      `json:"match_prefix_list,omitempty"`
	MatchASPath      *string      `json:"match_as_path,omitempty"`
	MatchCommunity   *string      `json:"match_community,omitempty"`
	SetLocalPref     *int         `json:"set_local_pref,omitempty"`
	SetMED           *int         `json:"set_med,omitempty"`
	SetASPathPrepend *string      `json:"set_as_path_prepend,omitempty"`
	SetCommunity     *string      `json:"set_community,omitempty"`
	SetNextHop       *string      `json:"set_next_hop,omitempty"`
	SetWeight        *int         `json:"set_weight,omitempty"`
}

// RouteMap represents a route-map for policy-based routing.
type RouteMap struct {
	Name    string          `json:"name"`
	Entries []RouteMapEntry `json:"entries,omitempty"`
}

// Validate checks the RouteMap and all its entries for correctness.
func (rm *RouteMap) Validate() error {
	for _, entry := range rm.Entries {
		if entry.Sequence <= 0 {
			return fmt.Errorf("route-map entry sequence must be a positive integer, got %d", entry.Sequence)
		}
	}
	return nil
}

// STPConfig represents Spanning Tree Protocol configuration.
type STPConfig struct {
	Mode               STPMode `json:"mode,omitempty"`
	Priority           int     `json:"priority,omitempty"`
	RootPrimaryVLANs   []int   `json:"root_primary_vlans,omitempty"`
	RootSecondaryVLANs []int   `json:"root_secondary_vlans,omitempty"`
	PortfastDefault    bool    `json:"portfast_default,omitempty"`
	BPDUGuardDefault   bool    `json:"bpduguard_default,omitempty"`
}

// NewSTPConfig creates a new STPConfig with default values.
func NewSTPConfig() *STPConfig {
	return &STPConfig{
		Mode:     STPModeRapidPVST,
		Priority: 32768,
	}
}

// Validate checks the STP configuration for correctness.
func (c *STPConfig) Validate() error {
	// Currently no validation rules, but method exists for consistency.
	return nil
}

// DeviceConfig represents a complete device configuration.
type DeviceConfig struct {
	Hostname     string        `json:"hostname"`
	Vendor       Vendor        `json:"vendor"`
	Interfaces   []Interface   `json:"interfaces,omitempty"`
	VLANs        []VLAN        `json:"vlans,omitempty"`
	ACLs         []ACL         `json:"acls,omitempty"`
	StaticRoutes []StaticRoute `json:"static_routes,omitempty"`
	OSPF         *OSPFConfig   `json:"ospf,omitempty"`
	EIGRP        *EIGRPConfig  `json:"eigrp,omitempty"`
	BGP          *BGPConfig    `json:"bgp,omitempty"`
	STP          *STPConfig    `json:"stp,omitempty"`
	PrefixLists  []PrefixList  `json:"prefix_lists,omitempty"`
	RouteMaps    []RouteMap    `json:"route_maps,omitempty"`
	EnableSecret *string       `json:"enable_secret,omitempty"`
	DomainName   *string       `json:"domain_name,omitempty"`
	DNSServers   []string      `json:"dns_servers,omitempty"`
	NTPServers   []string      `json:"ntp_servers,omitempty"`
	BannerMOTD   string        `json:"banner_motd,omitempty"`
}

// Validate checks the entire device configuration for correctness.
func (c *DeviceConfig) Validate() error {
	var errs []error
	for _, iface := range c.Interfaces {
		if err := iface.Validate(); err != nil {
			errs = append(errs, fmt.Errorf("invalid interface (%s): %w", iface.Name, err))
		}
	}
	for _, vlan := range c.VLANs {
		if err := vlan.Validate(); err != nil {
			errs = append(errs, fmt.Errorf("invalid VLAN (%d): %w", vlan.ID, err))
		}
	}
	for _, acl := range c.ACLs {
		if err := acl.Validate(); err != nil {
			errs = append(errs, fmt.Errorf("invalid ACL (%s): %w", acl.Name, err))
		}
	}
	for _, route := range c.StaticRoutes {
		if err := route.Validate(); err != nil {
			errs = append(errs, fmt.Errorf("invalid static route (%s): %w", route.Destination, err))
		}
	}
	for _, pl := range c.PrefixLists {
		if err := pl.Validate(); err != nil {
			errs = append(errs, fmt.Errorf("invalid prefix-list (%s): %w", pl.Name, err))
		}
	}
	for _, rm := range c.RouteMaps {
		if err := rm.Validate(); err != nil {
			errs = append(errs, fmt.Errorf("invalid route-map (%s): %w", rm.Name, err))
		}
	}
	if c.OSPF != nil {
		if err := c.OSPF.Validate(); err != nil {
			errs = append(errs, fmt.Errorf("invalid OSPF config: %w", err))
		}
	}
	if c.EIGRP != nil {
		if err := c.EIGRP.Validate(); err != nil {
			errs = append(errs, fmt.Errorf("invalid EIGRP config: %w", err))
		}
	}
	if c.BGP != nil {
		if err := c.BGP.Validate(); err != nil {
			errs = append(errs, fmt.Errorf("invalid BGP config: %w", err))
		}
	}
	if c.STP != nil {
		if err := c.STP.Validate(); err != nil {
			errs = append(errs, fmt.Errorf("invalid STP config: %w", err))
		}
	}
	return errors.Join(errs...)
}
