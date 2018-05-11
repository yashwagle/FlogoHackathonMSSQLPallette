package mssqlpallete

import (
	"errors"
	"strings"

	"github.com/yashwagle/goLibrary/MSQLPackage"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var activityLog = logger.GetLogger("activity-tibco-mssql")

const (
	methodSelect = "DQL"
	methodDML    = "DML"
	methodCreate = "DDL"

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
	query = strings.TrimSpace(query)

	switch method {
	case methodSelect: //Select Queries
		if strings.TrimPrefix(query, " ") != "Select" {
			err := errors.New("Not Select Query")
			activityLog.Errorf(err.Error())
			return false, err
		}
		op, err := MSQLPackage.FireQuery(username, password, host, port, dbname, query)
		if err != nil {
			activityLog.Errorf(err.Error())
			return false, err
		}

		context.SetOutput("output", op)
		return true, nil

	case methodDML: //DML queries
		if strings.TrimPrefix(query, " ") != "Update" && strings.TrimPrefix(query, " ") != "Delete" && strings.TrimPrefix(query, " ") != "Insert" {
			err := errors.New("Not DQL Query")
			activityLog.Errorf(err.Error())
			return false, err
		}
		op, err := MSQLPackage.UpdateQuery(username, password, host, port, dbname, query)
		if err != nil {
			activityLog.Errorf(err.Error())
			return false, err
		}
		context.SetOutput("output", op)
		return true, nil

	case methodCreate: //Create Query
		op, err := MSQLPackage.CreateQuery(username, password, host, port, dbname, query)
		if err != nil {
			activityLog.Errorf(err.Error())
			return false, err
		}
		context.SetOutput("output", op)
		return true, nil

	}

	return true, nil
}
