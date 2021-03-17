package sdk

import lcpackger "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/lifecycle"

// LCPackageCC 使用 lifecycle 方式打包链码
func LCPackageCC(reqparam CCParam) ([]byte, error) {
	desc := &lcpackger.Descriptor{
		Path:  reqparam.Path,
		Type:  getCCType(reqparam.Lang),
		Label: reqparam.Label,
	}
	err := desc.Validate()
	if err != nil {
		return nil, err
	}
	return lcpackger.NewCCPackage(desc)
}
