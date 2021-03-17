package sdk

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"

	"github.com/tinywell/fabtool/pkg/core"
)

// NewSDK ...
func NewSDK(configYaml string) (*fabsdk.FabricSDK, error) {
	configProvider := config.FromFile(configYaml)
	sdk, err := fabsdk.New(configProvider)
	if err != nil {
		return nil, err
	}
	return sdk, nil
}

// Service ...
type Service struct {
	sdk    *fabsdk.FabricSDK
	logger core.Logger
}

// NewService ...
func NewService(sdk *fabsdk.FabricSDK, logger core.Logger) (*Service, error) {
	return &Service{
		sdk:    sdk,
		logger: logger,
	}, nil
}
