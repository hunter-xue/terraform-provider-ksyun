// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package bws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opAssociateBandWidthShare = "AssociateBandWidthShare"

// AssociateBandWidthShareRequest generates a "ksc/request.Request" representing the
// client's request for the AssociateBandWidthShare operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See AssociateBandWidthShare for more information on using the AssociateBandWidthShare
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the AssociateBandWidthShareRequest method.
//    req, resp := client.AssociateBandWidthShareRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/bws-2016-03-04/AssociateBandWidthShare
func (c *Bws) AssociateBandWidthShareRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAssociateBandWidthShare,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// AssociateBandWidthShare API operation for bws.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for bws's
// API operation AssociateBandWidthShare for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/bws-2016-03-04/AssociateBandWidthShare
func (c *Bws) AssociateBandWidthShare(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AssociateBandWidthShareRequest(input)
	return out, req.Send()
}

// AssociateBandWidthShareWithContext is the same as AssociateBandWidthShare with the addition of
// the ability to pass a context and additional request options.
//
// See AssociateBandWidthShare for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Bws) AssociateBandWidthShareWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AssociateBandWidthShareRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateBandWidthShare = "CreateBandWidthShare"

// CreateBandWidthShareRequest generates a "ksc/request.Request" representing the
// client's request for the CreateBandWidthShare operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See CreateBandWidthShare for more information on using the CreateBandWidthShare
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the CreateBandWidthShareRequest method.
//    req, resp := client.CreateBandWidthShareRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/bws-2016-03-04/CreateBandWidthShare
func (c *Bws) CreateBandWidthShareRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateBandWidthShare,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// CreateBandWidthShare API operation for bws.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for bws's
// API operation CreateBandWidthShare for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/bws-2016-03-04/CreateBandWidthShare
func (c *Bws) CreateBandWidthShare(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateBandWidthShareRequest(input)
	return out, req.Send()
}

// CreateBandWidthShareWithContext is the same as CreateBandWidthShare with the addition of
// the ability to pass a context and additional request options.
//
// See CreateBandWidthShare for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Bws) CreateBandWidthShareWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateBandWidthShareRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteBandWidthShare = "DeleteBandWidthShare"

// DeleteBandWidthShareRequest generates a "ksc/request.Request" representing the
// client's request for the DeleteBandWidthShare operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DeleteBandWidthShare for more information on using the DeleteBandWidthShare
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DeleteBandWidthShareRequest method.
//    req, resp := client.DeleteBandWidthShareRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/bws-2016-03-04/DeleteBandWidthShare
func (c *Bws) DeleteBandWidthShareRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteBandWidthShare,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteBandWidthShare API operation for bws.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for bws's
// API operation DeleteBandWidthShare for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/bws-2016-03-04/DeleteBandWidthShare
func (c *Bws) DeleteBandWidthShare(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteBandWidthShareRequest(input)
	return out, req.Send()
}

// DeleteBandWidthShareWithContext is the same as DeleteBandWidthShare with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteBandWidthShare for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Bws) DeleteBandWidthShareWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteBandWidthShareRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeBandWidthShares = "DescribeBandWidthShares"

// DescribeBandWidthSharesRequest generates a "ksc/request.Request" representing the
// client's request for the DescribeBandWidthShares operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DescribeBandWidthShares for more information on using the DescribeBandWidthShares
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DescribeBandWidthSharesRequest method.
//    req, resp := client.DescribeBandWidthSharesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/bws-2016-03-04/DescribeBandWidthShares
func (c *Bws) DescribeBandWidthSharesRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeBandWidthShares,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeBandWidthShares API operation for bws.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for bws's
// API operation DescribeBandWidthShares for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/bws-2016-03-04/DescribeBandWidthShares
func (c *Bws) DescribeBandWidthShares(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeBandWidthSharesRequest(input)
	return out, req.Send()
}

// DescribeBandWidthSharesWithContext is the same as DescribeBandWidthShares with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBandWidthShares for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Bws) DescribeBandWidthSharesWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeBandWidthSharesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDisassociateBandWidthShare = "DisassociateBandWidthShare"

// DisassociateBandWidthShareRequest generates a "ksc/request.Request" representing the
// client's request for the DisassociateBandWidthShare operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DisassociateBandWidthShare for more information on using the DisassociateBandWidthShare
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DisassociateBandWidthShareRequest method.
//    req, resp := client.DisassociateBandWidthShareRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/bws-2016-03-04/DisassociateBandWidthShare
func (c *Bws) DisassociateBandWidthShareRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDisassociateBandWidthShare,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DisassociateBandWidthShare API operation for bws.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for bws's
// API operation DisassociateBandWidthShare for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/bws-2016-03-04/DisassociateBandWidthShare
func (c *Bws) DisassociateBandWidthShare(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DisassociateBandWidthShareRequest(input)
	return out, req.Send()
}

// DisassociateBandWidthShareWithContext is the same as DisassociateBandWidthShare with the addition of
// the ability to pass a context and additional request options.
//
// See DisassociateBandWidthShare for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Bws) DisassociateBandWidthShareWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DisassociateBandWidthShareRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyBandWidthShare = "ModifyBandWidthShare"

// ModifyBandWidthShareRequest generates a "ksc/request.Request" representing the
// client's request for the ModifyBandWidthShare operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ModifyBandWidthShare for more information on using the ModifyBandWidthShare
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ModifyBandWidthShareRequest method.
//    req, resp := client.ModifyBandWidthShareRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/bws-2016-03-04/ModifyBandWidthShare
func (c *Bws) ModifyBandWidthShareRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyBandWidthShare,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyBandWidthShare API operation for bws.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for bws's
// API operation ModifyBandWidthShare for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/bws-2016-03-04/ModifyBandWidthShare
func (c *Bws) ModifyBandWidthShare(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyBandWidthShareRequest(input)
	return out, req.Send()
}

// ModifyBandWidthShareWithContext is the same as ModifyBandWidthShare with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyBandWidthShare for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Bws) ModifyBandWidthShareWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyBandWidthShareRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}
