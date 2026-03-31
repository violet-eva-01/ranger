// Package client @author: Violet-Eva @date  : 2026/3/30 @notes :
package client

import (
	"fmt"

	"github.com/violet-eva-01/ranger/types"
)

func (c *Client) GetServiceDefs() ([]types.ServiceDef, error) {
	var (
		sd         []types.ServiceDef
		startIndex = 0
		pageSize   = 1000
	)
	for {
		var pd types.PluginsDefinitions
		if err := c.RequestToStruct("GET", fmt.Sprintf("/plugins/definitions?startIndex=%d&pageSize=%d", startIndex, pageSize),
			nil, &pd); err != nil {
			return nil, err
		}
		sd = append(sd, pd.ServiceDefs...)
		if pd.ResultSize == 0 || (pd.ResultSize > 0 && pd.ResultSize < pd.PageSize) {
			break
		} else {
			startIndex += pageSize
		}
	}
	return sd, nil
}
