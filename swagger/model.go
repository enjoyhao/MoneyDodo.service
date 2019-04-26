package swagger

import (
	"github.com/money-hub/MoneyDodo.service/model"
)

// Create User request
// swagger:parameters swaggCreateUserReq
type swaggCreateUserReq struct {
	// in:body
	Body model.User
}

// HTTP status code 200 and user model in data
// swagger:response swaggUserResp
type swaggUserResp struct {
	// in:body
	Body struct {
		// Boolean true/false
		Status bool `json:"status"`
		// Detailed error message
		Errinfo string `json:"errinfo"`
		// User model
		Data *model.User `json:"data"`
	}
}

// HTTP status code 200 and an array of user models in data
// swagger:response swaggUsersResp
type swaggUsersResp struct {
	// in:body
	Body struct {
		// Boolean true/false
		Status bool `json:"status"`
		// Detailed error message
		Errinfo string `json:"errinfo"`
		// User model
		Data []model.User `json:"data"`
	}
}

// HTTP status code 200 and task model in data
// swagger:response swaggTaskResp
type swaggTaskResp struct {
	// in:body
	Body struct {
		// Boolean true/false
		Status bool `json:"status"`
		// Detailed error message
		Errinfo string `json:"errinfo"`
		// Task model
		Data *model.Task `json:"data"`
	}
}

// HTTP status code 200 and an array of task models in data
// swagger:response swaggTasksResp
type swaggTasksResp struct {
	// in:body
	Body struct {
		// Boolean true/false
		Status bool `json:"status"`
		// Detailed error message
		Errinfo string `json:"errinfo"`
		// Tasks model
		Data []model.Task `json:"data"`
	}
}

// Error Fail
// swagger:response swaggBadReq
type swaggBadReq struct {
	// in:body
	Body struct {
		// HTTP Status Code 200
		Status bool `json:"status"`
		// Detailed error message
		Errinfo string `json:"errinfo"`
	}
}

// HTTP status code 200 and no return value
// swagger:response swaggNoReturnValue
type swaggNoReturnValue struct {
	// in:body
	Body struct {
		// HTTP Status Code 200
		Status bool `json:"status"`
		// Detailed error message
		Errinfo string `json:"errinfo"`
	}
}