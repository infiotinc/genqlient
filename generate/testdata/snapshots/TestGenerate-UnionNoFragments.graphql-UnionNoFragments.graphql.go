// Code generated by github.com/infiotinc/genqlient, DO NOT EDIT.

package test

import (
	"encoding/json"
	"fmt"

	"github.com/infiotinc/genqlient/graphql"
)

// UnionNoFragmentsQueryRandomLeafArticle includes the requested fields of the GraphQL type Article.
type UnionNoFragmentsQueryRandomLeafArticle struct {
	Typename string `json:"__typename"`
}

// GetTypename returns UnionNoFragmentsQueryRandomLeafArticle.Typename, and is useful for accessing the field via an interface.
func (v *UnionNoFragmentsQueryRandomLeafArticle) GetTypename() string { return v.Typename }

// UnionNoFragmentsQueryRandomLeafLeafContent includes the requested fields of the GraphQL interface LeafContent.
//
// UnionNoFragmentsQueryRandomLeafLeafContent is implemented by the following types:
// UnionNoFragmentsQueryRandomLeafArticle
// UnionNoFragmentsQueryRandomLeafVideo
// The GraphQL type's documentation follows.
//
// LeafContent represents content items that can't have child-nodes.
type UnionNoFragmentsQueryRandomLeafLeafContent interface {
	implementsGraphQLInterfaceUnionNoFragmentsQueryRandomLeafLeafContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
}

func (v *UnionNoFragmentsQueryRandomLeafArticle) implementsGraphQLInterfaceUnionNoFragmentsQueryRandomLeafLeafContent() {
}
func (v *UnionNoFragmentsQueryRandomLeafVideo) implementsGraphQLInterfaceUnionNoFragmentsQueryRandomLeafLeafContent() {
}

func __unmarshalUnionNoFragmentsQueryRandomLeafLeafContent(b []byte, v *UnionNoFragmentsQueryRandomLeafLeafContent) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(UnionNoFragmentsQueryRandomLeafArticle)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(UnionNoFragmentsQueryRandomLeafVideo)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing LeafContent.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for UnionNoFragmentsQueryRandomLeafLeafContent: "%v"`, tn.TypeName)
	}
}

func __marshalUnionNoFragmentsQueryRandomLeafLeafContent(v *UnionNoFragmentsQueryRandomLeafLeafContent) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *UnionNoFragmentsQueryRandomLeafArticle:
		typename = "Article"

		result := struct {
			TypeName string `json:"__typename"`
			*UnionNoFragmentsQueryRandomLeafArticle
		}{typename, v}
		return json.Marshal(result)
	case *UnionNoFragmentsQueryRandomLeafVideo:
		typename = "Video"

		result := struct {
			TypeName string `json:"__typename"`
			*UnionNoFragmentsQueryRandomLeafVideo
		}{typename, v}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for UnionNoFragmentsQueryRandomLeafLeafContent: "%T"`, v)
	}
}

// UnionNoFragmentsQueryRandomLeafVideo includes the requested fields of the GraphQL type Video.
type UnionNoFragmentsQueryRandomLeafVideo struct {
	Typename string `json:"__typename"`
}

// GetTypename returns UnionNoFragmentsQueryRandomLeafVideo.Typename, and is useful for accessing the field via an interface.
func (v *UnionNoFragmentsQueryRandomLeafVideo) GetTypename() string { return v.Typename }

// UnionNoFragmentsQueryResponse is returned by UnionNoFragmentsQuery on success.
type UnionNoFragmentsQueryResponse struct {
	RandomLeaf UnionNoFragmentsQueryRandomLeafLeafContent `json:"-"`
}

// GetRandomLeaf returns UnionNoFragmentsQueryResponse.RandomLeaf, and is useful for accessing the field via an interface.
func (v *UnionNoFragmentsQueryResponse) GetRandomLeaf() UnionNoFragmentsQueryRandomLeafLeafContent {
	return v.RandomLeaf
}

func (v *UnionNoFragmentsQueryResponse) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*UnionNoFragmentsQueryResponse
		RandomLeaf json.RawMessage `json:"randomLeaf"`
		graphql.NoUnmarshalJSON
	}
	firstPass.UnionNoFragmentsQueryResponse = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.RandomLeaf
		src := firstPass.RandomLeaf
		if len(src) != 0 && string(src) != "null" {
			err = __unmarshalUnionNoFragmentsQueryRandomLeafLeafContent(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"unable to unmarshal UnionNoFragmentsQueryResponse.RandomLeaf: %w", err)
			}
		}
	}
	return nil
}

type __premarshalUnionNoFragmentsQueryResponse struct {
	RandomLeaf json.RawMessage `json:"randomLeaf"`
}

func (v *UnionNoFragmentsQueryResponse) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *UnionNoFragmentsQueryResponse) __premarshalJSON() (*__premarshalUnionNoFragmentsQueryResponse, error) {
	var retval __premarshalUnionNoFragmentsQueryResponse

	{

		dst := &retval.RandomLeaf
		src := v.RandomLeaf
		var err error
		*dst, err = __marshalUnionNoFragmentsQueryRandomLeafLeafContent(
			&src)
		if err != nil {
			return nil, fmt.Errorf(
				"unable to marshal UnionNoFragmentsQueryResponse.RandomLeaf: %w", err)
		}
	}
	return &retval, nil
}

// The query or mutation executed by UnionNoFragmentsQuery.
const UnionNoFragmentsQuery_Operation = `
query UnionNoFragmentsQuery {
	randomLeaf {
		__typename
	}
}
`

func UnionNoFragmentsQuery(
	client_ graphql.Client,
) (*UnionNoFragmentsQueryResponse, error) {
	req_ := &graphql.Request{
		OpName: "UnionNoFragmentsQuery",
		Query:  UnionNoFragmentsQuery_Operation,
	}
	var err_ error

	var data_ UnionNoFragmentsQueryResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return &data_, err_
}

