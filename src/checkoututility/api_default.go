/*
 * Adyen Checkout Utility API
 *
 * A web service containing utility functions available for merchants integrating with Checkout APIs. ## Authentication Each request to the Checkout Utility API must be signed with an API key. For this, obtain an API Key from your Customer Area, as described in [How to get the API key](https://docs.adyen.com/user-management/how-to-get-the-api-key). Then set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: Your_Checkout_API_key\" \\ ... ``` Note that when going live, you need to generate a new API Key to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Checkout API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v1/originKeys ```
 *
 * API version: 1
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package checkoututility

import (
	_context "context"
	_nethttp "net/http"

	"github.com/adyen/adyen-go-api-library/v2/src/common"
)

// CheckoutUtility CheckoutUtility service
type CheckoutUtility common.Service

/*
OriginKeys Create originKey values for one or more merchant domains.
This operation takes the origin domains and returns a JSON object containing the corresponding origin keys for the domains.
 * @param request CheckoutUtilityRequest - reference of CheckoutUtilityRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return CheckoutUtilityResponse
*/
func (a CheckoutUtility) OriginKeys(req *CheckoutUtilityRequest, ctxs ..._context.Context) (CheckoutUtilityResponse, *_nethttp.Response, error) {
	res := &CheckoutUtilityResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/originKeys", ctxs...)
	return *res, httpRes, err
}
