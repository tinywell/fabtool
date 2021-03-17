package sdk

import (
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric/common/policydsl"
)

// LCPackageCC 使用 lifecycle 方式打包链码
func (s *Service) LCPackageCC(reqparam CCParam) ([]byte, error) {
	return LCPackageCC(reqparam)
}

// LCInstallCC 使用 lifecycle 方式安装链码
func (s *Service) LCInstallCC(reqparam CCParam) (string, error) {
	ctx := s.sdk.Context(fabsdk.WithUser(reqparam.User.User), fabsdk.WithOrg(reqparam.User.Org))
	client, err := resmgmt.New(ctx)
	if err != nil {
		return "", err
	}
	req := resmgmt.LifecycleInstallCCRequest{
		Label:   reqparam.Label,
		Package: reqparam.Package,
	}
	opts := []resmgmt.RequestOption{}
	if len(reqparam.Peers) > 0 {
		opts = append(opts, resmgmt.WithTargetEndpoints(reqparam.Peers...))
	}
	resp, err := client.LifecycleInstallCC(req, opts...)
	if err != nil {
		return "", err
	}
	resLog := "合约安装结果：\n"
	for index, r := range resp {
		resLog += fmt.Sprintf("\t %d: status=%d target=%s package ID=%s\n", index+1, r.Status, r.Target, r.PackageID)
	}
	s.logger.Info(resLog)
	return resp[0].PackageID, nil
}

// LCQueryInstalledCC 查询通过 lifecycle 安装的链码信息
func (s *Service) LCQueryInstalledCC(reqparam CCParam) ([]*LCChaincodeInfo, error) {
	ctx := s.sdk.Context(fabsdk.WithUser(reqparam.User.User), fabsdk.WithOrg(reqparam.User.Org))
	client, err := resmgmt.New(ctx)
	if err != nil {
		return nil, err
	}
	opts := []resmgmt.RequestOption{}
	if len(reqparam.Peers) > 0 {
		opts = append(opts, resmgmt.WithTargetEndpoints(reqparam.Peers...))
	}
	resp, err := client.LifecycleQueryInstalledCC(opts...)
	if err != nil {
		return nil, err
	}
	return respToLCCCInfo(resp), nil
}

// LCGetInstalledCC 下载 lifecycle 方式安装的链码包
func (s *Service) LCGetInstalledCC(reqparam CCParam) ([]byte, error) {
	ctx := s.sdk.Context(fabsdk.WithUser(reqparam.User.User), fabsdk.WithOrg(reqparam.User.Org))
	client, err := resmgmt.New(ctx)
	if err != nil {
		return nil, err
	}
	opts := []resmgmt.RequestOption{}
	if len(reqparam.Peers) > 0 {
		opts = append(opts, resmgmt.WithTargetEndpoints(reqparam.Peers...))
	}
	return client.LifecycleGetInstalledCCPackage(reqparam.PackageID, opts...)
}

// LCApproveCC 使用 lifecycle 方式为组织审批链码
func (s *Service) LCApproveCC(reqparam CCParam) error {
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
	req := resmgmt.LifecycleApproveCCRequest{
		Name:            reqparam.Name,
		Version:         reqparam.Version,
		PackageID:       reqparam.PackageID,
		Sequence:        reqparam.Sequence,
		SignaturePolicy: policy,
		InitRequired:    reqparam.Init,
	}
	opts := []resmgmt.RequestOption{}
	if len(reqparam.Peers) > 0 {
		opts = append(opts, resmgmt.WithTargetEndpoints(reqparam.Peers...))
	}
	txid, err := client.LifecycleApproveCC(reqparam.Channel, req, opts...)
	if err != nil {
		return err
	}
	s.logger.Debugf("approve cc txid=%s", string(txid))
	return nil
}

// LCQueryApproved 查询组织提交的对链码的审批信息
func (s *Service) LCQueryApproved(reqparam CCParam) (*LCApprovedInfo, error) {
	ctx := s.sdk.Context(fabsdk.WithUser(reqparam.User.User), fabsdk.WithOrg(reqparam.User.Org))
	client, err := resmgmt.New(ctx)
	if err != nil {
		return nil, err
	}
	req := resmgmt.LifecycleQueryApprovedCCRequest{
		Name:     reqparam.Name,
		Sequence: reqparam.Sequence,
	}
	opts := []resmgmt.RequestOption{}
	if len(reqparam.Peers) > 0 {
		opts = append(opts, resmgmt.WithTargetEndpoints(reqparam.Peers...))
	}
	resp, err := client.LifecycleQueryApprovedCC(reqparam.Channel, req, opts...)
	if err != nil {
		return nil, err
	}
	return respToApproveInfo(resp), nil
}

// LCCheckCommit 查询链码的审批情况
func (s *Service) LCCheckCommit(reqparam CCParam) ([]Approval, error) {
	ctx := s.sdk.Context(fabsdk.WithUser(reqparam.User.User), fabsdk.WithOrg(reqparam.User.Org))
	client, err := resmgmt.New(ctx)
	if err != nil {
		return nil, err
	}
	opts := []resmgmt.RequestOption{}
	if len(reqparam.Peers) > 0 {
		opts = append(opts, resmgmt.WithTargetEndpoints(reqparam.Peers...))
	}
	var policy *common.SignaturePolicyEnvelope
	if len(reqparam.Policy) == 0 {
		policy = policydsl.AcceptAllPolicy
	} else {
		policy, err = policydsl.FromString(reqparam.Policy)
		if err != nil {
			return nil, err
		}
	}
	req := resmgmt.LifecycleCheckCCCommitReadinessRequest{
		Name:            reqparam.Name,
		Version:         reqparam.Version,
		Sequence:        reqparam.Sequence,
		InitRequired:    reqparam.Init,
		SignaturePolicy: policy,
	}
	resp, err := client.LifecycleCheckCCCommitReadiness(reqparam.Channel, req, opts...)
	if err != nil {
		return nil, err
	}
	return respToApprovels(resp), nil

}

// LCCommitCC 使用 lifecycle 方式提交链码
func (s *Service) LCCommitCC(reqparam CCParam) error {
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
	opts := []resmgmt.RequestOption{}
	if len(reqparam.Peers) > 0 {
		opts = append(opts, resmgmt.WithTargetEndpoints(reqparam.Peers...))
	}
	req := resmgmt.LifecycleCommitCCRequest{
		Name:            reqparam.Name,
		Version:         reqparam.Version,
		Sequence:        reqparam.Sequence,
		SignaturePolicy: policy,
		InitRequired:    reqparam.Init,
	}
	txid, err := client.LifecycleCommitCC(reqparam.Channel, req, opts...)
	if err != nil {
		return err
	}
	s.logger.Infof("commit cc txid: %s", string(txid))
	return nil
}

// LCQueryCommited 查询已提交的链码信息
func (s *Service) LCQueryCommitted(reqparam CCParam) ([]*LCCommitedInfo, error) {
	ctx := s.sdk.Context(fabsdk.WithUser(reqparam.User.User), fabsdk.WithOrg(reqparam.User.Org))
	client, err := resmgmt.New(ctx)
	if err != nil {
		return nil, err
	}
	req := resmgmt.LifecycleQueryCommittedCCRequest{
		Name: reqparam.Name,
	}
	opts := []resmgmt.RequestOption{}
	if len(reqparam.Peers) > 0 {
		opts = append(opts, resmgmt.WithTargetEndpoints(reqparam.Peers...))
	}
	resp, err := client.LifecycleQueryCommittedCC(reqparam.Channel, req, opts...)
	if err != nil {
		return nil, err
	}
	return respToCommittedInfo(resp), nil
}

func respToLCCCInfo(resp []resmgmt.LifecycleInstalledCC) []*LCChaincodeInfo {
	ci := make([]*LCChaincodeInfo, 0, len(resp))
	for _, r := range resp {
		ref := make(map[string]string)
		for k, v := range r.References {
			refs := make([]string, 0, len(v))
			for _, rf := range v {
				rfstr := rf.Name + "." + rf.Version
				refs = append(refs, rfstr)
			}
			ref[k] = strings.Join(refs, "|")
		}
		c := &LCChaincodeInfo{
			PackageID:  r.PackageID,
			Label:      r.Label,
			References: ref,
		}
		ci = append(ci, c)
	}
	return ci
}

func respToApproveInfo(resp resmgmt.LifecycleApprovedChaincodeDefinition) *LCApprovedInfo {
	return &LCApprovedInfo{
		Name:                resp.Name,
		Version:             resp.Version,
		Sequence:            resp.Sequence,
		EndorsementPlugin:   resp.EndorsementPlugin,
		ValidationPlugin:    resp.ValidationPlugin,
		SignaturePolicy:     resp.SignaturePolicy.String(),
		ChannelConfigPolicy: resp.ChannelConfigPolicy,
		InitRequired:        resp.InitRequired,
		PackageID:           resp.PackageID,
	}
}

func respToApprovels(resp resmgmt.LifecycleCheckCCCommitReadinessResponse) []Approval {
	return transApprovels(resp.Approvals)
}

func transApprovels(approvels map[string]bool) []Approval {
	res := make([]Approval, 0, len(approvels))
	for k, v := range approvels {
		ap := Approval(fmt.Sprintf("%s:%t", k, v))
		res = append(res, ap)
	}
	return res
}

func respToCommittedInfo(resp []resmgmt.LifecycleChaincodeDefinition) []*LCCommitedInfo {
	res := make([]*LCCommitedInfo, 0, len(resp))
	for _, r := range resp {
		ci := &LCCommitedInfo{
			Name:                r.Name,
			Version:             r.Version,
			Sequence:            r.Sequence,
			EndorsementPlugin:   r.EndorsementPlugin,
			ValidationPlugin:    r.ValidationPlugin,
			SignaturePolicy:     r.SignaturePolicy.String(),
			ChannelConfigPolicy: r.ChannelConfigPolicy,
			InitRequired:        r.InitRequired,
			Approvals:           transApprovels(r.Approvals),
		}
		res = append(res, ci)
	}
	return res
}
