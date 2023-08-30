package api

import (
	"github.com/csvwolf/ker.go/model"
)

// GetDomainList 获取域名列表 https://doc.hostker.dev/dns/domainList.html
func (ker *Ker) GetDomainList() (*model.DomainListResponse, error) {
	const action = "DomainDNS.getDomainList"
	var (
		response *model.DomainListResponse
		err      error
	)

	if response, err = fetch[model.DomainListRequest, model.DomainListResponse](
		ker, &model.DomainListRequest{Action: action},
	); err != nil {
		return nil, err
	}

	return response, nil
}

// GetDomainRecords 获取解析记录列表 https://doc.hostker.dev/dns/getRecords.html
func (ker *Ker) GetDomainRecords(domain string) (*model.RecordsResponse, error) {
	const action = "DomainDNS.getRecords"
	var (
		response *model.RecordsResponse
		err      error
	)
	if response, err = fetch[model.RecordsRequest, model.RecordsResponse](
		ker, &model.RecordsRequest{Action: action, Domain: domain},
	); err != nil {
		return nil, err
	}

	return response, nil
}

// CreateDomainRecord 添加解析记录 https://doc.hostker.dev/dns/createRecord.html
func (ker *Ker) CreateDomainRecord(domain string, params *model.Record) (*model.CreateRecordResponse, error) {
	const action = "DomainDNS.createRecord"
	var (
		response *model.CreateRecordResponse
		err      error
		request  = &model.CreateRecordRequest{
			Action: action,
			Domain: domain,
			Header: params.Header,
			Type:   params.Type,
			Ttl:    params.Ttl,
			Value:  params.Value,
		}
	)
	if response, err = fetch[model.CreateRecordRequest, model.CreateRecordResponse](ker, request); err != nil {
		return nil, err
	}
	return response, nil
}

// UpdateDomainRecord 修改解析记录（注意 value 填写完整） https://doc.hostker.dev/dns/updateRecord.html
func (ker *Ker) UpdateDomainRecord(domain string, oldValue *model.Record, newValue *model.Record) (*model.UpdateRecordResponse, error) {
	const action = "DomainDNS.updateRecord"
	var (
		response *model.UpdateRecordResponse
		err      error
		request  = &model.UpdateRecordRequest{
			Action:    action,
			Domain:    domain,
			OldRecord: oldValue,
			NewRecord: newValue,
		}
	)
	if response, err = fetch[model.UpdateRecordRequest, model.UpdateRecordResponse](ker, request); err != nil {
		return nil, err
	}
	return response, nil
}

// DeleteDomainRecord 删除解析记录（请注意填写完整信息） https://doc.hostker.dev/dns/deleteRecord.html
func (ker *Ker) DeleteDomainRecord(domain string, record *model.Record) (*model.DeleteRecordResponse, error) {
	const action = "DomainDNS.deleteRecord"
	var (
		response *model.DeleteRecordResponse
		err      error
		request  = &model.DeleteRecordRequest{
			Action: action,
			Domain: domain,
			Header: record.Header,
			Type:   record.Type,
			Value:  record.Value,
			Ttl:    record.Ttl,
		}
	)
	if response, err = fetch[model.DeleteRecordRequest, model.DeleteRecordResponse](ker, request); err != nil {
		return nil, err
	}
	return response, nil
}
