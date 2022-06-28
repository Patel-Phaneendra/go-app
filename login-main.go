//go:generate go run github.com/vektra/mockery/v2 --all --with-expecter --inpackage

package main

import (
	"github.hdfcbank.com/HDFCBANK/mb-microservices-template/login/cmd/prepare"
)

func main() {
	prepare.Prepare()
}
