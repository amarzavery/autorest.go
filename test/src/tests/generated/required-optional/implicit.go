package optionalgroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ImplicitClient is the test Infrastructure for AutoRest
type ImplicitClient struct {
	ManagementClient
}

// NewImplicitClient creates an instance of the ImplicitClient client.
func NewImplicitClient(requiredGlobalPath string, requiredGlobalQuery string, optionalGlobalQuery *int32) ImplicitClient {
	return NewImplicitClientWithBaseURI(DefaultBaseURI, requiredGlobalPath, requiredGlobalQuery, optionalGlobalQuery)
}

// NewImplicitClientWithBaseURI creates an instance of the ImplicitClient client.
func NewImplicitClientWithBaseURI(baseURI string, requiredGlobalPath string, requiredGlobalQuery string, optionalGlobalQuery *int32) ImplicitClient {
	return ImplicitClient{NewWithBaseURI(baseURI, requiredGlobalPath, requiredGlobalQuery, optionalGlobalQuery)}
}

// GetOptionalGlobalQuery test implicitly optional query parameter
func (client ImplicitClient) GetOptionalGlobalQuery() (result Error, err error) {
	req, err := client.GetOptionalGlobalQueryPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "optionalgroup.ImplicitClient", "GetOptionalGlobalQuery", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetOptionalGlobalQuerySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "optionalgroup.ImplicitClient", "GetOptionalGlobalQuery", resp, "Failure sending request")
		return
	}

	result, err = client.GetOptionalGlobalQueryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "optionalgroup.ImplicitClient", "GetOptionalGlobalQuery", resp, "Failure responding to request")
	}

	return
}

// GetOptionalGlobalQueryPreparer prepares the GetOptionalGlobalQuery request.
func (client ImplicitClient) GetOptionalGlobalQueryPreparer() (*http.Request, error) {
	queryParameters := map[string]interface{}{}
	if client.OptionalGlobalQuery != nil {
		queryParameters["optional-global-query"] = autorest.Encode("query", *client.OptionalGlobalQuery)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/reqopt/global/optional/query"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetOptionalGlobalQuerySender sends the GetOptionalGlobalQuery request. The method will close the
// http.Response Body if it receives an error.
func (client ImplicitClient) GetOptionalGlobalQuerySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetOptionalGlobalQueryResponder handles the response to the GetOptionalGlobalQuery request. The method always
// closes the http.Response Body.
func (client ImplicitClient) GetOptionalGlobalQueryResponder(resp *http.Response) (result Error, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetRequiredGlobalPath test implicitly required path parameter
func (client ImplicitClient) GetRequiredGlobalPath() (result Error, err error) {
	req, err := client.GetRequiredGlobalPathPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "optionalgroup.ImplicitClient", "GetRequiredGlobalPath", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetRequiredGlobalPathSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "optionalgroup.ImplicitClient", "GetRequiredGlobalPath", resp, "Failure sending request")
		return
	}

	result, err = client.GetRequiredGlobalPathResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "optionalgroup.ImplicitClient", "GetRequiredGlobalPath", resp, "Failure responding to request")
	}

	return
}

// GetRequiredGlobalPathPreparer prepares the GetRequiredGlobalPath request.
func (client ImplicitClient) GetRequiredGlobalPathPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"required-global-path": autorest.Encode("path", client.RequiredGlobalPath),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/reqopt/global/required/path/{required-global-path}", pathParameters))
	return preparer.Prepare(&http.Request{})
}

// GetRequiredGlobalPathSender sends the GetRequiredGlobalPath request. The method will close the
// http.Response Body if it receives an error.
func (client ImplicitClient) GetRequiredGlobalPathSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetRequiredGlobalPathResponder handles the response to the GetRequiredGlobalPath request. The method always
// closes the http.Response Body.
func (client ImplicitClient) GetRequiredGlobalPathResponder(resp *http.Response) (result Error, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetRequiredGlobalQuery test implicitly required query parameter
func (client ImplicitClient) GetRequiredGlobalQuery() (result Error, err error) {
	req, err := client.GetRequiredGlobalQueryPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "optionalgroup.ImplicitClient", "GetRequiredGlobalQuery", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetRequiredGlobalQuerySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "optionalgroup.ImplicitClient", "GetRequiredGlobalQuery", resp, "Failure sending request")
		return
	}

	result, err = client.GetRequiredGlobalQueryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "optionalgroup.ImplicitClient", "GetRequiredGlobalQuery", resp, "Failure responding to request")
	}

	return
}

// GetRequiredGlobalQueryPreparer prepares the GetRequiredGlobalQuery request.
func (client ImplicitClient) GetRequiredGlobalQueryPreparer() (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"required-global-query": autorest.Encode("query", client.RequiredGlobalQuery),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/reqopt/global/required/query"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetRequiredGlobalQuerySender sends the GetRequiredGlobalQuery request. The method will close the
// http.Response Body if it receives an error.
func (client ImplicitClient) GetRequiredGlobalQuerySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetRequiredGlobalQueryResponder handles the response to the GetRequiredGlobalQuery request. The method always
// closes the http.Response Body.
func (client ImplicitClient) GetRequiredGlobalQueryResponder(resp *http.Response) (result Error, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetRequiredPath test implicitly required path parameter
//
func (client ImplicitClient) GetRequiredPath(pathParameter string) (result Error, err error) {
	req, err := client.GetRequiredPathPreparer(pathParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "optionalgroup.ImplicitClient", "GetRequiredPath", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetRequiredPathSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "optionalgroup.ImplicitClient", "GetRequiredPath", resp, "Failure sending request")
		return
	}

	result, err = client.GetRequiredPathResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "optionalgroup.ImplicitClient", "GetRequiredPath", resp, "Failure responding to request")
	}

	return
}

// GetRequiredPathPreparer prepares the GetRequiredPath request.
func (client ImplicitClient) GetRequiredPathPreparer(pathParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"pathParameter": autorest.Encode("path", pathParameter),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/reqopt/implicit/required/path/{pathParameter}", pathParameters))
	return preparer.Prepare(&http.Request{})
}

// GetRequiredPathSender sends the GetRequiredPath request. The method will close the
// http.Response Body if it receives an error.
func (client ImplicitClient) GetRequiredPathSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetRequiredPathResponder handles the response to the GetRequiredPath request. The method always
// closes the http.Response Body.
func (client ImplicitClient) GetRequiredPathResponder(resp *http.Response) (result Error, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutOptionalBody test implicitly optional body parameter
//
func (client ImplicitClient) PutOptionalBody(bodyParameter string) (result autorest.Response, err error) {
	req, err := client.PutOptionalBodyPreparer(bodyParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "optionalgroup.ImplicitClient", "PutOptionalBody", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutOptionalBodySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "optionalgroup.ImplicitClient", "PutOptionalBody", resp, "Failure sending request")
		return
	}

	result, err = client.PutOptionalBodyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "optionalgroup.ImplicitClient", "PutOptionalBody", resp, "Failure responding to request")
	}

	return
}

// PutOptionalBodyPreparer prepares the PutOptionalBody request.
func (client ImplicitClient) PutOptionalBodyPreparer(bodyParameter string) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/reqopt/implicit/optional/body"))
	if len(bodyParameter) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(bodyParameter))
	}
	return preparer.Prepare(&http.Request{})
}

// PutOptionalBodySender sends the PutOptionalBody request. The method will close the
// http.Response Body if it receives an error.
func (client ImplicitClient) PutOptionalBodySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// PutOptionalBodyResponder handles the response to the PutOptionalBody request. The method always
// closes the http.Response Body.
func (client ImplicitClient) PutOptionalBodyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutOptionalHeader test implicitly optional header parameter
//
func (client ImplicitClient) PutOptionalHeader(queryParameter string) (result autorest.Response, err error) {
	req, err := client.PutOptionalHeaderPreparer(queryParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "optionalgroup.ImplicitClient", "PutOptionalHeader", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutOptionalHeaderSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "optionalgroup.ImplicitClient", "PutOptionalHeader", resp, "Failure sending request")
		return
	}

	result, err = client.PutOptionalHeaderResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "optionalgroup.ImplicitClient", "PutOptionalHeader", resp, "Failure responding to request")
	}

	return
}

// PutOptionalHeaderPreparer prepares the PutOptionalHeader request.
func (client ImplicitClient) PutOptionalHeaderPreparer(queryParameter string) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/reqopt/implicit/optional/header"))
	if len(queryParameter) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("queryParameter", autorest.String(queryParameter)))
	}
	return preparer.Prepare(&http.Request{})
}

// PutOptionalHeaderSender sends the PutOptionalHeader request. The method will close the
// http.Response Body if it receives an error.
func (client ImplicitClient) PutOptionalHeaderSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// PutOptionalHeaderResponder handles the response to the PutOptionalHeader request. The method always
// closes the http.Response Body.
func (client ImplicitClient) PutOptionalHeaderResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutOptionalQuery test implicitly optional query parameter
//
func (client ImplicitClient) PutOptionalQuery(queryParameter string) (result autorest.Response, err error) {
	req, err := client.PutOptionalQueryPreparer(queryParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "optionalgroup.ImplicitClient", "PutOptionalQuery", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutOptionalQuerySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "optionalgroup.ImplicitClient", "PutOptionalQuery", resp, "Failure sending request")
		return
	}

	result, err = client.PutOptionalQueryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "optionalgroup.ImplicitClient", "PutOptionalQuery", resp, "Failure responding to request")
	}

	return
}

// PutOptionalQueryPreparer prepares the PutOptionalQuery request.
func (client ImplicitClient) PutOptionalQueryPreparer(queryParameter string) (*http.Request, error) {
	queryParameters := map[string]interface{}{}
	if len(queryParameter) > 0 {
		queryParameters["queryParameter"] = autorest.Encode("query", queryParameter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/reqopt/implicit/optional/query"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// PutOptionalQuerySender sends the PutOptionalQuery request. The method will close the
// http.Response Body if it receives an error.
func (client ImplicitClient) PutOptionalQuerySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// PutOptionalQueryResponder handles the response to the PutOptionalQuery request. The method always
// closes the http.Response Body.
func (client ImplicitClient) PutOptionalQueryResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
