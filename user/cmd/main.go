// Copyright 2019 money-hub. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// money-hub MoneyDodo/user
//
// This documentation describes example APIs found under https://github.com/money-hub/MoneyDodo.service
//
//     Schemes: http
//     Version: 1.0.0
//     License: MIT http://opensource.org/licenses/MIT
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
package main

import (
	_ "github.com/money-hub/MoneyDodo.service/swagger"
	service "github.com/money-hub/MoneyDodo.service/user/cmd/service"
)

func main() {
	service.Run()
}
