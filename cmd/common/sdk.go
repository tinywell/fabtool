package common

import (
	"github.com/tinywell/fabtool/common/logging"
	"github.com/tinywell/fabtool/pkg/sdk"
)

// SDKInit ...
func SDKInit(config string) (*sdk.Service, error) {
	newsdk, err := sdk.NewSDK(config)
	if err != nil {
		return nil, err
	}
	logger := logging.DefaultLogger()
	ser, err := sdk.NewService(newsdk, logger)
	if err != nil {
		return nil, err
	}
	return ser, nil
}
