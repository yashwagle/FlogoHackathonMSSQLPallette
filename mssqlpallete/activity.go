package mssqlpallete

import (
	"github.com/yashwagle/goLibrary/MSQLPackage"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var activityLog = logger.GetLogger("activity-tibco-mssql")

const (
	methodSelect = "Select"
	methodDML    = "DML"
	methodCreate = "Create"

	ipMethod   = "method"
	ipHost     = "host"
	ipPort     = "port"
	ipDBname   = "dbname"
	ipusername = "username"
	ippassword = "password"
	ipquery    = "query"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {
	method, _ := context.GetInput(ipMethod).(string)
	host, _ := context.GetInput(ipHost).(string)
	port, _ := context.GetInput(ipPort).(string)
	username, _ := context.GetInput(ipusername).(string)
	password, _ := context.GetInput(ippassword).(string)
	dbname, _ := context.GetInput(ipDBname).(string)
	query, _ := context.GetInput(ipquery).(string)

	switch method {
	case methodSelect:
		op, err := MSQLPackage.FireQuery(username, password, host, port, dbname, query)
		if err != nil {
			activityLog.Debugf(err.Error())
			return false, err
		}

		context.SetOutput("output", op)
		return true, nil
	case methodDML:
		op, err := MSQLPackage.UpdateQuery(username, password, host, port, dbname, query)
		if err != nil {
			activityLog.Debugf(err.Error())
			return false, err
		}
		context.SetOutput("output", op)
		return true, nil
	}
	// do eval

	return true, nil
}
