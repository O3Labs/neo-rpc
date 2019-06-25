package neorpc

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type NEORPCInterface interface {
	GetAccountState(address string) (GetAccountStateResponse, error)
	GetBlock(index uint) (GetBlockResponse, error)
	GetBlocks(from uint, to uint) (GetBlocksResponse, error)
	GetRawTransactionWithDetail(hash string) (GetRawTransactionWithDetailResponse, error)
	GetNEP5Balances(address string) (GetNEP5BalancesResponse, error)
	GetAllTransactions(address string) (GetAllTransactionsResponse, error)
	GetContractState(hash string) (GetContractStateResponse, error)
	GetContract(hash string) (GetContractResponse, error)
	GetApplicationLog(txID string) (GetApplicationLogResponse, error)
	GetUnclaimed(address string) (GetUnclaimedResponse, error)
	//generic call to return map[string] interface
	Call(method string, params []interface{}) (map[string]interface{}, error)
}

type NEORPCClient struct {
	Endpoint   url.URL
	httpClient *http.Client
}

//make sure all method interface is implemented
var _ NEORPCInterface = (*NEORPCClient)(nil)

var doOnce sync.Once

func Client(endpoint string) NEORPCClient {

	doOnce.Do(func() {
		log.Printf("run once init endpoint")
		u, err := url.Parse(endpoint)
		if err != nil {
			panic(err)
		}
		var netTransport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 8 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 8 * time.Second,
		}

		var netClient = &http.Client{
			Timeout:   time.Second * 20,
			Transport: netTransport,
		}
		SingletonClient = NEORPCClient{Endpoint: *u, httpClient: netClient}
	})
	return SingletonClient
}

func (n NEORPCClient) makeRequest(method string, params []interface{}, out interface{}) error {
	request := NewRequest(method, params)

	jsonValue, _ := json.Marshal(request)
	req, err := http.NewRequest("POST", n.Endpoint.String(), bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/json")
	res, err := n.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		// Drain and close the body to let the Transport reuse the connection
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
	}()

	err = json.NewDecoder(res.Body).Decode(&out)
	if err != nil {
		return err
	}

	return nil
}

func (n NEORPCClient) GetBlock(index uint) (GetBlockResponse, error) {
	response := GetBlockResponse{}
	params := []interface{}{index, 1}
	err := n.makeRequest("getblock", params, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (n NEORPCClient) GetBlocks(from uint, to uint) (GetBlocksResponse, error) {
	response := GetBlocksResponse{}
	params := []interface{}{from, to}
	err := n.makeRequest("getblocks", params, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (n NEORPCClient) GetRawTransactionWithDetail(hash string) (GetRawTransactionWithDetailResponse, error) {
	response := GetRawTransactionWithDetailResponse{}
	params := []interface{}{hash}
	err := n.makeRequest("getrawtransactionwithdetail", params, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (n NEORPCClient) GetAccountState(address string) (GetAccountStateResponse, error) {
	response := GetAccountStateResponse{}
	params := []interface{}{address}
	err := n.makeRequest("getaccountstate", params, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (n NEORPCClient) GetNEP5Balances(address string) (GetNEP5BalancesResponse, error) {
	response := GetNEP5BalancesResponse{}
	params := []interface{}{address}
	err := n.makeRequest("getnep5balances", params, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (n NEORPCClient) GetAllTransactions(address string) (GetAllTransactionsResponse, error) {
	response := GetAllTransactionsResponse{}
	params := []interface{}{address}
	err := n.makeRequest("getalltransactions", params, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (n NEORPCClient) GetContractState(hash string) (GetContractStateResponse, error) {
	response := GetContractStateResponse{}
	params := []interface{}{hash}
	err := n.makeRequest("getcontractstate", params, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (n NEORPCClient) GetContract(hash string) (GetContractResponse, error) {
	response := GetContractResponse{}
	params := []interface{}{hash}
	err := n.makeRequest("getcontract", params, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (n *NEORPCClient) GetApplicationLog(txID string) (GetApplicationLogResponse, error) {
	response := GetApplicationLogResponse{}
	params := []interface{}{txID}
	err := n.makeRequest("getapplicationlog", params, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (n NEORPCClient) GetUnclaimed(address string) (GetUnclaimedResponse, error) {
	response := GetUnclaimedResponse{}
	params := []interface{}{address}
	err := n.makeRequest("getunclaimed", params, &response)

	if err != nil {
		return response, err
	}
	return response, nil
}

func (n NEORPCClient) Call(method string, params []interface{}) (map[string]interface{}, error) {
	response := map[string]interface{}{}
	err := n.makeRequest(method, params, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
