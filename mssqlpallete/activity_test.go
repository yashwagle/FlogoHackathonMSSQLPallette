package mssqlpallete

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}
		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	ip := `Insert into master.dbo.customer (name,phonenumber,dob,address) values ('Yash','123','1995-03-23','PQR');`
	//setup attrs
	method := "DML"
	tc.SetInput("method", method)
	tc.SetInput("query", ip)
	tc.SetInput("host", `localhost`)
	tc.SetInput("port", "1433")
	tc.SetInput("username", "sa")
	tc.SetInput("password", "Tibco2018")
	tc.SetInput("dbname", "master")

	act.Eval(tc)
	result := tc.GetOutput("output")
	expectedop := `{"numberOfRowsAffected":"1"}`
	resultString, _ := json.Marshal(result)
	assert.Equal(t, expectedop, string(resultString))

	//check result attr
}
