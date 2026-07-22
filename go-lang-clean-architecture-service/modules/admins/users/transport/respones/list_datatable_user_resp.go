package respones

type ListDatatableUserResp struct {
	Draw            int                  `json:"draw"`
	RecordsTotal    int64                `json:"recordsTotal"`
	RecordsFiltered int64                `json:"recordsFiltered"`
	Data            *[]DataTableUserResp `json:"data"`
}
