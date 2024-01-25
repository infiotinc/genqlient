// Code generated by github.com/infiotinc/genqlient, DO NOT EDIT.

package test

import (
	"encoding/json"
	"fmt"

	"github.com/infiotinc/genqlient/graphql"
	"github.com/infiotinc/genqlient/internal/testutil"
)

// Role is a type a user may have.
type Role string

const (
	// What is a student?
	//
	// A student is primarily a person enrolled in a school or other educational institution and who is under learning with goals of acquiring knowledge, developing professions and achieving employment at desired field. In the broader sense, a student is anyone who applies themselves to the intensive intellectual engagement with some matter necessary to master it as part of some practical affair in which such mastery is basic or decisive.
	//
	// (from [Wikipedia](https://en.wikipedia.org/wiki/Student))
	RoleStudent Role = "STUDENT"
	// Teacher is a teacher, who teaches the students.
	RoleTeacher Role = "TEACHER"
)

// StructOptionResponse is returned by StructOption on success.
type StructOptionResponse struct {
	Root StructOptionRootTopic `json:"root"`
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User StructOptionUser `json:"user"`
}

// GetRoot returns StructOptionResponse.Root, and is useful for accessing the field via an interface.
func (v *StructOptionResponse) GetRoot() StructOptionRootTopic { return v.Root }

// GetUser returns StructOptionResponse.User, and is useful for accessing the field via an interface.
func (v *StructOptionResponse) GetUser() StructOptionUser { return v.User }

// StructOptionRootTopic includes the requested fields of the GraphQL type Topic.
type StructOptionRootTopic struct {
	// ID is documented in the Content interface.
	Id       testutil.ID                            `json:"id"`
	Children []StructOptionRootTopicChildrenContent `json:"children"`
}

// GetId returns StructOptionRootTopic.Id, and is useful for accessing the field via an interface.
func (v *StructOptionRootTopic) GetId() testutil.ID { return v.Id }

// GetChildren returns StructOptionRootTopic.Children, and is useful for accessing the field via an interface.
func (v *StructOptionRootTopic) GetChildren() []StructOptionRootTopicChildrenContent {
	return v.Children
}

// StructOptionRootTopicChildrenContent includes the requested fields of the GraphQL type Content.
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type StructOptionRootTopicChildrenContent struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id     testutil.ID                                     `json:"id"`
	Parent StructOptionRootTopicChildrenContentParentTopic `json:"parent"`
}

// GetTypename returns StructOptionRootTopicChildrenContent.Typename, and is useful for accessing the field via an interface.
func (v *StructOptionRootTopicChildrenContent) GetTypename() string { return v.Typename }

// GetId returns StructOptionRootTopicChildrenContent.Id, and is useful for accessing the field via an interface.
func (v *StructOptionRootTopicChildrenContent) GetId() testutil.ID { return v.Id }

// GetParent returns StructOptionRootTopicChildrenContent.Parent, and is useful for accessing the field via an interface.
func (v *StructOptionRootTopicChildrenContent) GetParent() StructOptionRootTopicChildrenContentParentTopic {
	return v.Parent
}

// StructOptionRootTopicChildrenContentParentTopic includes the requested fields of the GraphQL type Topic.
type StructOptionRootTopicChildrenContentParentTopic struct {
	// ID is documented in the Content interface.
	Id                testutil.ID                                                               `json:"id"`
	Children          []StructOptionRootTopicChildrenContentParentTopicChildrenContent          `json:"children"`
	InterfaceChildren []StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent `json:"-"`
}

// GetId returns StructOptionRootTopicChildrenContentParentTopic.Id, and is useful for accessing the field via an interface.
func (v *StructOptionRootTopicChildrenContentParentTopic) GetId() testutil.ID { return v.Id }

// GetChildren returns StructOptionRootTopicChildrenContentParentTopic.Children, and is useful for accessing the field via an interface.
func (v *StructOptionRootTopicChildrenContentParentTopic) GetChildren() []StructOptionRootTopicChildrenContentParentTopicChildrenContent {
	return v.Children
}

// GetInterfaceChildren returns StructOptionRootTopicChildrenContentParentTopic.InterfaceChildren, and is useful for accessing the field via an interface.
func (v *StructOptionRootTopicChildrenContentParentTopic) GetInterfaceChildren() []StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent {
	return v.InterfaceChildren
}

func (v *StructOptionRootTopicChildrenContentParentTopic) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*StructOptionRootTopicChildrenContentParentTopic
		InterfaceChildren []json.RawMessage `json:"interfaceChildren"`
		graphql.NoUnmarshalJSON
	}
	firstPass.StructOptionRootTopicChildrenContentParentTopic = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.InterfaceChildren
		src := firstPass.InterfaceChildren
		*dst = make(
			[]StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			if len(src) != 0 && string(src) != "null" {
				err = __unmarshalStructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent(
					src, dst)
				if err != nil {
					return fmt.Errorf(
						"unable to unmarshal StructOptionRootTopicChildrenContentParentTopic.InterfaceChildren: %w", err)
				}
			}
		}
	}
	return nil
}

type __premarshalStructOptionRootTopicChildrenContentParentTopic struct {
	Id testutil.ID `json:"id"`

	Children []StructOptionRootTopicChildrenContentParentTopicChildrenContent `json:"children"`

	InterfaceChildren []json.RawMessage `json:"interfaceChildren"`
}

func (v *StructOptionRootTopicChildrenContentParentTopic) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *StructOptionRootTopicChildrenContentParentTopic) __premarshalJSON() (*__premarshalStructOptionRootTopicChildrenContentParentTopic, error) {
	var retval __premarshalStructOptionRootTopicChildrenContentParentTopic

	retval.Id = v.Id
	retval.Children = v.Children
	{

		dst := &retval.InterfaceChildren
		src := v.InterfaceChildren
		*dst = make(
			[]json.RawMessage,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			var err error
			*dst, err = __marshalStructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent(
				&src)
			if err != nil {
				return nil, fmt.Errorf(
					"unable to marshal StructOptionRootTopicChildrenContentParentTopic.InterfaceChildren: %w", err)
			}
		}
	}
	return &retval, nil
}

// StructOptionRootTopicChildrenContentParentTopicChildrenContent includes the requested fields of the GraphQL type Content.
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type StructOptionRootTopicChildrenContentParentTopicChildrenContent struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

// GetTypename returns StructOptionRootTopicChildrenContentParentTopicChildrenContent.Typename, and is useful for accessing the field via an interface.
func (v *StructOptionRootTopicChildrenContentParentTopicChildrenContent) GetTypename() string {
	return v.Typename
}

// GetId returns StructOptionRootTopicChildrenContentParentTopicChildrenContent.Id, and is useful for accessing the field via an interface.
func (v *StructOptionRootTopicChildrenContentParentTopicChildrenContent) GetId() testutil.ID {
	return v.Id
}

// StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenArticle includes the requested fields of the GraphQL type Article.
type StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

// GetTypename returns StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenArticle.Typename, and is useful for accessing the field via an interface.
func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenArticle) GetTypename() string {
	return v.Typename
}

// GetId returns StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenArticle.Id, and is useful for accessing the field via an interface.
func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenArticle) GetId() testutil.ID {
	return v.Id
}

// StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent includes the requested fields of the GraphQL interface Content.
//
// StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent is implemented by the following types:
// StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenArticle
// StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenTopic
// StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent interface {
	implementsGraphQLInterfaceStructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() testutil.ID
}

func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenArticle) implementsGraphQLInterfaceStructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent() {
}
func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenTopic) implementsGraphQLInterfaceStructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent() {
}
func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo) implementsGraphQLInterfaceStructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent() {
}

func __unmarshalStructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent(b []byte, v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent) error {
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
		*v = new(StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenArticle)
		return json.Unmarshal(b, *v)
	case "Topic":
		*v = new(StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenTopic)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent: "%v"`, tn.TypeName)
	}
}

func __marshalStructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent(v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenArticle:
		typename = "Article"

		result := struct {
			TypeName string `json:"__typename"`
			*StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenArticle
		}{typename, v}
		return json.Marshal(result)
	case *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenTopic:
		typename = "Topic"

		result := struct {
			TypeName string `json:"__typename"`
			*StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenTopic
		}{typename, v}
		return json.Marshal(result)
	case *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo:
		typename = "Video"

		premarshaled, err := v.__premarshalJSON()
		if err != nil {
			return nil, err
		}
		result := struct {
			TypeName string `json:"__typename"`
			*__premarshalStructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo
		}{typename, premarshaled}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent: "%T"`, v)
	}
}

// StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenTopic includes the requested fields of the GraphQL type Topic.
type StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

// GetTypename returns StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenTopic.Typename, and is useful for accessing the field via an interface.
func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenTopic) GetTypename() string {
	return v.Typename
}

// GetId returns StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenTopic.Id, and is useful for accessing the field via an interface.
func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenTopic) GetId() testutil.ID {
	return v.Id
}

// StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo includes the requested fields of the GraphQL type Video.
type StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id          testutil.ID `json:"id"`
	VideoFields `json:"-"`
}

// GetTypename returns StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo.Typename, and is useful for accessing the field via an interface.
func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo) GetTypename() string {
	return v.Typename
}

// GetId returns StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo.Id, and is useful for accessing the field via an interface.
func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo) GetId() testutil.ID {
	return v.Id
}

// GetDuration returns StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo.Duration, and is useful for accessing the field via an interface.
func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo) GetDuration() int {
	return v.VideoFields.Duration
}

func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo
		graphql.NoUnmarshalJSON
	}
	firstPass.StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.VideoFields)
	if err != nil {
		return err
	}
	return nil
}

type __premarshalStructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo struct {
	Typename string `json:"__typename"`

	Id testutil.ID `json:"id"`

	Duration int `json:"duration"`
}

func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo) __premarshalJSON() (*__premarshalStructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo, error) {
	var retval __premarshalStructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo

	retval.Typename = v.Typename
	retval.Id = v.Id
	retval.Duration = v.VideoFields.Duration
	return &retval, nil
}

// StructOptionUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type StructOptionUser struct {
	Roles []Role `json:"roles"`
}

// GetRoles returns StructOptionUser.Roles, and is useful for accessing the field via an interface.
func (v *StructOptionUser) GetRoles() []Role { return v.Roles }

// VideoFields includes the GraphQL fields of Video requested by the fragment VideoFields.
type VideoFields struct {
	Duration int `json:"duration"`
}

// GetDuration returns VideoFields.Duration, and is useful for accessing the field via an interface.
func (v *VideoFields) GetDuration() int { return v.Duration }

// The query or mutation executed by StructOption.
const StructOption_Operation = `
query StructOption {
	root {
		id
		children {
			__typename
			id
			parent {
				id
				children {
					__typename
					id
				}
				interfaceChildren: children {
					__typename
					id
					... VideoFields
				}
			}
		}
	}
	user {
		roles
	}
}
fragment VideoFields on Video {
	duration
}
`

func StructOption(
	client_ graphql.Client,
) (*StructOptionResponse, error) {
	req_ := &graphql.Request{
		OpName: "StructOption",
		Query:  StructOption_Operation,
	}
	var err_ error

	var data_ StructOptionResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return &data_, err_
}

