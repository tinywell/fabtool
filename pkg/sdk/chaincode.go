package sdk

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-protos-go/common"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/resource"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric/common/policydsl"
)

// InstallCC 合约安装
func (s *Service) InstallCC(reqparam CCParam) error {
	ctx := s.sdk.Context(fabsdk.WithUser(reqparam.User.User), fabsdk.WithOrg(reqparam.User.Org))
	client, err := resmgmt.New(ctx)
	if err != nil {
		return err
	}

	pkg := &resource.CCPackage{
		Type: getCCType(reqparam.Lang),
		Code: reqparam.Package,
	}
	req := resmgmt.InstallCCRequest{
		Name:    reqparam.Name,
		Path:    reqparam.Path,
		Version: reqparam.Version,
		Package: pkg,
	}
	opts := []resmgmt.RequestOption{}
	if len(reqparam.Peers) > 0 {
		opts = append(opts, resmgmt.WithTargetEndpoints(reqparam.Peers...))
	}
	resp, err := client.InstallCC(req, opts...)
	if err != nil {
		return err
	}

	resLog := "合约安装结果：\n"
	for index, r := range resp {
		resLog += fmt.Sprintf("\t %d: status=%d target=%s info=%s\n", index+1, r.Status, r.Target, r.Info)
	}
	s.logger.Info(resLog)
	return nil
}

// InstantiateCC 合约实例化
func (s *Service) InstantiateCC(reqparam CCParam) error {
	ctx := s.sdk.Context(fabsdk.WithUser(reqparam.User.User), fabsdk.WithOrg(reqparam.User.Org))
	client, err := resmgmt.New(ctx)
	if err != nil {
		return err
	}
	var policy *common.SignaturePolicyEnvelope
	if len(reqparam.Policy) == 0 {
		policy = policydsl.AcceptAllPolicy
	} else {
		policy, err = policydsl.FromString(reqparam.Policy)
		if err != nil {
			return err
		}
	}

	req := resmgmt.InstantiateCCRequest{
		Name:    reqparam.Name,
		Path:    reqparam.Path,
		Version: reqparam.Version,
		Lang:    getCCType(reqparam.Lang),
		Args:    reqparam.Args,
		Policy:  policy,
	}
	opts := []resmgmt.RequestOption{}
	_, err = client.InstantiateCC(reqparam.Channel, req, opts...)
	if err != nil {
		return err
	}
	return nil
}

// QueryInstalledCC 查询节点上已安装合约
func (s *Service) QueryInstalledCC(reqparam CCParam) ([]*ChaincodeInfo, error) {
	ctx := s.sdk.Context(fabsdk.WithUser(reqparam.User.User), fabsdk.WithOrg(reqparam.User.Org))
	client, err := resmgmt.New(ctx)
	if err != nil {
		return nil, err
	}
	opts := []resmgmt.RequestOption{}
	if len(reqparam.Peers) > 0 {
		opts = append(opts, resmgmt.WithTargetEndpoints(reqparam.Peers...))
	} else {
		return nil, errors.New("no targets")
	}
	resp, err := client.QueryInstalledChaincodes(opts...)
	if err != nil {
		return nil, err
	}
	return respToCCInfo(resp), nil

}

func getCCType(cclang string) pb.ChaincodeSpec_Type {
	var cctype pb.ChaincodeSpec_Type
	lang := strings.ToUpper(cclang)
	if l, ok := pb.ChaincodeSpec_Type_value[lang]; ok {
		cctype = pb.ChaincodeSpec_Type(l)
	} else if lang == "go" {
		cctype = pb.ChaincodeSpec_GOLANG
	} else {
		cctype = pb.ChaincodeSpec_UNDEFINED
	}
	return cctype
}

func respToCCInfo(resp *pb.ChaincodeQueryResponse) []*ChaincodeInfo {
	ci := make([]*ChaincodeInfo, 0, len(resp.Chaincodes))
	for _, r := range resp.Chaincodes {
		c := &ChaincodeInfo{
			Name:    r.Name,
			Version: r.Version,
			Path:    r.Path,
			Input:   r.Input,
			Escc:    r.Escc,
			Vscc:    r.Vscc,
			ID:      r.Id,
		}
		ci = append(ci, c)
	}
	return ci
}
