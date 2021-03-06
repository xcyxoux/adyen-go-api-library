/*
 * Adyen Payment API
 *
 * A set of API endpoints that allow you to initiate, settle, and modify payments on the Adyen payments platform. You can use the API to accept card payments (including One-Click and 3D Secure), bank transfers, ewallets, and many other payment methods.  To learn more about the API, visit [Classic integration](https://docs.adyen.com/classic-integration).  ## Authentication To connect to the Payments API, you must use your basic authentication credentials. For this, create your web service user, as described in [How to get the WS user password](https://docs.adyen.com/user-management/how-to-get-the-web-service-ws-user-password). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@Company.YourCompany\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Payments API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://pal-test.adyen.com/pal/servlet/Payment/v52/authorise ```
 *
 * API version: 52
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package payments

// ThreeDS2RequestData struct for ThreeDS2RequestData
type ThreeDS2RequestData struct {
	// Required for [authentication-only integration](https://docs.adyen.com/checkout/3d-secure/other-3ds-flows/authentication-only). The acquiring BIN enrolled for 3D Secure 2. This string should match the value that you will use in the authorisation. Use 123456 on the Test platform.
	AcquirerBIN string `json:"acquirerBIN,omitempty"`
	// Required for [authentication-only integration](https://docs.adyen.com/checkout/3d-secure/other-3ds-flows/authentication-only). The merchantId that is enrolled for 3D Secure 2 by the merchant's acquirer. This string should match the value that you will use in the authorisation. Use 123456 on the Test platform.
	AcquirerMerchantID string `json:"acquirerMerchantID,omitempty"`
	// If set to true, you will only perform the [3D Secure 2 authentication](https://docs.adyen.com/checkout/3d-secure/other-3ds-flows/authentication-only), and not the payment authorisation.
	AuthenticationOnly bool `json:"authenticationOnly,omitempty"`
	// Possibility to specify a preference for receiving a challenge from the issuer. Allowed values: * `noPreference` * `requestNoChallenge` * `requestChallenge` * `requestChallengeAsMandate`
	ChallengeIndicator string `json:"challengeIndicator,omitempty"`
	// The environment of the shopper. Allowed values: * `app` * `browser`
	DeviceChannel       string               `json:"deviceChannel"`
	DeviceRenderOptions *DeviceRenderOptions `json:"deviceRenderOptions,omitempty"`
	// Required for merchants that have been enrolled for 3D Secure 2 by another party than Adyen, mostly [authentication-only integrations](https://docs.adyen.com/checkout/3d-secure/other-3ds-flows/authentication-only). The `mcc` is a four-digit code with which the previously given `acquirerMerchantID` is registered at the scheme.
	Mcc string `json:"mcc,omitempty"`
	// Required for [authentication-only integration](https://docs.adyen.com/checkout/3d-secure/other-3ds-flows/authentication-only). The merchant name that the issuer presents to the shopper if they get a challenge. We recommend to use the same value that you will use in the authorization. Maximum length is 40 characters. > Optional for a [full 3D Secure 2 integration](https://docs.adyen.com/checkout/3d-secure/native-3ds2/api-integration). Use this field if you are enrolled for 3D Secure 2 with us and want to override the merchant name already configured on your account.
	MerchantName string `json:"merchantName,omitempty"`
	// The `messageVersion` value indicating the 3D Secure 2 protocol version.
	MessageVersion string `json:"messageVersion,omitempty"`
	// URL to where the issuer should send the `CRes`. Required if you are not using components for `channel` **Web** or if you are using classic integration `deviceChannel` **browser**.
	NotificationURL string `json:"notificationURL,omitempty"`
	// The `sdkAppID` value as received from the 3D Secure 2 SDK. Required for `deviceChannel` set to **app**.
	SdkAppID string `json:"sdkAppID,omitempty"`
	// The `sdkEncData` value as received from the 3D Secure 2 SDK. Required for `deviceChannel` set to **app**.
	SdkEncData     string          `json:"sdkEncData,omitempty"`
	SdkEphemPubKey *SDKEphemPubKey `json:"sdkEphemPubKey,omitempty"`
	// The maximum amount of time in minutes for the 3D Secure 2 authentication process. Optional and only for `deviceChannel` set to **app**. Defaults to **60** minutes.
	SdkMaxTimeout int32 `json:"sdkMaxTimeout,omitempty"`
	// The `sdkReferenceNumber` value as received from the 3D Secure 2 SDK. Only for `deviceChannel` set to **app**.
	SdkReferenceNumber string `json:"sdkReferenceNumber,omitempty"`
	// The `sdkTransID` value as received from the 3D Secure 2 SDK. Only for `deviceChannel` set to **app**.
	SdkTransID string `json:"sdkTransID,omitempty"`
	// Completion indicator for the device fingerprinting.
	ThreeDSCompInd string `json:"threeDSCompInd,omitempty"`
	// Required for [authentication-only integration](https://docs.adyen.com/checkout/3d-secure/other-3ds-flows/authentication-only) for Visa. Unique 3D Secure requestor identifier assigned by the Directory Server when you enrol for 3D Secure 2.
	ThreeDSRequestorID string `json:"threeDSRequestorID,omitempty"`
	// Required for [authentication-only integration](https://docs.adyen.com/checkout/3d-secure/other-3ds-flows/authentication-only) for Visa. Unique 3D Secure requestor name assigned by the Directory Server when you enrol for 3D Secure 2.
	ThreeDSRequestorName string `json:"threeDSRequestorName,omitempty"`
	// URL of the (customer service) website that will be shown to the shopper in case of technical errors during the 3D Secure 2 process.
	ThreeDSRequestorURL string `json:"threeDSRequestorURL,omitempty"`
	// Identify the type of the transaction being authenticated.
	TransactionType string `json:"transactionType,omitempty"`
	// The `whiteListStatus` value returned from a previous 3D Secure 2 transaction, only applicable for 3D Secure 2 protocol version 2.2.0.
	WhiteListStatus string `json:"whiteListStatus,omitempty"`
}
