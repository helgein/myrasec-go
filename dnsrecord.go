package myrasec

import (
	"fmt"

	"github.com/Myra-Security-GmbH/myrasec-go/pkg/types"
)

//
// DNSRecord ...
//
type DNSRecord struct {
	ID               int             `json:"id,omitempty"`
	Created          *types.DateTime `json:"created,omitempty"`
	Modified         *types.DateTime `json:"modified,omitempty"`
	Name             string          `json:"name"`
	TTL              int             `json:"ttl"`
	RecordType       string          `json:"recordType"`
	AlternativeCNAME string          `json:"alternativeCname,omitempty"`
	Active           bool            `json:"active"`
	Value            string          `json:"value"`
	Priority         int             `json:"priority,omitempty"`
	Port             int             `json:"port,omitempty"`
}

//
// ListDNSRecords returns a slice containing all visible DNS records for a domain
//
func (api *API) ListDNSRecords(domainName string) ([]DNSRecord, error) {
	if _, ok := methods["listDNSRecords"]; !ok {
		return nil, fmt.Errorf("Passed action [%s] is not supported", "listDNSRecords")
	}
	definition := methods["listDNSRecords"]
	definition.Action = fmt.Sprintf(definition.Action, domainName, 1)

	result, err := api.call(definition)
	if err != nil {
		return nil, err
	}
	var records []DNSRecord
	for _, v := range *result.(*[]DNSRecord) {
		records = append(records, v)
	}

	return records, nil
}

//
// CreateDNSRecord creates a new DNS record using the MYRA API
//
func (api *API) CreateDNSRecord(record *DNSRecord, domainName string) (*DNSRecord, error) {
	if _, ok := methods["createDNSRecord"]; !ok {
		return nil, fmt.Errorf("Passed action [%s] is not supported", "createDNSRecord")
	}
	definition := methods["createDNSRecord"]
	definition.Action = fmt.Sprintf(definition.Action, domainName)

	result, err := api.call(definition, record)
	if err != nil {
		return nil, err
	}
	return result.(*DNSRecord), nil
}

//
// UpdateDNSRecord updates the passed DNS record using the MYRA API
//
func (api *API) UpdateDNSRecord(record *DNSRecord, domainName string) (*DNSRecord, error) {
	if _, ok := methods["updateDNSRecord"]; !ok {
		return nil, fmt.Errorf("Passed action [%s] is not supported", "updateDNSRecord")
	}
	definition := methods["updateDNSRecord"]
	definition.Action = fmt.Sprintf(definition.Action, domainName)

	result, err := api.call(definition, record)
	if err != nil {
		return nil, err
	}
	return result.(*DNSRecord), nil
}

//
// DeleteDNSRecord deletes the passed DNS record using the MYRA API
//
func (api *API) DeleteDNSRecord(record *DNSRecord, domainName string) (*DNSRecord, error) {
	if _, ok := methods["deleteDNSRecord"]; !ok {
		return nil, fmt.Errorf("Passed action [%s] is not supported", "deleteDNSRecord")
	}
	definition := methods["deleteDNSRecord"]
	definition.Action = fmt.Sprintf(definition.Action, domainName)

	result, err := api.call(definition, record)
	if err != nil {
		return nil, err
	}
	return result.(*DNSRecord), nil
}
