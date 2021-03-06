// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pi

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// PI provides the API operation methods for making requests to
// AWS Performance Insights. See this package's package overview docs
// for details on the service.
//
// PI methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type PI struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*PI)

// Used for custom request initialization logic
var initRequest func(*PI, *aws.Request)

// Service information constants
const (
	ServiceName = "pi"        // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the PI client with a config.
//
// Example:
//     // Create a PI client from just a config.
//     svc := pi.New(myConfig)
func New(config aws.Config) *PI {
	var signingName string
	signingName = "pi"
	signingRegion := config.Region

	svc := &PI{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2018-02-27",
				JSONVersion:   "1.1",
				TargetPrefix:  "PerformanceInsightsv20180227",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a PI operation and runs any
// custom request initialization.
func (c *PI) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
