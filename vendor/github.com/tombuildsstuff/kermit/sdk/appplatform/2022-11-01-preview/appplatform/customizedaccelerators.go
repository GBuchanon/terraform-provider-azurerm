package appplatform

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
)

// CustomizedAcceleratorsClient is the REST API for Azure Spring Apps
type CustomizedAcceleratorsClient struct {
	BaseClient
}

// NewCustomizedAcceleratorsClient creates an instance of the CustomizedAcceleratorsClient client.
func NewCustomizedAcceleratorsClient(subscriptionID string) CustomizedAcceleratorsClient {
	return NewCustomizedAcceleratorsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCustomizedAcceleratorsClientWithBaseURI creates an instance of the CustomizedAcceleratorsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewCustomizedAcceleratorsClientWithBaseURI(baseURI string, subscriptionID string) CustomizedAcceleratorsClient {
	return CustomizedAcceleratorsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update the customized accelerator.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// applicationAcceleratorName - the name of the application accelerator.
// customizedAcceleratorName - the name of the customized accelerator.
// customizedAcceleratorResource - the customized accelerator for the create or update operation
func (client CustomizedAcceleratorsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string, customizedAcceleratorResource CustomizedAcceleratorResource) (result CustomizedAcceleratorsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomizedAcceleratorsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: customizedAcceleratorResource,
			Constraints: []validation.Constraint{{Target: "customizedAcceleratorResource.Properties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "customizedAcceleratorResource.Properties.GitRepository", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "customizedAcceleratorResource.Properties.GitRepository.URL", Name: validation.Null, Rule: true, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("appplatform.CustomizedAcceleratorsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, applicationAcceleratorName, customizedAcceleratorName, customizedAcceleratorResource)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.CustomizedAcceleratorsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.CustomizedAcceleratorsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client CustomizedAcceleratorsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string, customizedAcceleratorResource CustomizedAcceleratorResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationAcceleratorName": autorest.Encode("path", applicationAcceleratorName),
		"customizedAcceleratorName":  autorest.Encode("path", customizedAcceleratorName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"serviceName":                autorest.Encode("path", serviceName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}/customizedAccelerators/{customizedAcceleratorName}", pathParameters),
		autorest.WithJSON(customizedAcceleratorResource),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client CustomizedAcceleratorsClient) CreateOrUpdateSender(req *http.Request) (future CustomizedAcceleratorsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client CustomizedAcceleratorsClient) CreateOrUpdateResponder(resp *http.Response) (result CustomizedAcceleratorResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete the customized accelerator.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// applicationAcceleratorName - the name of the application accelerator.
// customizedAcceleratorName - the name of the customized accelerator.
func (client CustomizedAcceleratorsClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string) (result CustomizedAcceleratorsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomizedAcceleratorsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, applicationAcceleratorName, customizedAcceleratorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.CustomizedAcceleratorsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.CustomizedAcceleratorsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client CustomizedAcceleratorsClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationAcceleratorName": autorest.Encode("path", applicationAcceleratorName),
		"customizedAcceleratorName":  autorest.Encode("path", customizedAcceleratorName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"serviceName":                autorest.Encode("path", serviceName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}/customizedAccelerators/{customizedAcceleratorName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client CustomizedAcceleratorsClient) DeleteSender(req *http.Request) (future CustomizedAcceleratorsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client CustomizedAcceleratorsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the customized accelerator.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// applicationAcceleratorName - the name of the application accelerator.
// customizedAcceleratorName - the name of the customized accelerator.
func (client CustomizedAcceleratorsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string) (result CustomizedAcceleratorResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomizedAcceleratorsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, applicationAcceleratorName, customizedAcceleratorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.CustomizedAcceleratorsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appplatform.CustomizedAcceleratorsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.CustomizedAcceleratorsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client CustomizedAcceleratorsClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationAcceleratorName": autorest.Encode("path", applicationAcceleratorName),
		"customizedAcceleratorName":  autorest.Encode("path", customizedAcceleratorName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"serviceName":                autorest.Encode("path", serviceName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}/customizedAccelerators/{customizedAcceleratorName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CustomizedAcceleratorsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CustomizedAcceleratorsClient) GetResponder(resp *http.Response) (result CustomizedAcceleratorResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List handle requests to list all customized accelerators.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// applicationAcceleratorName - the name of the application accelerator.
func (client CustomizedAcceleratorsClient) List(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string) (result CustomizedAcceleratorResourceCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomizedAcceleratorsClient.List")
		defer func() {
			sc := -1
			if result.carc.Response.Response != nil {
				sc = result.carc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, serviceName, applicationAcceleratorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.CustomizedAcceleratorsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.carc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appplatform.CustomizedAcceleratorsClient", "List", resp, "Failure sending request")
		return
	}

	result.carc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.CustomizedAcceleratorsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.carc.hasNextLink() && result.carc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client CustomizedAcceleratorsClient) ListPreparer(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationAcceleratorName": autorest.Encode("path", applicationAcceleratorName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"serviceName":                autorest.Encode("path", serviceName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}/customizedAccelerators", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client CustomizedAcceleratorsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client CustomizedAcceleratorsClient) ListResponder(resp *http.Response) (result CustomizedAcceleratorResourceCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client CustomizedAcceleratorsClient) listNextResults(ctx context.Context, lastResults CustomizedAcceleratorResourceCollection) (result CustomizedAcceleratorResourceCollection, err error) {
	req, err := lastResults.customizedAcceleratorResourceCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "appplatform.CustomizedAcceleratorsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "appplatform.CustomizedAcceleratorsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.CustomizedAcceleratorsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client CustomizedAcceleratorsClient) ListComplete(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string) (result CustomizedAcceleratorResourceCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomizedAcceleratorsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, serviceName, applicationAcceleratorName)
	return
}

// Validate check the customized accelerator are valid.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// applicationAcceleratorName - the name of the application accelerator.
// customizedAcceleratorName - the name of the customized accelerator.
// properties - customized accelerator properties to be validated
func (client CustomizedAcceleratorsClient) Validate(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string, properties CustomizedAcceleratorProperties) (result CustomizedAcceleratorValidateResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomizedAcceleratorsClient.Validate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: properties,
			Constraints: []validation.Constraint{{Target: "properties.GitRepository", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "properties.GitRepository.URL", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("appplatform.CustomizedAcceleratorsClient", "Validate", err.Error())
	}

	req, err := client.ValidatePreparer(ctx, resourceGroupName, serviceName, applicationAcceleratorName, customizedAcceleratorName, properties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.CustomizedAcceleratorsClient", "Validate", nil, "Failure preparing request")
		return
	}

	resp, err := client.ValidateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appplatform.CustomizedAcceleratorsClient", "Validate", resp, "Failure sending request")
		return
	}

	result, err = client.ValidateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.CustomizedAcceleratorsClient", "Validate", resp, "Failure responding to request")
		return
	}

	return
}

// ValidatePreparer prepares the Validate request.
func (client CustomizedAcceleratorsClient) ValidatePreparer(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string, properties CustomizedAcceleratorProperties) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationAcceleratorName": autorest.Encode("path", applicationAcceleratorName),
		"customizedAcceleratorName":  autorest.Encode("path", customizedAcceleratorName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"serviceName":                autorest.Encode("path", serviceName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	properties.ProvisioningState = ""
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}/customizedAccelerators/{customizedAcceleratorName}/validate", pathParameters),
		autorest.WithJSON(properties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidateSender sends the Validate request. The method will close the
// http.Response Body if it receives an error.
func (client CustomizedAcceleratorsClient) ValidateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ValidateResponder handles the response to the Validate request. The method always
// closes the http.Response Body.
func (client CustomizedAcceleratorsClient) ValidateResponder(resp *http.Response) (result CustomizedAcceleratorValidateResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}