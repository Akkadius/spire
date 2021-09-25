package crudcontrollers

type BulkFetchByIdsGetRequest struct {
	IDs []int `json:"ids"`
}
