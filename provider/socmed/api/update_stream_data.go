package api

import "github.com/hinha/sometor/provider"

// Show keyword data api handler
type UpdateStreamData struct {
	streamProvider provider.SocmedKeywordAPI
}

// NewUpdateStreamData show data stream keyword handler object
func NewUpdateStreamData(streamProvider provider.SocmedKeywordAPI) *UpdateStreamData {
	return &UpdateStreamData{streamProvider: streamProvider}
}

// Path return api path
func (s *UpdateStreamData) Path() string {
	return "/stream/:media/show"
}

// Method return api method
func (s *UpdateStreamData) Method() string {
	return "GET"
}

// Handle request create keyword
func (s *UpdateStreamData) Handle(context provider.APIContext) {

}
