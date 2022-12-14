/*
GitHub v3 REST API

GitHub's v3 REST API.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github

import (
	"encoding/json"
)

// TimelineLineCommentedEvent Timeline Line Commented Event
type TimelineLineCommentedEvent struct {
	Event *string `json:"event,omitempty"`
	NodeId *string `json:"node_id,omitempty"`
	Comments []PullRequestReviewComment `json:"comments,omitempty"`
}

// NewTimelineLineCommentedEvent instantiates a new TimelineLineCommentedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineLineCommentedEvent() *TimelineLineCommentedEvent {
	this := TimelineLineCommentedEvent{}
	return &this
}

// NewTimelineLineCommentedEventWithDefaults instantiates a new TimelineLineCommentedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineLineCommentedEventWithDefaults() *TimelineLineCommentedEvent {
	this := TimelineLineCommentedEvent{}
	return &this
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *TimelineLineCommentedEvent) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineLineCommentedEvent) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *TimelineLineCommentedEvent) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *TimelineLineCommentedEvent) SetEvent(v string) {
	o.Event = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *TimelineLineCommentedEvent) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineLineCommentedEvent) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *TimelineLineCommentedEvent) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *TimelineLineCommentedEvent) SetNodeId(v string) {
	o.NodeId = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *TimelineLineCommentedEvent) GetComments() []PullRequestReviewComment {
	if o == nil || o.Comments == nil {
		var ret []PullRequestReviewComment
		return ret
	}
	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineLineCommentedEvent) GetCommentsOk() ([]PullRequestReviewComment, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *TimelineLineCommentedEvent) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given []PullRequestReviewComment and assigns it to the Comments field.
func (o *TimelineLineCommentedEvent) SetComments(v []PullRequestReviewComment) {
	o.Comments = v
}

func (o TimelineLineCommentedEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.NodeId != nil {
		toSerialize["node_id"] = o.NodeId
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	return json.Marshal(toSerialize)
}

type NullableTimelineLineCommentedEvent struct {
	value *TimelineLineCommentedEvent
	isSet bool
}

func (v NullableTimelineLineCommentedEvent) Get() *TimelineLineCommentedEvent {
	return v.value
}

func (v *NullableTimelineLineCommentedEvent) Set(val *TimelineLineCommentedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineLineCommentedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineLineCommentedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineLineCommentedEvent(val *TimelineLineCommentedEvent) *NullableTimelineLineCommentedEvent {
	return &NullableTimelineLineCommentedEvent{value: val, isSet: true}
}

func (v NullableTimelineLineCommentedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineLineCommentedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


