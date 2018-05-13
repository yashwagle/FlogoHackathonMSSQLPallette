package mssqlpallete

import (
	"errors"
	"strings"

	"github.com/yashwagle/goLibrary/mssqlpackage"

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
	iptimeout  = "timeout"
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
	timeout, _ := context.GetInput(iptimeout).(int)
	query = strings.TrimSpace(query)
	if timeout < 0 {
		err := errors.New("negative timeout not allowed")
		activityLog.Errorf(err.Error())
		return false, err
	}
	if timeout == 0 {
		timeout = 15
	}

	switch method {
	case methodSelect: //Select Queries
		operation := strings.Split(query, " ")[0]
		if strings.ToLower(operation) != "select" {
			err := errors.New("Not Select Query" + operation)
			activityLog.Errorf(err.Error())
			return false, err
		}
		op, err := mssqlpackage.FireQuery(username, password, host, port, dbname, query, timeout)
		if err != nil {
			activityLog.Errorf(err.Error())
			return false, err
		}

		context.SetOutput("output", op)
		return true, nil

	case methodDML: //DML queries
		operation := strings.Split(query, " ")[0]
		if strings.ToLower(operation) != "update" && strings.ToLower(operation) != "delete" && strings.ToLower(operation) != "insert" {
			err := errors.New("Not DML Query " + operation)
			activityLog.Errorf(err.Error())
			return false, err
		}
		op, err := mssqlpackage.UpdateQuery(username, password, host, port, dbname, query, timeout)
		if err != nil {
			activityLog.Errorf(err.Error())
			return false, err
		}
		context.SetOutput("output", op)
		return true, nil

	case methodCreate: //Create Query
		operation := strings.Split(query, " ")[0]
		op, err := mssqlpackage.CreateQuery(username, password, host, port, dbname, query, timeout)
		if strings.ToLower(operation) != "create" && strings.ToLower(operation) != "drop" && strings.ToLower(operation) != "alter" && strings.ToLower(operation) != "truncate" && strings.ToLower(operation) != "comment" {
			err := errors.New("Not DML Query " + operation)
			activityLog.Errorf(err.Error())
			return false, err
		}
		if err != nil {
			activityLog.Errorf(err.Error())
			return false, err
		}
		context.SetOutput("output", op)
		return true, nil

	}

	return true, nil
}
