package controllers

import "../core"

var DataBase *core.DataBase = nil

func init() {
	DataBase = core.InitDataBase()
}