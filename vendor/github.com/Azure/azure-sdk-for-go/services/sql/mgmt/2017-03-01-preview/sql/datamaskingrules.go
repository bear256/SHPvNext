package sql

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// DataMaskingRulesClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type DataMaskingRulesClient struct {
	BaseClient
}

// NewDataMaskingRulesClient creates an instance of the DataMaskingRulesClient client.
func NewDataMaskingRulesClient(subscriptionID string) DataMaskingRulesClient {
	return NewDataMaskingRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDataMaskingRulesClientWithBaseURI creates an instance of the DataMaskingRulesClient client.
func NewDataMaskingRulesClientWithBaseURI(baseURI string, subscriptionID string) DataMaskingRulesClient {
	return DataMaskingRulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a database data masking rule.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of the
// database. dataMaskingPolicyName is the name of the database for which the data masking rule applies.
// dataMaskingRuleName is the name of the data masking rule. parameters is the required parameters for creating or
// updating a data masking rule.
func (client DataMaskingRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, dataMaskingPolicyName string, dataMaskingRuleName string, parameters DataMaskingRule) (result DataMaskingRule, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.DataMaskingRuleProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.DataMaskingRuleProperties.SchemaName", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.DataMaskingRuleProperties.TableName", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.DataMaskingRuleProperties.ColumnName", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "sql.DataMaskingRulesClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, databaseName, dataMaskingPolicyName, dataMaskingRuleName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DataMaskingRulesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.DataMaskingRulesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DataMaskingRulesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DataMaskingRulesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, dataMaskingPolicyName string, dataMaskingRuleName string, parameters DataMaskingRule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":          autorest.Encode("path", databaseName),
		"dataMaskingPolicyName": autorest.Encode("path", dataMaskingPolicyName),
		"dataMaskingRuleName":   autorest.Encode("path", dataMaskingRuleName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"serverName":            autorest.Encode("path", serverName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/dataMaskingPolicies/{dataMaskingPolicyName}/rules/{dataMaskingRuleName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DataMaskingRulesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DataMaskingRulesClient) CreateOrUpdateResponder(resp *http.Response) (result DataMaskingRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDatabase gets a list of database data masking rules.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of the
// database. dataMaskingPolicyName is the name of the database for which the data masking rule applies.
func (client DataMaskingRulesClient) ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string, dataMaskingPolicyName string) (result DataMaskingRuleListResult, err error) {
	req, err := client.ListByDatabasePreparer(ctx, resourceGroupName, serverName, databaseName, dataMaskingPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DataMaskingRulesClient", "ListByDatabase", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDatabaseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.DataMaskingRulesClient", "ListByDatabase", resp, "Failure sending request")
		return
	}

	result, err = client.ListByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DataMaskingRulesClient", "ListByDatabase", resp, "Failure responding to request")
	}

	return
}

// ListByDatabasePreparer prepares the ListByDatabase request.
func (client DataMaskingRulesClient) ListByDatabasePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, dataMaskingPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":          autorest.Encode("path", databaseName),
		"dataMaskingPolicyName": autorest.Encode("path", dataMaskingPolicyName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"serverName":            autorest.Encode("path", serverName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/dataMaskingPolicies/{dataMaskingPolicyName}/rules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDatabaseSender sends the ListByDatabase request. The method will close the
// http.Response Body if it receives an error.
func (client DataMaskingRulesClient) ListByDatabaseSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByDatabaseResponder handles the response to the ListByDatabase request. The method always
// closes the http.Response Body.
func (client DataMaskingRulesClient) ListByDatabaseResponder(resp *http.Response) (result DataMaskingRuleListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
