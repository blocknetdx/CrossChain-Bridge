package dcrm

import (
	"encoding/json"
	"fmt"

	"github.com/fsn-dev/crossChain-Bridge/common"
	"github.com/fsn-dev/crossChain-Bridge/rpc/client"
)

func newWrongStatusError(subject, errInfo string) error {
	return fmt.Errorf("[%v] Wrong status, %v", subject, errInfo)
}

func wrapPostError(method string, err error) error {
	return fmt.Errorf("[post] %v error, %v", method, err)
}

func httpPost(result interface{}, method string, params ...interface{}) error {
	return client.RpcPost(&result, dcrmRpcAddress, method, params...)
}

func GetEnode() (string, error) {
	var result GetEnodeResp
	err := httpPost(&result, "dcrm_getEnode")
	if err != nil {
		return "", wrapPostError("dcrm_getEnode", err)
	}
	if result.Status != "Success" {
		return "", newWrongStatusError("GetEnode", result.Error)
	}
	return result.Data.Enode, nil
}

func GetSignNonce() (uint64, error) {
	var result DataResultResp
	err := httpPost(&result, "dcrm_getSignNonce", keyWrapper.Address.String())
	if err != nil {
		return 0, wrapPostError("dcrm_getSignNonce", err)
	}
	if result.Status != "Success" {
		return 0, newWrongStatusError("GetSignNonce", result.Error)
	}
	bi, err := common.GetBigIntFromStr(result.Data.Result)
	if err != nil {
		return 0, newWrongStatusError("GetSignNonce", err.Error())
	}
	return bi.Uint64(), nil
}

func GetSignStatus(key string) (*SignStatus, error) {
	var result DataResultResp
	err := httpPost(&result, "dcrm_getSignStatus", key)
	if err != nil {
		return nil, wrapPostError("dcrm_getSignStatus", err)
	}
	if result.Status != "Success" {
		return nil, newWrongStatusError("GetSignStatus", "responce error "+result.Error)
	}
	data := result.Data.Result
	var signStatus SignStatus
	json.Unmarshal([]byte(data), &signStatus)
	if signStatus.Status != "Success" {
		return nil, newWrongStatusError("GetSignStatus", "sign status error "+signStatus.Error)
	}
	return &signStatus, nil
}

func GetCurNodeSignInfo() ([]*SignInfoData, error) {
	var result SignInfoResp
	err := httpPost(&result, "dcrm_getCurNodeSignInfo", keyWrapper.Address.String())
	if err != nil {
		return nil, wrapPostError("dcrm_getCurNodeSignInfo", err)
	}
	if result.Status != "Success" {
		return nil, newWrongStatusError("GetCurNodeSignInfo", result.Error)
	}
	return result.Data, nil
}

func Sign(raw string) (string, error) {
	var result DataResultResp
	err := httpPost(&result, "dcrm_sign", raw)
	if err != nil {
		return "", wrapPostError("dcrm_sign", err)
	}
	if result.Status != "Success" {
		return "", newWrongStatusError("Sign", result.Error)
	}
	return result.Data.Result, nil
}

func AcceptSign(raw string) (string, error) {
	var result DataResultResp
	err := httpPost(&result, "dcrm_acceptSign", raw)
	if err != nil {
		return "", wrapPostError("dcrm_acceptSign", err)
	}
	if result.Status != "Success" {
		return "", newWrongStatusError("AcceptSign", result.Error)
	}
	return result.Data.Result, nil
}
