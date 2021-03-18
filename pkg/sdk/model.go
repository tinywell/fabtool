package sdk

// CCParam ...
type CCParam struct {
	User      Admin
	Channel   string
	Name      string
	Version   string
	Path      string
	Package   []byte
	Args      [][]byte
	Lang      string // 合约语言类型
	Policy    string
	Peers     []string
	Label     string
	PackageID string
	Sequence  int64
	Init      bool
}

// ExecuteReq 合约执行请求参数
type ExecuteReq struct {
	User         Admin
	Channel      string
	Chaincode    string
	Fcn          string
	Args         [][]byte
	TransientMap map[string][]byte
	Peers        []string
}

// Admin ...
type Admin struct {
	User string
	Org  string
}

// ChaincodeInfo contains general information about an installed/instantiated
// chaincode
type ChaincodeInfo struct {
	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Path    string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Input   string `protobuf:"bytes,4,opt,name=input,proto3" json:"input,omitempty"`
	Escc    string `protobuf:"bytes,5,opt,name=escc,proto3" json:"escc,omitempty"`
	Vscc    string `protobuf:"bytes,6,opt,name=vscc,proto3" json:"vscc,omitempty"`
	ID      []byte `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
}

// LCChaincodeInfo ...
type LCChaincodeInfo struct {
	PackageID  string
	Label      string
	References map[string]string
}

// LCApprovedInfo contains information about the approved chaincode
type LCApprovedInfo struct {
	Name                string
	Version             string
	Sequence            int64
	EndorsementPlugin   string
	ValidationPlugin    string
	SignaturePolicy     string
	ChannelConfigPolicy string
	InitRequired        bool
	PackageID           string
}

// Approval ...
type Approval string

// LCCommitedInfo ...
type LCCommitedInfo struct {
	Name                string
	Version             string
	Sequence            int64
	EndorsementPlugin   string
	ValidationPlugin    string
	SignaturePolicy     string
	ChannelConfigPolicy string
	InitRequired        bool
	Approvals           []Approval
}

// CCResult 合约执行结果
type CCResult struct {
	TxID      string
	Status    int32
	ValidCode string
	Payload   []byte
	Message   string
}
