package gosura

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	CREATE_SELECT_PERMISSION_TYPE string = `create_select_permission`
)

type CreateSelectPermission struct {
	Arguments CreateSelectPermissionArgs `json:"args"`
	QueryType string                     `json:"type"`
}

type CreateSelectPermissionArgs struct {
	Table      string            `json:"table"`
	Role       string            `json:"role"`
	Permission *SelectPermission `json:"permission"`
	Comment    string            `json:"comment,omitempty"`
}

type CreateSelectPermissionResponse map[string]interface{}

func (t *CreateSelectPermission) SetArgs(args interface{}) error {
	switch args.(type) {
	case CreateSelectPermissionArgs:
		t.Arguments = args.(CreateSelectPermissionArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *CreateSelectPermission) SetVersion(_ int) {}

func (t *CreateSelectPermission) SetType(name string) {
	t.QueryType = name
}

func (t *CreateSelectPermission) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *CreateSelectPermission) Method() string {
	return http.MethodPost
}

func (t *CreateSelectPermission) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var createSelectPermissionResponse CreateSelectPermissionResponse
	if err := json.Unmarshal(body, &createSelectPermissionResponse); err != nil {
		return nil, err
	}
	return createSelectPermissionResponse, nil
}

func NewCreateSelectPermissionQuery() Query {
	query := CreateSelectPermission{
		QueryType: CREATE_SELECT_PERMISSION_TYPE,
	}

	return Query(&query)
}
