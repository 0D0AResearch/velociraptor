// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	proto1 "www.velocidex.com/golang/velociraptor/actions/proto"
	_ "www.velocidex.com/golang/velociraptor/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Information about Velociraptor. This is a psuedo config item - it
// will never be read from config files, but can be used in VQL
// queries. It will be written to config files as metadata about the
// version read from.
type Version struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Commit               string   `protobuf:"bytes,4,opt,name=commit,proto3" json:"commit,omitempty"`
	BuildTime            string   `protobuf:"bytes,5,opt,name=build_time,json=buildTime,proto3" json:"build_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{0}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Version) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Version) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Version) GetCommit() string {
	if m != nil {
		return m.Commit
	}
	return ""
}

func (m *Version) GetBuildTime() string {
	if m != nil {
		return m.BuildTime
	}
	return ""
}

// The client's state which are persisted in the writeback file.
type Writeback struct {
	PrivateKey             string   `protobuf:"bytes,7,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	HuntLastTimestamp      uint64   `protobuf:"varint,13,opt,name=hunt_last_timestamp,json=huntLastTimestamp,proto3" json:"hunt_last_timestamp,omitempty"`
	LastServerSerialNumber uint64   `protobuf:"varint,14,opt,name=last_server_serial_number,json=lastServerSerialNumber,proto3" json:"last_server_serial_number,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Writeback) Reset()         { *m = Writeback{} }
func (m *Writeback) String() string { return proto.CompactTextString(m) }
func (*Writeback) ProtoMessage()    {}
func (*Writeback) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{1}
}

func (m *Writeback) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Writeback.Unmarshal(m, b)
}
func (m *Writeback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Writeback.Marshal(b, m, deterministic)
}
func (m *Writeback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Writeback.Merge(m, src)
}
func (m *Writeback) XXX_Size() int {
	return xxx_messageInfo_Writeback.Size(m)
}
func (m *Writeback) XXX_DiscardUnknown() {
	xxx_messageInfo_Writeback.DiscardUnknown(m)
}

var xxx_messageInfo_Writeback proto.InternalMessageInfo

func (m *Writeback) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *Writeback) GetHuntLastTimestamp() uint64 {
	if m != nil {
		return m.HuntLastTimestamp
	}
	return 0
}

func (m *Writeback) GetLastServerSerialNumber() uint64 {
	if m != nil {
		return m.LastServerSerialNumber
	}
	return 0
}

// Configuration for the windows installer.
type WindowsInstallerConfig struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	InstallPath          string   `protobuf:"bytes,2,opt,name=install_path,json=installPath,proto3" json:"install_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WindowsInstallerConfig) Reset()         { *m = WindowsInstallerConfig{} }
func (m *WindowsInstallerConfig) String() string { return proto.CompactTextString(m) }
func (*WindowsInstallerConfig) ProtoMessage()    {}
func (*WindowsInstallerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{2}
}

func (m *WindowsInstallerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WindowsInstallerConfig.Unmarshal(m, b)
}
func (m *WindowsInstallerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WindowsInstallerConfig.Marshal(b, m, deterministic)
}
func (m *WindowsInstallerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WindowsInstallerConfig.Merge(m, src)
}
func (m *WindowsInstallerConfig) XXX_Size() int {
	return xxx_messageInfo_WindowsInstallerConfig.Size(m)
}
func (m *WindowsInstallerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_WindowsInstallerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_WindowsInstallerConfig proto.InternalMessageInfo

func (m *WindowsInstallerConfig) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *WindowsInstallerConfig) GetInstallPath() string {
	if m != nil {
		return m.InstallPath
	}
	return ""
}

type ClientConfig struct {
	Labels               []string                `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty"`
	ServerUrls           []string                `protobuf:"bytes,8,rep,name=server_urls,json=serverUrls,proto3" json:"server_urls,omitempty"`
	CaCertificate        string                  `protobuf:"bytes,11,opt,name=ca_certificate,json=caCertificate,proto3" json:"ca_certificate,omitempty"`
	Nonce                string                  `protobuf:"bytes,12,opt,name=nonce,proto3" json:"nonce,omitempty"`
	WritebackLinux       string                  `protobuf:"bytes,9,opt,name=writeback_linux,json=writebackLinux,proto3" json:"writeback_linux,omitempty"`
	WritebackWindows     string                  `protobuf:"bytes,10,opt,name=writeback_windows,json=writebackWindows,proto3" json:"writeback_windows,omitempty"`
	MaxPoll              uint64                  `protobuf:"varint,15,opt,name=max_poll,json=maxPoll,proto3" json:"max_poll,omitempty"`
	WindowsInstaller     *WindowsInstallerConfig `protobuf:"bytes,16,opt,name=windows_installer,json=windowsInstaller,proto3" json:"windows_installer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ClientConfig) Reset()         { *m = ClientConfig{} }
func (m *ClientConfig) String() string { return proto.CompactTextString(m) }
func (*ClientConfig) ProtoMessage()    {}
func (*ClientConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{3}
}

func (m *ClientConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientConfig.Unmarshal(m, b)
}
func (m *ClientConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientConfig.Marshal(b, m, deterministic)
}
func (m *ClientConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientConfig.Merge(m, src)
}
func (m *ClientConfig) XXX_Size() int {
	return xxx_messageInfo_ClientConfig.Size(m)
}
func (m *ClientConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ClientConfig proto.InternalMessageInfo

func (m *ClientConfig) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ClientConfig) GetServerUrls() []string {
	if m != nil {
		return m.ServerUrls
	}
	return nil
}

func (m *ClientConfig) GetCaCertificate() string {
	if m != nil {
		return m.CaCertificate
	}
	return ""
}

func (m *ClientConfig) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *ClientConfig) GetWritebackLinux() string {
	if m != nil {
		return m.WritebackLinux
	}
	return ""
}

func (m *ClientConfig) GetWritebackWindows() string {
	if m != nil {
		return m.WritebackWindows
	}
	return ""
}

func (m *ClientConfig) GetMaxPoll() uint64 {
	if m != nil {
		return m.MaxPoll
	}
	return 0
}

func (m *ClientConfig) GetWindowsInstaller() *WindowsInstallerConfig {
	if m != nil {
		return m.WindowsInstaller
	}
	return nil
}

type APIConfig struct {
	BindAddress          string   `protobuf:"bytes,1,opt,name=bind_address,json=bindAddress,proto3" json:"bind_address,omitempty"`
	BindPort             uint32   `protobuf:"varint,2,opt,name=bind_port,json=bindPort,proto3" json:"bind_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *APIConfig) Reset()         { *m = APIConfig{} }
func (m *APIConfig) String() string { return proto.CompactTextString(m) }
func (*APIConfig) ProtoMessage()    {}
func (*APIConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{4}
}

func (m *APIConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIConfig.Unmarshal(m, b)
}
func (m *APIConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIConfig.Marshal(b, m, deterministic)
}
func (m *APIConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIConfig.Merge(m, src)
}
func (m *APIConfig) XXX_Size() int {
	return xxx_messageInfo_APIConfig.Size(m)
}
func (m *APIConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_APIConfig.DiscardUnknown(m)
}

var xxx_messageInfo_APIConfig proto.InternalMessageInfo

func (m *APIConfig) GetBindAddress() string {
	if m != nil {
		return m.BindAddress
	}
	return ""
}

func (m *APIConfig) GetBindPort() uint32 {
	if m != nil {
		return m.BindPort
	}
	return 0
}

type GUIConfig struct {
	BindAddress          string   `protobuf:"bytes,1,opt,name=bind_address,json=bindAddress,proto3" json:"bind_address,omitempty"`
	BindPort             uint32   `protobuf:"varint,2,opt,name=bind_port,json=bindPort,proto3" json:"bind_port,omitempty"`
	InternalCidr         []string `protobuf:"bytes,3,rep,name=internal_cidr,json=internalCidr,proto3" json:"internal_cidr,omitempty"`
	VpnCidr              []string `protobuf:"bytes,4,rep,name=vpn_cidr,json=vpnCidr,proto3" json:"vpn_cidr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GUIConfig) Reset()         { *m = GUIConfig{} }
func (m *GUIConfig) String() string { return proto.CompactTextString(m) }
func (*GUIConfig) ProtoMessage()    {}
func (*GUIConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{5}
}

func (m *GUIConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GUIConfig.Unmarshal(m, b)
}
func (m *GUIConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GUIConfig.Marshal(b, m, deterministic)
}
func (m *GUIConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GUIConfig.Merge(m, src)
}
func (m *GUIConfig) XXX_Size() int {
	return xxx_messageInfo_GUIConfig.Size(m)
}
func (m *GUIConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GUIConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GUIConfig proto.InternalMessageInfo

func (m *GUIConfig) GetBindAddress() string {
	if m != nil {
		return m.BindAddress
	}
	return ""
}

func (m *GUIConfig) GetBindPort() uint32 {
	if m != nil {
		return m.BindPort
	}
	return 0
}

func (m *GUIConfig) GetInternalCidr() []string {
	if m != nil {
		return m.InternalCidr
	}
	return nil
}

func (m *GUIConfig) GetVpnCidr() []string {
	if m != nil {
		return m.VpnCidr
	}
	return nil
}

type CAConfig struct {
	PrivateKey           string   `protobuf:"bytes,1,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CAConfig) Reset()         { *m = CAConfig{} }
func (m *CAConfig) String() string { return proto.CompactTextString(m) }
func (*CAConfig) ProtoMessage()    {}
func (*CAConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{6}
}

func (m *CAConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CAConfig.Unmarshal(m, b)
}
func (m *CAConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CAConfig.Marshal(b, m, deterministic)
}
func (m *CAConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CAConfig.Merge(m, src)
}
func (m *CAConfig) XXX_Size() int {
	return xxx_messageInfo_CAConfig.Size(m)
}
func (m *CAConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CAConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CAConfig proto.InternalMessageInfo

func (m *CAConfig) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

type FrontendConfig struct {
	BindAddress          string   `protobuf:"bytes,1,opt,name=bind_address,json=bindAddress,proto3" json:"bind_address,omitempty"`
	BindPort             uint32   `protobuf:"varint,2,opt,name=bind_port,json=bindPort,proto3" json:"bind_port,omitempty"`
	Certificate          string   `protobuf:"bytes,3,opt,name=certificate,proto3" json:"certificate,omitempty"`
	PrivateKey           string   `protobuf:"bytes,4,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	ClientLeaseTime      uint32   `protobuf:"varint,5,opt,name=client_lease_time,json=clientLeaseTime,proto3" json:"client_lease_time,omitempty"`
	DnsName              string   `protobuf:"bytes,6,opt,name=dns_name,json=dnsName,proto3" json:"dns_name,omitempty"`
	ArtifactsPath        string   `protobuf:"bytes,7,opt,name=artifacts_path,json=artifactsPath,proto3" json:"artifacts_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrontendConfig) Reset()         { *m = FrontendConfig{} }
func (m *FrontendConfig) String() string { return proto.CompactTextString(m) }
func (*FrontendConfig) ProtoMessage()    {}
func (*FrontendConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{7}
}

func (m *FrontendConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrontendConfig.Unmarshal(m, b)
}
func (m *FrontendConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrontendConfig.Marshal(b, m, deterministic)
}
func (m *FrontendConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrontendConfig.Merge(m, src)
}
func (m *FrontendConfig) XXX_Size() int {
	return xxx_messageInfo_FrontendConfig.Size(m)
}
func (m *FrontendConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FrontendConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FrontendConfig proto.InternalMessageInfo

func (m *FrontendConfig) GetBindAddress() string {
	if m != nil {
		return m.BindAddress
	}
	return ""
}

func (m *FrontendConfig) GetBindPort() uint32 {
	if m != nil {
		return m.BindPort
	}
	return 0
}

func (m *FrontendConfig) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

func (m *FrontendConfig) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *FrontendConfig) GetClientLeaseTime() uint32 {
	if m != nil {
		return m.ClientLeaseTime
	}
	return 0
}

func (m *FrontendConfig) GetDnsName() string {
	if m != nil {
		return m.DnsName
	}
	return ""
}

func (m *FrontendConfig) GetArtifactsPath() string {
	if m != nil {
		return m.ArtifactsPath
	}
	return ""
}

type DatastoreConfig struct {
	Implementation       string   `protobuf:"bytes,1,opt,name=implementation,proto3" json:"implementation,omitempty"`
	Location             string   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	FilestoreDirectory   string   `protobuf:"bytes,3,opt,name=filestore_directory,json=filestoreDirectory,proto3" json:"filestore_directory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatastoreConfig) Reset()         { *m = DatastoreConfig{} }
func (m *DatastoreConfig) String() string { return proto.CompactTextString(m) }
func (*DatastoreConfig) ProtoMessage()    {}
func (*DatastoreConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{8}
}

func (m *DatastoreConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatastoreConfig.Unmarshal(m, b)
}
func (m *DatastoreConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatastoreConfig.Marshal(b, m, deterministic)
}
func (m *DatastoreConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatastoreConfig.Merge(m, src)
}
func (m *DatastoreConfig) XXX_Size() int {
	return xxx_messageInfo_DatastoreConfig.Size(m)
}
func (m *DatastoreConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DatastoreConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DatastoreConfig proto.InternalMessageInfo

func (m *DatastoreConfig) GetImplementation() string {
	if m != nil {
		return m.Implementation
	}
	return ""
}

func (m *DatastoreConfig) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *DatastoreConfig) GetFilestoreDirectory() string {
	if m != nil {
		return m.FilestoreDirectory
	}
	return ""
}

type FlowsConfig struct {
	InterrogateAdditionalQueries []*proto1.VQLRequest `protobuf:"bytes,1,rep,name=interrogate_additional_queries,json=interrogateAdditionalQueries,proto3" json:"interrogate_additional_queries,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}             `json:"-"`
	XXX_unrecognized             []byte               `json:"-"`
	XXX_sizecache                int32                `json:"-"`
}

func (m *FlowsConfig) Reset()         { *m = FlowsConfig{} }
func (m *FlowsConfig) String() string { return proto.CompactTextString(m) }
func (*FlowsConfig) ProtoMessage()    {}
func (*FlowsConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{9}
}

func (m *FlowsConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowsConfig.Unmarshal(m, b)
}
func (m *FlowsConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowsConfig.Marshal(b, m, deterministic)
}
func (m *FlowsConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowsConfig.Merge(m, src)
}
func (m *FlowsConfig) XXX_Size() int {
	return xxx_messageInfo_FlowsConfig.Size(m)
}
func (m *FlowsConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowsConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FlowsConfig proto.InternalMessageInfo

func (m *FlowsConfig) GetInterrogateAdditionalQueries() []*proto1.VQLRequest {
	if m != nil {
		return m.InterrogateAdditionalQueries
	}
	return nil
}

// Event generators for the client. Note: These are sent from the
// server to the client and are not embedded in the client.
type ClientEvents struct {
	Artifacts            []string `protobuf:"bytes,1,rep,name=artifacts,proto3" json:"artifacts,omitempty"`
	Version              uint64   `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientEvents) Reset()         { *m = ClientEvents{} }
func (m *ClientEvents) String() string { return proto.CompactTextString(m) }
func (*ClientEvents) ProtoMessage()    {}
func (*ClientEvents) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{10}
}

func (m *ClientEvents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientEvents.Unmarshal(m, b)
}
func (m *ClientEvents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientEvents.Marshal(b, m, deterministic)
}
func (m *ClientEvents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientEvents.Merge(m, src)
}
func (m *ClientEvents) XXX_Size() int {
	return xxx_messageInfo_ClientEvents.Size(m)
}
func (m *ClientEvents) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientEvents.DiscardUnknown(m)
}

var xxx_messageInfo_ClientEvents proto.InternalMessageInfo

func (m *ClientEvents) GetArtifacts() []string {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

func (m *ClientEvents) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type Config struct {
	Version              *Version         `protobuf:"bytes,8,opt,name=version,proto3" json:"version,omitempty"`
	Client               *ClientConfig    `protobuf:"bytes,1,opt,name=Client,proto3" json:"Client,omitempty"`
	Events               *ClientEvents    `protobuf:"bytes,10,opt,name=Events,proto3" json:"Events,omitempty"`
	API                  *APIConfig       `protobuf:"bytes,2,opt,name=API,proto3" json:"API,omitempty"`
	GUI                  *GUIConfig       `protobuf:"bytes,3,opt,name=GUI,proto3" json:"GUI,omitempty"`
	CA                   *CAConfig        `protobuf:"bytes,4,opt,name=CA,proto3" json:"CA,omitempty"`
	Frontend             *FrontendConfig  `protobuf:"bytes,5,opt,name=Frontend,proto3" json:"Frontend,omitempty"`
	Datastore            *DatastoreConfig `protobuf:"bytes,6,opt,name=Datastore,proto3" json:"Datastore,omitempty"`
	Flows                *FlowsConfig     `protobuf:"bytes,7,opt,name=Flows,proto3" json:"Flows,omitempty"`
	Writeback            *Writeback       `protobuf:"bytes,9,opt,name=writeback,proto3" json:"writeback,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{11}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *Config) GetClient() *ClientConfig {
	if m != nil {
		return m.Client
	}
	return nil
}

func (m *Config) GetEvents() *ClientEvents {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *Config) GetAPI() *APIConfig {
	if m != nil {
		return m.API
	}
	return nil
}

func (m *Config) GetGUI() *GUIConfig {
	if m != nil {
		return m.GUI
	}
	return nil
}

func (m *Config) GetCA() *CAConfig {
	if m != nil {
		return m.CA
	}
	return nil
}

func (m *Config) GetFrontend() *FrontendConfig {
	if m != nil {
		return m.Frontend
	}
	return nil
}

func (m *Config) GetDatastore() *DatastoreConfig {
	if m != nil {
		return m.Datastore
	}
	return nil
}

func (m *Config) GetFlows() *FlowsConfig {
	if m != nil {
		return m.Flows
	}
	return nil
}

func (m *Config) GetWriteback() *Writeback {
	if m != nil {
		return m.Writeback
	}
	return nil
}

func init() {
	proto.RegisterType((*Version)(nil), "proto.Version")
	proto.RegisterType((*Writeback)(nil), "proto.Writeback")
	proto.RegisterType((*WindowsInstallerConfig)(nil), "proto.WindowsInstallerConfig")
	proto.RegisterType((*ClientConfig)(nil), "proto.ClientConfig")
	proto.RegisterType((*APIConfig)(nil), "proto.APIConfig")
	proto.RegisterType((*GUIConfig)(nil), "proto.GUIConfig")
	proto.RegisterType((*CAConfig)(nil), "proto.CAConfig")
	proto.RegisterType((*FrontendConfig)(nil), "proto.FrontendConfig")
	proto.RegisterType((*DatastoreConfig)(nil), "proto.DatastoreConfig")
	proto.RegisterType((*FlowsConfig)(nil), "proto.FlowsConfig")
	proto.RegisterType((*ClientEvents)(nil), "proto.ClientEvents")
	proto.RegisterType((*Config)(nil), "proto.Config")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_3eaf2c85e69e9ea4) }

var fileDescriptor_3eaf2c85e69e9ea4 = []byte{
	// 2247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x58, 0xcd, 0x6f, 0x1c, 0x49,
	0x15, 0x57, 0xdb, 0x8e, 0x3f, 0x6a, 0x6c, 0x27, 0xae, 0xec, 0x86, 0x21, 0x4a, 0x48, 0xed, 0xb0,
	0x10, 0x67, 0x59, 0xb5, 0xed, 0x49, 0x36, 0xbb, 0x89, 0x58, 0x65, 0xc7, 0x93, 0x0f, 0xcc, 0x3a,
	0x91, 0xd3, 0xb1, 0x63, 0xb4, 0x80, 0x86, 0x9a, 0xee, 0xf2, 0x4c, 0x25, 0xd5, 0x55, 0x93, 0xaa,
	0x1a, 0x8f, 0x2d, 0x24, 0x04, 0x9c, 0x40, 0xdc, 0x90, 0xf8, 0x94, 0x38, 0xb0, 0xc0, 0x6d, 0x0f,
	0xdc, 0xb9, 0x72, 0xe0, 0xaf, 0xe0, 0x00, 0x17, 0xb8, 0xc1, 0x9d, 0x03, 0xaa, 0x57, 0xd5, 0xdd,
	0x33, 0x93, 0x20, 0x38, 0xec, 0x81, 0x4b, 0x62, 0x57, 0xbf, 0xf7, 0x7b, 0xbf, 0xf7, 0xea, 0x7d,
	0x95, 0xd1, 0x72, 0xaa, 0xe4, 0x11, 0xef, 0xc5, 0x03, 0xad, 0xac, 0xc2, 0x67, 0xe0, 0xbf, 0x8b,
	0xb7, 0x47, 0xa3, 0x51, 0x7c, 0xcc, 0x84, 0x4a, 0x79, 0xc6, 0x4e, 0xe2, 0x54, 0xe5, 0x1b, 0x3d,
	0x25, 0xa8, 0xec, 0x6d, 0xf8, 0x43, 0x4d, 0x07, 0x56, 0xe9, 0x0d, 0x10, 0xde, 0x30, 0x2c, 0xa7,
	0xd2, 0xf2, 0xd4, 0x43, 0x5c, 0x7c, 0xff, 0x7f, 0xd3, 0xa5, 0xa9, 0xe5, 0x4a, 0x9a, 0x80, 0x71,
	0xfc, 0x42, 0x78, 0xf5, 0xc6, 0x3f, 0x66, 0xd0, 0xc2, 0x53, 0xa6, 0x0d, 0x57, 0x12, 0xc7, 0x68,
	0x4e, 0xd2, 0x9c, 0xd5, 0x23, 0x12, 0xad, 0x2f, 0x6d, 0x5f, 0xfc, 0xcb, 0xbf, 0xfe, 0xfa, 0xc7,
	0xe8, 0x35, 0x8c, 0xf7, 0xfb, 0x8c, 0xa4, 0x82, 0x33, 0x69, 0xaf, 0x1a, 0xe2, 0x04, 0xe2, 0x04,
	0xe4, 0xf0, 0x23, 0x54, 0xcb, 0x98, 0x49, 0x35, 0x1f, 0x38, 0xec, 0xfa, 0x0c, 0xa8, 0xbd, 0x0d,
	0x6a, 0x5f, 0xc4, 0x6f, 0x4e, 0xa8, 0x09, 0x25, 0x7b, 0x64, 0x4c, 0x98, 0x18, 0xab, 0xb9, 0xec,
	0x25, 0xe3, 0x00, 0x98, 0xa2, 0x85, 0x63, 0x4f, 0xa5, 0x3e, 0x0b, 0x58, 0x0f, 0x00, 0xab, 0x85,
	0xef, 0x4c, 0x60, 0x05, 0x99, 0x00, 0x41, 0x86, 0xc6, 0xfd, 0x4b, 0x49, 0x11, 0x97, 0xe2, 0xbb,
	0x3b, 0x35, 0x69, 0x9f, 0x39, 0xbe, 0x05, 0x2e, 0xbe, 0x83, 0xe6, 0x53, 0x95, 0xe7, 0xdc, 0xd6,
	0xe7, 0xc0, 0xc2, 0x55, 0xb0, 0xf0, 0x06, 0xbe, 0x32, 0x61, 0xa1, 0xc7, 0x2d, 0xf1, 0x62, 0xc1,
	0x48, 0x9c, 0x04, 0x35, 0xdc, 0x42, 0xa8, 0x3b, 0xe4, 0x22, 0xeb, 0x58, 0x9e, 0xb3, 0xfa, 0x19,
	0x00, 0x69, 0x00, 0xc8, 0x25, 0x7c, 0xf1, 0xb0, 0xcf, 0x24, 0xb1, 0x25, 0x12, 0x19, 0x51, 0x43,
	0x9c, 0xb4, 0x8d, 0x93, 0x25, 0xd0, 0xda, 0xe7, 0x39, 0x6b, 0x7c, 0x32, 0x8b, 0x96, 0x0e, 0x35,
	0xb7, 0xac, 0x4b, 0xd3, 0xe7, 0x38, 0x41, 0xb5, 0x81, 0xe6, 0xc7, 0xd4, 0xb2, 0xce, 0x73, 0x76,
	0x5a, 0x5f, 0x00, 0xc4, 0x2d, 0x40, 0xfc, 0x12, 0xbe, 0x36, 0x41, 0x2b, 0xc8, 0x91, 0xe7, 0xec,
	0x94, 0x70, 0x49, 0xf6, 0xee, 0x3d, 0x24, 0x4c, 0xa6, 0x2a, 0x03, 0x82, 0x28, 0x7c, 0xfd, 0x90,
	0x9d, 0xe2, 0x6f, 0xa1, 0xf3, 0xfd, 0xa1, 0xb4, 0x1d, 0x41, 0x8d, 0x05, 0xa2, 0xc6, 0xd2, 0x7c,
	0x50, 0x5f, 0x21, 0xd1, 0xfa, 0xdc, 0xf6, 0x26, 0x60, 0xbf, 0x85, 0xd7, 0x1d, 0xb6, 0x93, 0x20,
	0x4e, 0x96, 0x94, 0x62, 0xc4, 0xf6, 0xb9, 0x29, 0x3c, 0xd0, 0x54, 0xc6, 0xc9, 0x9a, 0x13, 0xd8,
	0xa5, 0xc6, 0xee, 0x17, 0x32, 0xf8, 0x6f, 0x11, 0xfa, 0x2c, 0xa0, 0x1b, 0xa6, 0x8f, 0x99, 0x76,
	0xff, 0x71, 0x2a, 0x3a, 0x72, 0x98, 0x77, 0x99, 0xae, 0xaf, 0x82, 0xa1, 0x4f, 0x22, 0xb0, 0xf4,
	0xbb, 0x08, 0x7f, 0x1c, 0x95, 0xb6, 0x52, 0xa6, 0x2d, 0x3f, 0xe2, 0xa9, 0xf3, 0xc3, 0x2b, 0x11,
	0xaf, 0x44, 0x46, 0x8c, 0x18, 0x3a, 0x22, 0x47, 0x5a, 0xe5, 0x10, 0x43, 0x8f, 0x1c, 0x93, 0x36,
	0x30, 0x31, 0x44, 0xb3, 0xa3, 0xa1, 0x61, 0xc4, 0x2a, 0x92, 0x2a, 0x29, 0x59, 0x6a, 0xdd, 0x8f,
	0x5e, 0xc8, 0x90, 0x11, 0xb7, 0x7d, 0xa2, 0x44, 0xc6, 0xf4, 0x24, 0xaa, 0x21, 0xd4, 0x90, 0x01,
	0xd5, 0x96, 0xa8, 0xa3, 0x20, 0x0c, 0xe1, 0xd3, 0xca, 0x52, 0xc8, 0x43, 0xa8, 0x82, 0x54, 0x89,
	0x38, 0xb9, 0xe0, 0xf8, 0x3d, 0x01, 0x91, 0x27, 0x00, 0xf1, 0x08, 0x10, 0x1a, 0x3f, 0x9a, 0x41,
	0x17, 0x0e, 0xb9, 0xcc, 0xd4, 0xc8, 0xec, 0x48, 0x63, 0xa9, 0x10, 0x4c, 0xb7, 0xa1, 0x88, 0xf1,
	0x43, 0xb4, 0xec, 0x30, 0x79, 0xca, 0x3a, 0x63, 0x85, 0xf3, 0x16, 0xb8, 0xfd, 0x26, 0x6e, 0x38,
	0xa7, 0xdd, 0xb9, 0xb3, 0x5d, 0x78, 0xc4, 0x53, 0xef, 0x82, 0x66, 0xd4, 0xb2, 0x38, 0xa9, 0x85,
	0xb3, 0x47, 0xae, 0x9e, 0x7e, 0x1d, 0xa1, 0x65, 0xee, 0x4d, 0x74, 0x06, 0xd4, 0xf6, 0x43, 0x45,
	0x7d, 0x07, 0xf0, 0x4e, 0xf0, 0xf1, 0x61, 0x9f, 0x69, 0x46, 0x4c, 0x5f, 0x0d, 0x45, 0x06, 0x80,
	0x5d, 0x2e, 0xa9, 0x3e, 0x25, 0x5d, 0x46, 0x82, 0x1a, 0xcb, 0xee, 0x10, 0x67, 0x35, 0xfc, 0x0a,
	0x49, 0x4c, 0x65, 0x46, 0x52, 0x35, 0xe0, 0xcc, 0x8c, 0xeb, 0x58, 0xe5, 0xef, 0x59, 0xa8, 0xd4,
	0x07, 0xc2, 0x89, 0x05, 0x35, 0x33, 0xce, 0x36, 0x4e, 0x6a, 0xe1, 0x78, 0x8f, 0xda, 0x7e, 0xe3,
	0xef, 0x0b, 0x68, 0xd9, 0xdf, 0x48, 0x88, 0xc1, 0xf7, 0x22, 0x34, 0x2f, 0x68, 0x97, 0x09, 0x53,
	0x9f, 0x27, 0xb3, 0xeb, 0x4b, 0xdb, 0x1c, 0xe8, 0xa6, 0x98, 0xb6, 0x88, 0xe0, 0x06, 0x02, 0xef,
	0xbf, 0x8f, 0x57, 0x46, 0x9f, 0x9a, 0x98, 0xec, 0x3b, 0x02, 0x54, 0x08, 0x35, 0x32, 0xc4, 0x30,
	0xc1, 0x52, 0xcb, 0x32, 0xd2, 0xd3, 0x6a, 0x38, 0x30, 0x4e, 0x2b, 0x0d, 0x17, 0x6f, 0x95, 0xf3,
	0xd0, 0x52, 0xdd, 0x63, 0xd6, 0x49, 0x70, 0x09, 0x19, 0x6b, 0xe2, 0x24, 0x18, 0xc6, 0x14, 0xd5,
	0x42, 0x1e, 0x0e, 0xb5, 0x30, 0xf5, 0x45, 0xe0, 0xf1, 0x01, 0xf0, 0xb8, 0x8d, 0xdf, 0xab, 0x78,
	0x84, 0x04, 0x38, 0x48, 0x76, 0x27, 0xc8, 0x8c, 0xb8, 0x10, 0xc4, 0xfa, 0xa0, 0x54, 0xb9, 0x15,
	0x27, 0xc8, 0xcb, 0x1f, 0x68, 0x61, 0xf0, 0x0f, 0x23, 0xb4, 0x9a, 0xd2, 0xce, 0x58, 0xfa, 0xd6,
	0x6b, 0x70, 0x3b, 0x5d, 0x30, 0xf3, 0x0d, 0xfc, 0x91, 0x8b, 0x7b, 0xbb, 0x75, 0xd5, 0x4c, 0xa4,
	0x38, 0xd4, 0xa7, 0xe7, 0xbc, 0x77, 0xef, 0x61, 0x4c, 0xf6, 0x5f, 0xb2, 0x3a, 0x34, 0x96, 0x50,
	0x79, 0x5a, 0x70, 0x33, 0xbc, 0x27, 0x59, 0xe6, 0xf3, 0x19, 0x2e, 0xa8, 0xdd, 0x8a, 0x93, 0x95,
	0x94, 0xb6, 0x2b, 0x50, 0xfc, 0x87, 0x08, 0x9d, 0x91, 0x4a, 0xa6, 0xac, 0xbe, 0x0c, 0x14, 0x7e,
	0xeb, 0x0b, 0xed, 0x57, 0x11, 0xfe, 0x65, 0xd4, 0x22, 0xa6, 0x4f, 0x35, 0xcb, 0x08, 0x08, 0xbc,
	0xe4, 0xe5, 0x40, 0x33, 0xc3, 0xa4, 0xf5, 0x44, 0x82, 0x49, 0xf8, 0xa0, 0xd9, 0x33, 0xe7, 0x76,
	0x11, 0xf3, 0x51, 0x9f, 0xa7, 0x7d, 0x92, 0x29, 0x22, 0x95, 0x2d, 0x94, 0x3c, 0x1f, 0x57, 0x3d,
	0x46, 0x11, 0x6e, 0x49, 0xee, 0xa8, 0x77, 0x19, 0x61, 0x79, 0x97, 0x65, 0xc1, 0x49, 0x3b, 0xde,
	0xaf, 0xfc, 0xa8, 0x1b, 0x6a, 0x48, 0xb0, 0x38, 0xf1, 0x9c, 0x71, 0x8a, 0xce, 0x8e, 0x8a, 0xee,
	0xd7, 0x11, 0x5c, 0x0e, 0x4f, 0xea, 0x4b, 0xe0, 0xc6, 0x6d, 0xf0, 0xe2, 0x06, 0x6e, 0xb6, 0x88,
	0xcb, 0x7e, 0x32, 0x82, 0x74, 0x07, 0x89, 0x8a, 0x97, 0xe3, 0x0a, 0xea, 0x90, 0xb9, 0x82, 0x18,
	0x0b, 0x75, 0xb4, 0x5a, 0x42, 0xee, 0x3a, 0x79, 0xcc, 0xd1, 0x5a, 0x65, 0x64, 0xe4, 0xab, 0xb7,
	0x8e, 0xc0, 0xcc, 0x97, 0xc1, 0xcc, 0x4d, 0x7c, 0x63, 0xc2, 0x4c, 0x90, 0xf9, 0xaf, 0x86, 0xce,
	0x95, 0xb0, 0xa1, 0x27, 0xe0, 0xaf, 0xa3, 0xc5, 0x9c, 0x9e, 0x74, 0x06, 0x4a, 0x88, 0xfa, 0x59,
	0x68, 0x7c, 0x45, 0xe6, 0x5d, 0x77, 0x41, 0xce, 0xe9, 0x09, 0xcf, 0x87, 0x79, 0xd1, 0xe5, 0x20,
	0x0b, 0x53, 0x25, 0x33, 0x43, 0xba, 0xcc, 0x8e, 0x18, 0x93, 0xc5, 0x35, 0x38, 0x00, 0x13, 0x37,
	0x67, 0x6f, 0x6e, 0x6e, 0x26, 0x0b, 0x39, 0x3d, 0xd9, 0x53, 0x42, 0xe0, 0x6f, 0xa3, 0xb5, 0xc0,
	0xac, 0x53, 0x94, 0xb8, 0xae, 0x9f, 0x23, 0xd1, 0x7a, 0xad, 0x79, 0xd9, 0x4f, 0xf0, 0xf8, 0xd5,
	0xbd, 0x69, 0x7b, 0x03, 0x48, 0x5c, 0xc3, 0x57, 0xdb, 0xe3, 0xb7, 0x50, 0xb4, 0xa2, 0xc2, 0xdd,
	0x12, 0xd4, 0x79, 0x36, 0x05, 0xd4, 0xf8, 0x73, 0x84, 0x96, 0x5a, 0x7b, 0x3b, 0xa1, 0xd0, 0x7f,
	0x11, 0xa1, 0xe5, 0x2e, 0x97, 0x59, 0x87, 0x66, 0x99, 0x66, 0xc6, 0x84, 0x6e, 0x77, 0x0c, 0x76,
	0x06, 0x58, 0xb6, 0xfc, 0x31, 0x94, 0x2b, 0x97, 0x19, 0xe9, 0x25, 0x7b, 0x6d, 0xc2, 0x64, 0x36,
	0x50, 0xdc, 0x27, 0x1c, 0x37, 0x45, 0xef, 0x1a, 0x9a, 0x21, 0x15, 0xe2, 0x94, 0x28, 0x29, 0xa0,
	0x7b, 0x6d, 0x35, 0xdf, 0x8d, 0x37, 0xe3, 0xcd, 0x78, 0xeb, 0x6d, 0xa2, 0x6c, 0x9f, 0xe9, 0x11,
	0x37, 0xcc, 0x9d, 0x9b, 0xa1, 0x86, 0x6e, 0x39, 0xd0, 0x6a, 0xc0, 0xb4, 0x70, 0x45, 0x92, 0xba,
	0x23, 0x6e, 0xe3, 0xa4, 0xe6, 0x8c, 0x04, 0x9b, 0xf8, 0x1d, 0xb4, 0x04, 0xd4, 0x06, 0x4a, 0x5b,
	0xe8, 0x9a, 0x2b, 0xdb, 0x75, 0xe0, 0x85, 0xf1, 0xb9, 0x3d, 0xa5, 0x6d, 0x49, 0xca, 0x95, 0xf5,
	0xa2, 0xfb, 0xc9, 0x9d, 0x36, 0xbe, 0x3f, 0x87, 0x96, 0x1e, 0x1c, 0x14, 0x0e, 0xfe, 0xfc, 0xd5,
	0x0e, 0x0e, 0x01, 0x48, 0xe1, 0x7c, 0xda, 0xc1, 0x07, 0x07, 0x3b, 0xff, 0xdf, 0xfe, 0xe1, 0x9f,
	0x45, 0x68, 0x85, 0x4b, 0xcb, 0xb4, 0xa4, 0xa2, 0x93, 0xf2, 0x4c, 0xd7, 0x67, 0xa1, 0x35, 0x6a,
	0xd0, 0x15, 0xf8, 0x59, 0x7b, 0xe7, 0x6e, 0x42, 0x82, 0xbb, 0x0c, 0x9a, 0x6e, 0x21, 0x4e, 0x24,
	0xb3, 0x23, 0xa5, 0x9f, 0x1b, 0xb2, 0xce, 0xe2, 0x5e, 0x4c, 0xb6, 0x6e, 0x35, 0xe3, 0xad, 0x9b,
	0xef, 0x39, 0x37, 0x36, 0xb6, 0x6e, 0x5e, 0x8b, 0xc9, 0x21, 0x23, 0x30, 0x8e, 0x9d, 0xdb, 0x6e,
	0x10, 0xf7, 0xd5, 0x08, 0x32, 0xab, 0x00, 0xd8, 0x60, 0x27, 0x01, 0x89, 0xa7, 0xae, 0xfa, 0x97,
	0x8b, 0x0f, 0x6d, 0x9e, 0x69, 0xfc, 0x83, 0x08, 0x2d, 0x1e, 0x0f, 0xa4, 0x27, 0x35, 0x07, 0xa4,
	0x72, 0x20, 0xd5, 0xc3, 0xec, 0x65, 0x52, 0xc7, 0x03, 0xf9, 0xa9, 0xf3, 0x59, 0x38, 0x1e, 0x48,
	0x47, 0xa5, 0xf1, 0x9b, 0x08, 0x2d, 0xb6, 0x5b, 0x21, 0x07, 0x7e, 0x1a, 0x4d, 0xae, 0x63, 0x53,
	0x29, 0xe0, 0x0a, 0x7a, 0x7c, 0x0b, 0x0b, 0xd5, 0xd4, 0x6e, 0xbd, 0xa2, 0xd1, 0x73, 0xe3, 0x3b,
	0x64, 0xaa, 0xb4, 0x66, 0x66, 0xa0, 0x64, 0xd9, 0xd5, 0xd9, 0xc4, 0x88, 0xe0, 0x32, 0xac, 0x38,
	0xf1, 0xe4, 0x7c, 0x99, 0x58, 0xe9, 0x1a, 0xff, 0x9c, 0x43, 0xab, 0xf7, 0xb5, 0x92, 0x96, 0xc9,
	0x2c, 0x70, 0xfd, 0xee, 0xab, 0xf3, 0xf5, 0x9b, 0x40, 0xf6, 0x10, 0x1f, 0x4c, 0xe7, 0xeb, 0x51,
	0x50, 0x1f, 0x4b, 0xda, 0x62, 0xb7, 0x1a, 0xdb, 0xa8, 0x20, 0x80, 0x85, 0x04, 0x24, 0xa7, 0x12,
	0x82, 0x1c, 0x29, 0x4d, 0x9e, 0xa9, 0xae, 0xf9, 0x74, 0xf2, 0xf2, 0x11, 0xaa, 0x8d, 0x0f, 0xd2,
	0xd9, 0xc9, 0x87, 0xc3, 0xd7, 0xde, 0xd9, 0xbc, 0x45, 0xc6, 0xe6, 0x9d, 0x8b, 0x72, 0xc9, 0x3d,
	0x6c, 0x85, 0xc9, 0x38, 0x00, 0x3e, 0x9a, 0xbc, 0x34, 0xbf, 0xda, 0xdf, 0x03, 0xbc, 0x3b, 0xf8,
	0xfd, 0xff, 0x70, 0x69, 0x53, 0x90, 0x53, 0x37, 0xe8, 0x3c, 0xce, 0xa9, 0x9d, 0xdc, 0xab, 0x9f,
	0xa2, 0x35, 0x3f, 0x1f, 0x3a, 0x82, 0x51, 0xc3, 0xaa, 0x37, 0xc0, 0x4a, 0xb5, 0xf4, 0x7d, 0x45,
	0x8d, 0xfc, 0x6b, 0xc7, 0x2a, 0x02, 0x52, 0xc5, 0xfc, 0xcd, 0x99, 0x31, 0xb4, 0xc7, 0x4c, 0x9c,
	0x9c, 0xf5, 0x27, 0xbb, 0xee, 0xab, 0x5b, 0xa9, 0xf1, 0x07, 0x68, 0x31, 0x93, 0xc6, 0xef, 0x90,
	0xf3, 0x40, 0xfe, 0x0b, 0x00, 0x77, 0x05, 0x5f, 0x76, 0xe4, 0xef, 0x3e, 0x7a, 0x32, 0xb1, 0x47,
	0x16, 0xcc, 0xe3, 0x64, 0x21, 0x93, 0x06, 0x56, 0xc7, 0x27, 0x68, 0x95, 0xba, 0x78, 0xd0, 0xd4,
	0x1a, 0xbf, 0x3b, 0x2e, 0x4c, 0x06, 0xd5, 0x2d, 0x6f, 0x8e, 0x12, 0xcd, 0x32, 0xee, 0x06, 0x01,
	0x15, 0xa4, 0x54, 0x00, 0xaa, 0x8a, 0x66, 0xc9, 0x4a, 0x79, 0x04, 0xbb, 0xde, 0xef, 0x67, 0xd0,
	0xd9, 0xbb, 0xd4, 0x52, 0x63, 0x95, 0x66, 0x21, 0xe9, 0x0c, 0x5a, 0xe5, 0xf9, 0x40, 0xb0, 0x9c,
	0x49, 0xbf, 0x41, 0x87, 0xac, 0xfb, 0x10, 0x0c, 0xdd, 0xc3, 0x9b, 0xd3, 0x4b, 0x6f, 0x56, 0x00,
	0x90, 0x49, 0x3d, 0x67, 0x74, 0x68, 0x58, 0xdc, 0x5c, 0xbb, 0xcf, 0x05, 0xdb, 0xa6, 0x86, 0x39,
	0x53, 0x4f, 0x9c, 0x64, 0x32, 0x65, 0x02, 0xef, 0xa2, 0xc5, 0x62, 0x4f, 0x0d, 0x3b, 0xf1, 0xe4,
	0x23, 0xa6, 0xd8, 0x61, 0xd7, 0x95, 0x86, 0x99, 0x7e, 0xcd, 0x27, 0xf2, 0x98, 0xf1, 0x38, 0x29,
	0x11, 0xf0, 0x47, 0xe8, 0xfc, 0x11, 0x17, 0x0c, 0xce, 0x3b, 0x19, 0xd7, 0x2c, 0xb5, 0x4a, 0x9f,
	0x86, 0x2c, 0xbc, 0x06, 0xc0, 0x9f, 0xc7, 0x6f, 0x40, 0xd6, 0x84, 0xa0, 0x79, 0xfe, 0xc3, 0x81,
	0x0b, 0x12, 0xcb, 0x08, 0x68, 0xc7, 0x09, 0x2e, 0x51, 0xee, 0x16, 0x20, 0x8d, 0x8f, 0x23, 0x54,
	0xbb, 0xef, 0x16, 0xda, 0x10, 0xae, 0x1f, 0x47, 0xe8, 0x73, 0xd0, 0x81, 0xb4, 0xea, 0xb9, 0xf4,
	0xac, 0xc2, 0xdf, 0x79, 0x31, 0x64, 0x9a, 0x33, 0x57, 0xb5, 0xb3, 0xeb, 0xb5, 0xe6, 0x5a, 0x98,
	0xe6, 0x4f, 0x1f, 0xef, 0x26, 0xec, 0xc5, 0x90, 0x19, 0xbb, 0x7d, 0x0b, 0xa8, 0x5c, 0xc7, 0x5b,
	0xad, 0xea, 0xce, 0x82, 0x92, 0x23, 0xc5, 0x8d, 0x19, 0xc2, 0xf4, 0xd8, 0xa9, 0xf0, 0xc9, 0x91,
	0xb3, 0x1d, 0x27, 0x97, 0xc6, 0x6c, 0x56, 0xda, 0x8f, 0xbd, 0x72, 0xe3, 0x4f, 0x51, 0xb1, 0xc3,
	0xdf, 0x3b, 0x76, 0xe5, 0x8f, 0x19, 0x5a, 0x2a, 0x6f, 0x1e, 0xf8, 0x8c, 0x3d, 0xbd, 0xbd, 0x88,
	0xdf, 0x87, 0xba, 0x8c, 0xa4, 0x4a, 0x84, 0x45, 0xbd, 0x7c, 0xa8, 0x85, 0xfc, 0xf6, 0xef, 0x70,
	0xdb, 0x67, 0x86, 0x55, 0xa9, 0x15, 0x27, 0x15, 0x32, 0x7e, 0x5c, 0xbd, 0xef, 0x67, 0x60, 0x51,
	0x7a, 0x17, 0x8c, 0x6c, 0xe1, 0x8d, 0x96, 0x24, 0x5c, 0xa6, 0x1a, 0xae, 0xdc, 0xa1, 0x14, 0x4f,
	0x7c, 0xf7, 0x0a, 0xf0, 0x06, 0xdd, 0x6f, 0xda, 0x07, 0xc6, 0x54, 0xef, 0xf9, 0xc6, 0x4f, 0xce,
	0xa0, 0xf9, 0x10, 0xea, 0xfb, 0x15, 0xfa, 0x22, 0x2c, 0x48, 0xab, 0x45, 0x48, 0xfd, 0xe9, 0xf6,
	0x25, 0xb0, 0x76, 0x01, 0xbf, 0x16, 0x0e, 0x08, 0x97, 0xbe, 0xc8, 0xb9, 0x9f, 0x07, 0xe1, 0x4f,
	0x04, 0x5f, 0x45, 0xf3, 0x3e, 0x38, 0x90, 0xd9, 0xb5, 0xe6, 0xf9, 0x00, 0x33, 0xfe, 0xea, 0xd9,
	0xbe, 0x0c, 0x58, 0x9f, 0xc1, 0xaf, 0xfb, 0xd3, 0xe9, 0x55, 0x37, 0x20, 0xe0, 0x7d, 0x34, 0xef,
	0xe3, 0x07, 0xbb, 0xe7, 0x34, 0x96, 0xff, 0x54, 0xfd, 0x0d, 0x22, 0x84, 0x1a, 0x5e, 0x21, 0xe0,
	0xb4, 0x8f, 0x71, 0x58, 0x49, 0xe3, 0x24, 0x60, 0xe1, 0x3d, 0x34, 0xdb, 0xda, 0xdb, 0x81, 0x18,
	0xd6, 0x9a, 0xe7, 0x02, 0x64, 0xb9, 0xa8, 0x55, 0x35, 0x3f, 0xb9, 0xf9, 0xb9, 0x16, 0x0e, 0x3b,
	0x59, 0x6b, 0x6f, 0x6c, 0x6f, 0x49, 0x1c, 0x14, 0x6e, 0xa0, 0xd9, 0x07, 0x07, 0x3b, 0x50, 0x02,
	0x15, 0x62, 0xb9, 0x19, 0x25, 0xee, 0x23, 0xbe, 0x82, 0x66, 0xda, 0x2d, 0xe8, 0xad, 0xb5, 0xe6,
	0xd9, 0xc2, 0x8f, 0x30, 0x37, 0x93, 0x99, 0x76, 0x0b, 0x6f, 0xa1, 0xc5, 0x62, 0x42, 0x41, 0x53,
	0xac, 0x35, 0x5f, 0x0f, 0x62, 0x93, 0x83, 0x2b, 0x29, 0xc5, 0xf0, 0x0d, 0xb4, 0x54, 0x36, 0x18,
	0xe8, 0x7c, 0xb5, 0xe6, 0x85, 0xa0, 0x33, 0xd5, 0x78, 0x92, 0x4a, 0x10, 0xaf, 0xa3, 0x33, 0x50,
	0x63, 0xd0, 0xe3, 0x6a, 0x4d, 0x5c, 0x58, 0xa9, 0xea, 0x2e, 0xf1, 0x02, 0x78, 0x84, 0x96, 0xca,
	0x7d, 0x1d, 0x5e, 0x19, 0x95, 0x77, 0xe5, 0x5f, 0x60, 0xaa, 0x54, 0x6f, 0x03, 0x2b, 0x6b, 0x8a,
	0xd6, 0x55, 0x3e, 0x64, 0x60, 0xf7, 0x27, 0xd4, 0xf8, 0x46, 0x50, 0x3e, 0x74, 0x4a, 0x68, 0xe8,
	0x08, 0x71, 0x52, 0xd9, 0xea, 0xce, 0x83, 0x91, 0xeb, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xc1,
	0x4b, 0x6d, 0x13, 0xef, 0x13, 0x00, 0x00,
}
