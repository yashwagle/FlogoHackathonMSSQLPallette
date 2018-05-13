# MSSQL Connector

This activity allows you to connect to Microsoft SQL Server. It takes DB parameters and query as input. It can be used to execute Select, Update and Create Queries. The output is a JSON schema based on the type of query.

## Installation

### Flogo CLI

```
flogo install github.com/yashwagle/HelloWorld/mssqlpallete
```

### Third-party libraries used
- #### go-mssqldb :
go-mssqldb is a driver written in GOLANG to connect to MS-SQL database. The purpose of this is to get the query output and then parse it to give the proper output.


### Schema

```
"inputs":[
  {
    "name": "host",
    "type": "string",
    "required": true
  },
  {
    "name": "method",
    "type": "string",
    "allowed": [
      "DQL",
      "DML",
      "DDL"
    ],
    "value": "DQL",
    "required": true
  },

  {
    "name": "port",
    "type": "string",
    "required": true
  },
  {
    "name": "dbname",
    "type": "string",
    "required": true
  },
  {
    "name": "username",
    "type": "string",
    "required": true
  },
  {
    "name": "password",
    "type": "string",
    "required": true
  },
  {
    "name": "timeout",
    "type": "int",
    "required": true
  },
  {
    "name": "query",
    "type": "string",
    "value":0,
    "required": false
  }
],
"outputs": [
  {
    "name": "output",
    "type": "object"
  }
]
}
```

### Activity Input


| Name | Required | Type | Description |
| ---- | -------- | ---- |------------ |
| host | True | String | Name of the Database Server |
| method  | True | String | Type of Query DML,DQL,DDL |
| port  | True | String | Port of the Database |
| dbname  | True | String | Name of the database |
| username  | True | String | Database username |
| password  | True | String | Database password |
| timeout  | False | Integer | Timeout for the string |
| Query  | True | String | Query to be Executed |


### Activity Output


| Name | Type | Description |
| ---- | ---- | ----------- |
| Output | Object | JSON output depending on the type of query |

### Example :
This activity will give the response in a following way,


DML Operation

```
{"numberOfRowsAffected":"1"}

```

DQL Operation

```
{
	"rows": [{
		"row": [{
			"column": {
				"name": "name",
				"value": "John"
			}
		}, {
			"column": {
				"name": "dob",
				"value": "1995-03-27 00:00:00 +0000 UTC"
			}
		}, {
			"column": {
				"name": "address",
				"value": "NY"
			}
		}, {
			"column": {
				"name": "phonenumber",
				"value": "123456789"
			}
		}]
	}]
}
```

DDL Operation

```
{"Query Status":"Operation Successful"}
```
