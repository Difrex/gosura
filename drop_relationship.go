package gosura

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	DROP_RELATIONSHIP_TYPE string = `drop_relationship`
)

type DropRelationship struct {
	Arguments DropRelationshipArgs `json:"args"`
	Ver       int                  `json:"version"`
	QueryType string               `json:"type"`
}

type DropRelationshipArgs struct {
	Table        string `json:"table"`
	Relationship string `json:"relationship"`
	Cascade      bool   `json:"cascade"`
}

type DropRelationshipResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *DropRelationship) SetArgs(args interface{}) error {
	switch args.(type) {
	case DropRelationshipArgs:
		t.Arguments = args.(DropRelationshipArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *DropRelationship) SetVersion(version int) {
	t.Ver = version
}

func (t *DropRelationship) SetType(name string) {
	t.QueryType = name
}

func (t *DropRelationship) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *DropRelationship) Method() string {
	return http.MethodPost
}

func (t *DropRelationship) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var dropRelationshipResponse DropRelationshipResponse
	if err := json.Unmarshal(body, &dropRelationshipResponse); err != nil {
		return nil, err
	}
	return dropRelationshipResponse, nil
}

func NewDropRelationshipQuery() Query {
	query := DropRelationship{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: DROP_RELATIONSHIP_TYPE,
	}

	return Query(&query)
}
