// Package client @author: Violet-Eva @date  : 2025/8/31 @notes :
package client

import (
	"github.com/violet-eva-01/ranger/client/functions"
	"strings"
)

type PluginsDefinitions struct {
	StartIndex  int          `json:"startIndex"`
	PageSize    int          `json:"pageSize"`
	TotalCount  int          `json:"totalCount"`
	ResultSize  int          `json:"resultSize"`
	QueryTimeMS int64        `json:"queryTimeMS"`
	ServiceDefs []ServiceDef `json:"serviceDefs"`
}

func (p *PluginsDefinitions) GetServiceTypesIds() []ServiceTypeId {
	var sti []ServiceTypeId
	for _, sd := range p.ServiceDefs {
		index := functions.FindIndex(strings.ToLower(sd.Name), serviceTypeName)
		if index >= 0 {
			var tmpSTI ServiceTypeId
			tmpSTI.ServiceTypeId = index
			tmpSTI.ServiceType = ServiceType(index)
			sti = append(sti, tmpSTI)
		}
	}
	return sti
}
