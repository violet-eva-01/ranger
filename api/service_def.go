package api

type ServiceDef struct {
	Id          int    `json:"id"`
	Guid        string `json:"guid"`
	IsEnabled   bool   `json:"isEnabled"`
	CreateTime  int64  `json:"createTime"`
	UpdateTime  int64  `json:"updateTime"`
	Version     int    `json:"version"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	ImplClass   string `json:"implClass"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Options     struct {
		EnableDenyAndExceptionsInPolicies string `json:"enableDenyAndExceptionsInPolicies"`
		UiPages                           string `json:"ui.pages,omitempty"`
	} `json:"options"`
	Configs []struct {
		ItemId            int    `json:"itemId"`
		Name              string `json:"name"`
		Type              string `json:"type"`
		SubType           string `json:"subType,omitempty"`
		Mandatory         bool   `json:"mandatory"`
		ValidationRegEx   string `json:"validationRegEx,omitempty"`
		ValidationMessage string `json:"validationMessage,omitempty"`
		UiHint            string `json:"uiHint,omitempty"`
		Label             string `json:"label,omitempty"`
		DefaultValue      string `json:"defaultValue,omitempty"`
	} `json:"configs"`
	Resources []struct {
		ItemId             int    `json:"itemId"`
		Name               string `json:"name"`
		Type               string `json:"type"`
		Level              int    `json:"level"`
		Mandatory          bool   `json:"mandatory"`
		LookupSupported    bool   `json:"lookupSupported"`
		RecursiveSupported bool   `json:"recursiveSupported"`
		ExcludesSupported  bool   `json:"excludesSupported"`
		Matcher            string `json:"matcher,omitempty"`
		MatcherOptions     struct {
			WildCard          string `json:"wildCard,omitempty"`
			IgnoreCase        string `json:"ignoreCase,omitempty"`
			PathSeparatorChar string `json:"pathSeparatorChar,omitempty"`
		} `json:"matcherOptions"`
		ValidationRegEx        string   `json:"validationRegEx,omitempty"`
		ValidationMessage      string   `json:"validationMessage,omitempty"`
		UiHint                 string   `json:"uiHint,omitempty"`
		Label                  string   `json:"label"`
		Description            string   `json:"description"`
		AccessTypeRestrictions []string `json:"accessTypeRestrictions"`
		IsValidLeaf            bool     `json:"isValidLeaf"`
		Parent                 string   `json:"parent,omitempty"`
	} `json:"resources"`
	AccessTypes []struct {
		ItemId        int      `json:"itemId"`
		Name          string   `json:"name"`
		Label         string   `json:"label"`
		ImpliedGrants []string `json:"impliedGrants"`
	} `json:"accessTypes"`
	PolicyConditions []struct {
		ItemId           int    `json:"itemId"`
		Name             string `json:"name"`
		Evaluator        string `json:"evaluator"`
		EvaluatorOptions struct {
			ScriptTemplate string `json:"scriptTemplate,omitempty"`
			EngineName     string `json:"engineName,omitempty"`
			UiIsMultiline  string `json:"ui.isMultiline,omitempty"`
		} `json:"evaluatorOptions"`
		ValidationRegEx   string `json:"validationRegEx,omitempty"`
		ValidationMessage string `json:"validationMessage,omitempty"`
		UiHint            string `json:"uiHint,omitempty"`
		Label             string `json:"label"`
		Description       string `json:"description"`
	} `json:"policyConditions"`
	ContextEnrichers []struct {
		ItemId          int    `json:"itemId"`
		Name            string `json:"name"`
		Enricher        string `json:"enricher"`
		EnricherOptions struct {
			TagRetrieverClassName       string `json:"tagRetrieverClassName"`
			TagRefresherPollingInterval string `json:"tagRefresherPollingInterval"`
		} `json:"enricherOptions"`
	} `json:"contextEnrichers"`
	Enums []struct {
		ItemId   int    `json:"itemId"`
		Name     string `json:"name"`
		Elements []struct {
			ItemId int    `json:"itemId"`
			Name   string `json:"name"`
			Label  string `json:"label"`
		} `json:"elements"`
		DefaultIndex int `json:"defaultIndex"`
	} `json:"enums"`
	DataMaskDef struct {
		MaskTypes []struct {
			ItemId          int    `json:"itemId"`
			Name            string `json:"name"`
			Label           string `json:"label"`
			Description     string `json:"description"`
			Transformer     string `json:"transformer,omitempty"`
			DataMaskOptions struct {
			} `json:"dataMaskOptions"`
		} `json:"maskTypes"`
		AccessTypes []struct {
			ItemId        int           `json:"itemId"`
			Name          string        `json:"name"`
			Label         string        `json:"label"`
			ImpliedGrants []interface{} `json:"impliedGrants"`
		} `json:"accessTypes"`
		Resources []struct {
			ItemId             int    `json:"itemId"`
			Name               string `json:"name"`
			Type               string `json:"type"`
			Level              int    `json:"level"`
			Mandatory          bool   `json:"mandatory"`
			LookupSupported    bool   `json:"lookupSupported"`
			RecursiveSupported bool   `json:"recursiveSupported"`
			ExcludesSupported  bool   `json:"excludesSupported"`
			Matcher            string `json:"matcher"`
			MatcherOptions     struct {
				WildCard               string `json:"wildCard"`
				IgnoreCase             string `json:"ignoreCase"`
				IsValidLeaf            string `json:"__isValidLeaf,omitempty"`
				AccessTypeRestrictions string `json:"__accessTypeRestrictions,omitempty"`
			} `json:"matcherOptions"`
			ValidationRegEx        string   `json:"validationRegEx"`
			ValidationMessage      string   `json:"validationMessage"`
			UiHint                 string   `json:"uiHint"`
			Label                  string   `json:"label"`
			Description            string   `json:"description"`
			AccessTypeRestrictions []string `json:"accessTypeRestrictions"`
			IsValidLeaf            bool     `json:"isValidLeaf"`
			Parent                 string   `json:"parent,omitempty"`
		} `json:"resources"`
	} `json:"dataMaskDef"`
	RowFilterDef struct {
		AccessTypes []struct {
			ItemId        int           `json:"itemId"`
			Name          string        `json:"name"`
			Label         string        `json:"label"`
			ImpliedGrants []interface{} `json:"impliedGrants"`
		} `json:"accessTypes"`
		Resources []struct {
			ItemId             int    `json:"itemId"`
			Name               string `json:"name"`
			Type               string `json:"type"`
			Level              int    `json:"level"`
			Mandatory          bool   `json:"mandatory"`
			LookupSupported    bool   `json:"lookupSupported"`
			RecursiveSupported bool   `json:"recursiveSupported"`
			ExcludesSupported  bool   `json:"excludesSupported"`
			Matcher            string `json:"matcher"`
			MatcherOptions     struct {
				WildCard   string `json:"wildCard"`
				IgnoreCase string `json:"ignoreCase"`
			} `json:"matcherOptions"`
			ValidationRegEx        string   `json:"validationRegEx"`
			ValidationMessage      string   `json:"validationMessage"`
			UiHint                 string   `json:"uiHint"`
			Label                  string   `json:"label"`
			Description            string   `json:"description"`
			AccessTypeRestrictions []string `json:"accessTypeRestrictions"`
			IsValidLeaf            bool     `json:"isValidLeaf"`
			Parent                 string   `json:"parent,omitempty"`
		} `json:"resources"`
	} `json:"rowFilterDef"`
	CreatedBy string `json:"createdBy,omitempty"`
	UpdatedBy string `json:"updatedBy,omitempty"`
}
