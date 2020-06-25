/*
 * Adyen for Platforms: Account API
 *
 * The Account API provides endpoints for managing account-related entities on your platform. These related entities include account holders, accounts, bank accounts, shareholders, and KYC-related documents. The management operations include actions such as creation, retrieval, updating, and deletion of them.  For more information, refer to our [documentation](https://docs.adyen.com/marketpay). ## Authentication To connect to the Account API, you must use basic authentication credentials of your web service user. If you don't have one, please contact the [Adyen Support Team](https://support.adyen.com/hc/en-us/requests/new). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@MarketPlace.YourMarketPlace\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning The Account API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://cal-test.adyen.com/cal/services/Account/v5/createAccountHolder ```
 *
 * API version: 5
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package platformsaccount
// ViasPhoneNumber struct for ViasPhoneNumber
type ViasPhoneNumber struct {
	// The two-character country code of the phone number. >The permitted country codes are defined in ISO-3166-1 alpha-2 (e.g. 'NL').
	PhoneCountryCode string `json:"phoneCountryCode"`
	// The phone number. >The inclusion of the phone number country code is not necessary.
	PhoneNumber string `json:"phoneNumber"`
	// The type of the phone number. >The following values are permitted: `Landline`, `Mobile`, `SIP`, `Fax`.
	PhoneType string `json:"phoneType,omitempty"`
}