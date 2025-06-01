package response

type BaseRS struct {
	APIVersion string      `json:"apiVersion"`
	Data       interface{} `json:"data"`
	Error      interface{} `json:"error"`
}

type ItemsRs struct {
	Items interface{} `json:"items"`
}

type PageRs struct {
	Items            interface{} `json:"items"`
	TotalItems       int         `json:"totalItems"`
	ItemsPerPage     interface{} `json:"itemsPerPage"`
	StartIndex       int         `json:"startIndex"`
	CurrentItemCount int         `json:"currentItemCount"`
	LastPage         bool        `json:"lastPage"`
}
