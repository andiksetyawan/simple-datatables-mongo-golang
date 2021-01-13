package main

type Query struct {
	SEcho          int64  `form:"sEcho"`
	IColumns       string `form:"iColumns"`
	SColumns       string `form:"sColumns"`
	IDisplayStart  int64  `form:"iDisplayStart"`
	IDisplayLength int64  `form:"iDisplayLength"`
	MDataProp0     string `form:"mDataProp_0"`
	SSearch0       string `form:"sSearch_0"`
	BRegex0        string `form:"bRegex_0"`
	BSearchable0   string `form:"bSearchable_0"`
	BSortable0     string `form:"bSortable_0"`
	MDataProp1     string `form:"mDataProp_1"`
	SSearch1       string `form:"sSearch_1"`
	BRegex1        string `form:"bRegex_1"`
	BSearchable1   string `form:"bSearchable_1"`
	BSortable1     string `form:"bSortable_1"`
	MDataProp2     string `form:"mDataProp_2"`
	SSearch2       string `form:"sSearch_2"`
	BRegex2        string `form:"bRegex_2"`
	BSearchable2   string `form:"bSearchable_2"`
	BSortable2     string `form:"bSortable_2"`
	MDataProp3     string `form:"mDataProp_3"`
	SSearch3       string `form:"sSearch_3"`
	BRegex3        string `form:"bRegex_3"`
	BSearchable3   string `form:"bSearchable_3"`
	BSortable3     string `form:"bSortable_3"`
	SSearch        string `form:"sSearch"`
	BRegex         string `form:"bRegex"`
	ISortCol0      string `form:"iSortCol_0"`
	SSortDir0      string `form:"sSortDir_0"`
	ISortingCols   int    `form:"iSortingCols"`
	NAMING_FAILED  string `form:"_"`
}

type Data struct {
	SEcho                int64       `json:"sEcho"`
	ITotalRecords        int64       `json:"iTotalRecords"`
	ITotalDisplayRecords int64       `json:"iTotalDisplayRecords"`
	Data                 interface{} `json:"data"`
}
