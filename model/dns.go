package model

import "github.com/csvwolf/ker.go/constant"

type Record struct {
	Header string              `json:"header"` // 域名头部，空为@
	Type   constant.DomainType `json:"type"`   // 类型
	Ttl    uint32              `json:"ttl"`    // 生命周期
	Value  string              `json:"value"`  // 记录值
}

// DomainList
type DomainListRequest struct {
	Action string `json:"action"`
}

type DomainListResponseItem struct {
	// 域名
	Domain string `json:"domain"`
	// Type 释义：https://doc.hostker.dev/dns/
	Type uint8 `json:"type"`
}

type DomainListResponse struct {
	Domains []*DomainListResponseItem `json:"domains"`
}

// Records
type RecordsRequest struct {
	Action string `json:"action"`
	Domain string `json:"domain"`
}

type RecordsResponse struct {
	Records []*Record `json:"records"`
}

// CreateRecord
type CreateRecordRequest struct {
	Action string              `json:"action"`
	Domain string              `json:"domain"`
	Header string              `json:"header"`
	Type   constant.DomainType `json:"type"`
	Ttl    uint32              `json:"ttl"`
	Value  string              `json:"value"`
}

type CreateRecordResponse = []struct{}

// UpdateRecord
type UpdateRecordRequest struct {
	Action    string  `json:"action"`
	Domain    string  `json:"domain"`
	OldRecord *Record `json:"old_record"`
	NewRecord *Record `json:"new_record"`
}

type UpdateRecordResponse = []struct{}

// DeleteRecord
type DeleteRecordRequest struct {
	Action string              `json:"action"`
	Domain string              `json:"domain"`
	Header string              `json:"header"`
	Type   constant.DomainType `json:"type"`
	Ttl    uint32              `json:"ttl"`
	Value  string              `json:"value"`
}

type DeleteRecordResponse = []struct{}
