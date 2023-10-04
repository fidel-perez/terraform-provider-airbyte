// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type JobsResponse struct {
	Data     []JobResponse `json:"data"`
	Next     *string       `json:"next,omitempty"`
	Previous *string       `json:"previous,omitempty"`
}

func (o *JobsResponse) GetData() []JobResponse {
	if o == nil {
		return []JobResponse{}
	}
	return o.Data
}

func (o *JobsResponse) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *JobsResponse) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}
