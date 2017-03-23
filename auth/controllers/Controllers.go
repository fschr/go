package controllers

import "github.com/fschr/go/auth/core"

var DataBase *core.DataBase = nil

func init() {
	DataBase = core.InitDataBase()
}
