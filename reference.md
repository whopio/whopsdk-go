# Reference
## AccessTokens
<details><summary><code>client.AccessTokens.Create(request) -> *whopsdk.AccessToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a short-lived access token for authenticating API requests. When using API key authentication, provide company_id or user_id. When using OAuth, the user is derived from the token. Use this token with Whop's web and mobile embedded components.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateAccessTokensRequest{}
client.AccessTokens.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `*string` — The unique identifier of the company to generate the token for, starting with 'biz_'. The API key must have permission to access this company.
    
</dd>
</dl>

<dl>
<dd>

**expiresAt:** `*time.Time` — The expiration timestamp for the access token. Defaults to 1 hour from now, with a maximum of 3 hours.
    
</dd>
</dl>

<dl>
<dd>

**scopedActions:** `[]string` — An array of permission scopes to grant to the access token. If empty or omitted, all permissions from the authenticating credential are inherited. Must be a subset of the credential's permissions.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — The unique identifier of the user to generate the token for, starting with 'user_'. The API key must have permission to access this user.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## AccountLinks
<details><summary><code>client.AccountLinks.Create(request) -> *whopsdk.AccountLink</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Generate a URL that directs a sub-merchant to their account portal, such as the hosted payouts dashboard or the KYC onboarding flow.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateAccountLinksRequest{
    CompanyID: "biz_xxxxxxxxxxxxxx",
    RefreshURL: "refresh_url",
    ReturnURL: "return_url",
    UseCase: whopsdk.AccountLinkUseCasesAccountOnboarding,
}
client.AccountLinks.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `string` — The unique identifier of the company to generate the link for, starting with 'biz_'. Must be a sub-merchant of the API key's company.
    
</dd>
</dl>

<dl>
<dd>

**refreshURL:** `string` — The URL to redirect the user to if the session expires and needs to be re-authenticated, such as 'https://example.com/refresh'.
    
</dd>
</dl>

<dl>
<dd>

**returnURL:** `string` — The URL to redirect the user to when they want to return to your site, such as 'https://example.com/return'.
    
</dd>
</dl>

<dl>
<dd>

**useCase:** `*whopsdk.AccountLinkUseCases` — The purpose of the account link, such as hosted payouts portal or hosted KYC onboarding.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounts
<details><summary><code>client.Accounts.List() -> *whopsdk.ListAccountsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists accounts visible to the credential. User tokens return the user's business accounts; Account API keys return the requesting account and its connected accounts. Pass `parent_account_id` to return only that parent account's connected accounts.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListAccountsRequest{}
client.Accounts.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**first:** `*int` — The number of accounts to return (default 10, max 50).
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor; returns accounts after this position.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — The number of accounts to return from the end of the range.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — A cursor; returns accounts before this position.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListAccountsRequestOrder` — The field to sort accounts by. `volume` requires `stats:read` on the parent account.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListAccountsRequestDirection` — Sort direction.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.ListAccountsRequestStatus` — Return only accounts with this status: `active` (includes accounts that have not entered payments review) or `suspended`.
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` — Free-text filter on account title or ID. `%` and `_` match literally.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — Return only accounts created after this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — Return only accounts created before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**volumeMin:** `*float64` — Return only accounts whose lifetime USD volume is at least this value. Requires `stats:read` on the parent account.
    
</dd>
</dl>

<dl>
<dd>

**volumeMax:** `*float64` — Return only accounts whose lifetime USD volume is at most this value. Requires `stats:read` on the parent account.
    
</dd>
</dl>

<dl>
<dd>

**parentAccountID:** `*string` — For platforms: the parent account ID whose direct connected accounts to return. Requires `payout:account:read` on the parent account.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.Create(request) -> *whopsdk.Account</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an account. User tokens create business accounts; Account API keys create connected accounts. Tax fields (`tax_remitted_by`, `tax_type`, `product_tax_code_id`, `business_address`, `tax_identifiers`) are configured with Update Account, not at creation.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateAccountsRequest{}
client.Accounts.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**affiliateCode:** `*string` — The username, if any, of the partner who referred this account
    
</dd>
</dl>

<dl>
<dd>

**blueprintID:** `*string` — The blueprint App ID, prefixed `app_`. Creates a hosted website for the account and queues its deployment asynchronously; the Account response does not report deployment completion.
    
</dd>
</dl>

<dl>
<dd>

**country:** `*string` — The ISO 3166-1 alpha-2 country code where the account's business is located (e.g. `US`). Defaults to the parent account's country for connected accounts.
    
</dd>
</dl>

<dl>
<dd>

**email:** `*string` — The email address of the account owner. Required for Account API key requests.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Arbitrary key/value metadata to store on the account.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The display name of the account. Defaults to `metadata.external_id` or the owner's email when omitted.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.Me() -> *whopsdk.Account</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the account associated with the current Account API key.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounts.Me(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.Retrieve(ID) -> *whopsdk.Account</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a single account by ID or public route when it is visible to the credential, including its crypto wallet. The reserved id `me` retrieves the account associated with the current Account API key; user tokens have no single account, so they must address one by ID or route.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveAccountsRequest{
    ID: "id",
}
client.Accounts.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Account ID, prefixed `biz_`, its public route, or `me` for the account associated with the current API key.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.Update(ID, request) -> *whopsdk.Account</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an account. User tokens can update business accounts; Account API keys can update connected accounts. The reserved id `me` — accepted on Retrieve Account — resolves to the requesting account, which an Account API key cannot edit, so updates must name the connected account by its `biz_` id.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateAccountsRequest{
    ID: "id",
}
client.Accounts.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Account ID, prefixed `biz_`.
    
</dd>
</dl>

<dl>
<dd>

**affiliateApplicationRequired:** `*bool` — Whether prospective affiliates must submit an application before promoting this account.
    
</dd>
</dl>

<dl>
<dd>

**affiliateInstructions:** `*string` — Guidelines shown to affiliates promoting this account.
    
</dd>
</dl>

<dl>
<dd>

**bannerImage:** `*whopsdk.UpdateAccountsRequestBannerImage` — Account banner image, used as the cover photo when creating a Whop-managed Facebook page. Image files up to 10 MB, except `image/gif`. Pass a JSON object containing an `id` from [Create File](/api-reference/files/create-file).
    
</dd>
</dl>

<dl>
<dd>

**businessAddress:** `*whopsdk.UpdateAccountsRequestBusinessAddress` — Account business address used to calculate tax. A complete address in a supported country is required when `tax_remitted_by` is `self`.
    
</dd>
</dl>

<dl>
<dd>

**businessName:** `*string` — The legal business name used with the account's tax address.
    
</dd>
</dl>

<dl>
<dd>

**businessType:** `*whopsdk.UpdateAccountsRequestBusinessType` — High-level business category for the account. See the [business types and industries glossary](/api-reference/beta/accounts/account#business-types-and-industries-glossary) for valid values.
    
</dd>
</dl>

<dl>
<dd>

**collectVatID:** `*bool` — Whether checkout shows a VAT/tax ID field for buyers to optionally enter. Does not require a VAT ID to purchase.
    
</dd>
</dl>

<dl>
<dd>

**country:** `*string` — Country where the account is located.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — Account promotional description. When creating a Whop-managed Facebook page, it is truncated to 155 characters and used as the About text.
    
</dd>
</dl>

<dl>
<dd>

**featuredAffiliateProductID:** `*string` — The ID of the product to feature for affiliates. Pass `null` to clear.
    
</dd>
</dl>

<dl>
<dd>

**homePreferences:** `[]*whopsdk.UpdateAccountsRequestHomePreferencesItem` — Public account home page preferences.
    
</dd>
</dl>

<dl>
<dd>

**industryGroup:** `*whopsdk.UpdateAccountsRequestIndustryGroup` — Account industry group. See the [business types and industries glossary](/api-reference/beta/accounts/account#business-types-and-industries-glossary) for valid values.
    
</dd>
</dl>

<dl>
<dd>

**industryType:** `*string` — Specific industry vertical for the account. See the [business types and industries glossary](/api-reference/beta/accounts/account#business-types-and-industries-glossary) for valid values.
    
</dd>
</dl>

<dl>
<dd>

**invoicePrefix:** `*string` — Prefix used for account invoices.
    
</dd>
</dl>

<dl>
<dd>

**logo:** `*whopsdk.UpdateAccountsRequestLogo` — Account logo, used as the profile picture when creating a Whop-managed Facebook page. Image files up to 5 MB. Pass a JSON object containing an `id` from [Create File](/api-reference/files/create-file).
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Arbitrary key/value metadata to store on the account.
    
</dd>
</dl>

<dl>
<dd>

**onboardingType:** `*whopsdk.UpdateAccountsRequestOnboardingType` — The type of onboarding the account has completed.
    
</dd>
</dl>

<dl>
<dd>

**opengraphImage:** `*whopsdk.UpdateAccountsRequestOpengraphImage` — Open Graph preview media used when the account is shared. Image and video files up to 5 MB. Pass a JSON object containing an `id` from [Create File](/api-reference/files/create-file).
    
</dd>
</dl>

<dl>
<dd>

**opengraphImageVariant:** `*whopsdk.UpdateAccountsRequestOpengraphImageVariant` — The account Open Graph image variant.
    
</dd>
</dl>

<dl>
<dd>

**otherBusinessDescription:** `*string` — The description of the business type when business_type is other.
    
</dd>
</dl>

<dl>
<dd>

**otherIndustryDescription:** `*string` — The description of the industry type when industry_type is other.
    
</dd>
</dl>

<dl>
<dd>

**productTaxCodeID:** `*string` — ID of the tax classification code applied by default to the account's products. See the available [product categories](https://docs.numeral.com/essentials/product-categories).
    
</dd>
</dl>

<dl>
<dd>

**require2Fa:** `*bool` — Whether the account requires authorized users to have two-factor authentication enabled.
    
</dd>
</dl>

<dl>
<dd>

**route:** `*string` — The unique URL slug for the account.
    
</dd>
</dl>

<dl>
<dd>

**sendCustomerEmails:** `*bool` — Whether Whop sends transactional emails to customers on behalf of this account.
    
</dd>
</dl>

<dl>
<dd>

**showJoinedWhops:** `*bool` — Whether the account appears in joined whops on other accounts.
    
</dd>
</dl>

<dl>
<dd>

**showReviewsDtc:** `*bool` — Whether reviews are displayed on direct-to-consumer product pages.
    
</dd>
</dl>

<dl>
<dd>

**showUserDirectory:** `*bool` — Whether the account shows users in the user directory.
    
</dd>
</dl>

<dl>
<dd>

**socialLinks:** `[]map[string]any` — The full list of social links to display for the account.
    
</dd>
</dl>

<dl>
<dd>

**storePageConfig:** `*whopsdk.UpdateAccountsRequestStorePageConfig` — Account store page display configuration.
    
</dd>
</dl>

<dl>
<dd>

**targetAudience:** `*string` — The target audience for this account.
    
</dd>
</dl>

<dl>
<dd>

**taxCollectionEnabledStates:** `[]*whopsdk.UpdateAccountsRequestTaxCollectionEnabledStatesItem` — US state codes (50 states plus `DC`) where the account collects tax. Replaces the full set on update. Only settable when `tax_remitted_by` is `self`.
    
</dd>
</dl>

<dl>
<dd>

**taxIdentifiers:** `[]*whopsdk.UpdateAccountsRequestTaxIdentifiersItem` — Account tax/VAT registrations to add or update. When `tax_remitted_by` is `self`, tax is calculated and collected only in the countries where the account holds a registration.
    
</dd>
</dl>

<dl>
<dd>

**taxRemittedBy:** `*whopsdk.UpdateAccountsRequestTaxRemittedBy` — Determines whether Whop or the account calculates and remits tax. The account must provide a supported-country business address when it self-remits.
    
</dd>
</dl>

<dl>
<dd>

**taxType:** `*whopsdk.UpdateAccountsRequestTaxType` — Determines whether tax is included in the listed price or added at checkout.
    
</dd>
</dl>

<dl>
<dd>

**threeDsLevel:** `*whopsdk.UpdateAccountsRequestThreeDsLevel` — Account-level 3D Secure behavior. Set `mandate_challenge` to require cardholder verification on supported card payments, or `null` to use the standard checkout flow.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The display name of the account.
    
</dd>
</dl>

<dl>
<dd>

**useLogoAsOpengraphImageFallback:** `*bool` — Whether the account uses its logo as the fallback Open Graph image.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.FormCompany(ID, request) -> *whopsdk.FormCompanyAccountsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Starts an LLC or C-Corp formation for a business account. Defaults to an LLC; set `entity_type` to `c_corp` to form a C-Corp, which additionally requires `share_structure` and officer `roles` on every founder. On submission, the application is validated and the response returns a hosted checkout URL. Once paid, the filing is submitted. Track progress through the account's [`company_formation`](/api-reference/beta/accounts/retrieve-account) field on Retrieve Account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.FormCompanyAccountsRequest{
    ID: "id",
    BusinessAddress: &whopsdk.FormCompanyAccountsRequestBusinessAddress{
        City: "Austin",
        Country: "US",
        Line1: "4180 Burnet Rd",
        Line2: whopsdk.String(
            "Suite 2",
        ),
        PostalCode: "78756",
        State: "TX",
    },
    BusinessName: "Shine Time Auto Detailing",
    BusinessPhone: whopsdk.String(
        "+15125550142",
    ),
    BusinessType: "brick_and_mortar",
    BusinessWebsite: whopsdk.String(
        "https://shinetime.example",
    ),
    EntitySuffix: whopsdk.FormCompanyAccountsRequestEntitySuffixLlc.Ptr(),
    EntityType: whopsdk.FormCompanyAccountsRequestEntityTypeLlc.Ptr(),
    ExpediteEin: whopsdk.Bool(
        true,
    ),
    FormationState: whopsdk.FormCompanyAccountsRequestFormationStateTx,
    Founders: []*whopsdk.FormCompanyAccountsRequestFoundersItem{
        &whopsdk.FormCompanyAccountsRequestFoundersItem{
            Address: &whopsdk.FormCompanyAccountsRequestFoundersItemAddress{
                City: "Austin",
                Country: "US",
                Line1: "907 Ridgemont Dr",
                Line2: whopsdk.String(
                    "Apt 4",
                ),
                PostalCode: "78704",
                State: "TX",
            },
            DateOfBirth: whopsdk.String(
                "1988-03-14",
            ),
            Email: "marcus@shinetime.example",
            FirstName: "Marcus",
            IsPrimary: true,
            LastName: "Webb",
            OwnershipPercentage: whopsdk.Float64(
                100,
            ),
            Phone: "+15125550142",
            Roles: []whopsdk.FormCompanyAccountsRequestFoundersItemRolesItem{
                whopsdk.FormCompanyAccountsRequestFoundersItemRolesItemPresident,
            },
            Ssn: whopsdk.String(
                "123-45-6789",
            ),
        },
    },
    IndustryGroup: "automotive",
    IndustryType: "car_wash",
    ShareStructure: &whopsdk.FormCompanyAccountsRequestShareStructure{
        NumberOfShares: 123,
        Value: 123,
    },
    UseRegisteredAgent: whopsdk.Bool(
        true,
    ),
}
client.Accounts.FormCompany(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Account ID, prefixed `biz_`.
    
</dd>
</dl>

<dl>
<dd>

**businessAddress:** `*whopsdk.FormCompanyAccountsRequestBusinessAddress` — Company mailing address. Required unless `use_registered_agent` is `true`.
    
</dd>
</dl>

<dl>
<dd>

**businessName:** `string` — Legal name for the new company.
    
</dd>
</dl>

<dl>
<dd>

**businessPhone:** `*string` — Business phone number in E.164 format, for example `+12125550100`. Required unless `use_registered_agent` is `true`.
    
</dd>
</dl>

<dl>
<dd>

**businessType:** `string` — High-level business category, from the Whop business taxonomy. Valid values are listed on [business types and industries glossary](/api-reference/beta/accounts/account#business-types-and-industries-glossary).
    
</dd>
</dl>

<dl>
<dd>

**businessWebsite:** `*string` — Company website URL.
    
</dd>
</dl>

<dl>
<dd>

**entitySuffix:** `*whopsdk.FormCompanyAccountsRequestEntitySuffix` — Legal entity ending appended to `business_name`. LLC formations accept `LLC`, `L.L.C`, `L.L.C.` or `Limited Liability Company` and default to `LLC`; C-Corp formations accept `Inc`, `Inc.`, `Incorporated`, `Corp.`, `Corporation`, `C Corp`, `C Corporation`, `CCorp` or `Company` and default to `Inc.`. Unrecognized values fall back to the default for the entity type.
    
</dd>
</dl>

<dl>
<dd>

**entityType:** `*whopsdk.FormCompanyAccountsRequestEntityType` — Legal entity type to form. Defaults to `llc`.
    
</dd>
</dl>

<dl>
<dd>

**expediteEin:** `*bool` — Request expedited EIN processing for an additional fee. Available only when no founder supplies an SSN.
    
</dd>
</dl>

<dl>
<dd>

**formationState:** `*whopsdk.FormCompanyAccountsRequestFormationState` — Two-letter code of the US state (or `DC`) to form the company in.
    
</dd>
</dl>

<dl>
<dd>

**founders:** `[]*whopsdk.FormCompanyAccountsRequestFoundersItem` — The company's founders. Exactly one must be marked `is_primary` — the responsible party for the filing.
    
</dd>
</dl>

<dl>
<dd>

**industryGroup:** `string` — Industry group, from the Whop business taxonomy. Valid values are listed on [business types and industries glossary](/api-reference/beta/accounts/account#business-types-and-industries-glossary).
    
</dd>
</dl>

<dl>
<dd>

**industryType:** `string` — Specific industry vertical, from the Whop business taxonomy. Valid values are listed on [business types and industries glossary](/api-reference/beta/accounts/account#business-types-and-industries-glossary).
    
</dd>
</dl>

<dl>
<dd>

**shareStructure:** `*whopsdk.FormCompanyAccountsRequestShareStructure` — Authorized share structure. Required when `entity_type` is `c_corp`; ignored for LLCs.
    
</dd>
</dl>

<dl>
<dd>

**useRegisteredAgent:** `*bool` — Use the registered agent's address as the company address instead of `business_address`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.TransferOwnership(ID, request) -> *whopsdk.TransferOwnershipAccountsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Transfers ownership of the account to another user, identified by user ID or email address. If the recipient already holds the owner role, ownership moves immediately; otherwise they get an invite and ownership moves when they accept.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.TransferOwnershipAccountsRequest{
    ID: "id",
    Identifier: "marcus@shinetime.example",
}
client.Accounts.TransferOwnership(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Account ID, prefixed `biz_`.
    
</dd>
</dl>

<dl>
<dd>

**asPartner:** `*bool` — If true, the current owner is credited as the account's Whop partner, earning partner commission on its sales. Requires the current owner to already be an enrolled Whop partner. Skipped if the account already has an active partner.
    
</dd>
</dl>

<dl>
<dd>

**identifier:** `string` — The user to transfer ownership to: a user ID (`user_*`) or an email address. An email address with no Whop account yet is sent an invite to create one.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ad Campaigns
<details><summary><code>client.AdCampaigns.List() -> *whopsdk.ListAdCampaignsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the ad campaigns for an account, with stats over the requested window.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListAdCampaignsRequest{}
client.AdCampaigns.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The account the campaigns belong to. Defaults to the account-scoped key's own account.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.ListAdCampaignsRequestStatus` — Only return campaigns with this status.
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` — Filter campaigns by a title or ID substring.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListAdCampaignsRequestOrder` — The field to sort by. Defaults to created_at. Stat columns (spend, impressions, …) rank over the stats_from/stats_to window across the whole list, not just the current page. results, cost_per_result and return_on_ad_spend rank by the same Whop pixel-attributed values the response reports.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListAdCampaignsRequestDirection` — The sort direction. Defaults to desc.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Only return campaigns created before this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Only return campaigns created after this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**statsFrom:** `*string` — Start of the stats window. Defaults to all-time.
    
</dd>
</dl>

<dl>
<dd>

**statsTo:** `*string` — End of the stats window. Defaults to now.
    
</dd>
</dl>

<dl>
<dd>

**timeZone:** `*string` — IANA timezone (e.g. America/New_York) the stats window is interpreted in. Bare stats_from/stats_to dates resolve to day boundaries on this clock. Defaults to UTC.
    
</dd>
</dl>

<dl>
<dd>

**attributionModel:** `*whopsdk.ListAdCampaignsRequestAttributionModel` — Attribution model the conversion stats count under (defaults to last_touch). Under both models a journey with any whop ad touch attributes to whop; the model picks which whop touch credits the entity and which non-whop source wins otherwise.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of campaigns to return.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor to fetch the page after (from page_info.end_cursor).
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — The number of campaigns to return from the end of the range.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Cursor to fetch the page before (from page_info.start_cursor).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AdCampaigns.Create(request) -> *whopsdk.AdCampaign</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an ad campaign for an account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateAdCampaignsRequest{
    Objective: whopsdk.CreateAdCampaignsRequestObjectiveAwareness,
    Platform: whopsdk.CreateAdCampaignsRequestPlatformMeta,
    Title: "Now hiring mobile detailers — Austin",
}
client.AdCampaigns.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The account to create the campaign under. Defaults to the account-scoped key's own account.
    
</dd>
</dl>

<dl>
<dd>

**bidType:** `*whopsdk.CreateAdCampaignsRequestBidType` — How delivery bids in the ad auction: `minimum_cost` gets the most results for the budget, `average_target` holds an average cost per result, `maximum_target` never bids above a cap. Only for campaigns that own the budget.
    
</dd>
</dl>

<dl>
<dd>

**budgetAmount:** `*float64` — The campaign's budget, in the ad account's currency. Required when budget_optimization is `ad_campaign`; omit when each ad group sets its own budget.
    
</dd>
</dl>

<dl>
<dd>

**budgetOptimization:** `*whopsdk.CreateAdCampaignsRequestBudgetOptimization` — Which level owns the budget: the whole campaign (`ad_campaign`) or each ad group individually (`ad_group`). Defaults to `ad_group`.
    
</dd>
</dl>

<dl>
<dd>

**budgetType:** `*whopsdk.CreateAdCampaignsRequestBudgetType` — Whether the budget is spent per day (`daily`) or over the campaign's full run (`lifetime`). Defaults to `daily`.
    
</dd>
</dl>

<dl>
<dd>

**desiredCostPerResult:** `*float64` — Cost per result to aim for (`average_target`) or never exceed (`maximum_target`). Only for campaigns that own the budget.
    
</dd>
</dl>

<dl>
<dd>

**endsAt:** `*string` — When the campaign stops delivering, as an ISO 8601 timestamp. Only for campaigns that own the budget.
    
</dd>
</dl>

<dl>
<dd>

**objective:** `*whopsdk.CreateAdCampaignsRequestObjective` — The goal the campaign optimizes toward.
    
</dd>
</dl>

<dl>
<dd>

**platform:** `*whopsdk.CreateAdCampaignsRequestPlatform` — The ad network the campaign runs on.
    
</dd>
</dl>

<dl>
<dd>

**specialAdCategories:** `[]*whopsdk.CreateAdCampaignsRequestSpecialAdCategoriesItem` — Regulated categories the campaign falls under. Ads in these categories are subject to extra targeting restrictions.
    
</dd>
</dl>

<dl>
<dd>

**startsAt:** `*string` — When the campaign starts delivering, as an ISO 8601 timestamp. Only for campaigns that own the budget.
    
</dd>
</dl>

<dl>
<dd>

**title:** `string` — The title of the campaign.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AdCampaigns.Retrieve(ID) -> *whopsdk.AdCampaign</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a single ad campaign with stats over the requested window.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveAdCampaignsRequest{
    ID: "id",
}
client.AdCampaigns.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ad campaign ID.
    
</dd>
</dl>

<dl>
<dd>

**statsFrom:** `*string` — Start of the stats window.
    
</dd>
</dl>

<dl>
<dd>

**statsTo:** `*string` — End of the stats window.
    
</dd>
</dl>

<dl>
<dd>

**timeZone:** `*string` — IANA timezone the stats window is interpreted in. Defaults to UTC.
    
</dd>
</dl>

<dl>
<dd>

**attributionModel:** `*whopsdk.RetrieveAdCampaignsRequestAttributionModel` — Attribution model the conversion stats count under (defaults to last_touch). Under both models a journey with any whop ad touch attributes to whop; the model picks which whop touch credits the entity and which non-whop source wins otherwise.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AdCampaigns.Delete(ID) -> *whopsdk.DeleteAdCampaignsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes an ad campaign and archives it on the ad platform (cascades to ad groups and ads).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteAdCampaignsRequest{
    ID: "id",
}
client.AdCampaigns.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ad campaign ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AdCampaigns.Update(ID, request) -> *whopsdk.AdCampaign</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an ad campaign's editable fields (title, budget, schedule, bid strategy, special ad categories, and, before launch, budget type and budget optimization), and launches a draft campaign by setting status to active. Objective and desired cost per result are fixed at creation and cannot be changed.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateAdCampaignsRequest{
    ID: "id",
}
client.AdCampaigns.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ad campaign ID.
    
</dd>
</dl>

<dl>
<dd>

**bidType:** `*whopsdk.UpdateAdCampaignsRequestBidType` — How delivery bids in the ad auction: `minimum_cost` gets the most results for the budget, `average_target` holds an average cost per result, `maximum_target` never bids above a cap. Switching to `minimum_cost` clears the cap amounts stored on the campaign's ad groups. Only for campaigns that own the budget.
    
</dd>
</dl>

<dl>
<dd>

**budgetAmount:** `*float64` — The campaign budget, in the account's currency. Interpreted as daily or lifetime per the campaign's budget type, including a budget_type sent in the same request.
    
</dd>
</dl>

<dl>
<dd>

**budgetOptimization:** `*whopsdk.UpdateAdCampaignsRequestBudgetOptimization` — Which level owns the budget: the whole campaign (`ad_campaign`) or each ad group individually (`ad_group`). Only changeable before the campaign is live on the ad network; switching to `ad_campaign` requires budget_amount in the same request, and switching to `ad_group` clears the campaign budget.
    
</dd>
</dl>

<dl>
<dd>

**budgetType:** `*whopsdk.UpdateAdCampaignsRequestBudgetType` — Whether `budget_amount` is spent per day (`daily`) or over the campaign's full run (`lifetime`). Only changeable while the campaign is a draft; send budget_amount in the same request so the amount lands on the new type.
    
</dd>
</dl>

<dl>
<dd>

**endsAt:** `*string` — When the campaign stops delivering, as an ISO 8601 timestamp. Only for campaigns that own the budget.
    
</dd>
</dl>

<dl>
<dd>

**specialAdCategories:** `[]*whopsdk.UpdateAdCampaignsRequestSpecialAdCategoriesItem` — Regulated categories the campaign falls under. Editable on any campaign, draft or launched; pass an empty array to clear.
    
</dd>
</dl>

<dl>
<dd>

**startsAt:** `*string` — When the campaign starts delivering, as an ISO 8601 timestamp. Only for campaigns that own the budget.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.UpdateAdCampaignsRequestStatus` — Set to active to launch a draft campaign (moderates and pushes it live). Live-campaign pause and resume use the pause and unpause actions.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The name of the campaign.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AdCampaigns.Duplicate(ID, request) -> *whopsdk.DuplicateAdCampaignsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates copies of the campaign in `duplicating` status and returns them; each copy transitions to `draft` once duplication completes. Poll each returned campaign until it leaves `duplicating` — a copy that could not be completed is deleted and returns 404.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DuplicateAdCampaignsRequest{
    ID: "id",
}
client.AdCampaigns.Duplicate(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ad campaign ID.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — Number of copies to create (1-10). Defaults to 1.
    
</dd>
</dl>

<dl>
<dd>

**preserveEngagement:** `*bool` — Whether the copied ads keep the original posts' engagement (likes, comments, shares). Defaults to false.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AdCampaigns.Pause(ID) -> *whopsdk.AdCampaign</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Pauses an active ad campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.PauseAdCampaignsRequest{
    ID: "id",
}
client.AdCampaigns.Pause(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ad campaign ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AdCampaigns.RetryPayment(ID) -> *whopsdk.AdCampaign</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retries billing for an ad campaign whose payment previously failed.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetryPaymentAdCampaignsRequest{
    ID: "id",
}
client.AdCampaigns.RetryPayment(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ad campaign ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AdCampaigns.Unpause(ID) -> *whopsdk.AdCampaign</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Resumes a paused ad campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UnpauseAdCampaignsRequest{
    ID: "id",
}
client.AdCampaigns.Unpause(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ad campaign ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ad Groups
<details><summary><code>client.AdGroups.List() -> *whopsdk.ListAdGroupsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists ad groups for the account, newest first.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListAdGroupsRequest{
    AdCampaignIDs: []*string{
        whopsdk.String(
            "adcamp_xxxxxxxxxxxxxx",
        ),
    },
}
client.AdGroups.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — Account whose ad groups to list. Defaults to the authenticated account.
    
</dd>
</dl>

<dl>
<dd>

**adCampaignID:** `*string` — Filter to ad groups in this campaign.
    
</dd>
</dl>

<dl>
<dd>

**adCampaignIDs:** `*string` — Filter to ad groups in these campaigns (max 100). Repeat the parameter for each id (ad_campaign_ids=a&ad_campaign_ids=b).
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.ListAdGroupsRequestStatus` — Filter to ad groups with this status.
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` — Filter ad groups by a title or ID substring.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListAdGroupsRequestOrder` — The field to sort by. Defaults to created_at. Stat columns (spend, impressions, …) rank over the stats_from/stats_to window across the whole list, not just the current page. results, cost_per_result and return_on_ad_spend rank by the same Whop pixel-attributed values the response reports.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListAdGroupsRequestDirection` — The sort direction. Defaults to desc.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Only return ad groups created before this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Only return ad groups created after this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**statsFrom:** `*string` — Start of the stats window. Defaults to all-time.
    
</dd>
</dl>

<dl>
<dd>

**statsTo:** `*string` — End of the stats window. Defaults to now.
    
</dd>
</dl>

<dl>
<dd>

**timeZone:** `*string` — IANA timezone (e.g. America/New_York) the stats window is interpreted in. Bare stats_from/stats_to dates resolve to day boundaries on this clock. Defaults to UTC.
    
</dd>
</dl>

<dl>
<dd>

**attributionModel:** `*whopsdk.ListAdGroupsRequestAttributionModel` — Attribution model the conversion stats count under (defaults to last_touch). Under both models a journey with any whop ad touch attributes to whop; the model picks which whop touch credits the entity and which non-whop source wins otherwise.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of ad groups to return.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor to fetch the page after (from page_info.end_cursor).
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — The number of ad groups to return from the end of the range.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Cursor to fetch the page before (from page_info.start_cursor).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AdGroups.Create(request) -> *whopsdk.AdGroup</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an ad group (ad set) in a campaign.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateAdGroupsRequest{
    AdCampaignID: "adcamp_xxxxxxxxxxxxxx",
}
client.AdGroups.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**adCampaignID:** `string` — The ad campaign to create the ad group in, prefixed `adcamp_`.
    
</dd>
</dl>

<dl>
<dd>

**audiences:** `*whopsdk.AdGroupAudiencesBody` — Saved audiences to deliver to or exclude. Can't be combined with demographics.automatic.
    
</dd>
</dl>

<dl>
<dd>

**bidType:** `*whopsdk.CreateAdGroupsRequestBidType` — How delivery bids are set in the ad auction. Target-based strategies use `desired_cost_per_result`.
    
</dd>
</dl>

<dl>
<dd>

**budgetAmount:** `*float64` — This ad group's budget, in the ad account's currency. Omit when the budget is set on the campaign instead.
    
</dd>
</dl>

<dl>
<dd>

**budgetType:** `*whopsdk.CreateAdGroupsRequestBudgetType` — Whether budget_amount is spent per day (`daily`) or over the ad group's full run (`lifetime`).
    
</dd>
</dl>

<dl>
<dd>

**conversionEvent:** `*whopsdk.ConversionEvent` 
    
</dd>
</dl>

<dl>
<dd>

**conversionLocation:** `*whopsdk.CreateAdGroupsRequestConversionLocation` — Where the outcome being optimized for occurs, such as a website visit, social-profile visit, messaging conversation, ad interaction, or lead-form submission. The lead form itself is set on the ad.
    
</dd>
</dl>

<dl>
<dd>

**demographics:** `*whopsdk.AdGroupDemographicsBody` — Age, gender, and automatic-audience targeting.
    
</dd>
</dl>

<dl>
<dd>

**desiredCostPerResult:** `*float64` — Cost per result to aim for (`average_target`) or never exceed (`maximum_target`).
    
</dd>
</dl>

<dl>
<dd>

**detailedTargeting:** `*whopsdk.AdGroupDetailedTargetingBody` — Interest, behavior, and demographic targeting, using categories from the ad platform's targeting taxonomy. Entries across interests, behaviors, and demographics are OR'd together (anyone matching any entry is reached), matching Ads Manager's detailed-targeting box. At most 100 entries per section. Can't be combined with demographics.automatic, and unavailable to campaigns with special_ad_categories. Send the complete intended state — a section you omit is cleared.
    
</dd>
</dl>

<dl>
<dd>

**devices:** `*whopsdk.AdGroupDevicesBody` — Device platforms and operating systems to target.
    
</dd>
</dl>

<dl>
<dd>

**dynamicCreative:** `*bool` — Let the ad platform automatically mix and match this ad group's creatives and copy to find the best-performing combinations. Set at creation; can't be changed afterward.
    
</dd>
</dl>

<dl>
<dd>

**endsAt:** `*string` — When the ad group stops delivering, as an ISO 8601 timestamp. Omit to run until paused.
    
</dd>
</dl>

<dl>
<dd>

**frequencyCap:** `*whopsdk.CreateAdGroupsRequestFrequencyCap` — Cap on how often one person sees ads from this ad group. Only available on campaigns with the `awareness` objective.
    
</dd>
</dl>

<dl>
<dd>

**languages:** `[]string` — Languages to target, as ISO 639 codes such as `en` or `es`. Empty or omitted targets all languages.
    
</dd>
</dl>

<dl>
<dd>

**messageApps:** `[]*whopsdk.CreateAdGroupsRequestMessageAppsItem` — Apps the conversation opens in. Required when setting `conversion_location` to `messaging`, and rejected unless the ad group's conversion location is `messaging`.
    
</dd>
</dl>

<dl>
<dd>

**minimumDailySpend:** `*float64` — Minimum the ad group tries to spend each day.
    
</dd>
</dl>

<dl>
<dd>

**optimizationGoal:** `*whopsdk.CreateAdGroupsRequestOptimizationGoal` — The result the ad group's delivery is optimized to get the most of.
    
</dd>
</dl>

<dl>
<dd>

**placements:** `*whopsdk.CreateAdGroupsRequestPlacements` 

`automatic` to let the ad platform choose placements, or the list of platforms and positions to target. Omit a platform's positions to target all of them.

Valid positions per platform:

- `facebook`: `feed`, `right_hand_column`, `marketplace`, `search`, `profile_feed`, `notification`, `story`, `instream_video`, `facebook_reels`, `facebook_reels_overlay`, `biz_disco_feed`
- `instagram`: `stream`, `story`, `explore`, `explore_home`, `reels`, `profile_feed`, `profile_reels`, `ig_search`
- `messenger`: `story`
- `audience_network`: `classic`, `rewarded_video`
- `threads`: `threads_stream`
- `whatsapp`: `status`
    
</dd>
</dl>

<dl>
<dd>

**regions:** `*whopsdk.AdGroupRegionsBody` — Locations to target and exclude.
    
</dd>
</dl>

<dl>
<dd>

**startsAt:** `*string` — When the ad group starts delivering, as an ISO 8601 timestamp. Omit to start as soon as it's active.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.CreateAdGroupsRequestStatus` — Initial status (default: `active`).
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The display name of the ad group.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AdGroups.EstimateReach(request) -> *whopsdk.ReachEstimate</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Estimates how many people a draft targeting spec can reach, before an ad group is created. The body takes the same targeting fields as creating an ad group — `regions`, `demographics`, `detailed_targeting`, `audiences`, `languages`, and `devices` — and nothing is persisted.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.EstimateReachAdGroupsRequest{
    Platform: whopsdk.EstimateReachAdGroupsRequestPlatformMeta,
}
client.AdGroups.EstimateReach(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — Account to estimate on behalf of. Defaults to the authenticated account.
    
</dd>
</dl>

<dl>
<dd>

**audiences:** `*whopsdk.AdGroupAudiencesBody` — Saved audiences to deliver to or exclude. Can't be combined with demographics.automatic.
    
</dd>
</dl>

<dl>
<dd>

**demographics:** `*whopsdk.AdGroupDemographicsBody` — Age, gender, and automatic-audience targeting.
    
</dd>
</dl>

<dl>
<dd>

**detailedTargeting:** `*whopsdk.AdGroupDetailedTargetingBody` — Interest, behavior, and demographic targeting, using categories from the ad platform's targeting taxonomy. At most 100 entries per section.
    
</dd>
</dl>

<dl>
<dd>

**devices:** `*whopsdk.AdGroupDevicesBody` — Device platforms and operating systems to target.
    
</dd>
</dl>

<dl>
<dd>

**languages:** `[]string` — Languages to target, as ISO 639 codes such as `en` or `es`. Empty or omitted targets all languages.
    
</dd>
</dl>

<dl>
<dd>

**platform:** `*whopsdk.EstimateReachAdGroupsRequestPlatform` — The ad network the estimate runs on.
    
</dd>
</dl>

<dl>
<dd>

**regions:** `*whopsdk.AdGroupRegionsBody` — Locations to target and exclude.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AdGroups.SearchTargetingOptions() -> *whopsdk.SearchTargetingOptionsAdGroupsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Searches the ad platform's targeting taxonomy for options to target an ad group with. Each result comes back in the exact shape the ad-group body accepts for its `type`, so it can be used in `detailed_targeting`, `regions`, or `languages` as-is. A blank `query` browses the small fixed lists (behaviors, browse demographic categories, languages); interests, work employers, job titles, schools, majors, and locations need a search term.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.SearchTargetingOptionsAdGroupsRequest{
    Platform: whopsdk.SearchTargetingOptionsAdGroupsRequestPlatformMeta,
}
client.AdGroups.SearchTargetingOptions(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — Account to search on behalf of. Defaults to the authenticated account.
    
</dd>
</dl>

<dl>
<dd>

**platform:** `*whopsdk.SearchTargetingOptionsAdGroupsRequestPlatform` — The ad network whose targeting taxonomy to search.
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` — The search term. Blank browses the fixed lists; interests, work employers, job titles, schools, majors, and locations return nothing without one.
    
</dd>
</dl>

<dl>
<dd>

**types:** `*whopsdk.SearchTargetingOptionsAdGroupsRequestTypesItem` — Kinds of targeting options to search. Defaults to all of them.
    
</dd>
</dl>

<dl>
<dd>

**locationTypes:** `*whopsdk.SearchTargetingOptionsAdGroupsRequestLocationTypesItem` — Narrow location results to these kinds of places. Only applies when `types` includes `locations`.
    
</dd>
</dl>

<dl>
<dd>

**country:** `*string` — Narrow location results to one country, as an ISO 3166-1 code such as `US`. Only applies when `types` includes `locations`.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of results per requested type.
    
</dd>
</dl>

<dl>
<dd>

**specialAdCategories:** `*whopsdk.SearchTargetingOptionsAdGroupsRequestSpecialAdCategoriesItem` — The campaign's declared special ad categories. Under `housing`, `employment`, or `financial_products` the ad platform allows interests only, drawn from a short approved list, so results are narrowed to what such a campaign can launch with and other kinds return nothing. Blank `query` browses that approved list instead of the usual fixed lists.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AdGroups.Retrieve(ID) -> *whopsdk.AdGroup</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a single ad group.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveAdGroupsRequest{
    ID: "id",
}
client.AdGroups.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ad group ID.
    
</dd>
</dl>

<dl>
<dd>

**statsFrom:** `*string` — Start of the stats window.
    
</dd>
</dl>

<dl>
<dd>

**statsTo:** `*string` — End of the stats window.
    
</dd>
</dl>

<dl>
<dd>

**timeZone:** `*string` — IANA timezone the stats window is interpreted in. Defaults to UTC.
    
</dd>
</dl>

<dl>
<dd>

**attributionModel:** `*whopsdk.RetrieveAdGroupsRequestAttributionModel` — Attribution model the conversion stats count under (defaults to last_touch). Under both models a journey with any whop ad touch attributes to whop; the model picks which whop touch credits the entity and which non-whop source wins otherwise.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AdGroups.Delete(ID) -> *whopsdk.DeleteAdGroupsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes an ad group.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteAdGroupsRequest{
    ID: "id",
}
client.AdGroups.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ad group ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AdGroups.Update(ID, request) -> *whopsdk.AdGroup</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an ad group's editable fields. Only the keys you send are changed.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateAdGroupsRequest{
    ID: "id",
}
client.AdGroups.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ad group ID.
    
</dd>
</dl>

<dl>
<dd>

**audiences:** `*whopsdk.AdGroupAudiencesBody` — Saved audiences to deliver to or exclude. Can't be combined with demographics.automatic.
    
</dd>
</dl>

<dl>
<dd>

**bidType:** `*whopsdk.UpdateAdGroupsRequestBidType` — How delivery bids are set in the ad auction. Target-based strategies use `desired_cost_per_result`.
    
</dd>
</dl>

<dl>
<dd>

**budgetAmount:** `*float64` — This ad group's budget, in the ad account's currency. Omit when the budget is set on the campaign instead.
    
</dd>
</dl>

<dl>
<dd>

**budgetType:** `*whopsdk.UpdateAdGroupsRequestBudgetType` — Whether budget_amount is spent per day (`daily`) or over the ad group's full run (`lifetime`).
    
</dd>
</dl>

<dl>
<dd>

**conversionEvent:** `*whopsdk.ConversionEvent` 
    
</dd>
</dl>

<dl>
<dd>

**conversionLocation:** `*whopsdk.UpdateAdGroupsRequestConversionLocation` — Where the outcome being optimized for occurs, such as a website visit, social-profile visit, messaging conversation, ad interaction, or lead-form submission. The lead form itself is set on the ad.
    
</dd>
</dl>

<dl>
<dd>

**demographics:** `*whopsdk.AdGroupDemographicsBody` — Age, gender, and automatic-audience targeting.
    
</dd>
</dl>

<dl>
<dd>

**desiredCostPerResult:** `*float64` — Cost per result to aim for (`average_target`) or never exceed (`maximum_target`).
    
</dd>
</dl>

<dl>
<dd>

**detailedTargeting:** `*whopsdk.AdGroupDetailedTargetingBody` — Interest, behavior, and demographic targeting, using categories from the ad platform's targeting taxonomy. Entries across interests, behaviors, and demographics are OR'd together (anyone matching any entry is reached), matching Ads Manager's detailed-targeting box. At most 100 entries per section. Can't be combined with demographics.automatic, and unavailable to campaigns with special_ad_categories. Send the complete intended state — a section you omit is cleared.
    
</dd>
</dl>

<dl>
<dd>

**devices:** `*whopsdk.AdGroupDevicesBody` — Device platforms and operating systems to target.
    
</dd>
</dl>

<dl>
<dd>

**endsAt:** `*string` — When the ad group stops delivering, as an ISO 8601 timestamp. Omit to run until paused.
    
</dd>
</dl>

<dl>
<dd>

**frequencyCap:** `*whopsdk.UpdateAdGroupsRequestFrequencyCap` — Cap on how often one person sees ads from this ad group. Only available on campaigns with the `awareness` objective.
    
</dd>
</dl>

<dl>
<dd>

**languages:** `[]string` — Languages to target, as ISO 639 codes such as `en` or `es`. Empty or omitted targets all languages.
    
</dd>
</dl>

<dl>
<dd>

**messageApps:** `[]*whopsdk.UpdateAdGroupsRequestMessageAppsItem` — Apps the conversation opens in. Required when setting `conversion_location` to `messaging`, and rejected unless the ad group's conversion location is `messaging`.
    
</dd>
</dl>

<dl>
<dd>

**minimumDailySpend:** `*float64` — Minimum the ad group tries to spend each day.
    
</dd>
</dl>

<dl>
<dd>

**optimizationGoal:** `*whopsdk.UpdateAdGroupsRequestOptimizationGoal` — The result the ad group's delivery is optimized to get the most of.
    
</dd>
</dl>

<dl>
<dd>

**placements:** `*whopsdk.UpdateAdGroupsRequestPlacements` 

`automatic` to let the ad platform choose placements, or the list of platforms and positions to target. Omit a platform's positions to target all of them.

Valid positions per platform:

- `facebook`: `feed`, `right_hand_column`, `marketplace`, `search`, `profile_feed`, `notification`, `story`, `instream_video`, `facebook_reels`, `facebook_reels_overlay`, `biz_disco_feed`
- `instagram`: `stream`, `story`, `explore`, `explore_home`, `reels`, `profile_feed`, `profile_reels`, `ig_search`
- `messenger`: `story`
- `audience_network`: `classic`, `rewarded_video`
- `threads`: `threads_stream`
- `whatsapp`: `status`
    
</dd>
</dl>

<dl>
<dd>

**regions:** `*whopsdk.AdGroupRegionsBody` — Locations to target and exclude.
    
</dd>
</dl>

<dl>
<dd>

**startsAt:** `*string` — When the ad group starts delivering, as an ISO 8601 timestamp. Omit to start as soon as it's active.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.UpdateAdGroupsRequestStatus` — Initial status (default: `active`).
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The display name of the ad group.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AdGroups.Duplicate(ID, request) -> *whopsdk.DuplicateAdGroupsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates copies of the ad group in `duplicating` status and returns them — into its own campaign, or into target_ad_campaign_id (which must belong to the same account and be compatible with the ad group's targeting and goals); each copy transitions to its final status (matching the source's active/paused state) once duplication completes. Poll each returned ad group until it leaves `duplicating` — a copy that could not be completed is deleted and returns 404.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DuplicateAdGroupsRequest{
    ID: "id",
}
client.AdGroups.Duplicate(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ad group ID.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — Number of copies to create (1-10). Defaults to 1.
    
</dd>
</dl>

<dl>
<dd>

**preserveEngagement:** `*bool` — Whether the copied ads keep the original posts' engagement (likes, comments, shares). Defaults to false.
    
</dd>
</dl>

<dl>
<dd>

**targetAdCampaignID:** `*string` — Campaign to duplicate into. Defaults to the ad group's own campaign.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AdGroups.Pause(ID) -> *whopsdk.AdGroup</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Pauses delivery of an ad group.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.PauseAdGroupsRequest{
    ID: "id",
}
client.AdGroups.Pause(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ad group ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AdGroups.Unpause(ID) -> *whopsdk.AdGroup</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Resumes delivery of a paused ad group.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UnpauseAdGroupsRequest{
    ID: "id",
}
client.AdGroups.Unpause(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ad group ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## AdReports
<details><summary><code>client.AdReports.Retrieve() -> *whopsdk.AdReport</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Performance report for a company, ad campaigns, ad groups, or ads. Always returns aggregate `summary` totals summed across the scope. Set `granularity` to additionally get a time series, or set `breakdown` (`campaign`/`ad_group`/`ad`) to additionally get per-entity rows inside the requested scope. Exactly one of `companyId`, `adCampaignIds`, `adGroupIds`, or `adIds` must be provided.

Required permissions:
 - `ad_campaign:stats:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveAdReportsRequest{
    CompanyID: whopsdk.String(
        "biz_xxxxxxxxxxxxxx",
    ),
    From: whopsdk.MustParseDateTime(
        "2023-12-01T05:00:00Z",
    ),
    To: whopsdk.MustParseDateTime(
        "2023-12-01T05:00:00Z",
    ),
}
client.AdReports.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**adCampaignIDs:** `*string` — Scope the report to these ad campaigns (max 100); stats are summed across them. Mutually exclusive with `companyId`, `adGroupIds`, and `adIds`.
    
</dd>
</dl>

<dl>
<dd>

**adGroupIDs:** `*string` — Scope the report to these ad groups (max 100); stats are summed across them. Mutually exclusive with `companyId`, `adCampaignIds`, and `adIds`.
    
</dd>
</dl>

<dl>
<dd>

**adIDs:** `*string` — Scope the report to these ads (max 100); stats are summed across them. Mutually exclusive with `companyId`, `adCampaignIds`, and `adGroupIds`.
    
</dd>
</dl>

<dl>
<dd>

**breakdown:** `*whopsdk.AdReportBreakdownLevels` 
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `*string` — The unique identifier of a company. Mutually exclusive with `adCampaignIds`, `adGroupIds`, and `adIds`. Use with `breakdown` to fan out across every campaign, ad group, or ad in the company without paging.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*string` — ISO 4217 currency code to report `spend` in. Defaults to the company's ads reporting currency.
    
</dd>
</dl>

<dl>
<dd>

**from:** `time.Time` — Inclusive start of the reporting window.
    
</dd>
</dl>

<dl>
<dd>

**granularity:** `*whopsdk.Granularities` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `time.Time` — Inclusive end of the reporting window.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ads
<details><summary><code>client.Ads.List() -> *whopsdk.ListAdsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the ads for an account, with stats over the requested window.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListAdsRequest{
    AdCampaignIDs: []*string{
        whopsdk.String(
            "adcamp_xxxxxxxxxxxxxx",
        ),
    },
    AdGroupIDs: []*string{
        whopsdk.String(
            "adgrp_xxxxxxxxxxxxxx",
        ),
    },
}
client.Ads.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The account the ads belong to. Defaults to the account-scoped key's own account.
    
</dd>
</dl>

<dl>
<dd>

**adCampaignID:** `*string` — Only return ads in this ad campaign.
    
</dd>
</dl>

<dl>
<dd>

**adCampaignIDs:** `*string` — Only return ads in these ad campaigns (max 100). Repeat the parameter for each id (ad_campaign_ids=a&ad_campaign_ids=b).
    
</dd>
</dl>

<dl>
<dd>

**adGroupID:** `*string` — Only return ads in this ad group.
    
</dd>
</dl>

<dl>
<dd>

**adGroupIDs:** `*string` — Only return ads in these ad groups (max 100). Repeat the parameter for each id (ad_group_ids=a&ad_group_ids=b).
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.ListAdsRequestStatus` — Only return ads with this status.
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` — Filter ads by a title or ID substring.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListAdsRequestOrder` — The field to sort by. Defaults to created_at. Stat columns (spend, impressions, …) rank over the stats_from/stats_to window across the whole list, not just the current page. results, cost_per_result and return_on_ad_spend rank by the same Whop pixel-attributed values the response reports.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListAdsRequestDirection` — The sort direction. Defaults to desc.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Only return ads created before this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Only return ads created after this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**statsFrom:** `*string` — Start of the stats window. Defaults to all-time.
    
</dd>
</dl>

<dl>
<dd>

**statsTo:** `*string` — End of the stats window. Defaults to now.
    
</dd>
</dl>

<dl>
<dd>

**timeZone:** `*string` — IANA timezone (e.g. America/New_York) the stats window is interpreted in. Bare stats_from/stats_to dates resolve to day boundaries on this clock. Defaults to UTC.
    
</dd>
</dl>

<dl>
<dd>

**attributionModel:** `*whopsdk.ListAdsRequestAttributionModel` — Attribution model the conversion stats count under (defaults to last_touch). Under both models a journey with any whop ad touch attributes to whop; the model picks which whop touch credits the entity and which non-whop source wins otherwise.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of ads to return.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor to fetch the page after (from page_info.end_cursor).
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — The number of ads to return from the end of the range.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Cursor to fetch the page before (from page_info.start_cursor).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ads.Create(request) -> *whopsdk.Ad</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an ad in an ad group.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateAdsRequest{}
client.Ads.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**adGroup:** `map[string]any` — An inline ad group to create (same shape as POST /ad_groups, including ad_campaign_id). Creates the ad group and the ad together. Provide this OR ad_group_id.
    
</dd>
</dl>

<dl>
<dd>

**adGroupID:** `*string` — The existing ad group to create the ad in. Provide this OR ad_group, not both.
    
</dd>
</dl>

<dl>
<dd>

**callToAction:** `*whopsdk.CreateAdsRequestCallToAction` — The call-to-action button shown on the ad.
    
</dd>
</dl>

<dl>
<dd>

**creatives:** `[]*whopsdk.CreateAdsRequestCreativesItem` — The ad's creative assets. Each entry is an uploaded file id with an optional format; omit format for the original asset. Two or more entries with no format become a carousel (2-10 attachments), in order, sharing the ad's copy.
    
</dd>
</dl>

<dl>
<dd>

**descriptions:** `[]string` — The description variants shown on the ad.
    
</dd>
</dl>

<dl>
<dd>

**existingPostID:** `*string` — Promote a post you already published instead of uploading creatives — a Facebook post or Instagram media id. Mutually exclusive with creatives. Pair with post_source.
    
</dd>
</dl>

<dl>
<dd>

**headlines:** `[]string` — The headline variants shown on the ad.
    
</dd>
</dl>

<dl>
<dd>

**leadForm:** `*whopsdk.CreateAdsRequestLeadForm` — Instant lead form for the ad. Only allowed when the ad group's conversion_location is an instant-form destination (instant_forms, instant_forms_and_messenger, website_and_instant_forms). Mutually exclusive with lead_form_id.
    
</dd>
</dl>

<dl>
<dd>

**leadFormID:** `*string` — Use an existing instant form instead of creating one — the form's platform ID, from a form already on the ad's Facebook page. Only allowed when the ad group's conversion_location is an instant-form destination. Mutually exclusive with lead_form.
    
</dd>
</dl>

<dl>
<dd>

**messagingConfig:** `*whopsdk.CreateAdsRequestMessagingConfig` — Click-to-message welcome copy: the greeting (message) and the ice-breaker prompt (keyword).
    
</dd>
</dl>

<dl>
<dd>

**multiAdvertiserAds:** `*bool` — Whether the ad can appear alongside other advertisers' ads in the same unit. Defaults to true.
    
</dd>
</dl>

<dl>
<dd>

**postSource:** `*whopsdk.CreateAdsRequestPostSource` — Identifies the network that owns `existing_post_id`. The source is inferred from the ID shape when omitted.
    
</dd>
</dl>

<dl>
<dd>

**primaryTexts:** `[]string` — The primary text variants shown in the ad body.
    
</dd>
</dl>

<dl>
<dd>

**socialAccounts:** `[]*whopsdk.CreateAdsRequestSocialAccountsItem` — The social accounts the ad runs under — a connected Facebook page and, optionally, an Instagram profile.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The display name of the ad.
    
</dd>
</dl>

<dl>
<dd>

**url:** `*string` — The URL the ad links to. Query parameters are merged into url_parameters, so the stored URL is always bare.
    
</dd>
</dl>

<dl>
<dd>

**urlParameters:** `map[string]any` — Query parameters to append to the destination URL, keyed by parameter name. Merged with any query string on `url`. Whop adds its own click-attribution parameters; those are reserved and rejected if you set them (utm_meta_ad_id, utm_meta_adset_id, utm_meta_campaign_id, utm_source, utm_placement, utm_medium, utm_content, utm_adset, utm_whop, wacid, wasid, waid, tw_source, tw_adid).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ads.Retrieve(ID) -> *whopsdk.Ad</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a single ad with stats over the requested window.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveAdsRequest{
    ID: "id",
}
client.Ads.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ad ID.
    
</dd>
</dl>

<dl>
<dd>

**statsFrom:** `*string` — Start of the stats window.
    
</dd>
</dl>

<dl>
<dd>

**statsTo:** `*string` — End of the stats window.
    
</dd>
</dl>

<dl>
<dd>

**timeZone:** `*string` — IANA timezone the stats window is interpreted in. Defaults to UTC.
    
</dd>
</dl>

<dl>
<dd>

**attributionModel:** `*whopsdk.RetrieveAdsRequestAttributionModel` — Attribution model the conversion stats count under (defaults to last_touch). Under both models a journey with any whop ad touch attributes to whop; the model picks which whop touch credits the entity and which non-whop source wins otherwise.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ads.Delete(ID) -> *whopsdk.DeleteAdsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes an ad.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteAdsRequest{
    ID: "id",
}
client.Ads.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ad ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ads.Update(ID, request) -> *whopsdk.Ad</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an ad's editable fields.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateAdsRequest{
    ID: "id",
}
client.Ads.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ad ID.
    
</dd>
</dl>

<dl>
<dd>

**callToAction:** `*whopsdk.UpdateAdsRequestCallToAction` — The call-to-action button shown on the ad.
    
</dd>
</dl>

<dl>
<dd>

**creatives:** `[]*whopsdk.UpdateAdsRequestCreativesItem` — The ad's creative assets. Each entry is an uploaded file id with an optional format; omit format for the original asset. Replaces a live ad's creative on the platform. Two or more entries with no format replace it with a carousel (2-10 attachments), in order, sharing the ad's copy.
    
</dd>
</dl>

<dl>
<dd>

**descriptions:** `[]string` — The description variants shown on the ad.
    
</dd>
</dl>

<dl>
<dd>

**existingPostID:** `*string` — Promote a post you already published instead of uploading creatives — a Facebook post or Instagram media id. Mutually exclusive with creatives. Pair with post_source.
    
</dd>
</dl>

<dl>
<dd>

**headlines:** `[]string` — The headline variants shown on the ad.
    
</dd>
</dl>

<dl>
<dd>

**leadForm:** `*whopsdk.UpdateAdsRequestLeadForm` — Instant lead form for the ad. Only allowed when the ad group's conversion_location is an instant-form destination (instant_forms, instant_forms_and_messenger, website_and_instant_forms). Mutually exclusive with lead_form_id.
    
</dd>
</dl>

<dl>
<dd>

**leadFormID:** `*string` — Use an existing instant form instead of creating one — the form's platform ID, from a form already on the ad's Facebook page. Only allowed when the ad group's conversion_location is an instant-form destination. Mutually exclusive with lead_form. Replaces a stored lead_form.
    
</dd>
</dl>

<dl>
<dd>

**messagingConfig:** `*whopsdk.UpdateAdsRequestMessagingConfig` — Click-to-message welcome copy: the greeting (message) and the ice-breaker prompt (keyword).
    
</dd>
</dl>

<dl>
<dd>

**multiAdvertiserAds:** `*bool` — Whether the ad can appear alongside other advertisers' ads in the same unit. Defaults to true.
    
</dd>
</dl>

<dl>
<dd>

**postSource:** `*whopsdk.UpdateAdsRequestPostSource` — Identifies the network that owns `existing_post_id`. The source is inferred from the ID shape when omitted.
    
</dd>
</dl>

<dl>
<dd>

**primaryTexts:** `[]string` — The primary text variants shown in the ad body.
    
</dd>
</dl>

<dl>
<dd>

**socialAccounts:** `[]*whopsdk.UpdateAdsRequestSocialAccountsItem` — The social accounts the ad runs under — a connected Facebook page and, optionally, an Instagram profile.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The display name of the ad.
    
</dd>
</dl>

<dl>
<dd>

**url:** `*string` — The URL the ad links to. Query parameters are merged into url_parameters, so the stored URL is always bare.
    
</dd>
</dl>

<dl>
<dd>

**urlParameters:** `map[string]any` — Query parameters to append to the destination URL, keyed by parameter name. Merged with any query string on `url`. Whop adds its own click-attribution parameters; those are reserved and rejected if you set them (utm_meta_ad_id, utm_meta_adset_id, utm_meta_campaign_id, utm_source, utm_placement, utm_medium, utm_content, utm_adset, utm_whop, wacid, wasid, waid, tw_source, tw_adid).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ads.Duplicate(ID, request) -> *whopsdk.DuplicateAdsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Copies the ad into its own ad group, or into target_ad_group_id (which must belong to the same account and be compatible with the ad). Copies keep the source ad's active/paused state.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DuplicateAdsRequest{
    ID: "id",
}
client.Ads.Duplicate(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ad ID.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — Number of copies to create (1-10). Defaults to 1.
    
</dd>
</dl>

<dl>
<dd>

**preserveEngagement:** `*bool` — Whether the copies keep the original post's engagement (likes, comments, shares). Defaults to false.
    
</dd>
</dl>

<dl>
<dd>

**targetAdGroupID:** `*string` — Ad group to duplicate into. Defaults to the ad's own ad group.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ads.Pause(ID) -> *whopsdk.Ad</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Pauses an active ad.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.PauseAdsRequest{
    ID: "id",
}
client.Ads.Pause(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ad ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ads.Unpause(ID) -> *whopsdk.Ad</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Resumes a paused ad.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UnpauseAdsRequest{
    ID: "id",
}
client.Ads.Unpause(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ad ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Affiliates
<details><summary><code>client.Affiliates.List() -> *whopsdk.ListAffiliatesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of affiliates for the actor in context, with optional filtering by status, search, and sorting.

Required permissions:
 - `affiliate:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListAffiliatesRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    CompanyID: "biz_xxxxxxxxxxxxxx",
}
client.Affiliates.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `string` — The unique identifier of the company to list affiliates for.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.Direction` 
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.AffiliatesSortableColumns` 
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` — Search affiliates by username.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.Status` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Affiliates.Create(request) -> *whopsdk.Affiliate</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates or finds an affiliate for a company and user.

Required permissions:
 - `affiliate:create`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateAffiliatesRequest{
    CompanyID: "biz_xxxxxxxxxxxxxx",
    UserIdentifier: "user_identifier",
}
client.Affiliates.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `string` — The ID of the company to create the affiliate for.
    
</dd>
</dl>

<dl>
<dd>

**userIdentifier:** `string` — The user identifier (username, email, user ID, or Discord ID).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Affiliates.Retrieve(ID) -> *whopsdk.Affiliate</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing affiliate.

Required permissions:
 - `affiliate:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveAffiliatesRequest{
    ID: "aff_xxxxxxxxxxxxxx",
}
client.Affiliates.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the affiliate.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Affiliates.Archive(ID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Archives an existing Affiliate

Required permissions:
 - `affiliate:update`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ArchiveAffiliatesRequest{
    ID: "aff_xxxxxxxxxxxxxx",
}
client.Affiliates.Archive(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The internal ID of the affiliate to archive.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Affiliates.Unarchive(ID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Unarchives an existing Affiliate

Required permissions:
 - `affiliate:update`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UnarchiveAffiliatesRequest{
    ID: "aff_xxxxxxxxxxxxxx",
}
client.Affiliates.Unarchive(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The internal ID of the affiliate to archive.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## AiChats
<details><summary><code>client.AiChats.List() -> *whopsdk.ListAiChatsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of AI chat threads for the current authenticated user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListAiChatsRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
}
client.AiChats.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**onlyActiveCrons:** `*bool` — When true, returns only chats with an active cron schedule
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AiChats.Create(request) -> *whopsdk.AiChat</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new AI chat thread and send the first message to the AI agent.

Required permissions:
 - `ai_chat:create`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateAiChatsRequest{
    MessageText: "message_text",
}
client.AiChats.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**currentCompanyID:** `*string` — The unique identifier of the company to set as context for the AI chat (e.g., "biz_XXXXX").
    
</dd>
</dl>

<dl>
<dd>

**messageAttachments:** `[]*whopsdk.CreateAiChatsRequestMessageAttachmentsItem` — A list of previously uploaded file attachments to include with the first message.
    
</dd>
</dl>

<dl>
<dd>

**messageSource:** `*whopsdk.AiChatMessageSourceTypes` — The source of the message.
    
</dd>
</dl>

<dl>
<dd>

**messageText:** `string` — The text content of the first message to send to the AI agent.
    
</dd>
</dl>

<dl>
<dd>

**suggestionType:** `*string` — The type of suggestion prompt that was clicked, when message_source is 'suggestion'.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — An optional display title for the AI chat thread (e.g., "Help with billing").
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AiChats.Retrieve(ID) -> *whopsdk.AiChat</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing AI chat.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveAiChatsRequest{
    ID: "aich_xxxxxxxxxxxxx",
}
client.AiChats.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the AI chat to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AiChats.Delete(ID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete an AI chat thread so it no longer appears in the user's chat list.

Required permissions:
 - `ai_chat:delete`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteAiChatsRequest{
    ID: "aich_xxxxxxxxxxxxx",
}
client.AiChats.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the AI chat to delete (e.g., "ai_chat_XXXXX").
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AiChats.Update(ID, request) -> *whopsdk.AiChat</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update an AI chat's title, notification preferences, or associated company context.

Required permissions:
 - `ai_chat:update`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateAiChatsRequest{
    ID: "aich_xxxxxxxxxxxxx",
}
client.AiChats.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the AI chat to update (e.g., "ai_chat_XXXXX").
    
</dd>
</dl>

<dl>
<dd>

**currentCompanyID:** `*string` — The unique identifier of the company to set as context for the AI chat (e.g., "biz_XXXXX").
    
</dd>
</dl>

<dl>
<dd>

**notificationPreference:** `*whopsdk.AiChatNotificationPreferences` — The notification preference for the AI chat.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The new display title for the AI chat thread (e.g., "Help with billing").
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## API Keys
<details><summary><code>client.APIKeys.List() -> *whopsdk.ListAPIKeysResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the API keys of an account or app, newest first. Responses never include the full secret — only its obfuscated form.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListAPIKeysRequest{
    ResourceID: "resource_id",
    ResourceType: whopsdk.ListAPIKeysRequestResourceTypeAccount,
}
client.APIKeys.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**resourceID:** `string` — The account (`biz_`) or app (`app_`) tag to list API keys for.
    
</dd>
</dl>

<dl>
<dd>

**resourceType:** `*whopsdk.ListAPIKeysRequestResourceType` — The type of resource that owns the API keys.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*whopsdk.ListAPIKeysRequestCreatedBefore` — Only return API keys created before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*whopsdk.ListAPIKeysRequestCreatedAfter` — Only return API keys created after this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of API keys to return (default 20, max 100).
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor; returns API keys after this position.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — The number of API keys to return from the end of the range.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — A cursor; returns API keys before this position.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListAPIKeysRequestOrder` — The field to sort API keys by.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListAPIKeysRequestDirection` — Sort direction.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.APIKeys.Create(request) -> *whopsdk.APIKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an API key for an account or app. The response is the only place the full `secret_key` is returned — store it immediately. Requires a user session; API keys cannot manage API keys.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateAPIKeysRequest{
    Name: "Shine Time Booking (production)",
    Permissions: &whopsdk.CreateAPIKeysRequestPermissions{},
    ResourceID: "biz_xxxxxxxxxxxxxx",
    ResourceType: whopsdk.CreateAPIKeysRequestResourceTypeAccount,
}
client.APIKeys.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiVersionDate:** `*whopsdk.CreateAPIKeysRequestAPIVersionDate` — Dated API version used when requests authenticated with this key omit the `Api-Version-Date` header. New keys default to the latest version.
    
</dd>
</dl>

<dl>
<dd>

**expiresAt:** `*string` — When the API key should stop working, as an ISO 8601 timestamp. Omit (or pass `null` on update) for a key that never expires.
    
</dd>
</dl>

<dl>
<dd>

**ipAllowlist:** `[]string` — IPv4/IPv6 CIDR ranges allowed to use this key, for example `["203.0.113.0/24"]`. Empty or `null` allows any IP.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — A human-readable name for the API key, such as 'Production API Key'.
    
</dd>
</dl>

<dl>
<dd>

**permissions:** `*whopsdk.CreateAPIKeysRequestPermissions` — The permissions policy for the API key: explicit permission statements, or a system role to inherit from. Statements without a `resources` array default to the owning account (Account API keys) or every key-addressable resource (App API keys).
    
</dd>
</dl>

<dl>
<dd>

**resourceID:** `string` — The account (`biz_`) or app (`app_`) tag to create the API key for.
    
</dd>
</dl>

<dl>
<dd>

**resourceType:** `*whopsdk.CreateAPIKeysRequestResourceType` — The type of resource that will own this API key.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.APIKeys.ListPermissions() -> *whopsdk.ListPermissionsAPIKeysResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the catalog of permission actions that can be granted to users, apps, and API keys — the source for the dashboard's permission pickers. Small and returned in full on one page.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.APIKeys.ListPermissions(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.APIKeys.Retrieve(ID) -> *whopsdk.APIKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves an API key with its effective permission grants. The full secret is never returned — rotate the key if it was lost.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveAPIKeysRequest{
    ID: "id",
}
client.APIKeys.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — API key ID, prefixed `apik_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.APIKeys.Delete(ID) -> *whopsdk.DeleteAPIKeysResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Permanently revokes an API key; requests using its secret stop authenticating immediately. Default and agent-backend keys cannot be deleted.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteAPIKeysRequest{
    ID: "id",
}
client.APIKeys.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — API key ID, prefixed `apik_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.APIKeys.Update(ID, request) -> *whopsdk.APIKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an API key's name, permissions, API version, expiration, or IP allowlist. Fields that are omitted keep their current value; default keys cannot be modified.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateAPIKeysRequest{
    ID: "id",
}
client.APIKeys.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — API key ID, prefixed `apik_`.
    
</dd>
</dl>

<dl>
<dd>

**apiVersionDate:** `*whopsdk.UpdateAPIKeysRequestAPIVersionDate` — Dated API version used when requests authenticated with this key omit the `Api-Version-Date` header. New keys default to the latest version.
    
</dd>
</dl>

<dl>
<dd>

**expiresAt:** `*string` — When the API key should stop working, as an ISO 8601 timestamp. Omit (or pass `null` on update) for a key that never expires.
    
</dd>
</dl>

<dl>
<dd>

**ipAllowlist:** `[]string` — IPv4/IPv6 CIDR ranges allowed to use this key, for example `["203.0.113.0/24"]`. Empty or `null` allows any IP.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — A new human-readable name for the API key.
    
</dd>
</dl>

<dl>
<dd>

**permissions:** `*whopsdk.UpdateAPIKeysRequestPermissions` — The permissions policy for the API key: explicit permission statements, or a system role to inherit from. Statements without a `resources` array default to the owning account (Account API keys) or every key-addressable resource (App API keys).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.APIKeys.Rotate(ID) -> *whopsdk.APIKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Rotates the API key's secret, invalidating the previous secret immediately. The response is the only place the new `secret_key` is returned.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RotateAPIKeysRequest{
    ID: "id",
}
client.APIKeys.Rotate(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — API key ID, prefixed `apik_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## App Builds
<details><summary><code>client.AppBuilds.List() -> *whopsdk.ListAppBuildsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of build artifacts for an app, newest first, with optional platform, status, and creation-date filters.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListAppBuildsRequest{
    AppID: "app_id",
}
client.AppBuilds.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**appID:** `string` — The app to list builds for, prefixed `app_`.
    
</dd>
</dl>

<dl>
<dd>

**platform:** `*whopsdk.ListAppBuildsRequestPlatform` — Filter builds by target platform.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.ListAppBuildsRequestStatus` — Filter builds by review status.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*whopsdk.ListAppBuildsRequestCreatedBefore` — Only return builds created before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*whopsdk.ListAppBuildsRequestCreatedAfter` — Only return builds created after this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of builds to return (default 20, max 100).
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor; returns builds after this position.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — The number of builds to return from the end of the range.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — A cursor; returns builds before this position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AppBuilds.Create(request) -> *whopsdk.AppBuild</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Uploads a new build artifact for an app. Upload the file first (POST /files or a direct upload), then reference it here; iOS and Android take a .zip bundle, web takes a JavaScript file or a .zip archive of the hosted site.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateAppBuildsRequest{
    Attachment: &whopsdk.CreateAppBuildsRequestAttachment{},
    Checksum: "xxxxxxxxxxxxxxx",
    Platform: whopsdk.CreateAppBuildsRequestPlatformIos,
}
client.AppBuilds.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**aiPromptID:** `*string` — The AI prompt that generated this build, if applicable.
    
</dd>
</dl>

<dl>
<dd>

**appID:** `*string` — The app to create the build for, prefixed `app_`. Defaults to the app behind the presented credential.
    
</dd>
</dl>

<dl>
<dd>

**attachment:** `*whopsdk.CreateAppBuildsRequestAttachment` — The uploaded build file: `{ id }` for an existing file or `{ direct_upload_id }` for a completed direct upload.
    
</dd>
</dl>

<dl>
<dd>

**checksum:** `string` — A client-generated checksum of the build file, used to verify file integrity when unpacked.
    
</dd>
</dl>

<dl>
<dd>

**platform:** `*whopsdk.CreateAppBuildsRequestPlatform` — The target platform for the build.
    
</dd>
</dl>

<dl>
<dd>

**sourceAttachment:** `*whopsdk.CreateAppBuildsRequestSourceAttachment` — An optional compressed archive (.zip or .gz) of the source code that produced this build, stored alongside the build so it can be downloaded later. Referenced like `attachment`, and must be a different file.
    
</dd>
</dl>

<dl>
<dd>

**supportedAppViewTypes:** `[]*whopsdk.CreateAppBuildsRequestSupportedAppViewTypesItem` — The view types this build supports. Only list the ones its code implements.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AppBuilds.Retrieve(ID) -> *whopsdk.AppBuild</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing app build.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveAppBuildsRequest{
    ID: "id",
}
client.AppBuilds.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — App build ID, prefixed `abld_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AppBuilds.Promote(ID) -> *whopsdk.AppBuild</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Promotes a draft or approved app build to production so it becomes the active version served to users. Draft builds enter review first.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.PromoteAppBuildsRequest{
    ID: "id",
}
client.AppBuilds.Promote(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — App build ID, prefixed `abld_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Apps
<details><summary><code>client.Apps.List() -> *whopsdk.ListAppsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists apps on the Whop platform: the app store's live apps, or — with `account_id` and developer access to that account — every app the account owns. Requires authentication except for Whop's public app and website discovery lists. Public website discovery includes built official blueprints (verified apps with a product) and built, live community blueprints that Whop recommends.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListAppsRequest{}
client.Apps.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — Only return apps created by this account (`biz_` tag). With developer access to the account this includes its unlisted and hidden apps.
    
</dd>
</dl>

<dl>
<dd>

**appType:** `*whopsdk.ListAppsRequestAppType` — Filter apps by the type of end-user they are built for. Apps of type `website` are left out unless you ask for them by name.
    
</dd>
</dl>

<dl>
<dd>

**viewType:** `*whopsdk.ListAppsRequestViewType` — Only return apps supporting this view type, such as `dashboard` or `hub`.
    
</dd>
</dl>

<dl>
<dd>

**verified:** `*bool` — Only return apps whose Whop verification status matches this value. Omit this filter to include every verification status the caller can see.
    
</dd>
</dl>

<dl>
<dd>

**verifiedAppsOnly:** `*bool` — Legacy compatibility filter. Use `verified` for field equality. `true` returns verified apps; clients pinned before `2026-08-25-2` retain the earlier public website discovery behavior.
    
</dd>
</dl>

<dl>
<dd>

**recommended:** `*bool` — Only return apps Whop recommends (or, with `false`, only those it does not), independently of verification status.
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` — A search string matched against app names.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListAppsRequestOrder` — The field to sort apps by. Defaults to discoverable_at, showing the most recently published apps first. `template_usage` ranks Whop-verified apps first, then apps with a banner image, then by how many apps were created from each app as a template.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListAppsRequestDirection` — Sort direction.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of apps to return (default 20, max 100).
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor; returns apps after this position.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — The number of apps to return from the end of the range.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — A cursor; returns apps before this position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Apps.Create(request) -> *whopsdk.App</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Registers a new app on the Whop developer platform. Apps provide custom experiences that can be added to products.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateAppsRequest{
    Name: "Shine Time Booking",
}
client.Apps.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The account to create the app for (`biz_` tag). Defaults to the account behind the presented credential.
    
</dd>
</dl>

<dl>
<dd>

**appType:** `*whopsdk.CreateAppsRequestAppType` — The type of app to create. Defaults to `b2c_app`.
    
</dd>
</dl>

<dl>
<dd>

**baseURL:** `*string` — The base production URL where the app is hosted, such as `https://myapp.example.com`.
    
</dd>
</dl>

<dl>
<dd>

**icon:** `*whopsdk.CreateAppsRequestIcon` — The icon image for the app in PNG, JPEG, or GIF format, referencing an uploaded file: `{ id }` for an existing attachment or `{ direct_upload_id }` for a new direct upload.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — The display name for the app, shown to users on the app store and product pages.
    
</dd>
</dl>

<dl>
<dd>

**redirectURIs:** `[]string` — The whitelisted OAuth callback URLs that users are redirected to after authorizing the app.
    
</dd>
</dl>

<dl>
<dd>

**route:** `*string` — The subdomain route where the app's hosted web builds are served, such as `myapp` for myapp.whop.site.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Apps.UpdatePermissionsApp(AppID, request) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the permission requirements for an app

Required permissions:
 - `developer:update_app_authorization`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdatePermissionsAppRequest{
    AppID: "app_id",
    RequestedPermissions: []*whopsdk.UpdatePermissionsAppRequestRequestedPermissionsItem{
        &whopsdk.UpdatePermissionsAppRequestRequestedPermissionsItem{
            Action: "action",
            IsRequired: true,
            Justification: "justification",
        },
    },
}
client.Apps.UpdatePermissionsApp(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**appID:** `string` — The ID of the app the permission requirements are being updated for
    
</dd>
</dl>

<dl>
<dd>

**requestedPermissions:** `[]*whopsdk.UpdatePermissionsAppRequestRequestedPermissionsItem` — The permissions that the app will request off of users when a user installs the app.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Apps.Retrieve(ID) -> *whopsdk.App</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves an app by ID, claimed route, or proxy domain id. Credential fields (api_key, default_api_key, secrets) render `null` unless the caller has the corresponding developer permission on the owning account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveAppsRequest{
    ID: "id",
}
client.Apps.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — App ID (prefixed `app_`), the app's claimed route, or its proxy domain id.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Apps.Delete(ID) -> *whopsdk.DeleteAppsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes an app. The app stops resolving within seconds — a website's site stops serving, and any claimed subdomain is reserved for a month before it can be claimed again.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteAppsRequest{
    ID: "id",
}
client.Apps.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — App ID (prefixed `app_`), the app's claimed route, or its proxy domain id.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Apps.Update(ID, request) -> *whopsdk.App</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the settings, metadata, or status of an app. Fields that are omitted keep their current value.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateAppsRequest{
    ID: "id",
}
client.Apps.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — App ID (prefixed `app_`), the app's claimed route, or its proxy domain id.
    
</dd>
</dl>

<dl>
<dd>

**appStoreDescription:** `*string` — The detailed description shown on the app store's in-depth app view page.
    
</dd>
</dl>

<dl>
<dd>

**appType:** `*whopsdk.UpdateAppsRequestAppType` — The type of end-user the app is built for. Cannot be changed on an app whose type is already `website`.
    
</dd>
</dl>

<dl>
<dd>

**baseURL:** `*string` — The base production URL where the app is hosted. Set to `null` to take the app proxy offline.
    
</dd>
</dl>

<dl>
<dd>

**dashboardPath:** `*string` — The URL path for the account dashboard view.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — A short description of the app shown in listings and search results.
    
</dd>
</dl>

<dl>
<dd>

**discoverPath:** `*string` — The URL path for the discover view.
    
</dd>
</dl>

<dl>
<dd>

**experiencePath:** `*string` — The URL path for the member-facing hub view, such as `/experiences/[experienceId]`.
    
</dd>
</dl>

<dl>
<dd>

**icon:** `*whopsdk.UpdateAppsRequestIcon` — The icon image for the app in PNG, JPEG, or GIF format, referencing an uploaded file: `{ id }` for an existing attachment or `{ direct_upload_id }` for a new direct upload.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The display name for the app, shown to users on the app store and product pages.
    
</dd>
</dl>

<dl>
<dd>

**oauthClientType:** `*whopsdk.UpdateAppsRequestOauthClientType` — How the app authenticates at the OAuth token endpoint.
    
</dd>
</dl>

<dl>
<dd>

**openapiPath:** `*string` — The URL path to the app's OpenAPI spec file (requires the ai_chat capability).
    
</dd>
</dl>

<dl>
<dd>

**productionAndroidBuildID:** `*string` — The app build (`abld_` tag) to serve as the Android production build, or `null` to unassign it. Same rules as `production_web_build_id`.
    
</dd>
</dl>

<dl>
<dd>

**productionIosBuildID:** `*string` — The app build (`abld_` tag) to serve as the iOS production build, or `null` to unassign it. Same rules as `production_web_build_id`.
    
</dd>
</dl>

<dl>
<dd>

**productionWebBuildID:** `*string` — The app build (`abld_` tag) to serve as the web production build, or `null` to unassign it. The build must belong to this app, target web, and be in the draft or approved status; a draft build is queued for approval and takes over once approved. Requires the `developer:manage_builds` scope.
    
</dd>
</dl>

<dl>
<dd>

**redirectURIs:** `[]string` — The whitelisted OAuth callback URLs users are redirected to after authorizing the app.
    
</dd>
</dl>

<dl>
<dd>

**requiredScopes:** `[]string` — The OAuth scopes the app requests from users when they install it.
    
</dd>
</dl>

<dl>
<dd>

**route:** `*string` — The subdomain route where the app's hosted web builds are served.
    
</dd>
</dl>

<dl>
<dd>

**secrets:** `map[string]any` — Secrets to add or overwrite on the app, as an object of string values. Keys not included are left untouched; pass null or an empty string as the value to delete a secret. Encrypted at rest and injected into the app's hosted server runtime.
    
</dd>
</dl>

<dl>
<dd>

**skillsPath:** `*string` — The URL path to the app's skills directory (requires the ai_chat capability).
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.UpdateAppsRequestStatus` — Controls whether the app is published on Whop discovery or accessible only through its direct link. Publishing requires a name, icon, and description.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Apps.Deploy(ID, request) -> *whopsdk.AppDeployment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Builds the app's current source and ships it. Returns the run it started, so the caller can render progress from this response and then follow it on the app's `deployment` field. Only one deployment runs per app at a time — calling this while one is in flight reports that run rather than starting a second, and calling it with nothing to publish reports that instead of starting one.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeployAppsRequest{
    ID: "id",
}
client.Apps.Deploy(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The app to deploy, prefixed `app_`.
    
</dd>
</dl>

<dl>
<dd>

**draft:** `*bool` — Upload the build without making it live. Defaults to `false`, which deploys and promotes in one step.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Apps.Logs(ID) -> *whopsdk.LogsAppsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists a hosted app's server runtime logs, most recent first: console output, uncaught exceptions, and failed-request summaries captured on whop.site hosting. Logs are retained for 7 days.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.LogsAppsRequest{
    ID: "id",
}
client.Apps.Logs(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the app, which will look like app_*************.
    
</dd>
</dl>

<dl>
<dd>

**appBuildID:** `*string` — Only return logs from this build.
    
</dd>
</dl>

<dl>
<dd>

**level:** `*whopsdk.LogsAppsRequestLevel` — Only return console lines of this level.
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` — Only return logs whose message contains this text (case-insensitive).
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — Start of the time window as an ISO 8601 timestamp. Defaults to 7 days before created_before.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — End of the time window as an ISO 8601 timestamp. Defaults to now.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of log lines to return (max 500).
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor for fetching logs after a previous page.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — A cursor for fetching logs before a later page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Apps.UpdatePermissions(ID, request) -> *whopsdk.App</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Replaces the set of permissions the app requests from users when they install it. Requires a user session: the `developer:update_app_authorization` scope cannot be delegated to API keys. Sensitive permissions require step-up verification.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdatePermissionsAppsRequest{
    ID: "id",
    RequestedPermissions: []*whopsdk.UpdatePermissionsAppsRequestRequestedPermissionsItem{
        &whopsdk.UpdatePermissionsAppsRequestRequestedPermissionsItem{
            Action: "company:basic:read",
            IsRequired: true,
            Justification: "Reads basic account info to render the dashboard home.",
        },
    },
}
client.Apps.UpdatePermissions(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — App ID, prefixed `app_`.
    
</dd>
</dl>

<dl>
<dd>

**requestedPermissions:** `[]*whopsdk.UpdatePermissionsAppsRequestRequestedPermissionsItem` — The full set of permissions the app requests on install; permissions not listed are removed.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Audiences
<details><summary><code>client.Audiences.List() -> *whopsdk.ListAudiencesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists uploaded customer-list audiences for an account. Pass `audience_id` to return a specific audience.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListAudiencesRequest{
    AccountID: "account_id",
}
client.Audiences.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string` — Account ID, prefixed `biz_`.
    
</dd>
</dl>

<dl>
<dd>

**audienceID:** `*string` — Audience ID, prefixed `adaud_`, used to filter the response to one audience.
    
</dd>
</dl>

<dl>
<dd>

**audienceType:** `*whopsdk.ListAudiencesRequestAudienceType` — Filter by audience type: `custom` (uploaded lists) or `lookalike`.
    
</dd>
</dl>

<dl>
<dd>

**sourceType:** `*whopsdk.ListAudiencesRequestSourceType` — Filter by member source: `csv_upload` (uploaded lists) or `people_filter` (automatic audiences built from saved People filters).
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Number of audiences to return. Defaults to 20; maximum 100.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor for the next page of audiences.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Audiences.Create(request) -> *whopsdk.CreateAudiencesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an audience. Default (`audience_type` omitted or `custom`): creates one audience from an uploaded customer identity CSV file (`name`, `column_mapping`, and `file_id` required) and starts processing it; responds with the audience object. With `filters`: creates an audience from saved People filters (`name` required) — membership is built from the account's People data, and `auto_refresh` decides whether it keeps tracking the filters or keeps whoever matched at creation. With `audience_type: lookalike`: creates a ladder of Meta lookalike audiences from an existing ready custom audience (`source_audience_id`, `count`, and `percentage` required) — `count` equal similarity bands slicing the top `percentage`% (3 audiences at 6% = 0–2%, 2–4%, 4–6%), each returned as its own audience in a `{ data: [...] }` envelope.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateAudiencesRequest{
    AccountID: "biz_xxxxxxxxxxxxxx",
}
client.Audiences.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string` — Account ID, prefixed `biz_`.
    
</dd>
</dl>

<dl>
<dd>

**audienceType:** `*whopsdk.CreateAudiencesRequestAudienceType` — What to create. Defaults to `custom` (CSV upload).
    
</dd>
</dl>

<dl>
<dd>

**autoRefresh:** `*bool` — Filter audiences only, and set only at creation. `true` (the default) rebuilds membership from the filters twice a day. `false` keeps whoever matched at creation and never rebuilds.
    
</dd>
</dl>

<dl>
<dd>

**columnMapping:** `*whopsdk.CreateAudiencesRequestColumnMapping` — Custom audiences only. Maps supported identity fields to CSV column headers. Map at least one of `email` or `phone`.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — Lookalikes only. Number of lookalike audiences to create (1–6).
    
</dd>
</dl>

<dl>
<dd>

**fileID:** `*string` — Custom audiences only. The uploaded customer CSV — a file id (`file_...`) returned by `POST /files`.
    
</dd>
</dl>

<dl>
<dd>

**filters:** `map[string]any` — Filter audiences only. The People filters that define membership, keyed exactly as `GET /people` accepts them — for example `{"os": "iOS", "country": "US"}`. Date filters must be rolling windows — `first_seen_within_days` or `last_seen_within_days` — so the audience re-anchors on every refresh; fixed dates such as `first_seen_after` are rejected. Source values are canonical source paths (`whop:<campaign>:<group>:<ad>`, `ext:<platform>:...`, `referrer:<domain>`, `direct`), exact or with a trailing `:*` wildcard.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Audience display name. Required for custom audiences; lookalike names are generated from the source audience.
    
</dd>
</dl>

<dl>
<dd>

**percentage:** `*int` — Lookalikes only. Total similarity reach as a whole percent (1–20), sliced evenly across `count` — must be divisible by `count`.
    
</dd>
</dl>

<dl>
<dd>

**sourceAudienceID:** `*string` — Lookalikes only. The ready custom audience (`adaud_`) to build from; it needs at least 100 matched people.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Audiences.Delete(ID) -> *whopsdk.DeleteAudiencesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes an audience so it is no longer available for targeting.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteAudiencesRequest{
    ID: "id",
}
client.Audiences.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Audience ID, prefixed `adaud_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Audiences.Update(ID, request) -> *whopsdk.Audience</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Renames an audience. For an audience built from People filters that keeps itself up to date, pass `filters` to replace them, which rebuilds membership immediately. Whether an audience auto refreshes is set when it is created.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateAudiencesRequest{
    ID: "id",
}
client.Audiences.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Audience ID, prefixed `adaud_`.
    
</dd>
</dl>

<dl>
<dd>

**filters:** `map[string]any` — Replaces the People filters that define membership. The whole definition is replaced rather than merged, so send every filter you want to keep — a filter you leave out stops applying. Keys and values are the ones `GET /people` accepts, such as an `os` of `iOS` or a `country` of `US`, and at least one filter is required. Date filters must be rolling windows — `first_seen_within_days` or `last_seen_within_days` — so the audience re-anchors every time it rebuilds; fixed dates such as `first_seen_after` are rejected, as is `audience_id`. An array value holds at most 500 items, and each value at most 10 KB. Only an audience with a `source_type` of `people_filter` and `auto_refresh` of `true` accepts filters: an uploaded list has no filters to replace, and with auto refresh off the audience keeps the people it matched when it was built, so create a new audience instead.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — New audience display name. A blank value is ignored rather than clearing the name.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Audiences.AddPeople(ID, request) -> *whopsdk.Audience</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Adds users from a new CSV file to an existing uploaded custom audience. The file uses the audience's saved column mapping, processing happens in the background, and existing audience members remain unchanged.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.AddPeopleAudiencesRequest{
    ID: "id",
    FileID: "file_xxxxxxxxxxxxxx",
}
client.Audiences.AddPeople(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Audience ID, prefixed `adaud_`.
    
</dd>
</dl>

<dl>
<dd>

**fileID:** `string` — The new customer CSV — a file id (`file_...`) returned by `POST /files`. Its headers must match the audience's saved column mapping.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## AuthorizedUsers
<details><summary><code>client.AuthorizedUsers.List() -> *whopsdk.ListAuthorizedUsersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of authorized team members for a company, with optional filtering by user, role, and creation date.

Required permissions:
 - `company:authorized_user:read`
 - `member:email:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListAuthorizedUsersRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    CompanyID: whopsdk.String(
        "biz_xxxxxxxxxxxxxx",
    ),
    UserID: whopsdk.String(
        "user_xxxxxxxxxxxxx",
    ),
    CreatedBefore: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
    CreatedAfter: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
}
client.AuthorizedUsers.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `*string` — The unique identifier of the company to list authorized users for.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — Filter results to a specific user to check if they are an authorized team member.
    
</dd>
</dl>

<dl>
<dd>

**role:** `*whopsdk.AuthorizedUserRoles` 
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — Only return authorized users created before this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — Only return authorized users created after this timestamp.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AuthorizedUsers.Create(request) -> *whopsdk.AuthorizedUser</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a new authorized user to a company.

Required permissions:
 - `authorized_user:create`
 - `member:email:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateAuthorizedUsersRequest{
    CompanyID: "biz_xxxxxxxxxxxxxx",
    Role: whopsdk.GrantableAuthorizedUserRolesOwner,
    UserID: "user_xxxxxxxxxxxxx",
}
client.AuthorizedUsers.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `string` — The ID of the company to add the authorized user to.
    
</dd>
</dl>

<dl>
<dd>

**elevation:** `*whopsdk.CreateAuthorizedUsersRequestElevation` — Re-authentication proof required to perform this sensitive action.
    
</dd>
</dl>

<dl>
<dd>

**role:** `*whopsdk.GrantableAuthorizedUserRoles` — The role to assign to the authorized user within the company. Supported roles: 'moderator', 'sales_manager'.
    
</dd>
</dl>

<dl>
<dd>

**sendEmails:** `*bool` — Whether to send notification emails to the user on creation.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `string` — The ID of the user to add as an authorized user.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AuthorizedUsers.Retrieve(ID) -> *whopsdk.AuthorizedUser</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing authorized user.

Required permissions:
 - `company:authorized_user:read`
 - `member:email:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveAuthorizedUsersRequest{
    ID: "ausr_xxxxxxxxxxxxx",
}
client.AuthorizedUsers.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the authorized user to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AuthorizedUsers.Delete(ID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove an authorized user from a company.

Required permissions:
 - `authorized_user:delete`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteAuthorizedUsersRequest{
    ID: "ausr_xxxxxxxxxxxxx",
    CompanyID: whopsdk.String(
        "biz_xxxxxxxxxxxxxx",
    ),
}
client.AuthorizedUsers.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the authorized user or user to remove.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `*string` — The ID of the company the authorized user belongs to. Optional if the authorized user ID is provided.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Bounties
<details><summary><code>client.Bounties.List() -> *whopsdk.ListBountiesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists bounties visible to the credential — for an account API key, the account's bounties including scheduled drafts; for a user token, the bounties the user can see and work.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListBountiesRequest{}
client.Bounties.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — Scope the list to this account (`biz_` tag). Requires read access to the account; account API keys may pass their own account or a connected account.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — List the bounties this user participated in (`user_` tag). Must be the authenticated user.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.ListBountiesRequestStatus` — Filter by lifecycle state.
    
</dd>
</dl>

<dl>
<dd>

**businessGoalType:** `*whopsdk.ListBountiesRequestBusinessGoalType` — Filter by the poster's declared goal. Bounties created before the goal taxonomy carry no goal and never match this filter.
    
</dd>
</dl>

<dl>
<dd>

**country:** `*string` — Only bounties workable from this country, as an ISO 3166-1 alpha-2 code. Bounties with no country targeting are workable worldwide and always match.
    
</dd>
</dl>

<dl>
<dd>

**experienceID:** `*string` — Only bounties posted to this forum experience, prefixed `exp_`. An unknown experience, or one outside the caller's scope, matches nothing.
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` — Substring match on the bounty title or ID.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Only bounties created after this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Only bounties created before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListBountiesRequestOrder` — Sort field.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListBountiesRequestDirection` — Sort direction.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Number of bounties to return from the start of the window.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor to paginate forwards from.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Number of bounties to return from the end of the window.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Cursor to paginate backwards from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bounties.Create(request) -> *whopsdk.Bounty</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a bounty and escrows its reward pool. Publishes immediately, or as a scheduled draft when you set `publish_at`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateBountiesRequest{
    Description: "Record one continuous pass of a full interior detail, dash to trunk, on a customer vehicle.",
    GrossRewardAmount: 40,
    Title: "Record interior detailing passes",
}
client.Bounties.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**acceptedSubmissionsLimit:** `*int` — Number of submissions that can be accepted (winner slots). Defaults to 1. The escrowed total is `gross_reward_amount` times this limit and must be at least $5.
    
</dd>
</dl>

<dl>
<dd>

**acceptedSubmissionsPerUserLimit:** `*int` — How many winner slots one worker can win. Defaults to `1`. Wins plus proofs awaiting review never exceed this number, and a worker runs one attempt at a time. Cannot exceed `accepted_submissions_limit`.
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `*string` — Account whose balance funds the bounty pool (`biz_` tag). Defaults to the caller's personal balance. Requires permission to move the account's funds.
    
</dd>
</dl>

<dl>
<dd>

**allowedCountryCodes:** `[]string` — Countries whose residents can work the bounty, as ISO 3166 alpha-2 codes. Empty means worldwide.
    
</dd>
</dl>

<dl>
<dd>

**businessGoalType:** `*whopsdk.CreateBountiesRequestBusinessGoalType` — What the poster wants the work to achieve, declared once here.
    
</dd>
</dl>

<dl>
<dd>

**captureSpec:** `*whopsdk.CreateBountiesRequestCaptureSpec` — Per-bounty overrides of the served capture contract. Only accepted when `business_goal_type` is `data_capture`; omitted fields keep the platform defaults, and the resulting contract is echoed back as `capture_spec` on the bounty.
    
</dd>
</dl>

<dl>
<dd>

**description:** `string` — Full task instructions shown to workers.
    
</dd>
</dl>

<dl>
<dd>

**experienceID:** `*string` — Experience to host the bounty in (`exp_` tag). Any visibility — public for an open bounty, private for an invited one. Required unless account_id is set, in which case the bounty anchors in that account's public forum.
    
</dd>
</dl>

<dl>
<dd>

**frequency:** `*whopsdk.CreateBountiesRequestFrequency` — How often the schedule creates a new bounty. Each occurrence is a separate bounty. Defaults to `once`; only applies with `publish_at`.
    
</dd>
</dl>

<dl>
<dd>

**grossRewardAmount:** `float64` — Gross bounty-pool amount (USD) escrowed per accepted submission, in whole dollars. Platform fees and affiliate shares are paid from this amount.
    
</dd>
</dl>

<dl>
<dd>

**publishAt:** `*string` — ISO 8601 time to publish the bounty. When set, the bounty is created as a hidden draft and funded + published at this time instead of immediately.
    
</dd>
</dl>

<dl>
<dd>

**publishAtTimezone:** `*string` — IANA timezone for recurring occurrences. Required when publish_at is set.
    
</dd>
</dl>

<dl>
<dd>

**title:** `string` — Short name of the task shown to workers.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bounties.Retrieve(ID) -> *whopsdk.Bounty</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a bounty by ID. Authentication is optional: a request with no credential reads the bounty when it is publicly visible — published or completed, and not restricted to a private experience's members. Bounties outside the caller's scope, and bounties not publicly visible to an anonymous caller, return `404`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveBountiesRequest{
    ID: "id",
}
client.Bounties.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Bounty ID (`bnty_` tag).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bounties.Update(ID, request) -> *whopsdk.Bounty</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a bounty. A published bounty accepts title, description, and country targeting while it is still open with nothing under review. A scheduled (not-yet-published) draft additionally accepts the reward, winner slots, and schedule.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateBountiesRequest{
    ID: "id",
}
client.Bounties.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Bounty ID (`bnty_` tag).
    
</dd>
</dl>

<dl>
<dd>

**acceptedSubmissionsLimit:** `*int` — Scheduled drafts only. Number of submissions that can be accepted (winner slots).
    
</dd>
</dl>

<dl>
<dd>

**acceptedSubmissionsPerUserLimit:** `*int` — How many winner slots one worker can win. Defaults to `1`. Wins plus proofs awaiting review never exceed this number, and a worker runs one attempt at a time. Cannot exceed `accepted_submissions_limit`. Editable while the bounty is still open with nothing under review.
    
</dd>
</dl>

<dl>
<dd>

**allowedCountryCodes:** `[]string` — Replace the countries whose residents can work the bounty, as ISO 3166 alpha-2 codes. Empty means worldwide.
    
</dd>
</dl>

<dl>
<dd>

**businessGoalType:** `*whopsdk.UpdateBountiesRequestBusinessGoalType` — What the poster wants the work to achieve, declared once here.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — New full task instructions.
    
</dd>
</dl>

<dl>
<dd>

**frequency:** `*whopsdk.UpdateBountiesRequestFrequency` — Scheduled drafts only. How often the schedule creates a new bounty.
    
</dd>
</dl>

<dl>
<dd>

**grossRewardAmount:** `*float64` — Scheduled drafts only. Gross bounty-pool amount (USD) escrowed per accepted submission. The escrowed total (this times accepted_submissions_limit) must stay at least $5.
    
</dd>
</dl>

<dl>
<dd>

**publishAt:** `*string` — Scheduled drafts only. New ISO 8601 time to publish the draft. Must be in the future.
    
</dd>
</dl>

<dl>
<dd>

**publishAtTimezone:** `*string` — Scheduled drafts only. IANA timezone for recurring occurrences.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — New short name of the task.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bounties.Cancel(ID) -> *whopsdk.Bounty</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancels a bounty. With no in-flight work, it cancels immediately and refunds the funder. Otherwise it stops new submissions and cancels once the in-flight work resolves and pays out. Repeating the request is a no-op. A bounty that already paid out every slot returns `400`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CancelBountiesRequest{
    ID: "id",
}
client.Bounties.Cancel(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Bounty ID (`bnty_` tag).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Bounty Submissions
<details><summary><code>client.BountySubmissions.List() -> *whopsdk.ListBountySubmissionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists bounty submissions visible to the credential — for a user token, the submissions they authored plus those on bounties they posted; for an account API key, the submissions on the account's bounties. For the anonymous view of one bounty's reviewed work, use the submissions list under the bounty instead.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListBountySubmissionsRequest{}
client.BountySubmissions.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — Scope the list to submissions on this account's bounties (`biz_` tag). Requires read access to the account.
    
</dd>
</dl>

<dl>
<dd>

**bountyID:** `*string` — Only submissions on this bounty (`bnty_` tag).
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.ListBountySubmissionsRequestStatus` — Filter by lifecycle state.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Only submissions created after this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Only submissions created before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListBountySubmissionsRequestOrder` — Sort field.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListBountySubmissionsRequestDirection` — Sort direction.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Number of submissions to return from the start of the window.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor to paginate forwards from.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Number of submissions to return from the end of the window.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Cursor to paginate backwards from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BountySubmissions.Create(request) -> *whopsdk.BountySubmission</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a submission on a workforce bounty. Include a `deliverable` payload — any combination of links and uploaded files, with at least one of the two — and the submission goes straight to review; create is the only step. For `data_capture` bounties, omit the deliverable: this starts a claimed attempt whose proof accumulates server-side, and the separate submit endpoint sends it to review once complete. Requires a user credential — account API keys cannot author submissions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateBountySubmissionsRequest{
    BountyID: "bnty_xxxxxxxxxxxxxx",
}
client.BountySubmissions.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**affiliateCode:** `*string` — Affiliate code crediting the referrer, when the worker arrived through one.
    
</dd>
</dl>

<dl>
<dd>

**bountyID:** `string` — The bounty to submit to (`bnty_` tag).
    
</dd>
</dl>

<dl>
<dd>

**deliverable:** `*whopsdk.CreateBountySubmissionsRequestDeliverable` — The submitted work. Combine `urls`, `file_ids`, and `caption` freely; at least one link or file is required.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `*whopsdk.CreateBountySubmissionsRequestMetadata` — Optional capture metadata describing where and how the footage was recorded. Persisted on the submission. On a `data_capture` bounty every field except `fov` is required whenever metadata is provided.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BountySubmissions.Retrieve(ID) -> *whopsdk.BountySubmission</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves one bounty submission the credential can see — one the caller authored, or one on a bounty they posted or their account owns. Reading another member's work on an account's bounty takes `account_id`, the same way the list does.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveBountySubmissionsRequest{
    ID: "id",
}
client.BountySubmissions.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The bounty submission to act on (`btys_` tag).
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `*string` — Read the submission as this account (`biz_` tag), scoping the lookup to its bounties rather than the caller's own work. Requires read access to the account. Without it the lookup covers only what the credential owns — the submissions the caller authored plus those on bounties they posted.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BountySubmissions.Delete(ID) -> *whopsdk.DeleteBountySubmissionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancels the caller's own active attempt on a bounty and discards any accumulated capture clips. Only the worker who started the attempt can cancel it — account API keys cannot.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteBountySubmissionsRequest{
    ID: "id",
}
client.BountySubmissions.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The bounty submission to act on (`btys_` tag).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BountySubmissions.Submit(ID, request) -> *whopsdk.BountySubmission</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Submits a claimed attempt for review. A livestream attempt needs an ended proof stream and can attach an optional `deliverable` — links, files, and a caption in any combination; if the attempt already went to review when its stream ended, the payload attaches to it once, until reviewers start voting. A data capture attempt instead needs enough validated clip time and takes no payload. Only the worker who started the attempt can submit it — account API keys cannot.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.SubmitBountySubmissionsRequest{
    ID: "id",
}
client.BountySubmissions.Submit(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The claimed attempt to submit for review (`btys_` tag).
    
</dd>
</dl>

<dl>
<dd>

**deliverable:** `*whopsdk.SubmitBountySubmissionsRequestDeliverable` — Work to attach to the submission. Combine `urls`, `file_ids`, and `caption` freely; all are optional.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## CardTransactions
<details><summary><code>client.CardTransactions.List() -> *whopsdk.ListCardTransactionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists an account's card transactions, newest first. Defaults to the account the credential belongs to. Covers every card the owner has ever had, including canceled cards and spend that predates a re-application, and team members only see transactions on the cards assigned to them. Pass `transaction_ids` to fetch specific transactions instead of paging for them.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListCardTransactionsRequest{
    TransactionIDs: []*string{
        whopsdk.String(
            "citx_xxxxxxxxxxxxxx",
        ),
    },
    CardID: []*string{
        whopsdk.String(
            "icrd_xxxxxxxxxxxxxx",
        ),
    },
    CardholderID: []*string{
        whopsdk.String(
            "user_xxxxxxxxxxxxxx",
        ),
    },
}
client.CardTransactions.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The account whose card transactions to list, prefixed `biz_`. Defaults to the credential's account.
    
</dd>
</dl>

<dl>
<dd>

**transactionIDs:** `*string` — Return only these card transactions, each prefixed `citx_`. Repeat the parameter, or pass one comma-separated value.
    
</dd>
</dl>

<dl>
<dd>

**cardID:** `*string` — Return only transactions charged to these cards, each prefixed `icrd_`.
    
</dd>
</dl>

<dl>
<dd>

**cardholderID:** `*string` — Return only transactions on cards assigned to these users, each prefixed `user_`.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.ListCardTransactionsRequestStatus` — Return only transactions with this status.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Return only transactions authorized at or after this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Return only transactions authorized at or before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListCardTransactionsRequestOrder` — The field to sort by. Defaults to `created_at`.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListCardTransactionsRequestDirection` — The sort direction. Defaults to `desc`.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of card transactions to return.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor; returns card transactions after this position.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — The number of card transactions to return, counting back from the end.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — A cursor; returns card transactions before this position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CardTransactions.Retrieve(ID) -> *whopsdk.CardTransaction</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Fetches a single card transaction by its `citx_` identifier. The owner defaults to the account the credential belongs to.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveCardTransactionsRequest{
    ID: "id",
}
client.CardTransactions.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The card transaction ID, prefixed `citx_`.
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `*string` — The account that owns the transaction, prefixed `biz_`. Defaults to the credential's account.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Cards
<details><summary><code>client.Cards.List() -> *whopsdk.ListCardsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the Whop cards of an account or user, including ones still being set up. Team members only see the cards assigned to them.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListCardsRequest{}
client.Cards.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The owning account ID (a biz_ identifier). Provide this or user_id.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — The owning user ID (a user_ identifier). Provide this or account_id.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Cards.Create(request) -> *whopsdk.CreateCardsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Issue a virtual card, or apply for card issuing.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateCardsRequest{}
client.Cards.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The owning account ID (a biz_ identifier). Provide this or user_id.
    
</dd>
</dl>

<dl>
<dd>

**assignedUserID:** `*string` — The account member (a user_ identifier) to assign the card to. Required for business card issuing accounts.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — A display name for the card.
    
</dd>
</dl>

<dl>
<dd>

**spendLimit:** `*float64` — Spending limit amount, in dollars.
    
</dd>
</dl>

<dl>
<dd>

**spendLimitFrequency:** `*whopsdk.CreateCardsRequestSpendLimitFrequency` — The window the spend limit applies to.
    
</dd>
</dl>

<dl>
<dd>

**transactionLimit:** `*float64` — Per-transaction limit amount, in dollars.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — The owning user ID (a user_ identifier). Provide this or account_id.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Cards.Retrieve(ID) -> *whopsdk.RetrieveCardsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a single card.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveCardsRequest{
    ID: "id",
}
client.Cards.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Card ID to retrieve, prefixed `icrd_`.
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `*string` — The owning account ID (a biz_ identifier). Provide this or user_id.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — The owning user ID (a user_ identifier). Provide this or account_id.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Cards.Update(ID, request) -> *whopsdk.UpdateCardsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update, freeze, or cancel a card. Updating the card's name, billing address, or limits requires both `payout:account:update` and `company:balance:read`; a card's assigned holder may update their own card's pin and frozen state with any user token.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateCardsRequest{
    ID: "id",
}
client.Cards.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Card ID to retrieve, prefixed `icrd_`.
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `*string` — The owning account ID (a biz_ identifier). Provide this or user_id.
    
</dd>
</dl>

<dl>
<dd>

**billing:** `*whopsdk.UpdateCardsRequestBilling` — New billing address. Requires line1, city, region, postal_code, and country_code. On an invited card, passing billing alone (as the invited user) completes onboarding and starts card provisioning.
    
</dd>
</dl>

<dl>
<dd>

**canceled:** `*bool` — Pass `true` to permanently cancel the card. A canceled card cannot be uncanceled. Cannot be combined with other fields.
    
</dd>
</dl>

<dl>
<dd>

**frozen:** `*bool` — Pass `true` to freeze the card, `false` to unfreeze it. The assigned cardholder may freeze their own card without the payout:account:update scope.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — A display name for the card.
    
</dd>
</dl>

<dl>
<dd>

**pin:** `*string` — New 4-digit PIN. Can only be set on a card assigned to the acting user, who may set it without the payout:account:update scope.
    
</dd>
</dl>

<dl>
<dd>

**removeLimit:** `*bool` — Pass `true` to remove the spending limit (make the card unlimited).
    
</dd>
</dl>

<dl>
<dd>

**spendLimit:** `*float64` — Spending limit amount, in dollars.
    
</dd>
</dl>

<dl>
<dd>

**spendLimitFrequency:** `*whopsdk.UpdateCardsRequestSpendLimitFrequency` — The window the spend limit applies to.
    
</dd>
</dl>

<dl>
<dd>

**transactionLimit:** `*float64` — Per-transaction limit amount, in dollars.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — The owning user ID (a user_ identifier). Provide this or account_id.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ChatChannels
<details><summary><code>client.ChatChannels.List() -> *whopsdk.ListChatChannelsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of chat channels within a specific company, with optional filtering by product.

Required permissions:
 - `chat:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListChatChannelsRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    CompanyID: "biz_xxxxxxxxxxxxxx",
    ProductID: whopsdk.String(
        "prod_xxxxxxxxxxxxx",
    ),
}
client.ChatChannels.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `string` — The unique identifier of the company to list chat channels for.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `*string` — The unique identifier of a product to filter by. When set, only chat channels connected to this product are returned.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ChatChannels.Retrieve(ID) -> *whopsdk.ChatChannel</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing chat channel.

Required permissions:
 - `chat:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveChatChannelsRequest{
    ID: "id",
}
client.ChatChannels.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the chat channel or experience to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ChatChannels.Update(ID, request) -> *whopsdk.ChatChannel</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update moderation settings for a chat channel, such as who can post, banned words, and media restrictions.

Required permissions:
 - `chat:moderate`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateChatChannelsRequest{
    ID: "id",
}
client.ChatChannels.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the chat channel to update. Accepts either an experience ID (e.g. 'exp_xxxxx') or a chat channel ID.
    
</dd>
</dl>

<dl>
<dd>

**banMedia:** `*bool` — Whether media uploads such as images and videos are banned in this chat channel.
    
</dd>
</dl>

<dl>
<dd>

**banURLs:** `*bool` — Whether URLs and links are banned from being posted in this chat channel.
    
</dd>
</dl>

<dl>
<dd>

**bannedWords:** `[]string` — A list of words that are automatically blocked from messages in this chat channel. For example, ['spam', 'scam'].
    
</dd>
</dl>

<dl>
<dd>

**userPostsCooldownSeconds:** `*int` — The minimum number of seconds a user must wait between sending messages in this chat channel.
    
</dd>
</dl>

<dl>
<dd>

**whoCanPost:** `*whopsdk.WhoCanPostTypes` — Controls which roles are allowed to send messages in this chat channel.
    
</dd>
</dl>

<dl>
<dd>

**whoCanReact:** `*whopsdk.WhoCanReactTypes` — Controls which roles are allowed to add reactions to messages in this chat channel.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Checkout Configurations
<details><summary><code>client.CheckoutConfigurations.List() -> *whopsdk.ListCheckoutConfigurationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists checkout configurations for an account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListCheckoutConfigurationsRequest{
    AccountID: "account_id",
}
client.CheckoutConfigurations.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string` — Account ID, prefixed `biz_`.
    
</dd>
</dl>

<dl>
<dd>

**planID:** `*string` — Only return checkout configurations for this plan ID, prefixed `plan_`.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Only return checkout configurations created before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Only return checkout configurations created after this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListCheckoutConfigurationsRequestOrder` — Field used to sort checkout configurations.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListCheckoutConfigurationsRequestDirection` — Sort direction. Defaults to `desc`.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Number of checkout configurations to return.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor for the next page of results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CheckoutConfigurations.Create(request) -> *whopsdk.CreateCheckoutConfigurationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a reusable checkout configuration for an existing or inline plan.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateCheckoutConfigurationsRequest{
    AccountID: whopsdk.String(
        "biz_xxxxxxxxxxxxxx",
    ),
    PlanID: whopsdk.String(
        "plan_xxxxxxxxxxxxx",
    ),
}
client.CheckoutConfigurations.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — Account ID, prefixed `biz_`.
    
</dd>
</dl>

<dl>
<dd>

**affiliateCode:** `*string` — Affiliate code to apply to the checkout.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*string` — Currency used for setup-mode payment method availability.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Custom key-value metadata copied to payments and memberships.
    
</dd>
</dl>

<dl>
<dd>

**mode:** `*whopsdk.CreateCheckoutConfigurationsRequestMode` — Controls whether checkout charges the buyer immediately or saves payment details for later. Defaults to `payment`.
    
</dd>
</dl>

<dl>
<dd>

**paymentMethodConfiguration:** `*whopsdk.CreateCheckoutConfigurationsRequestPaymentMethodConfiguration` — Payment method overrides for this checkout. `null` uses the plan or platform defaults.
    
</dd>
</dl>

<dl>
<dd>

**plan:** `*whopsdk.CreateCheckoutConfigurationsRequestPlan` — Plan attributes used to create or find a plan for this checkout configuration. Mutually exclusive with `plan_id`.
    
</dd>
</dl>

<dl>
<dd>

**planID:** `*string` — Existing plan ID, prefixed `plan_`. Mutually exclusive with `plan`.
    
</dd>
</dl>

<dl>
<dd>

**redirectURL:** `*string` — URL customers are sent to after checkout.
    
</dd>
</dl>

<dl>
<dd>

**threeDsLevel:** `*whopsdk.CreateCheckoutConfigurationsRequestThreeDsLevel` — 3D Secure behavior for this checkout.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CheckoutConfigurations.Retrieve(ID) -> *whopsdk.RetrieveCheckoutConfigurationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a checkout configuration by ID. This endpoint is public so a checkout page can load from the configuration URL.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveCheckoutConfigurationsRequest{
    ID: "id",
}
client.CheckoutConfigurations.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the checkout configuration.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CheckoutConfigurations.Delete(ID) -> *whopsdk.DeleteCheckoutConfigurationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a checkout configuration so its checkout URL can no longer be used.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteCheckoutConfigurationsRequest{
    ID: "id",
}
client.CheckoutConfigurations.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the checkout configuration.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Companies
<details><summary><code>client.Companies.List() -> *whopsdk.ListCompaniesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of companies. When parent_company_id is provided, lists connected accounts under that platform. When omitted, lists companies the current user has access to.

Required permissions:
 - `company:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListCompaniesRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    CreatedBefore: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
    CreatedAfter: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
}
client.Companies.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**parentCompanyID:** `*string` — The unique identifier of the parent platform company. When provided, lists connected accounts under that platform. Omit to list the current user's own companies.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.Direction` 
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — Only return companies created before this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — Only return companies created after this timestamp.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.Create(request) -> *whopsdk.Company</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new company. Pass parent_company_id to create a connected account under a platform, or omit it to create a company for the current user.

Required permissions:
 - `company:create`
 - `company:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateCompaniesRequest{
    Title: "title",
}
client.Companies.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**country:** `*whopsdk.Countries` — The country the company is located in. Defaults to the parent company's country for connected accounts, or the owner's IP-derived country.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — A promotional pitch displayed to potential customers on the company's store page.
    
</dd>
</dl>

<dl>
<dd>

**email:** `*string` — The email address of the user who will own the connected account. Required when parent_company_id is provided.
    
</dd>
</dl>

<dl>
<dd>

**logo:** `*whopsdk.CreateCompaniesRequestLogo` — The company's logo image. Accepts PNG, JPEG, or GIF format.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — A key-value JSON object of custom metadata to store on the company.
    
</dd>
</dl>

<dl>
<dd>

**parentCompanyID:** `*string` — The unique identifier of the parent platform company. When provided, creates a connected account under that platform. Omit to create a company for the current user.
    
</dd>
</dl>

<dl>
<dd>

**sendCustomerEmails:** `*bool` — Whether Whop sends transactional emails to customers on behalf of this company. Only applies when creating a connected account.
    
</dd>
</dl>

<dl>
<dd>

**title:** `string` — The display name of the company shown to customers.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.Retrieve(ID) -> *whopsdk.Company</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing company.

Required permissions:
 - `company:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveCompaniesRequest{
    ID: "biz_xxxxxxxxxxxxxx",
}
client.Companies.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier or route slug of the company.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.Update(ID, request) -> *whopsdk.Company</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a company's title, description, logo, and other settings.

Required permissions:
 - `company:update`
 - `company:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateCompaniesRequest{
    ID: "biz_xxxxxxxxxxxxxx",
}
client.Companies.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the company to update.
    
</dd>
</dl>

<dl>
<dd>

**affiliateApplicationRequired:** `*bool` — Whether prospective affiliates must submit an application before they can promote this company.
    
</dd>
</dl>

<dl>
<dd>

**affiliateInstructions:** `*string` — Guidelines and instructions shown to affiliates explaining how to promote this company's products.
    
</dd>
</dl>

<dl>
<dd>

**bannerImage:** `*whopsdk.UpdateCompaniesRequestBannerImage` — The company's banner image. Accepts PNG or JPEG format.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — A promotional pitch displayed to potential customers on the company's store page.
    
</dd>
</dl>

<dl>
<dd>

**featuredAffiliateProductID:** `*string` — The ID of the product to feature on this company's affiliate page. Pass null to clear.
    
</dd>
</dl>

<dl>
<dd>

**logo:** `*whopsdk.UpdateCompaniesRequestLogo` — The company's logo image. Accepts PNG, JPEG, or GIF format.
    
</dd>
</dl>

<dl>
<dd>

**route:** `*string` — The unique URL slug for the company's store page. Must be lowercase and can include hyphens (e.g., 'my-company'). If not provided, the route will remain unchanged.
    
</dd>
</dl>

<dl>
<dd>

**sendCustomerEmails:** `*bool` — Whether Whop sends transactional emails (receipts, renewals, cancelations) to customers on behalf of this company.
    
</dd>
</dl>

<dl>
<dd>

**socialLinks:** `[]*whopsdk.UpdateCompaniesRequestSocialLinksItem` — The social media links to display on the company's store page. Pass the full list of desired social links — any existing links not included will be removed.
    
</dd>
</dl>

<dl>
<dd>

**targetAudience:** `*string` — The target audience for this company (e.g., 'beginner day traders aged 18-25 looking to learn options').
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The display name of the company shown to customers.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Companies.CreateAPIKey(ParentCompanyID, request) -> *whopsdk.CreateAPIKeyCompaniesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create an API key for a connected account (child company) owned by a parent company.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateAPIKeyCompaniesRequest{
    ParentCompanyID: "parent_company_id",
    ChildCompanyID: "child_company_id",
}
client.Companies.CreateAPIKey(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**parentCompanyID:** `string` — The unique identifier of the parent platform company (e.g. 'biz_xxx').
    
</dd>
</dl>

<dl>
<dd>

**childCompanyID:** `string` — The unique identifier of the connected account to create the API key for (e.g. 'biz_xxx').
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — A human-readable name for the API key, such as 'Production API Key'.
    
</dd>
</dl>

<dl>
<dd>

**permissions:** `[]*whopsdk.CreateAPIKeyCompaniesRequestPermissionsItem` — Granular permission statements defining which actions this API key can perform. Either permissions or role must be provided.
    
</dd>
</dl>

<dl>
<dd>

**role:** `*whopsdk.PermissionSystemRoles` — A system role to inherit permissions from (e.g. owner, admin, moderator). Either role or permissions must be provided.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## CompanyTokenTransactions
<details><summary><code>client.CompanyTokenTransactions.List() -> *whopsdk.ListCompanyTokenTransactionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of token transactions for a user or company, depending on the authenticated actor, with optional filtering by user and transaction type.

Required permissions:
 - `company_token_transaction:read`
 - `member:basic:read`
 - `company:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListCompanyTokenTransactionsRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    CompanyID: "biz_xxxxxxxxxxxxxx",
    UserID: whopsdk.String(
        "user_xxxxxxxxxxxxx",
    ),
}
client.CompanyTokenTransactions.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `string` — The unique identifier of the company to list token transactions for.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — Filter transactions to only those involving this specific user.
    
</dd>
</dl>

<dl>
<dd>

**transactionType:** `*whopsdk.CompanyTokenTransactionTypes` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CompanyTokenTransactions.Create(request) -> *whopsdk.CompanyTokenTransaction</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a token transaction to add, subtract, or transfer tokens for a member within a company.

Required permissions:
 - `company_token_transaction:create`
 - `member:basic:read`
 - `company:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateCompanyTokenTransactionsRequest{
    Transfer: &whopsdk.CreateCompanyTokenTransactionsRequestTransfer{
        Amount: 6.9,
        CompanyID: "biz_xxxxxxxxxxxxxx",
        DestinationUserID: "destination_user_id",
        UserID: "user_xxxxxxxxxxxxx",
    },
}
client.CompanyTokenTransactions.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*whopsdk.CreateCompanyTokenTransactionsRequest` — Parameters for CreateCompanyTokenTransaction
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CompanyTokenTransactions.Retrieve(ID) -> *whopsdk.CompanyTokenTransaction</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing company token transaction.

Required permissions:
 - `company_token_transaction:read`
 - `member:basic:read`
 - `company:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveCompanyTokenTransactionsRequest{
    ID: "id",
}
client.CompanyTokenTransactions.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the token transaction to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## CourseChapters
<details><summary><code>client.CourseChapters.List() -> *whopsdk.ListCourseChaptersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of chapters within a course, ordered by position.

Required permissions:
 - `courses:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListCourseChaptersRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    CourseID: "cors_xxxxxxxxxxxxx",
}
client.CourseChapters.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**courseID:** `string` — The unique identifier of the course to list chapters for.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CourseChapters.Create(request) -> *whopsdk.CourseChapter</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new chapter within a course to organize lessons into sections.

Required permissions:
 - `courses:update`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateCourseChaptersRequest{
    CourseID: "cors_xxxxxxxxxxxxx",
}
client.CourseChapters.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**courseID:** `string` — The unique identifier of the course to create the chapter in (e.g., "course_XXXXX").
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The display title of the chapter (e.g., "Module 1: Introduction").
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CourseChapters.Retrieve(ID) -> *whopsdk.CourseChapter</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing course chapter.

Required permissions:
 - `courses:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveCourseChaptersRequest{
    ID: "chap_xxxxxxxxxxxxx",
}
client.CourseChapters.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the chapter to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CourseChapters.Delete(ID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Permanently delete a chapter and all of its lessons from a course.

Required permissions:
 - `courses:update`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteCourseChaptersRequest{
    ID: "chap_xxxxxxxxxxxxx",
}
client.CourseChapters.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the chapter to delete (e.g., "chap_XXXXX").
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CourseChapters.Update(ID, request) -> *whopsdk.CourseChapter</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a chapter's title within a course.

Required permissions:
 - `courses:update`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateCourseChaptersRequest{
    ID: "chap_xxxxxxxxxxxxx",
    Title: "title",
}
client.CourseChapters.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the chapter to update (e.g., "chap_XXXXX").
    
</dd>
</dl>

<dl>
<dd>

**title:** `string` — The new display title of the chapter (e.g., "Module 1: Introduction").
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## CourseLessonInteractions
<details><summary><code>client.CourseLessonInteractions.List() -> *whopsdk.ListCourseLessonInteractionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of lesson interactions, filtered by lesson, course, user, or completion status.

Required permissions:
 - `courses:read`
 - `course_analytics:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListCourseLessonInteractionsRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    UserID: whopsdk.String(
        "user_xxxxxxxxxxxxx",
    ),
    LessonID: whopsdk.String(
        "lesn_xxxxxxxxxxxxx",
    ),
    CourseID: whopsdk.String(
        "cors_xxxxxxxxxxxxx",
    ),
}
client.CourseLessonInteractions.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — The unique identifier of the user to filter lesson interactions for.
    
</dd>
</dl>

<dl>
<dd>

**lessonID:** `*string` — The unique identifier of the lesson to filter interactions for.
    
</dd>
</dl>

<dl>
<dd>

**courseID:** `*string` — The unique identifier of the course to filter interactions for.
    
</dd>
</dl>

<dl>
<dd>

**completed:** `*bool` — Whether to filter for completed or in-progress lesson interactions.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CourseLessonInteractions.Retrieve(ID) -> *whopsdk.CourseLessonInteraction</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing course lesson interaction.

Required permissions:
 - `courses:read`
 - `course_analytics:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveCourseLessonInteractionsRequest{
    ID: "crsli_xxxxxxxxxxxx",
}
client.CourseLessonInteractions.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the lesson interaction to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## CourseLessons
<details><summary><code>client.CourseLessons.List() -> *whopsdk.ListCourseLessonsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of lessons within a course or chapter, ordered by position.

Required permissions:
 - `courses:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListCourseLessonsRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    CourseID: whopsdk.String(
        "cors_xxxxxxxxxxxxx",
    ),
    ChapterID: whopsdk.String(
        "chap_xxxxxxxxxxxxx",
    ),
}
client.CourseLessons.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**courseID:** `*string` — The unique identifier of the course to return all lessons across all chapters.
    
</dd>
</dl>

<dl>
<dd>

**chapterID:** `*string` — The unique identifier of a chapter to return only its lessons.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CourseLessons.Create(request) -> *whopsdk.CourseLesson</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new lesson within a course chapter. Lessons can contain video, text, or assessment content.

Required permissions:
 - `courses:update`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateCourseLessonsRequest{
    ChapterID: "chap_xxxxxxxxxxxxx",
    LessonType: whopsdk.LessonTypesText,
}
client.CourseLessons.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**chapterID:** `string` — The unique identifier of the chapter to create the lesson in (e.g., "chap_XXXXX").
    
</dd>
</dl>

<dl>
<dd>

**content:** `*string` — The Markdown content body of the lesson.
    
</dd>
</dl>

<dl>
<dd>

**daysFromCourseStartUntilUnlock:** `*int` — The number of days after a student starts the course before this lesson becomes accessible.
    
</dd>
</dl>

<dl>
<dd>

**embedID:** `*string` — The external video identifier for embedded content (e.g., a YouTube video ID or Loom share ID).
    
</dd>
</dl>

<dl>
<dd>

**embedType:** `*whopsdk.EmbedTypes` — The type of video embed for this lesson, such as YouTube or Loom.
    
</dd>
</dl>

<dl>
<dd>

**lessonType:** `*whopsdk.LessonTypes` — The content type of the lesson, such as video, text, quiz, or knowledge check.
    
</dd>
</dl>

<dl>
<dd>

**thumbnail:** `*whopsdk.CreateCourseLessonsRequestThumbnail` — The thumbnail image for the lesson in PNG, JPEG, or GIF format.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The display title of the lesson (e.g., "Getting Started with APIs").
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CourseLessons.Retrieve(ID) -> *whopsdk.CourseLesson</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing course lesson.

Required permissions:
 - `courses:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveCourseLessonsRequest{
    ID: "lesn_xxxxxxxxxxxxx",
}
client.CourseLessons.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the lesson to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CourseLessons.Delete(ID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Permanently delete a lesson and remove it from its chapter.

Required permissions:
 - `courses:update`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteCourseLessonsRequest{
    ID: "lesn_xxxxxxxxxxxxx",
}
client.CourseLessons.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the lesson to delete (e.g., "les_XXXXX").
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CourseLessons.Update(ID, request) -> *whopsdk.CourseLesson</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a lesson's content, type, visibility, assessment questions, or media attachments.

Required permissions:
 - `courses:update`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateCourseLessonsRequest{
    ID: "lesn_xxxxxxxxxxxxx",
}
client.CourseLessons.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the lesson to update (e.g., "les_XXXXX").
    
</dd>
</dl>

<dl>
<dd>

**assessmentCompletionRequirement:** `*whopsdk.UpdateCourseLessonsRequestAssessmentCompletionRequirement` — The passing criteria for quiz or knowledge check lessons, such as minimum grade or correct answers.
    
</dd>
</dl>

<dl>
<dd>

**assessmentQuestions:** `[]*whopsdk.UpdateCourseLessonsRequestAssessmentQuestionsItem` — The full list of assessment questions for quiz or knowledge check lessons. Replaces all existing questions.
    
</dd>
</dl>

<dl>
<dd>

**attachments:** `[]*whopsdk.UpdateCourseLessonsRequestAttachmentsItem` — File attachments for the lesson such as PDFs or documents. Replaces all existing attachments.
    
</dd>
</dl>

<dl>
<dd>

**content:** `*string` — The Markdown content body of the lesson.
    
</dd>
</dl>

<dl>
<dd>

**daysFromCourseStartUntilUnlock:** `*int` — The number of days after a student starts the course before this lesson becomes accessible.
    
</dd>
</dl>

<dl>
<dd>

**embedID:** `*string` — The external video identifier for embedded content (e.g., a YouTube video ID or Loom share ID).
    
</dd>
</dl>

<dl>
<dd>

**embedType:** `*whopsdk.EmbedTypes` — The type of video embed for this lesson, such as YouTube or Loom.
    
</dd>
</dl>

<dl>
<dd>

**lessonType:** `*whopsdk.LessonTypes` — The content type of the lesson, such as video, text, quiz, or knowledge check.
    
</dd>
</dl>

<dl>
<dd>

**mainPdf:** `*whopsdk.UpdateCourseLessonsRequestMainPdf` — The primary PDF document attached to this lesson for student reference.
    
</dd>
</dl>

<dl>
<dd>

**maxAttempts:** `*int` — The maximum number of attempts a student is allowed for assessment lessons.
    
</dd>
</dl>

<dl>
<dd>

**muxAssetID:** `*string` — The identifier of a Mux video asset to attach to this lesson (e.g., "mux_XXXXX").
    
</dd>
</dl>

<dl>
<dd>

**thumbnail:** `*whopsdk.UpdateCourseLessonsRequestThumbnail` — The thumbnail image for the lesson in PNG, JPEG, or GIF format.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The display title of the lesson (e.g., "Getting Started with APIs").
    
</dd>
</dl>

<dl>
<dd>

**visibility:** `*whopsdk.LessonVisibilities` — Controls whether this lesson is visible to students or hidden as a draft.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CourseLessons.MarkAsCompleted(LessonID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Mark a lesson as completed for the current user after they finish the content.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.MarkAsCompletedCourseLessonsRequest{
    LessonID: "lesson_id",
}
client.CourseLessons.MarkAsCompleted(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**lessonID:** `string` — The unique identifier of the lesson to mark as completed (e.g., "les_XXXXX").
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CourseLessons.Start(LessonID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Record that the current user has started viewing a lesson, creating progress tracking records.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.StartCourseLessonsRequest{
    LessonID: "lesson_id",
}
client.CourseLessons.Start(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**lessonID:** `string` — The unique identifier of the lesson the user is starting (e.g., "les_XXXXX").
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CourseLessons.SubmitAssessment(LessonID, request) -> *whopsdk.SubmitAssessmentCourseLessonsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Submit answers for a quiz or knowledge check lesson and receive a graded result.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.SubmitAssessmentCourseLessonsRequest{
    LessonID: "lesson_id",
    Answers: []*whopsdk.SubmitAssessmentCourseLessonsRequestAnswersItem{
        &whopsdk.SubmitAssessmentCourseLessonsRequestAnswersItem{
            QuestionID: "question_id",
        },
    },
}
client.CourseLessons.SubmitAssessment(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**lessonID:** `string` — The unique identifier of the quiz or knowledge check lesson to submit answers for (e.g., "les_XXXXX").
    
</dd>
</dl>

<dl>
<dd>

**answers:** `[]*whopsdk.SubmitAssessmentCourseLessonsRequestAnswersItem` — The list of answers to submit for each assessment question.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## CourseStudents
<details><summary><code>client.CourseStudents.List() -> *whopsdk.ListCourseStudentsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of students enrolled in a course, with optional name filtering.

Required permissions:
 - `courses:read`
 - `course_analytics:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListCourseStudentsRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    CourseID: "cors_xxxxxxxxxxxxx",
}
client.CourseStudents.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**courseID:** `string` — The unique identifier of the course to list enrolled students for.
    
</dd>
</dl>

<dl>
<dd>

**keyword:** `*string` — A search term to filter students by name or username.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CourseStudents.Retrieve(ID) -> *whopsdk.CourseStudent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing course student.

Required permissions:
 - `courses:read`
 - `course_analytics:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveCourseStudentsRequest{
    ID: "id",
}
client.CourseStudents.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the course student record to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Courses
<details><summary><code>client.Courses.List() -> *whopsdk.ListCoursesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of courses, filtered by either an experience or a company.

Required permissions:
 - `courses:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListCoursesRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    ExperienceID: whopsdk.String(
        "exp_xxxxxxxxxxxxxx",
    ),
    CompanyID: whopsdk.String(
        "biz_xxxxxxxxxxxxxx",
    ),
}
client.Courses.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**experienceID:** `*string` — The unique identifier of the experience to list courses for.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `*string` — The unique identifier of the company to list courses for.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Courses.Create(request) -> *whopsdk.Course</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new course within an experience, with optional chapters, lessons, and a certificate.

Required permissions:
 - `courses:update`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateCoursesRequest{
    ExperienceID: "exp_xxxxxxxxxxxxxx",
    Title: "title",
}
client.Courses.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**certificateAfterCompletionEnabled:** `*bool` — Whether the course awards students a PDF certificate after completing all lessons.
    
</dd>
</dl>

<dl>
<dd>

**experienceID:** `string` — The unique identifier of the experience to create the course in (e.g., "exp_XXXXX").
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — The decimal order position of the course within its experience. Use fractional values (e.g., "1.5") to place between existing courses.
    
</dd>
</dl>

<dl>
<dd>

**requireCompletingLessonsInOrder:** `*bool` — Whether students must complete each lesson sequentially before advancing to the next one.
    
</dd>
</dl>

<dl>
<dd>

**tagline:** `*string` — A short tagline displayed beneath the course title (e.g., "Master the fundamentals of design").
    
</dd>
</dl>

<dl>
<dd>

**thumbnail:** `*whopsdk.CreateCoursesRequestThumbnail` — The thumbnail image for the course in PNG, JPEG, or GIF format.
    
</dd>
</dl>

<dl>
<dd>

**title:** `string` — The display title of the course (e.g., "Introduction to Web Development").
    
</dd>
</dl>

<dl>
<dd>

**visibility:** `*whopsdk.CourseVisibilities` — Controls whether this course is visible to students or hidden as a draft.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Courses.Retrieve(ID) -> *whopsdk.Course</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing course.

Required permissions:
 - `courses:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveCoursesRequest{
    ID: "cors_xxxxxxxxxxxxx",
}
client.Courses.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the course to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Courses.Delete(ID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Permanently delete a course and all of its chapters, lessons, and student progress.

Required permissions:
 - `courses:update`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteCoursesRequest{
    ID: "cors_xxxxxxxxxxxxx",
}
client.Courses.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the course to delete (e.g., "course_XXXXX").
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Courses.Update(ID, request) -> *whopsdk.Course</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a course's title, description, visibility, thumbnail, or chapter ordering.

Required permissions:
 - `courses:update`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateCoursesRequest{
    ID: "cors_xxxxxxxxxxxxx",
}
client.Courses.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the course to update (e.g., "course_XXXXX").
    
</dd>
</dl>

<dl>
<dd>

**certificateAfterCompletionEnabled:** `*bool` — Whether the course awards students a PDF certificate after completing all lessons.
    
</dd>
</dl>

<dl>
<dd>

**chapters:** `[]*whopsdk.UpdateCoursesRequestChaptersItem` — A list of chapters with nested lessons to reorder or rename in bulk.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — A short description of the course displayed to students on the course page.
    
</dd>
</dl>

<dl>
<dd>

**language:** `*whopsdk.Languages` — The primary language spoken in the video content of the course.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — The decimal order position of the course within its experience. Use fractional values (e.g., "1.5") to place between existing courses.
    
</dd>
</dl>

<dl>
<dd>

**requireCompletingLessonsInOrder:** `*bool` — Whether students must complete each lesson sequentially before advancing to the next one.
    
</dd>
</dl>

<dl>
<dd>

**tagline:** `*string` — A short tagline displayed beneath the course title (e.g., "Master the fundamentals of design").
    
</dd>
</dl>

<dl>
<dd>

**thumbnail:** `*whopsdk.UpdateCoursesRequestThumbnail` — The thumbnail image for the course in PNG, JPEG, or GIF format.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The display title of the course (e.g., "Introduction to Web Development").
    
</dd>
</dl>

<dl>
<dd>

**visibility:** `*whopsdk.CourseVisibilities` — Controls whether this course is visible to students or hidden as a draft.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Deposits
<details><summary><code>client.Deposits.Create(request) -> *whopsdk.CreateDepositsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the deposit methods for an account, including crypto and bank transfer.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateDepositsRequest{
    Destination: &whopsdk.CreateDepositsRequestDestination{
        String: "destination",
    },
}
client.Deposits.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**amount:** `*float64` — Amount to prefill on hosted deposit page.
    
</dd>
</dl>

<dl>
<dd>

**destination:** `*whopsdk.CreateDepositsRequestDestination` — Destination account ID or wallet address. Object form is supported for compatibility. Any business resolves by its account ID without authentication; a user account resolves only for that same authenticated user.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Metadata to include with the deposit response.
    
</dd>
</dl>

<dl>
<dd>

**network:** `*whopsdk.CreateDepositsRequestNetwork` — Destination network override. Defaults to the destination wallet's own network.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Dispute alerts
<details><summary><code>client.DisputeAlerts.List() -> *whopsdk.ListDisputeAlertsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the dispute alerts and early fraud warnings across the accounts you can read.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListDisputeAlertsRequest{}
client.DisputeAlerts.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — Only alerts on this account's payments (`biz_` tag). Omit it to cover every account you can read.
    
</dd>
</dl>

<dl>
<dd>

**paymentID:** `*string` — Only alerts on this payment (`pay_` tag). A payment can carry several.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*whopsdk.ListDisputeAlertsRequestType` — Only alerts of this kind. `early_fraud_warning` for issuer fraud reports, `dispute_alert` for pre-dispute notices, `rapid_dispute_resolution` for Visa RDR cases the network already closed.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of alerts to return (default 20, max 100).
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor; returns alerts after this position.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — The number of alerts to return from the end of the range.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — A cursor; returns alerts before this position.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListDisputeAlertsRequestOrder` — The field to sort alerts by.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListDisputeAlertsRequestDirection` — Sort direction.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Only alerts Whop received before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Only alerts Whop received after this ISO 8601 timestamp.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DisputeAlerts.Retrieve(ID) -> *whopsdk.DisputeAlert</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a single dispute alert or early fraud warning by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveDisputeAlertsRequest{
    ID: "id",
}
client.DisputeAlerts.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The dispute alert ID, prefixed `dspa_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Disputes
<details><summary><code>client.Disputes.List() -> *whopsdk.ListDisputesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the disputes across the accounts you can read.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListDisputesRequest{}
client.Disputes.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — Only disputes filed against this account (`biz_` tag). Omit it to cover every account you can read.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of disputes to return (default 20, max 100).
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor; returns disputes after this position.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — The number of disputes to return from the end of the range.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — A cursor; returns disputes before this position.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListDisputesRequestOrder` — The field to sort disputes by.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListDisputesRequestDirection` — Sort direction.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.ListDisputesRequestStatusItem` — Only disputes in these statuses. Repeat the parameter to pass several — one paginated list covers all of them. Covers both chargebacks and inquiries at each stage. A `needs_response` dispute whose evidence deadline has passed reports and filters as `under_review` instead.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*string` — Only disputes in this three-letter ISO currency.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Only disputes opened before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Only disputes opened after this ISO 8601 timestamp.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Disputes.Summary() -> *whopsdk.SummaryDisputesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Totals up the same disputes the list returns, so you can build status tabs and totals without paging through them.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.SummaryDisputesRequest{}
client.Disputes.Summary(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**groups:** `*whopsdk.SummaryDisputesRequestGroupsItem` — Which breakdowns to return, keyed by these names under `groups`. Repeat the parameter to ask for several; omit it for all of them.
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `*string` — Only disputes filed against this account (`biz_` tag). Omit it to cover every account you can read.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.SummaryDisputesRequestStatusItem` — Only disputes in these statuses. Repeat the parameter to pass several. A `needs_response` dispute whose evidence deadline has passed reports and filters as `under_review` instead.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*string` — Only disputes in this three-letter ISO currency.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Only disputes opened before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Only disputes opened after this ISO 8601 timestamp.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Disputes.Retrieve(ID) -> *whopsdk.Dispute</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a single dispute.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveDisputesRequest{
    ID: "id",
}
client.Disputes.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The dispute ID (`dspt_` tag).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Disputes.Update(ID, request) -> *whopsdk.Dispute</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Edits a dispute's evidence, while it is still editable. Sending it is a separate call.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateDisputesRequest{
    ID: "id",
}
client.Disputes.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The dispute ID (`dspt_` tag).
    
</dd>
</dl>

<dl>
<dd>

**evidence:** `*whopsdk.UpdateDisputesRequestEvidence` — The evidence packet to send to the processor. Only the fields you provide are changed.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Disputes.Submit(ID) -> *whopsdk.Dispute</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Sends a dispute's evidence to the payment processor. This is final — it cannot be edited or sent again.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.SubmitDisputesRequest{
    ID: "id",
}
client.Disputes.Submit(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The dispute ID (`dspt_` tag).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Disputes.SubmitEvidenceDispute(ID) -> *whopsdk.Dispute</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Submit a payment dispute to the payment processor for review. Once submitted, no further edits can be made.

Required permissions:
 - `payment:dispute`
 - `plan:basic:read`
 - `access_pass:basic:read`
 - `company:basic:read`
 - `payment:basic:read`
 - `member:email:read`
 - `member:basic:read`
 - `member:phone:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.SubmitEvidenceDisputeRequest{
    ID: "dspt_xxxxxxxxxxxxx",
}
client.Disputes.SubmitEvidenceDispute(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the dispute to submit to the payment processor for review.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Disputes.UpdateEvidenceDispute(ID, request) -> *whopsdk.Dispute</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a dispute with evidence data to attempt to win the dispute.

Required permissions:
 - `payment:dispute`
 - `plan:basic:read`
 - `access_pass:basic:read`
 - `company:basic:read`
 - `payment:basic:read`
 - `member:email:read`
 - `member:basic:read`
 - `member:phone:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateEvidenceDisputeRequest{
    ID: "dspt_xxxxxxxxxxxxx",
}
client.Disputes.UpdateEvidenceDispute(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the dispute to update.
    
</dd>
</dl>

<dl>
<dd>

**accessActivityLog:** `*string` — An IP access activity log showing the customer used the service.
    
</dd>
</dl>

<dl>
<dd>

**billingAddress:** `*string` — The billing address associated with the customer's payment method.
    
</dd>
</dl>

<dl>
<dd>

**cancellationPolicyAttachment:** `*whopsdk.UpdateEvidenceDisputeRequestCancellationPolicyAttachment` — A file upload containing the company's cancellation policy document.
    
</dd>
</dl>

<dl>
<dd>

**cancellationPolicyDisclosure:** `*string` — The company's cancellation policy text to submit as evidence.
    
</dd>
</dl>

<dl>
<dd>

**customerCommunicationAttachment:** `*whopsdk.UpdateEvidenceDisputeRequestCustomerCommunicationAttachment` — A file upload containing evidence of customer communication. Must be a JPEG, PNG, GIF, or PDF.
    
</dd>
</dl>

<dl>
<dd>

**customerEmailAddress:** `*string` — The email address of the customer associated with the disputed payment.
    
</dd>
</dl>

<dl>
<dd>

**customerName:** `*string` — The full name of the customer associated with the disputed payment.
    
</dd>
</dl>

<dl>
<dd>

**notes:** `*string` — Additional notes or context to submit as part of the dispute evidence.
    
</dd>
</dl>

<dl>
<dd>

**productDescription:** `*string` — A description of the product or service that was provided to the customer.
    
</dd>
</dl>

<dl>
<dd>

**refundPolicyAttachment:** `*whopsdk.UpdateEvidenceDisputeRequestRefundPolicyAttachment` — A file upload containing the company's refund policy document.
    
</dd>
</dl>

<dl>
<dd>

**refundPolicyDisclosure:** `*string` — The company's refund policy text to submit as evidence.
    
</dd>
</dl>

<dl>
<dd>

**refundRefusalExplanation:** `*string` — An explanation of why the refund request was refused.
    
</dd>
</dl>

<dl>
<dd>

**serviceDate:** `*string` — The date when the product or service was delivered to the customer.
    
</dd>
</dl>

<dl>
<dd>

**uncategorizedAttachment:** `*whopsdk.UpdateEvidenceDisputeRequestUncategorizedAttachment` — A file upload for evidence that does not fit into the other categories.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Disputes.UploadEvidence(ID, request) -> *whopsdk.Dispute</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Replaces the full set of uploaded evidence documents on a dispute, beyond the four fixed evidence slots. Upload files through `POST /files` and reference them by `id`, or send the files as multipart file parts to upload and attach in one call. Send every document the packet should carry — up to 10, 10MB each and 25MB in total; an empty list removes them all. Accepted content types: application/pdf, application/json, image/jpeg, image/png, image/webp — any other type is rejected.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UploadEvidenceDisputesRequest{
    ID: "id",
    Documents: []*whopsdk.UploadEvidenceDisputesRequestDocumentsItem{
        &whopsdk.UploadEvidenceDisputesRequestDocumentsItem{
            DocumentType: whopsdk.UploadEvidenceDisputesRequestDocumentsItemDocumentTypeReturnPolicy,
        },
    },
}
client.Disputes.UploadEvidence(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The dispute ID (`dspt_` tag).
    
</dd>
</dl>

<dl>
<dd>

**documents:** `[]*whopsdk.UploadEvidenceDisputesRequestDocumentsItem` — The full set of evidence documents the dispute should carry. Replaces all previously uploaded documents.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## DmChannels
<details><summary><code>client.DmChannels.List() -> *whopsdk.ListDmChannelsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of DM channels for the currently authenticated user, sorted by most recently active.

Required permissions:
 - `dms:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListDmChannelsRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    CompanyID: whopsdk.String(
        "biz_xxxxxxxxxxxxxx",
    ),
}
client.DmChannels.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `*string` — The unique identifier of a company to filter DM channels by. Only returns channels scoped to this company.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DmChannels.Create(request) -> *whopsdk.DmChannel</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new DM channel between two or more users, optionally scoped to a specific company. Returns the existing channel if one already exists.

Required permissions:
 - `dms:channel:manage`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateDmChannelsRequest{
    WithUserIDs: []string{
        "with_user_ids",
    },
}
client.DmChannels.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `*string` — The unique identifier of the company to scope this DM channel to. When set, the channel is visible only within that company context.
    
</dd>
</dl>

<dl>
<dd>

**customName:** `*string` — A custom display name for the DM channel. For example, 'Project Discussion'.
    
</dd>
</dl>

<dl>
<dd>

**notificationsEnabled:** `*bool` — Whether Whop app notifications are enabled for this direct message channel. Webhooks still fire.
    
</dd>
</dl>

<dl>
<dd>

**withUserIDs:** `[]string` — The list of user identifiers to include in the DM channel. Each entry can be an email, username, or user ID (e.g. 'user_xxxxx').
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DmChannels.Retrieve(ID) -> *whopsdk.DmChannel</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing DM channel.

Required permissions (one of):
 - `dms:read`
 - `support_chat:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveDmChannelsRequest{
    ID: "id",
}
client.DmChannels.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the DM channel to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DmChannels.Delete(ID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Permanently delete a DM channel and all of its messages. Only an admin of the channel can perform this action.

Required permissions (one of):
 - `dms:channel:manage`
 - `support_chat:create`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteDmChannelsRequest{
    ID: "id",
}
client.DmChannels.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the DM channel to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DmChannels.Update(ID, request) -> *whopsdk.DmChannel</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the settings of an existing DM channel, such as its display name. Only an admin of the channel can perform this action.

Required permissions (one of):
 - `dms:channel:manage`
 - `support_chat:create`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateDmChannelsRequest{
    ID: "id",
}
client.DmChannels.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the DM channel to update.
    
</dd>
</dl>

<dl>
<dd>

**customName:** `*string` — A new custom display name for the DM channel. For example, 'Project Discussion'.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## DmMembers
<details><summary><code>client.DmMembers.List() -> *whopsdk.ListDmMembersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of members in a specific DM channel, sorted by the date they were added.

Required permissions (one of):
 - `dms:read`
 - `support_chat:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListDmMembersRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    ChannelID: "channel_id",
}
client.DmMembers.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**channelID:** `string` — The unique identifier of the DM channel to list members for.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DmMembers.Create(request) -> *whopsdk.DmMember</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a new user to an existing DM channel. Only an admin of the channel can add members.

Required permissions (one of):
 - `dms:message:manage`
 - `support_chat:message:create`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateDmMembersRequest{
    ChannelID: "channel_id",
    UserID: "user_xxxxxxxxxxxxx",
}
client.DmMembers.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**channelID:** `string` — The unique identifier of the DM channel to add the new member to.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `string` — The unique identifier of the user to add to the DM channel. For example, 'user_xxxxx'.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DmMembers.Retrieve(ID) -> *whopsdk.DmMember</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing DM member.

Required permissions (one of):
 - `dms:read`
 - `support_chat:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveDmMembersRequest{
    ID: "id",
}
client.DmMembers.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the DM channel member to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DmMembers.Delete(ID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove a user from a DM channel. An admin can remove any member, and a member can remove themselves.

Required permissions (one of):
 - `dms:read`
 - `support_chat:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteDmMembersRequest{
    ID: "id",
}
client.DmMembers.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the DM channel member to remove.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DmMembers.Update(ID, request) -> *whopsdk.DmMember</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a DM channel member's settings, such as their notification preferences or membership status.

Required permissions (one of):
 - `dms:read`
 - `support_chat:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateDmMembersRequest{
    ID: "id",
}
client.DmMembers.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the DM channel member to update.
    
</dd>
</dl>

<dl>
<dd>

**notificationPreference:** `*whopsdk.DmsFeedMemberNotificationPreferences` — The notification setting for this member, controlling how they receive alerts for new messages in this channel.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.DmsFeedMemberStatuses` — The membership status for this member in the DM channel.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Entries
<details><summary><code>client.Entries.List() -> *whopsdk.ListEntriesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of waitlist entries for a company, with optional filtering by product, plan, status, and creation date.

Required permissions:
 - `plan:waitlist:read`
 - `member:email:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListEntriesRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    CompanyID: "biz_xxxxxxxxxxxxxx",
    CreatedBefore: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
    CreatedAfter: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
}
client.Entries.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `string` — The unique identifier of the company to list waitlist entries for.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.Direction` 
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.EntriesSortableColumns` 
    
</dd>
</dl>

<dl>
<dd>

**productIDs:** `*string` — Filter entries to only those for specific products.
    
</dd>
</dl>

<dl>
<dd>

**planIDs:** `*string` — Filter entries to only those for specific plans.
    
</dd>
</dl>

<dl>
<dd>

**statuses:** `*whopsdk.EntryStatus` — Filter entries by their current status.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — Only return entries created before this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — Only return entries created after this timestamp.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entries.Retrieve(ID) -> *whopsdk.Entry</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing waitlist entry.

Required permissions:
 - `plan:waitlist:read`
 - `member:email:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveEntriesRequest{
    ID: "entry_xxxxxxxxxxxx",
}
client.Entries.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the waitlist entry to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entries.Approve(ID) -> *whopsdk.ApproveEntriesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Approve a pending waitlist entry, triggering the checkout process to grant the user access to the plan.

Required permissions:
 - `plan:waitlist:manage`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ApproveEntriesRequest{
    ID: "entry_xxxxxxxxxxxx",
}
client.Entries.Approve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the waitlist entry to approve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Entries.Deny(ID) -> *whopsdk.Entry</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deny a pending waitlist entry, preventing the user from gaining access to the plan.

Required permissions:
 - `plan:waitlist:manage`
 - `plan:basic:read`
 - `member:email:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DenyEntriesRequest{
    ID: "entry_xxxxxxxxxxxx",
}
client.Entries.Deny(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the waitlist entry to deny.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Events
<details><summary><code>client.Events.List() -> *whopsdk.ListEventsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists identity-linked events, most recent first by default. Pass identifier for one person's journey, or omit it to list events for an account within an explicit time range. Pass direction=asc to read a journey forwards from where it starts. Events are shaped like the POST /events intake: attribution in context, identity in user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListEventsRequest{}
client.Events.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `*string` — Any hard identifier of the person: a person ID (prsn_*), user ID, email, phone number, or a tracking cookie value (wuid, anonymous ID, fbp/fbc/ttp/ga). Omit to list recent events for the account.
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `*string` — Account ID, prefixed `biz_`. Optional for account API keys; required for credentials that can access multiple accounts.
    
</dd>
</dl>

<dl>
<dd>

**from:** `*time.Time` — Start of the time range as an ISO 8601 timestamp. Required when identifier is omitted.
    
</dd>
</dl>

<dl>
<dd>

**to:** `*time.Time` — End of the time range as an ISO 8601 timestamp. Required when identifier is omitted; otherwise defaults to now.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of events to return.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor for fetching events after a previous page.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — A cursor for fetching events before a later page.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListEventsRequestDirection` — The order events are returned in by time. Defaults to desc (most recent first); asc reads a journey forwards from where it starts. after and before always page forwards and backwards through that order.
    
</dd>
</dl>

<dl>
<dd>

**event:** `*string` — Full event names to filter by, comma-separated (payment.completed, pixel.lead, pixel.page, pixel.custom:<name>) — the same vocabulary the events / people metrics use.
    
</dd>
</dl>

<dl>
<dd>

**source:** `*string` — Canonical source path, exact or with a trailing :* prefix (whop:<campaign>:*, ext:meta:*, referrer:<domain>, direct). Restricts the list to conversion targets attributed to that source — the debuggability twin of a metric cell's source parameter.
    
</dd>
</dl>

<dl>
<dd>

**attributionModel:** `*whopsdk.ListEventsRequestAttributionModel` — Attribution model for the source filter (defaults to last_touch).
    
</dd>
</dl>

<dl>
<dd>

**country:** `*string` — Country codes to filter by, comma-separated.
    
</dd>
</dl>

<dl>
<dd>

**city:** `*string` — Cities to filter by, comma-separated.
    
</dd>
</dl>

<dl>
<dd>

**device:** `*string` — Device families to filter by, comma-separated (e.g. iPhone, Mac).
    
</dd>
</dl>

<dl>
<dd>

**browser:** `*string` — Browser families to filter by, comma-separated (e.g. Chrome, Mobile Safari).
    
</dd>
</dl>

<dl>
<dd>

**os:** `*string` — Operating system families to filter by, comma-separated (e.g. iOS, Windows).
    
</dd>
</dl>

<dl>
<dd>

**utmSource:** `*string` — utm_source values to filter by, comma-separated.
    
</dd>
</dl>

<dl>
<dd>

**hostname:** `*string` — Page hostnames to filter by, comma-separated.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Page paths to filter by, comma-separated.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Events.Create(request) -> *whopsdk.CreateEventsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Tracks a conversion or engagement event for an account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateEventsRequest{
    AccountID: "biz_xxxxxxxxxxxxxx",
    EventName: "coating_deposit_paid",
}
client.Events.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string` — The account to associate with this event.
    
</dd>
</dl>

<dl>
<dd>

**actionSource:** `*whopsdk.CreateEventsRequestActionSource` — Where the event originated.
    
</dd>
</dl>

<dl>
<dd>

**appBuildID:** `*string` — The build of the hosted app that served the page where the event occurred.
    
</dd>
</dl>

<dl>
<dd>

**appID:** `*string` — The hosted app that served the page where the event occurred.
    
</dd>
</dl>

<dl>
<dd>

**context:** `*whopsdk.CreateEventsRequestContext` — Tracking and attribution context.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*whopsdk.CreateEventsRequestCurrency` — ISO 4217 currency code.
    
</dd>
</dl>

<dl>
<dd>

**customName:** `*string` — Custom event name when event_name is 'custom'. Maximum 35 chars for this value.
    
</dd>
</dl>

<dl>
<dd>

**duration:** `*int` — For 'leave' events: milliseconds the visitor spent on the page.
    
</dd>
</dl>

<dl>
<dd>

**eventID:** `*string` — Client-provided identifier for deduplication. Generated if omitted.
    
</dd>
</dl>

<dl>
<dd>

**eventName:** `string` 

The type of event.

Use a standard event (lead, submit_application, contact, complete_registration, schedule, view_content, add_to_cart) or pass your own name directly for a custom event.
    
</dd>
</dl>

<dl>
<dd>

**eventTime:** `*time.Time` — When the event occurred. Defaults to now.
    
</dd>
</dl>

<dl>
<dd>

**planID:** `*string` — The plan associated with the event.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `*string` — The product associated with the event.
    
</dd>
</dl>

<dl>
<dd>

**referrerURL:** `*string` — The referring URL.
    
</dd>
</dl>

<dl>
<dd>

**resumed:** `*bool` — For 'page' events: true when the page was restored from the back/forward cache.
    
</dd>
</dl>

<dl>
<dd>

**source:** `*string` — For 'identify' events: where the identity was captured (url, form, manual, iframe).
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — For 'page' events: the document title.
    
</dd>
</dl>

<dl>
<dd>

**url:** `*string` — The URL where the event occurred.
    
</dd>
</dl>

<dl>
<dd>

**user:** `*whopsdk.CreateEventsRequestUser` — User identity and profile data.
    
</dd>
</dl>

<dl>
<dd>

**value:** `*float64` — Monetary value associated with the event.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Events.Pulse() -> *whopsdk.PulseEventsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a fully anonymized feed of recent platform-wide money movement, most recent first: purchases, affiliate commissions, card and ad spend, app revenue, off-platform sales, wallet deposits, card loads, claimed drops, transfers between accounts, and referral bonuses. Items carry only a `type`, the underlying event name, a USD amount, a coarse location under `user`, and a timestamp coarsened to the start of the minute; missing fields are omitted, not nulled. The payload is identical for every caller; no auth is required.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.PulseEventsRequest{}
client.Events.Pulse(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**event:** `*string` — Filter to one or more types, comma separated — for example `purchase,card_spend`. These are the item's `type`, not its `event_name`: several types share the `ledger_line.created` event name. Omit for every type in the feed. Values outside the feed's own set are rejected.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of events to return.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor for fetching events after a previous page.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — A cursor for fetching events before a later page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Events.ValidatePixel(request) -> *whopsdk.PixelValidation</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Checks whether the Whop pixel is installed for an account. Recent pixel events count as proof on their own, so an account that has sent data lately comes back installed without a `url`. Pass a `url` and events from that page settle it; conversion events are also read across the hostname because they commonly fire on a later confirmation page. If the requested page hasn't sent any events lately, it is fetched and read for the pixel and conversion events wired on it. `installed` is only true when the pixel was actually seen — in the account's events or in the page.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ValidatePixelEventsRequest{}
client.Events.ValidatePixel(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — Account to check. Defaults to the authenticated account.
    
</dd>
</dl>

<dl>
<dd>

**url:** `*string` — A page to read for the pixel, e.g. an ad destination. Omit it to check the account from its events alone.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Experiences
<details><summary><code>client.Experiences.List() -> *whopsdk.ListExperiencesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of experiences belonging to a company, with optional filtering by product and app.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListExperiencesRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    CompanyID: "biz_xxxxxxxxxxxxxx",
    ProductID: whopsdk.String(
        "prod_xxxxxxxxxxxxx",
    ),
    AppID: whopsdk.String(
        "app_xxxxxxxxxxxxxx",
    ),
    CreatedBefore: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
    CreatedAfter: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
}
client.Experiences.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `string` — The unique identifier of the company to list experiences for.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `*string` — Filter to only experiences attached to this product identifier.
    
</dd>
</dl>

<dl>
<dd>

**appID:** `*string` — Filter to only experiences powered by this app identifier.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — Only return experiences created before this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — Only return experiences created after this timestamp.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Experiences.Create(request) -> *whopsdk.Experience</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Required permissions:
 - `experience:create`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateExperiencesRequest{
    AppID: "app_xxxxxxxxxxxxxx",
    CompanyID: "biz_xxxxxxxxxxxxxx",
}
client.Experiences.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**appID:** `string` — The unique identifier of the app that powers this experience.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `string` — The unique identifier of the company to create this experience for.
    
</dd>
</dl>

<dl>
<dd>

**isPublic:** `*bool` — Whether the experience is publicly accessible without a membership.
    
</dd>
</dl>

<dl>
<dd>

**logo:** `*whopsdk.CreateExperiencesRequestLogo` — A logo image displayed alongside the experience name.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The display name of the experience. Defaults to the app's name if not provided.
    
</dd>
</dl>

<dl>
<dd>

**notificationsEnabled:** `*bool` — Whether Whop app notifications are enabled for this experience. Webhooks still fire.
    
</dd>
</dl>

<dl>
<dd>

**sectionID:** `*string` — The unique identifier of the section to place the experience in.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Experiences.Retrieve(ID) -> *whopsdk.Experience</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing experience.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveExperiencesRequest{
    ID: "exp_xxxxxxxxxxxxxx",
}
client.Experiences.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the experience.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Experiences.Delete(ID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Required permissions:
 - `experience:delete`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteExperiencesRequest{
    ID: "exp_xxxxxxxxxxxxxx",
}
client.Experiences.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the experience to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Experiences.Update(ID, request) -> *whopsdk.Experience</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Required permissions:
 - `experience:update`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateExperiencesRequest{
    ID: "exp_xxxxxxxxxxxxxx",
}
client.Experiences.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the experience to update.
    
</dd>
</dl>

<dl>
<dd>

**accessLevel:** `*whopsdk.ExperienceAccessLevels` — The access level of the experience.
    
</dd>
</dl>

<dl>
<dd>

**isPublic:** `*bool` — Whether the experience is publicly accessible without a membership.
    
</dd>
</dl>

<dl>
<dd>

**logo:** `*whopsdk.UpdateExperiencesRequestLogo` — A logo image displayed alongside the experience name.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The display name of the experience.
    
</dd>
</dl>

<dl>
<dd>

**notificationsEnabled:** `*bool` — Whether Whop app notifications are enabled for this experience. Webhooks still fire.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — The position of the experience within its section for display ordering.
    
</dd>
</dl>

<dl>
<dd>

**sectionID:** `*string` — The unique identifier of the section to move the experience into.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Experiences.Attach(ID, request) -> *whopsdk.Experience</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Attach an experience to a product, making it accessible to the product's customers.

Required permissions:
 - `experience:attach`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.AttachExperiencesRequest{
    ID: "exp_xxxxxxxxxxxxxx",
    ProductID: "prod_xxxxxxxxxxxxx",
}
client.Experiences.Attach(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the experience to attach.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `string` — The unique identifier of the product to attach the experience to.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Experiences.Detach(ID, request) -> *whopsdk.Experience</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Detach an experience from a product, removing customer access to it through that product.

Required permissions:
 - `experience:detach`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DetachExperiencesRequest{
    ID: "exp_xxxxxxxxxxxxxx",
    ProductID: "prod_xxxxxxxxxxxxx",
}
client.Experiences.Detach(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the experience to detach.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `string` — The unique identifier of the product to detach the experience from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Experiences.Duplicate(ID, request) -> *whopsdk.Experience</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Duplicates an existing experience. The name will be copied, unless provided. The new experience will be attached to the same products as the original experience.
If duplicating a Forum or Chat experience, the new experience will have the same settings as the original experience, e.g. who can post, who can comment, etc.
No content, e.g. posts, messages, lessons from within the original experience will be copied.


Required permissions:
 - `experience:create`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DuplicateExperiencesRequest{
    ID: "exp_xxxxxxxxxxxxxx",
}
client.Experiences.Duplicate(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the experience to duplicate.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The display name for the duplicated experience. Defaults to the original experience's name.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Exports
<details><summary><code>client.Exports.List() -> *whopsdk.ListExportsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the exports requested for an account, newest first. Only exports of resources the credential is allowed to export are returned.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListExportsRequest{}
client.Exports.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The account to list exports for, prefixed `biz_`. Defaults to the credential's account.
    
</dd>
</dl>

<dl>
<dd>

**resource:** `*whopsdk.ListExportsRequestResource` — Only return exports of this resource.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.ListExportsRequestStatus` — Only return exports in this status.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Only return exports created at or after this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Only return exports created at or before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListExportsRequestOrder` — The field to sort by.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListExportsRequestDirection` — The sort direction.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Exports.Create(request) -> *whopsdk.Export</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Starts an asynchronous export of a resource for an account. Returns the export in `pending`; poll `GET /exports/{id}` until `download_url` is set.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateExportsRequest{
    Resource: whopsdk.CreateExportsRequestResourceAdCampaigns,
}
client.Exports.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The account to export from, prefixed `biz_`. Defaults to the credential's account.
    
</dd>
</dl>

<dl>
<dd>

**columns:** `[]string` — Column keys to include. Empty means all columns for the resource.
    
</dd>
</dl>

<dl>
<dd>

**filters:** `map[string]any` — Resource-specific filters. For native REST resources (`payouts`, `transfers`, `products`) these are the resource's own list query params; for dashboard tables they mirror the dashboard table filters.
    
</dd>
</dl>

<dl>
<dd>

**resource:** `*whopsdk.CreateExportsRequestResource` — The resource to export, e.g. `payouts`, `receipts`, or `members`.
    
</dd>
</dl>

<dl>
<dd>

**timezone:** `*string` — IANA timezone for date columns, e.g. `America/New_York`. Defaults to `UTC`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Exports.Retrieve(ID) -> *whopsdk.Export</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Fetches an export's status and, once complete, its download link.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveExportsRequest{
    ID: "id",
}
client.Exports.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The export ID, prefixed `exprt_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FeeMarkups
<details><summary><code>client.FeeMarkups.List() -> *whopsdk.ListFeeMarkupsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of fee markups configured for a company. If the company is a platform account, returns the platform default markups.

Required permissions:
 - `company:update_child_fees`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListFeeMarkupsRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    CompanyID: "biz_xxxxxxxxxxxxxx",
}
client.FeeMarkups.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `string` — The unique identifier of the company to list fee markups for. Pass a platform account identifier to retrieve platform default markups.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FeeMarkups.Create(request) -> *whopsdk.FeeMarkup</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create or update a fee markup for a company. If a markup for the specified fee type already exists, it will be updated with the new values.

Required permissions:
 - `company:update_child_fees`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateFeeMarkupsRequest{
    CompanyID: "biz_xxxxxxxxxxxxxx",
    FeeType: whopsdk.FeeMarkupTypesCryptoWithdrawalMarkup,
}
client.FeeMarkups.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `string` — The unique identifier of the company to create or update the fee markup for.
    
</dd>
</dl>

<dl>
<dd>

**feeType:** `*whopsdk.FeeMarkupTypes` — The type of fee this markup applies to, such as processing or platform fees.
    
</dd>
</dl>

<dl>
<dd>

**fixedFeeUsd:** `*float64` — The fixed fee amount in USD to charge per transaction. Must be between 0 and 50.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Custom key-value metadata to attach to this fee markup.
    
</dd>
</dl>

<dl>
<dd>

**notes:** `*string` — Internal notes about this fee markup for record-keeping purposes.
    
</dd>
</dl>

<dl>
<dd>

**percentageFee:** `*float64` — The percentage fee to charge per transaction. Must be between 0 and 25.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FeeMarkups.Delete(ID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a fee markup configuration for a company. This removes the custom fee override and reverts to the parent company's default fees.

Required permissions:
 - `company:update_child_fees`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteFeeMarkupsRequest{
    ID: "id",
}
client.FeeMarkups.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the fee markup to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Files
<details><summary><code>client.Files.List() -> *whopsdk.ListFilesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the files with the given IDs, newest first — fetch a batch in one request instead of retrieving each file individually. Only files you created are returned; IDs that do not exist, or that another credential created, are omitted. A request for up to 100 IDs answers in a single page by default; a larger batch pages at up to 100 files per response — follow `page_info` with the same `file_ids` to walk the rest.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListFilesRequest{
    FileIDs: []*string{
        whopsdk.String(
            "file_xxxxxxxxxxxxx",
        ),
    },
}
client.Files.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fileIDs:** `*string` — The files to return, each prefixed `file_`. Repeat the parameter to pass several, up to 250 per request. Batches of up to 100 answer in one page by default; larger batches page at up to 100 per response.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListFilesRequestOrder` — The field to sort by.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListFilesRequestDirection` — The sort direction.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of files to return.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor; returns files after this position.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — The number of files to return from the end of the range.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — A cursor; returns files before this position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Files.Create(request) -> *whopsdk.File</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a file and returns a presigned destination to upload its bytes to. PUT the bytes to `upload_url` (single-part), or to each of `multipart_upload_urls` and then call Complete File Multipart Upload. Once the bytes land the file becomes `ready`, and its ID can be attached wherever a file is accepted — account legal documents, dispute evidence documents.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateFilesRequest{
    Filename: "terms.pdf",
}
client.Files.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**byteSize:** `*int` — The file's size in bytes. Required when `multipart` is `true`. Multipart uploads support at most 10,000 parts of 5MB each (about 50 GB).
    
</dd>
</dl>

<dl>
<dd>

**filename:** `string` — The name of the file including its extension, e.g. `terms.pdf`.
    
</dd>
</dl>

<dl>
<dd>

**multipart:** `*bool` — Upload the file in 5MB parts. Required for files larger than 5GB; useful above ~100MB. The file must be larger than 5MB.
    
</dd>
</dl>

<dl>
<dd>

**visibility:** `*whopsdk.CreateFilesRequestVisibility` — `public` files are served via an unsigned CDN URL — use for assets anyone may see. `private` files are served via a signed, expiring URL — use for sensitive documents. Defaults to `private`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Files.Retrieve(ID) -> *whopsdk.File</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a file you uploaded — poll it after uploading the bytes to see `upload_status` become `ready`. Only the creator can retrieve a file this way; a file attached to another resource is read through that resource.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveFilesRequest{
    ID: "id",
}
client.Files.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the file, prefixed `file_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Files.Complete(ID, request) -> *whopsdk.File</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Assembles the parts of a multipart upload after every part has been PUT to its presigned URL. Pass the `multipart_upload_id` from Create File and each part's `ETag` response header.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CompleteFilesRequest{
    ID: "id",
    MultipartParts: []*whopsdk.CompleteFilesRequestMultipartPartsItem{
        &whopsdk.CompleteFilesRequestMultipartPartsItem{
            Etag: "etag-1",
            PartNumber: 1,
        },
    },
    MultipartUploadID: "upload-id",
}
client.Files.Complete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the file, prefixed `file_`.
    
</dd>
</dl>

<dl>
<dd>

**multipartParts:** `[]*whopsdk.CompleteFilesRequestMultipartPartsItem` — Every uploaded part, in order.
    
</dd>
</dl>

<dl>
<dd>

**multipartUploadID:** `string` — The ID of the multipart upload, returned by Create File.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FinancialActivity
<details><summary><code>client.FinancialActivity.List() -> *whopsdk.ListFinancialActivityResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an account's or user's activity feed: every movement of money in or out.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListFinancialActivityRequest{}
client.FinancialActivity.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The owning account ID (a biz_ identifier). Provide this or user_id.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — The owning user ID (a user_ identifier). Provide this or account_id.
    
</dd>
</dl>

<dl>
<dd>

**includeOwnedAccounts:** `*bool` — When true, aggregates the authenticated user's personal ledger with the businesses they own (owner role with balance read) into one feed. Requires user_id to be the authenticated user; cannot be combined with account_id or the settlement-date filters. Each returned row includes the owning `account`.
    
</dd>
</dl>

<dl>
<dd>

**includeResource:** `*bool` — Whether to include the `resource` field in the response or not. Consider passing `false` if you need a fast response without as many rich details.
    
</dd>
</dl>

<dl>
<dd>

**lineTypes:** `*whopsdk.ListFinancialActivityRequestLineTypesItem` — Optional ledger line categories to include. Some categories (for example `onchain_deposit`, which covers inbound crypto deposits such as MoonPay onramps) are only returned when explicitly requested here.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListFinancialActivityRequestDirection` — Optional direction filter. `money_in` returns positive activity and `money_out` returns negative activity.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*string` — Optional currency code filter, for example `usd`.
    
</dd>
</dl>

<dl>
<dd>

**postedAfter:** `*time.Time` — Only include rows posted after this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**postedBefore:** `*time.Time` — Only include rows posted before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**availableAfter:** `*time.Time` — Only include rows whose funds became withdrawable on or after this `YYYY-MM-DD` settlement date (UTC), distinct from posted_at. Requires currency.
    
</dd>
</dl>

<dl>
<dd>

**availableBefore:** `*time.Time` — Only include rows whose funds became withdrawable on or before this `YYYY-MM-DD` settlement date (UTC). Set equal to available_after for a single day. Requires currency.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of rows to return.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Cursor returned by the previous page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FinancialReports
<details><summary><code>client.FinancialReports.Retrieve() -> *whopsdk.RetrieveFinancialReportsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a financial report — balance activity, income statement, or balance summary — for an account over a date range.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveFinancialReportsRequest{
    AccountID: "account_id",
    ReportType: whopsdk.RetrieveFinancialReportsRequestReportTypeBalanceSummary,
}
client.FinancialReports.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string` — The owning account ID (a biz_ identifier), or `global` for a platform-wide report across all ledger accounts (requires internal admin access).
    
</dd>
</dl>

<dl>
<dd>

**reportType:** `*whopsdk.RetrieveFinancialReportsRequestReportType` — The type of financial report to generate.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*string` — Filter rows to this currency, for example `usd`. Defaults to `usd` unless `in_currency` is provided.
    
</dd>
</dl>

<dl>
<dd>

**inCurrency:** `*string` — Aggregate all activity into this display currency via FX conversion.
    
</dd>
</dl>

<dl>
<dd>

**fromDate:** `*string` — Start of the report window as an ISO 8601 timestamp (UTC). Required for platform-wide (global) reports.
    
</dd>
</dl>

<dl>
<dd>

**toDate:** `*string` — End of the report window as an ISO 8601 timestamp (UTC). Required for platform-wide (global) reports.
    
</dd>
</dl>

<dl>
<dd>

**groupBy:** `*whopsdk.RetrieveFinancialReportsRequestGroupBy` — Grouping granularity for report rows.
    
</dd>
</dl>

<dl>
<dd>

**timezone:** `*string` — IANA timezone (for example `America/New_York`) used to bucket report periods and to interpret calendar-day boundaries for balance snapshots. Defaults to UTC. from_date/to_date remain exact instants regardless of this setting.
    
</dd>
</dl>

<dl>
<dd>

**lineTypes:** `*whopsdk.RetrieveFinancialReportsRequestLineTypesItem` — Account-level balance activity only: ledger line categories to include.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.RetrieveFinancialReportsRequestDirection` — Account-level balance activity only: include money moving in or money moving out.
    
</dd>
</dl>

<dl>
<dd>

**cumulative:** `*bool` — Platform-wide (global) reports only: when true, return cumulative balances as of to_date (all history, no lower bound) instead of activity within the period.
    
</dd>
</dl>

<dl>
<dd>

**scopeAccountID:** `*string` — Platform-wide (global) reports only: narrow the report to ledger lines on the ledger account owned by this account ID (a biz_ identifier). Ignored unless account_id is `global`.
    
</dd>
</dl>

<dl>
<dd>

**includePaymentFeeBreakdown:** `*bool` — Balance activity only: include payment costs grouped by payment method and provider.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ForumPosts
<details><summary><code>client.ForumPosts.List() -> *whopsdk.ListForumPostsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of forum posts within a specific experience, with optional filtering by parent post or pinned status.

Required permissions:
 - `forum:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListForumPostsRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    ExperienceID: "exp_xxxxxxxxxxxxxx",
}
client.ForumPosts.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**experienceID:** `string` — The unique identifier of the experience to list forum posts for.
    
</dd>
</dl>

<dl>
<dd>

**includeBountyAnchors:** `*bool` — Whether to include top-level bounty discussion anchors as rich forum items.
    
</dd>
</dl>

<dl>
<dd>

**parentID:** `*string` — The unique identifier of a parent post to list comments for. When set, returns replies to that post.
    
</dd>
</dl>

<dl>
<dd>

**pinned:** `*bool` — Whether to filter for only pinned posts. Set to true to return only pinned posts.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ForumPosts.Create(request) -> *whopsdk.ForumPost</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new forum post or comment within an experience. Supports text content, attachments, polls, paywalling, and pinning. Pass experience_id 'public' with a company_id to post to a company's public forum.

Required permissions:
 - `forum:post:create`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateForumPostsRequest{
    ExperienceID: "exp_xxxxxxxxxxxxxx",
}
client.ForumPosts.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**attachments:** `[]*whopsdk.CreateForumPostsRequestAttachmentsItem` — A list of file attachments to include with the post, such as images or videos.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `*string` — The unique identifier of the company whose public forum to post in. Required when experience_id is 'public'. For example, 'biz_xxxxx'.
    
</dd>
</dl>

<dl>
<dd>

**content:** `*string` — The main body of the post in Markdown format. For example, 'Check out this **update**'. Hidden if the post is paywalled and the viewer has not purchased access.
    
</dd>
</dl>

<dl>
<dd>

**experienceID:** `string` — The unique identifier of the experience to create this post in. For example, 'exp_xxxxx'. Pass 'public' along with company_id to automatically use the company's public forum.
    
</dd>
</dl>

<dl>
<dd>

**isMention:** `*bool` — Whether to send this post as a mention notification to all users in the experience who have mentions enabled.
    
</dd>
</dl>

<dl>
<dd>

**parentID:** `*string` — The unique identifier of the parent post to comment on. Omit this field to create a top-level post.
    
</dd>
</dl>

<dl>
<dd>

**paywallAmount:** `*float64` — The price to unlock this post in the specified paywall currency. For example, 5.00 for $5.00. When set, users must purchase access to view the post content.
    
</dd>
</dl>

<dl>
<dd>

**paywallCurrency:** `*whopsdk.Currencies` — The currency for the paywall price on this post. When set along with paywall_amount, users must purchase access to view the post content.
    
</dd>
</dl>

<dl>
<dd>

**pinned:** `*bool` — Whether this post should be pinned to the top of the forum.
    
</dd>
</dl>

<dl>
<dd>

**poll:** `*whopsdk.CreateForumPostsRequestPoll` — A poll to attach to this post, allowing members to vote on options.
    
</dd>
</dl>

<dl>
<dd>

**richContent:** `*string` — The rich content of the post in Tiptap JSON format. When provided, takes priority over the markdown content field for rendering.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The title of the post, displayed prominently at the top. Required for paywalled posts as it remains visible to non-purchasers.
    
</dd>
</dl>

<dl>
<dd>

**visibility:** `*whopsdk.ForumPostVisibilityTypes` — Controls who can see this forum post, such as members only or public.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ForumPosts.Retrieve(ID) -> *whopsdk.ForumPost</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing forum post.

Required permissions:
 - `forum:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveForumPostsRequest{
    ID: "id",
}
client.ForumPosts.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the forum post to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ForumPosts.Update(ID, request) -> *whopsdk.ForumPost</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Edit the content, attachments, pinned status, or visibility of an existing forum post or comment.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateForumPostsRequest{
    ID: "id",
}
client.ForumPosts.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the forum post to update.
    
</dd>
</dl>

<dl>
<dd>

**attachments:** `[]*whopsdk.UpdateForumPostsRequestAttachmentsItem` — A replacement list of file attachments for this post, such as images or videos.
    
</dd>
</dl>

<dl>
<dd>

**content:** `*string` — The updated body of the post in Markdown format. For example, 'Check out this **update**'. Hidden if the post is paywalled and the viewer has not purchased access.
    
</dd>
</dl>

<dl>
<dd>

**isPinned:** `*bool` — Whether this post should be pinned to the top of the forum. Only top-level posts can be pinned, not comments.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The updated title of the post, displayed prominently at the top. Required for paywalled posts as it remains visible to non-purchasers.
    
</dd>
</dl>

<dl>
<dd>

**visibility:** `*whopsdk.ForumPostVisibilityTypes` — Controls who can see this forum post, such as members only or public.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Forums
<details><summary><code>client.Forums.List() -> *whopsdk.ListForumsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of forums within a specific company, with optional filtering by product.

Required permissions:
 - `forum:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListForumsRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    CompanyID: "biz_xxxxxxxxxxxxxx",
    ProductID: whopsdk.String(
        "prod_xxxxxxxxxxxxx",
    ),
}
client.Forums.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `string` — The unique identifier of the company to list forums for.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `*string` — The unique identifier of a product to filter by. When set, only forums connected to this product are returned.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Forums.Retrieve(ID) -> *whopsdk.Forum</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing forum.

Required permissions:
 - `forum:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveForumsRequest{
    ID: "id",
}
client.Forums.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the forum or experience to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Forums.Update(ID, request) -> *whopsdk.Forum</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update moderation and notification settings for a forum, such as who can post, who can comment, and email notification preferences.

Required permissions:
 - `forum:moderate`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateForumsRequest{
    ID: "id",
}
client.Forums.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the forum to update. Accepts either an experience ID (e.g. 'exp_xxxxx') or a forum ID.
    
</dd>
</dl>

<dl>
<dd>

**bannedWords:** `[]string` — A list of words that are automatically blocked from posts in this forum. For example, ['spam', 'scam'].
    
</dd>
</dl>

<dl>
<dd>

**bannerImage:** `*whopsdk.UpdateForumsRequestBannerImage` — The banner image displayed at the top of the forum page. Pass null to remove the existing banner.
    
</dd>
</dl>

<dl>
<dd>

**emailNotificationPreference:** `*whopsdk.ForumEmailNotificationPreferences` — Controls how email notifications are sent to members when new posts are created in this forum.
    
</dd>
</dl>

<dl>
<dd>

**whoCanComment:** `*whopsdk.ForumWhoCanCommentTypes` — Controls which roles are allowed to comment on posts in this forum.
    
</dd>
</dl>

<dl>
<dd>

**whoCanPost:** `*whopsdk.ForumWhoCanPostTypes` — Controls which roles are allowed to create new posts in this forum.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## IdentityProfiles
<details><summary><code>client.IdentityProfiles.ListIdentityProfile() -> *whopsdk.ListIdentityProfileResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of identity profiles. When company_id is provided, lists IPs currently linked to that company's ledger. When omitted, lists IPs linked to any ledger the actor can read (including child companies under a parent).

Required permissions:
 - `identity:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListIdentityProfileRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    CompanyID: whopsdk.String(
        "biz_xxxxxxxxxxxxxx",
    ),
}
client.IdentityProfiles.ListIdentityProfile(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `*string` — The unique identifier of the company to filter to. When omitted, returns IPs across all ledgers the actor can read.
    
</dd>
</dl>

<dl>
<dd>

**profileType:** `*whopsdk.IdentityProfileKinds` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.IdentityProfileStatuses` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.IdentityProfiles.RetrieveIdentityProfile(ID) -> *whopsdk.IdentityProfile</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing identity profile.

Required permissions:
 - `identity:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveIdentityProfileRequest{
    ID: "idpf_xxxxxxxxxxxxx",
}
client.IdentityProfiles.RetrieveIdentityProfile(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the identity profile (idpf_xxx).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.IdentityProfiles.UnlinkIdentityProfile(ID) -> *whopsdk.IdentityProfile</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Unlinks an IdentityProfile from a LedgerAccount (flips the matching link to is_current=false).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UnlinkIdentityProfileRequest{
    ID: "idpf_xxxxxxxxxxxxx",
    LedgerAccountID: "ldgr_xxxxxxxxxxxxx",
}
client.IdentityProfiles.UnlinkIdentityProfile(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the IdentityProfile to unlink.
    
</dd>
</dl>

<dl>
<dd>

**ledgerAccountID:** `string` — The ID of the LedgerAccount to unlink the identity profile from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.IdentityProfiles.ListVerificationsIdentityProfile(ID) -> *whopsdk.ListVerificationsIdentityProfileResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of verifications attached to an identity profile, ordered by most recent first.

Required permissions:
 - `identity:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListVerificationsIdentityProfileRequest{
    ID: "idpf_xxxxxxxxxxxxx",
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
}
client.IdentityProfiles.ListVerificationsIdentityProfile(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the identity profile (idpf_xxx).
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Invoices
<details><summary><code>client.Invoices.List() -> *whopsdk.ListInvoicesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of invoices for a company, with optional filtering by product, status, collection method, and creation date.

Required permissions:
 - `invoice:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListInvoicesRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    CompanyID: whopsdk.String(
        "biz_xxxxxxxxxxxxxx",
    ),
    CreatedBefore: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
    CreatedAfter: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
}
client.Invoices.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `*string` — The unique identifier of the company to list invoices for.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.Direction` 
    
</dd>
</dl>

<dl>
<dd>

**productIDs:** `*string` — Filter invoices to only those associated with these specific product identifiers.
    
</dd>
</dl>

<dl>
<dd>

**collectionMethods:** `*whopsdk.InvoiceCollectionMethods` — Filter invoices by their collection method.
    
</dd>
</dl>

<dl>
<dd>

**statuses:** `*whopsdk.InvoiceStatuses` — Filter invoices by their current status.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.InvoicesSortableColumns` 
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — Only return invoices created before this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — Only return invoices created after this timestamp.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Invoices.Create(request) -> *whopsdk.Invoice</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create an invoice for a customer. The invoice can be charged automatically using a stored payment method, or sent to the customer for manual payment.

Required permissions:
 - `invoice:create`
 - `member:email:read`
 - `member:basic:read`
 - `payment:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateInvoicesRequest{
    CreateInvoicesRequestProduct: &whopsdk.CreateInvoicesRequestProduct{
        CollectionMethod: whopsdk.InvoiceCollectionMethodsSendInvoice,
        CompanyID: "biz_xxxxxxxxxxxxxx",
        Plan: &whopsdk.CreateInvoicesRequestProductPlan{},
        Product: &whopsdk.CreateInvoicesRequestProductProduct{
            Title: "title",
        },
    },
}
client.Invoices.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*whopsdk.CreateInvoicesRequest` — Parameters for CreateInvoice
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Invoices.Retrieve(ID) -> *whopsdk.Invoice</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing invoice.

Required permissions:
 - `invoice:basic:read`
 - `member:email:read`
 - `member:basic:read`
 - `payment:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveInvoicesRequest{
    ID: "inv_xxxxxxxxxxxxxx",
}
client.Invoices.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the invoice, or a secure token.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Invoices.Delete(ID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a draft invoice.

Required permissions:
 - `invoice:update`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteInvoicesRequest{
    ID: "inv_xxxxxxxxxxxxxx",
}
client.Invoices.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the draft invoice to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Invoices.Update(ID, request) -> *whopsdk.Invoice</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a draft invoice's details.

Required permissions:
 - `invoice:update`
 - `member:email:read`
 - `member:basic:read`
 - `payment:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateInvoicesRequest{
    ID: "inv_xxxxxxxxxxxxxx",
}
client.Invoices.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the invoice to update.
    
</dd>
</dl>

<dl>
<dd>

**automaticallyFinalizesAt:** `*time.Time` — The date and time when the invoice will be automatically finalized. For charge_automatically, triggers an automatic charge. For send_invoice, sends the invoice email at the specified time.
    
</dd>
</dl>

<dl>
<dd>

**billingAddress:** `*whopsdk.UpdateInvoicesRequestBillingAddress` — Inline billing address to create or update a mailing address for this invoice.
    
</dd>
</dl>

<dl>
<dd>

**chargeBuyerFee:** `*bool` — Whether to charge the customer a buyer fee on this invoice.
    
</dd>
</dl>

<dl>
<dd>

**collectionMethod:** `*whopsdk.InvoiceCollectionMethods` — How the invoice should be collected.
    
</dd>
</dl>

<dl>
<dd>

**customerName:** `*string` — The name of the customer.
    
</dd>
</dl>

<dl>
<dd>

**dueDate:** `*time.Time` — The date by which the invoice must be paid.
    
</dd>
</dl>

<dl>
<dd>

**emailAddress:** `*string` — The email address of the customer.
    
</dd>
</dl>

<dl>
<dd>

**lineItems:** `[]*whopsdk.UpdateInvoicesRequestLineItemsItem` — Line items that break down the invoice total. When provided, the sum of (quantity * unit_price) for all items must equal the plan price. Individual items may be negative to represent a credit, as long as the sum is not negative and clears the currency's minimum charge. Pass an empty list to remove the breakdown.
    
</dd>
</dl>

<dl>
<dd>

**mailingAddressID:** `*string` — The unique identifier of an existing mailing address to attach.
    
</dd>
</dl>

<dl>
<dd>

**memberID:** `*string` — The unique identifier of a member to assign as the customer.
    
</dd>
</dl>

<dl>
<dd>

**paymentMethodID:** `*string` — The unique identifier of the payment method to charge.
    
</dd>
</dl>

<dl>
<dd>

**plan:** `*whopsdk.UpdateInvoicesRequestPlan` — Updated plan attributes.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `*string` — The unique identifier of an existing product to attach to this invoice. Only allowed while the invoice is still a draft.
    
</dd>
</dl>

<dl>
<dd>

**subscriptionBillingAnchorAt:** `*time.Time` — The date that defines when the subscription billing cycle should start.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Invoices.MarkPaid(ID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Mark an open invoice as paid when payment was collected outside of Whop.

Required permissions:
 - `invoice:update`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.MarkPaidInvoicesRequest{
    ID: "inv_xxxxxxxxxxxxxx",
}
client.Invoices.MarkPaid(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the invoice to mark as paid.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Invoices.MarkUncollectible(ID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Mark an open invoice as uncollectible when payment is not expected.

Required permissions:
 - `invoice:update`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.MarkUncollectibleInvoicesRequest{
    ID: "inv_xxxxxxxxxxxxxx",
}
client.Invoices.MarkUncollectible(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the invoice to mark as uncollectible.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Invoices.Resend(ID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Resend the notification email for an existing invoice to the customer.

Required permissions:
 - `invoice:update`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ResendInvoicesRequest{
    ID: "inv_xxxxxxxxxxxxxx",
}
client.Invoices.Resend(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the invoice to resend.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Invoices.Void(ID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Void an open invoice so it can no longer be paid. Voiding is permanent and cannot be undone.

Required permissions:
 - `invoice:update`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.VoidInvoicesRequest{
    ID: "inv_xxxxxxxxxxxxxx",
}
client.Invoices.Void(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the invoice to void.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Leads
<details><summary><code>client.Leads.List() -> *whopsdk.ListLeadsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of leads for a company, with optional filtering by product and creation date.

Required permissions:
 - `lead:basic:read`
 - `member:email:read`
 - `access_pass:basic:read`
 - `member:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListLeadsRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    CompanyID: "biz_xxxxxxxxxxxxxx",
    CreatedAfter: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
    CreatedBefore: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
}
client.Leads.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `string` — The unique identifier of the company to list leads for.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — Only return leads created after this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — Only return leads created before this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**productIDs:** `*string` — Filter leads to only those associated with these specific product identifiers.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Leads.Create(request) -> *whopsdk.Lead</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Record a new lead for a company, capturing a potential customer's interest in a specific product.

Required permissions:
 - `lead:manage`
 - `member:email:read`
 - `access_pass:basic:read`
 - `member:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateLeadsRequest{
    CompanyID: "biz_xxxxxxxxxxxxxx",
}
client.Leads.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `string` — The unique identifier of the company to create the lead for, starting with 'biz_'.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — A JSON object of custom metadata to attach to the lead for tracking purposes.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `*string` — The unique identifier of the product the lead is interested in, starting with 'prod_'.
    
</dd>
</dl>

<dl>
<dd>

**referrer:** `*string` — The referral URL that brought the lead to the company, such as 'https://example.com/landing'.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — The unique identifier of the user to record as the lead. If authenticated as a user, that user is used automatically.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Leads.Retrieve(ID) -> *whopsdk.Lead</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing lead.

Required permissions:
 - `lead:basic:read`
 - `member:email:read`
 - `access_pass:basic:read`
 - `member:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveLeadsRequest{
    ID: "lead_xxxxxxxxxxxxx",
}
client.Leads.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the lead to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Leads.Update(ID, request) -> *whopsdk.Lead</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the metadata or referrer information on an existing lead record.

Required permissions:
 - `lead:manage`
 - `member:email:read`
 - `access_pass:basic:read`
 - `member:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateLeadsRequest{
    ID: "lead_xxxxxxxxxxxxx",
}
client.Leads.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the lead to update, starting with 'lead_'.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — A JSON object of custom metadata to set on the lead, replacing any existing metadata.
    
</dd>
</dl>

<dl>
<dd>

**referrer:** `*string` — The updated referral URL for the lead, such as 'https://example.com/landing'.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## LedgerAccounts
<details><summary><code>client.LedgerAccounts.Retrieve(ID) -> *whopsdk.LedgerAccount</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing ledger account.

Required permissions:
 - `company:balance:read`
 - `payout:account:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveLedgerAccountsRequest{
    ID: "ldgr_xxxxxxxxxxxxx",
}
client.LedgerAccounts.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The identifier to look up the ledger account. Accepts a user ID ('user_xxx'), company ID ('biz_xxx'), or ledger account ID ('ldgr_xxx').
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Media
<details><summary><code>client.Media.Generate(request) -> *whopsdk.MediaAsset</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Starts an AI media generation job billed from the account's balance. Generation is asynchronous — poll `GET /media/{id}` until the asset is `ready`, then use `file.id` anywhere attachments are accepted.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.GenerateMediaRequest{
    Prompt: "A 9:16 product showcase of a cordless power scrubber",
    Type: whopsdk.GenerateMediaRequestTypeVideo,
}
client.Media.Generate(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — Account ID, prefixed `biz_`. Defaults to the account the API key belongs to.
    
</dd>
</dl>

<dl>
<dd>

**durationSeconds:** `*int` — Video length in seconds. Video only; defaults to 5.
    
</dd>
</dl>

<dl>
<dd>

**prompt:** `string` — What to generate. Up to 2,000 characters.
    
</dd>
</dl>

<dl>
<dd>

**referenceMedia:** `[]string` — Optional reference image file IDs (`file_` prefixed), up to 4. For video, a single reference seeds the opening frame; multiple references guide subject and style instead.
    
</dd>
</dl>

<dl>
<dd>

**resolution:** `*whopsdk.GenerateMediaRequestResolution` — Video resolution. Video only; defaults to `1080p`. `1080p` is not supported by Seedance 2.0 Fast or Mini; `4k` is only supported by Seedance 2.0.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*whopsdk.GenerateMediaRequestType` — The kind of media to generate.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Media.Retrieve(ID) -> *whopsdk.MediaAsset</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a media asset by ID. Poll this while the asset is `processing`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveMediaRequest{
    ID: "id",
}
client.Media.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Media asset ID, prefixed `media_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Members
<details><summary><code>client.Members.List() -> *whopsdk.ListMembersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the members of an account. A member is one buyer's relationship with the account, regardless of how many memberships they hold.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListMembersRequest{}
client.Members.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The account to list members for (`biz_` tag). Defaults to the account the credential acts as.
    
</dd>
</dl>

<dl>
<dd>

**accessLevel:** `*whopsdk.ListMembersRequestAccessLevel` — Filter by what the member can reach on the account.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.ListMembersRequestStatus` — Filter by whether the member is still part of the account.
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` — Search members by name or username. An exact email address also matches when the credential holds the member:email:read scope.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Only members who joined after this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Only members who joined before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListMembersRequestOrder` — Sort field.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListMembersRequestDirection` — Sort direction.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Number of members to return from the start of the window.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor to paginate forwards from.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Number of members to return from the end of the window.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Cursor to paginate backwards from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Members.Retrieve(ID) -> *whopsdk.Member</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a member by ID. Accessible to the account and to the member's own user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveMembersRequest{
    ID: "id",
}
client.Members.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Member ID (`mber_` tag).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Memberships
<details><summary><code>client.Memberships.List() -> *whopsdk.ListMembershipsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists every membership the caller can read: an account API key its account's; a user credential their own plus those of every account they manage. `account_id` and `user_id` only narrow that list — values outside the caller's reach return fewer results, not an error.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListMembershipsRequest{}
client.Memberships.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — Narrow to one account (`biz_` tag). With read access to the account this lists all of its memberships; without, only the caller's own memberships in it.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — Narrow to one user's memberships (`user_` tag, or `me` for the caller). A user outside the caller's visible set returns an empty list.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.ListMembershipsRequestStatus` — Filter by billing state. `canceling` matches active memberships set to cancel at period end; `paused` matches memberships with payment collection paused.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `*string` — Filter to memberships of this product (`prod_` tag). Repeat as product_ids[] for several.
    
</dd>
</dl>

<dl>
<dd>

**planID:** `*string` — Filter to memberships of this plan (`plan_` tag). Repeat as plan_ids[] for several.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Only memberships created after this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Only memberships created before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListMembershipsRequestOrder` — Sort field.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListMembershipsRequestDirection` — Sort direction.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Number of memberships to return from the start of the window.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor to paginate forwards from.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Number of memberships to return from the end of the window.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Cursor to paginate backwards from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Memberships.Invite(request) -> *whopsdk.InviteMembershipsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Sends an email inviting one recipient to join the account through a free plan. Identify the recipient by exactly one of `user_id` or `email`. The invitation is bound to that recipient; after signing in, accepting it immediately grants the membership without checkout. This Experimental endpoint is available only to accounts enabled for membership invitations.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.InviteMembershipsRequestBody{
    InviteMembershipsRequestBodyUserID: &whopsdk.InviteMembershipsRequestBodyUserID{
        PlanID: "plan_xxxxxxxxxxxxxx",
        UserID: "user_xxxxxxxxxxxxxx",
    },
}
client.Memberships.Invite(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*whopsdk.InviteMembershipsRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Memberships.Retrieve(ID) -> *whopsdk.Membership</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a membership by ID or license key. Accessible to the account and to the membership's own user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveMembershipsRequest{
    ID: "id",
}
client.Memberships.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Membership ID (`mem_` tag), or a software license key.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Memberships.Update(ID, request) -> *whopsdk.Membership</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a membership: merge metadata key-value pairs, or toggle `cancel_at_period_end` — `true` schedules the cancellation for the end of the current billing period, `false` reverses a pending one.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateMembershipsRequest{
    ID: "id",
}
client.Memberships.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Membership ID (`mem_` tag), or a software license key.
    
</dd>
</dl>

<dl>
<dd>

**cancelAtPeriodEnd:** `*bool` — `true` cancels at the end of the current billing period (the customer keeps access until then); `false` reverses a pending cancellation.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Key-value pairs to merge into the membership's metadata. Pass an empty object to clear it.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Memberships.AddFreeDaysMembership(ID, request) -> *whopsdk.Membership</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add free days to extend a membership's current billing period, expiration date, or Stripe trial.

Required permissions:
 - `member:manage`
 - `member:email:read`
 - `member:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.AddFreeDaysMembershipRequest{
    ID: "mem_xxxxxxxxxxxxxx",
    FreeDays: 42,
}
client.Memberships.AddFreeDaysMembership(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the membership.
    
</dd>
</dl>

<dl>
<dd>

**freeDays:** `int` — The number of free days to add (1-1095). Extends the billing period, expiration date, or Stripe trial depending on plan type.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Memberships.Cancel(ID, request) -> *whopsdk.Membership</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancels a membership. Pass `cancel_at_period_end: true` to stop auto-renewal and keep access until the current billing period ends. Omit it (or pass `false`) to revoke access immediately. Buyers cannot cancel buy-now-pay-later (`splitit`, `sezzle`) or non-trial split-pay memberships.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CancelMembershipsRequest{
    ID: "id",
}
client.Memberships.Cancel(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Membership ID (`mem_` tag).
    
</dd>
</dl>

<dl>
<dd>

**cancelAtPeriodEnd:** `*bool` — `true` stops auto-renewal and keeps access until the current billing period ends. Omit or `false` revokes access immediately.
    
</dd>
</dl>

<dl>
<dd>

**reason:** `*string` — Free-form note recording why the membership was canceled.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Memberships.Extend(ID, request) -> *whopsdk.Membership</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Adds free days to a membership, extending its current billing period, expiration date, or trial depending on the plan type.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ExtendMembershipsRequest{
    ID: "id",
    Days: 7,
}
client.Memberships.Extend(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Membership ID (`mem_` tag).
    
</dd>
</dl>

<dl>
<dd>

**days:** `int` — Number of free days to add (1-1095).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Memberships.Pause(ID, request) -> *whopsdk.Membership</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Pauses a membership's recurring payment collection. The customer keeps access but is not charged until the membership is resumed.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.PauseMembershipsRequest{
    ID: "id",
}
client.Memberships.Pause(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Membership ID (`mem_` tag).
    
</dd>
</dl>

<dl>
<dd>

**until:** `*string` — ISO 8601 time to automatically resume payment collection. Must be in the future; only supported for memberships billed by Whop.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Memberships.Resume(ID) -> *whopsdk.Membership</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Resumes a previously paused membership's recurring payment collection. Billing resumes on the next cycle.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ResumeMembershipsRequest{
    ID: "id",
}
client.Memberships.Resume(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Membership ID (`mem_` tag).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Memberships.ResyncAccessMembership(ID) -> *whopsdk.Membership</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Re-run access fulfillment for a membership. Recomputes the member's content access on Whop, re-validates their Discord link (re-adding them to the server and re-assigning roles if needed), and re-fulfills TradingView indicator access. Telegram access is invite-based and cannot be resynced here. The outcome is written to the membership's logs.

Required permissions:
 - `membership:resync_access`
 - `member:email:read`
 - `member:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ResyncAccessMembershipRequest{
    ID: "mem_xxxxxxxxxxxxxx",
}
client.Memberships.ResyncAccessMembership(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the membership to resync access for.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Memberships.Transfer(ID) -> *whopsdk.TransferMembershipsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a one-use transfer URL for a membership. Opening the URL while logged into a different Whop account claims the membership onto that account. The membership's buyer can generate a link for their own membership with `membership:transfer` when the product allows transfers and the membership is `trialing`, `active`, or `completed`. An account credential with `membership:update` bypasses both restrictions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.TransferMembershipsRequest{
    ID: "id",
}
client.Memberships.Transfer(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Membership ID (`mem_` tag).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Memberships.UncancelMembership(ID) -> *whopsdk.Membership</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Reverse a pending cancellation for a membership that was scheduled to cancel at period end.

Required permissions:
 - `member:manage`
 - `member:email:read`
 - `member:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UncancelMembershipRequest{
    ID: "mem_xxxxxxxxxxxxxx",
}
client.Memberships.UncancelMembership(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the membership to uncancel.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Messages
<details><summary><code>client.Messages.List() -> *whopsdk.ListMessagesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of messages within a specific experience chat, DM, or group chat channel, sorted by creation time.

Required permissions (one of):
 - `chat:read`
 - `dms:read`
 - `support_chat:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListMessagesRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    ChannelID: "channel_id",
}
client.Messages.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**channelID:** `string` — The unique identifier of the channel or experience to list messages for.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.Direction` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Messages.Create(request) -> *whopsdk.Message</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Send a new message in an experience chat, DM, or group chat channel. Supports text content, attachments, polls, and replies.

Required permissions (one of):
 - `chat:message:create`
 - `dms:message:manage`
 - `livestream:chat:write`
 - `support_chat:message:create`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateMessagesRequest{
    ChannelID: "channel_id",
    Content: "content",
}
client.Messages.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**attachments:** `[]*whopsdk.CreateMessagesRequestAttachmentsItem` — A list of file attachments to include with the message, such as images or videos.
    
</dd>
</dl>

<dl>
<dd>

**autoDetectLinks:** `*bool` — Automatically detect URLs in the message and generate link previews.
    
</dd>
</dl>

<dl>
<dd>

**channelID:** `string` — The unique identifier of the channel or experience to send the message in. For example, 'exp_xxxxx' or 'feed_xxxxx'.
    
</dd>
</dl>

<dl>
<dd>

**content:** `string` — The body of the message in Markdown format. For example, 'Hello **world**'.
    
</dd>
</dl>

<dl>
<dd>

**poll:** `*whopsdk.CreateMessagesRequestPoll` — A poll to attach to this message, allowing recipients to vote on options.
    
</dd>
</dl>

<dl>
<dd>

**replyingToMessageID:** `*string` — The unique identifier of the message this is replying to, creating a threaded reply.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Messages.Retrieve(ID) -> *whopsdk.Message</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing message.

Required permissions (one of):
 - `chat:read`
 - `dms:read`
 - `livestream:chat:read`
 - `support_chat:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveMessagesRequest{
    ID: "id",
}
client.Messages.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the message to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Messages.Delete(ID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Permanently delete a message from an experience chat, DM, or group chat channel. Only the message author or a channel admin can delete a message.

Required permissions (one of):
 - `chat:message:create` and `chat:read`
 - `dms:message:manage` and `dms:read`
 - `livestream:chat:write` and `livestream:chat:read`
 - `support_chat:message:create` and `support_chat:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteMessagesRequest{
    ID: "id",
}
client.Messages.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the message to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Messages.Update(ID, request) -> *whopsdk.Message</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Edit the content, attachments, or pinned status of an existing message in an experience chat, DM, or group chat channel.

Required permissions (one of):
 - `chat:message:create`
 - `dms:message:manage`
 - `livestream:chat:write`
 - `support_chat:message:create`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateMessagesRequest{
    ID: "id",
}
client.Messages.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the message to update.
    
</dd>
</dl>

<dl>
<dd>

**attachments:** `[]*whopsdk.UpdateMessagesRequestAttachmentsItem` — A replacement list of file attachments for this message, such as images or videos.
    
</dd>
</dl>

<dl>
<dd>

**content:** `*string` — The updated body of the message in Markdown format. For example, 'Hello **world**'.
    
</dd>
</dl>

<dl>
<dd>

**isPinned:** `*bool` — Whether this message should be pinned to the top of the channel.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Notifications
<details><summary><code>client.Notifications.List() -> *whopsdk.ListNotificationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the authenticated user's notifications, newest first. Requires a user credential — an account API key has no notification feed. Without filters the feed spans every experience the user belongs to plus the teams they are a member of.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListNotificationsRequest{}
client.Notifications.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**unread:** `*bool` — Only return notifications created since the user last viewed their source.
    
</dd>
</dl>

<dl>
<dd>

**experienceID:** `*string` — Only return notifications from this experience (`exp_` tag).
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `*string` — Only return team notifications for this account (`biz_` tag).
    
</dd>
</dl>

<dl>
<dd>

**mentions:** `*bool` — Only return notifications that mention the user directly.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of notifications to return (default 20, max 100).
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor (a notification `id` from a previous page); returns notifications older than it.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Notifications.Create(request) -> *whopsdk.CreateNotificationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Queues a notification to every user of an experience or to an account's team, processed asynchronously. Every send is attributed to an app: use an app API key, or a credential acting on behalf of an app. Narrow the audience with `user_ids` to send a mention.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateNotificationsRequest{
    Content: "Drop off at 4180 Burnet Rd. Plan on two days for the full coating.",
    Title: "Your ceramic coating is booked",
}
client.Notifications.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — Account whose team members receive the notification (`biz_` tag). Exactly one of `experience_id` or `account_id` is required.
    
</dd>
</dl>

<dl>
<dd>

**content:** `string` — Main body text of the notification.
    
</dd>
</dl>

<dl>
<dd>

**experienceID:** `*string` — Experience whose users receive the notification (`exp_` tag). Exactly one of `experience_id` or `account_id` is required.
    
</dd>
</dl>

<dl>
<dd>

**iconUserID:** `*string` — User whose profile picture is used as the notification icon. Defaults to the experience or account avatar.
    
</dd>
</dl>

<dl>
<dd>

**restPath:** `*string` — Path segment appended to the generated deep link that opens your app, for example `/settings/billing`.
    
</dd>
</dl>

<dl>
<dd>

**subtitle:** `*string` — Optional secondary line displayed below the title.
    
</dd>
</dl>

<dl>
<dd>

**title:** `string` — Headline text of the notification.
    
</dd>
</dl>

<dl>
<dd>

**userIDs:** `[]string` — Optional `user_` tags narrowing the audience. When provided, only these users are notified (as a mention), provided they are in the targeted experience or account.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Notifications.Badges() -> *whopsdk.BadgesNotificationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the authenticated user's per-experience unread badge state. Requires a user credential. Returns one row per experience the user belongs to (or per requested experience).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.BadgesNotificationsRequest{
    ExperienceIDs: []*string{
        whopsdk.String(
            "exp_xxxxxxxxxxxxxx",
        ),
    },
}
client.Notifications.Badges(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**experienceIDs:** `*string` — Only return badges for these experiences (`exp_` tags).
    
</dd>
</dl>

<dl>
<dd>

**lastFetchedAt:** `*string` — The client's last fetched-at ISO 8601 timestamp, used to partially refresh badges after a websocket message.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Notifications.MarkRead(request) -> *whopsdk.MarkReadNotificationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Marks the authenticated user's notifications as read: one experience's (`experience_id`) or everything (`all: true`) — exactly one of the two. Requires a user credential. Responds with the refreshed badge rows for the affected scope.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.MarkReadNotificationsRequest{}
client.Notifications.MarkRead(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**all:** `*bool` — Pass `true` to mark every notification read. Exactly one of `experience_id` or `all` is required.
    
</dd>
</dl>

<dl>
<dd>

**experienceID:** `*string` — Experience to mark read (`exp_` tag). Exactly one of `experience_id` or `all` is required.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Notifications.Retrieve(ID) -> *whopsdk.Notification</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a single notification by id — either an `id` returned by List Notifications, or the ephemeral id delivered with a push/websocket event. Requires a user credential.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveNotificationsRequest{
    ID: "id",
}
client.Notifications.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — A notification `id` from List Notifications, or the id delivered with a push/websocket event.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Partners
<details><summary><code>client.Partners.Create() -> *whopsdk.CreatePartnersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Enrolls the calling user in the Whop partner program, making their partner businesses eligible for earnings. Idempotent — enrolling again keeps the original enrollment time.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Partners.Create(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Partners.Leaderboard() -> *whopsdk.LeaderboardPartnersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Ranks referrers by partner business earnings — all-time by default, or over the current day, month, year, or trailing 30 days. Authentication is optional: authenticated callers also get their own standing, anonymous callers get the rankings alone.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.LeaderboardPartnersRequest{}
client.Partners.Leaderboard(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**period:** `*whopsdk.LeaderboardPartnersRequestPeriod` — Time window for the rankings. `day`, `month`, and `year` count earnings since the start of the current calendar day, month, or year; `last_30_days` counts earnings over the trailing 30 days; `all_time` ranks lifetime earnings.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Partners.ReferredUsers() -> *whopsdk.ReferredUsersPartnersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the users the caller referred onto Whop (newest first), each with the second-tier earnings the caller has made from that user's businesses.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ReferredUsersPartnersRequest{}
client.Partners.ReferredUsers(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**hasBusinesses:** `*bool` — When true, only referred users who brought at least one business onto Whop.
    
</dd>
</dl>

<dl>
<dd>

**hasEarningBusinesses:** `*bool` — When true, only referred users with at least one business that has generated earnings.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Number of referred users to return from the start of the window.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor to fetch the page after (from page_info.end_cursor).
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Number of referred users to return from the end of the window.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Cursor to fetch the page before (from page_info.start_cursor).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Payment Method Domains
<details><summary><code>client.PaymentMethodDomains.List() -> *whopsdk.ListPaymentMethodDomainsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists payment method domains. Without `account_id`, returns the caller's own domains and those of every connected account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListPaymentMethodDomainsRequest{}
client.PaymentMethodDomains.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — Only domains registered for this account (`biz_` tag). Defaults to the caller's account plus its connected accounts.
    
</dd>
</dl>

<dl>
<dd>

**hostname:** `*string` — Only the domain with this exact hostname.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.ListPaymentMethodDomainsRequestStatus` — Only domains with this verification status.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*whopsdk.ListPaymentMethodDomainsRequestProvider` — Only domains registered with this wallet provider.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — Only domains created before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — Only domains created after this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListPaymentMethodDomainsRequestOrder` — Sort field.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListPaymentMethodDomainsRequestDirection` — Sort direction.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Number of domains to return from the start of the window.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor to paginate forwards from.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Number of domains to return from the end of the window.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Cursor to paginate backwards from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PaymentMethodDomains.Create(request) -> *whopsdk.PaymentMethodDomain</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Registers a hostname with the wallet provider and attempts verification inline. Returns `verified` when the provider fetched the domain-association file (for Apple Pay, `/.well-known/apple-developer-merchantid-domain-association`), or `pending` when it could not — host the file, then retry with the verify endpoint.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreatePaymentMethodDomainsRequest{
    Hostname: "pending.shinetime.example",
}
client.PaymentMethodDomains.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — Account to register the domain for (`biz_` tag). Defaults to the caller's account.
    
</dd>
</dl>

<dl>
<dd>

**hostname:** `string` — Hostname to register (e.g. `checkout.shinetime.example`).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PaymentMethodDomains.Retrieve(ID) -> *whopsdk.PaymentMethodDomain</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a payment method domain to check its verification status.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrievePaymentMethodDomainsRequest{
    ID: "id",
}
client.PaymentMethodDomains.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the payment method domain, prefixed `pmd_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PaymentMethodDomains.Delete(ID) -> *whopsdk.DeletePaymentMethodDomainsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Unregisters a payment method domain so its wallet payment methods stop rendering there.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeletePaymentMethodDomainsRequest{
    ID: "id",
}
client.PaymentMethodDomains.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the payment method domain, prefixed `pmd_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PaymentMethodDomains.Verify(ID) -> *whopsdk.PaymentMethodDomain</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Re-attempts provider verification of a pending domain once the association file is hosted. Fails with a `bad_request` explaining what to fix; verifying an already `verified` domain is a no-op.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.VerifyPaymentMethodDomainsRequest{
    ID: "id",
}
client.PaymentMethodDomains.Verify(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the payment method domain, prefixed `pmd_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## PaymentMethods
<details><summary><code>client.PaymentMethods.List() -> *whopsdk.ListPaymentMethodsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of payment methods for a member or company, or for the authenticated user when neither is given, with optional filtering by creation date. A payment method is a stored representation of how a customer intends to pay, such as a card, bank account, or digital wallet.

Required permissions:
 - `member:payment_methods:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListPaymentMethodsRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    MemberID: whopsdk.String(
        "mber_xxxxxxxxxxxxx",
    ),
    CompanyID: whopsdk.String(
        "biz_xxxxxxxxxxxxxx",
    ),
    CreatedBefore: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
    CreatedAfter: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
}
client.PaymentMethods.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**memberID:** `*string` — The unique identifier of the member to list payment methods for. Omit this and company_id to list your own saved payment methods.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `*string` — The unique identifier of the company. Provide either this or member_id, not both. Omit both to address your own saved payment methods.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.Direction` 
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — Only return payment methods created before this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — Only return payment methods created after this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**futureUsage:** `*whopsdk.FutureUsageTypes` 
    
</dd>
</dl>

<dl>
<dd>

**paymentMethodTypes:** `*whopsdk.PaymentMethodTypes` — Only return payment methods of these types. Pass the eligible `type` values from the payment method types catalogue so the list holds nothing the purchase cannot take. An empty list returns no payment methods.
    
</dd>
</dl>

<dl>
<dd>

**cardBrands:** `*whopsdk.CardBrands` — Only return cards on these networks, such as the networks the seller accepts. Payment methods that are not cards are unaffected.
    
</dd>
</dl>

<dl>
<dd>

**cardFundingTypes:** `*whopsdk.CardFundingTypes` — Only return cards funded this way. A card whose funding could not be determined is excluded, and payment methods that are not cards are unaffected.
    
</dd>
</dl>

<dl>
<dd>

**hasPayerDocument:** `*bool` — Filter cards by whether they carry the payer identity document their payment provider requires. Payment methods that are not cards are unaffected.
    
</dd>
</dl>

<dl>
<dd>

**expired:** `*bool` — Filter by expiry. Only a card can expire, so `false` keeps every payment method that is not past its expiration month and `true` returns expired cards alone.
    
</dd>
</dl>

<dl>
<dd>

**broken:** `*bool` — Filter by whether the stored credential has permanently stopped charging, such as a vault entry its provider closed.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PaymentMethods.Retrieve(ID) -> *whopsdk.PaymentMethod</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing payment method. Addresses a member's wallet when member_id or company_id is given, otherwise your own.

Required permissions:
 - `member:payment_methods:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrievePaymentMethodsRequest{
    ID: "payt_xxxxxxxxxxxxx",
    CompanyID: whopsdk.String(
        "biz_xxxxxxxxxxxxxx",
    ),
    MemberID: whopsdk.String(
        "mber_xxxxxxxxxxxxx",
    ),
}
client.PaymentMethods.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the payment method.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `*string` — The unique identifier of the company. Provide either this or member_id, not both. Omit both to address your own saved payment methods.
    
</dd>
</dl>

<dl>
<dd>

**memberID:** `*string` — The unique identifier of the member. Provide either this or company_id, not both. Omit both to address your own saved payment methods.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PaymentMethods.DeletePaymentMethod(ID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a saved payment method. Cannot delete a payment method attached to an active subscription.

Required permissions:
 - `member:payment_methods:manage`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeletePaymentMethodRequest{
    ID: "payt_xxxxxxxxxxxxx",
    CompanyID: whopsdk.String(
        "biz_xxxxxxxxxxxxxx",
    ),
    MemberID: whopsdk.String(
        "mber_xxxxxxxxxxxxx",
    ),
}
client.PaymentMethods.DeletePaymentMethod(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the payment method to delete.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `*string` — The unique identifier of the company. Provide either this or member_id, not both. Omit both to address your own saved payment methods.
    
</dd>
</dl>

<dl>
<dd>

**memberID:** `*string` — The unique identifier of the member. Provide either this or company_id, not both. Omit both to address your own saved payment methods.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Payments
<details><summary><code>client.Payments.List() -> *whopsdk.ListPaymentsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of payments for the actor in context, with optional filtering by product, plan, status, billing reason, currency, and creation date.

Required permissions:
 - `payment:basic:read`
 - `plan:basic:read`
 - `access_pass:basic:read`
 - `member:email:read`
 - `member:basic:read`
 - `member:phone:read`
 - `promo_code:basic:read`
 - `shipment:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListPaymentsRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    CompanyID: whopsdk.String(
        "biz_xxxxxxxxxxxxxx",
    ),
    CreatedBefore: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
    CreatedAfter: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
    UpdatedBefore: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
    UpdatedAfter: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
}
client.Payments.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `*string` — The unique identifier of the company to list payments for.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.Direction` 
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ReceiptV2Order` 
    
</dd>
</dl>

<dl>
<dd>

**productIDs:** `*string` — Filter payments to only those associated with these specific product identifiers.
    
</dd>
</dl>

<dl>
<dd>

**billingReasons:** `*whopsdk.BillingReasons` — Filter payments by their billing reason.
    
</dd>
</dl>

<dl>
<dd>

**currencies:** `*whopsdk.Currencies` — Filter payments by their currency code.
    
</dd>
</dl>

<dl>
<dd>

**planIDs:** `*string` — Filter payments to only those associated with these specific plan identifiers.
    
</dd>
</dl>

<dl>
<dd>

**statuses:** `*whopsdk.ReceiptStatus` — Filter payments by their current status.
    
</dd>
</dl>

<dl>
<dd>

**substatuses:** `*whopsdk.FriendlyReceiptStatus` — Filter payments by their current substatus for more granular filtering.
    
</dd>
</dl>

<dl>
<dd>

**includeFree:** `*bool` — Whether to include payments with a zero amount.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — Only return payments created before this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — Only return payments created after this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**updatedBefore:** `*time.Time` — Only return payments last updated before this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**updatedAfter:** `*time.Time` — Only return payments last updated after this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` — Search payments by user ID, membership ID, user email, name, or username. Email filtering requires the member:email:read permission.
    
</dd>
</dl>

<dl>
<dd>

**checkoutConfigurationIDs:** `*string` — Only return payments from these checkout configurations.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payments.Create(request) -> *whopsdk.CreatePaymentsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Charge an existing member off-session using one of their stored payment methods. You can provide an existing plan, or create a new one in-line. This endpoint will respond with a payment object immediately, but the payment is processed asynchronously in the background. Use webhooks to be notified when the payment succeeds or fails.

Required permissions:
 - `payment:charge`
 - `plan:create`
 - `access_pass:create`
 - `access_pass:update`
 - `plan:basic:read`
 - `access_pass:basic:read`
 - `member:email:read`
 - `member:basic:read`
 - `member:phone:read`
 - `promo_code:basic:read`
 - `shipment:basic:read`
 - `payment:dispute:read`
 - `payment:resolution_center_case:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreatePaymentsRequest{
    CreatePaymentsRequestZero: &whopsdk.CreatePaymentsRequestZero{
        CompanyID: "biz_xxxxxxxxxxxxxx",
        ConfirmationToken: "confirmation_token",
        Plan: &whopsdk.CreatePaymentsRequestZeroPlan{
            Currency: whopsdk.CurrenciesUsd,
        },
    },
}
client.Payments.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*whopsdk.CreatePaymentsRequest` — Parameters for CreatePayment
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payments.Retrieve(ID) -> *whopsdk.RetrievePaymentsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing payment.

Required permissions:
 - `payment:basic:read`
 - `plan:basic:read`
 - `access_pass:basic:read`
 - `member:email:read`
 - `member:basic:read`
 - `member:phone:read`
 - `promo_code:basic:read`
 - `shipment:basic:read`
 - `payment:dispute:read`
 - `payment:resolution_center_case:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrievePaymentsRequest{
    ID: "pay_xxxxxxxxxxxxxx",
}
client.Payments.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the payment.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payments.Capture(ID) -> *whopsdk.PaymentStatus</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Captures the full amount of a card payment created with `capture: false`. The payment must still be in `requires_capture` before `capture_expires_at`. Partial capture, multiple captures, capturing more than the authorized amount, and tips are not supported.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CapturePaymentsRequest{
    ID: "id",
}
client.Payments.Capture(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the payment.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payments.ListFees(ID) -> *whopsdk.ListFeesPaymentsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the list of fees associated with a specific payment, including platform fees and processing fees.

Required permissions:
 - `payment:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListFeesPaymentsRequest{
    ID: "pay_xxxxxxxxxxxxxx",
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
}
client.Payments.ListFees(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the payment to list fees for.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payments.Refund(ID, request) -> *whopsdk.Payment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Issue a full or partial refund for a payment. The refund is processed through the original payment processor and the membership status is updated accordingly.

Required permissions:
 - `payment:manage`
 - `plan:basic:read`
 - `access_pass:basic:read`
 - `member:email:read`
 - `member:basic:read`
 - `member:phone:read`
 - `promo_code:basic:read`
 - `shipment:basic:read`
 - `payment:dispute:read`
 - `payment:resolution_center_case:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RefundPaymentsRequest{
    ID: "pay_xxxxxxxxxxxxxx",
}
client.Payments.Refund(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the payment to refund.
    
</dd>
</dl>

<dl>
<dd>

**partialAmount:** `*float64` — The amount to refund. For multi-currency payments, this is in the charge currency (what the buyer paid). For single-currency, this is in the payment currency. If omitted, the full payment amount is refunded.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payments.Retry(ID) -> *whopsdk.Payment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retry a failed or pending payment. This re-attempts the charge using the original payment method and plan details.

Required permissions:
 - `payment:manage`
 - `plan:basic:read`
 - `access_pass:basic:read`
 - `member:email:read`
 - `member:basic:read`
 - `member:phone:read`
 - `promo_code:basic:read`
 - `shipment:basic:read`
 - `payment:dispute:read`
 - `payment:resolution_center_case:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetryPaymentsRequest{
    ID: "pay_xxxxxxxxxxxxxx",
}
client.Payments.Retry(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the payment to retry.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payments.Void(ID) -> *whopsdk.Payment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Void a payment that has not yet been settled. Voiding cancels the payment before it is captured by the payment processor.

Required permissions:
 - `payment:manage`
 - `plan:basic:read`
 - `access_pass:basic:read`
 - `member:email:read`
 - `member:basic:read`
 - `member:phone:read`
 - `promo_code:basic:read`
 - `shipment:basic:read`
 - `payment:dispute:read`
 - `payment:resolution_center_case:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.VoidPaymentsRequest{
    ID: "pay_xxxxxxxxxxxxxx",
}
client.Payments.Void(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the payment to void.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payments.UpdateReturnURL(PaymentID, request) -> *whopsdk.PaymentStatus</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Changes where the buyer lands after completing an off-site step, up until they return. Accepts either a secret key or the payment's own `client_secret`, so the surface that knows the final destination can set it.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateReturnURLPaymentsRequest{
    PaymentID: "payment_id",
    ReturnURL: "https://shinetime.example/checkout/thanks",
}
client.Payments.UpdateReturnURL(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentID:** `string` — The unique identifier of the payment.
    
</dd>
</dl>

<dl>
<dd>

**returnURL:** `string` — Where the buyer continues after completing an off-site step. Must be an absolute https URL without credentials (http is allowed for localhost), at most 2,048 characters.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payments.RetrieveStatus(PaymentID) -> *whopsdk.PaymentStatus</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves how far a payment has got and what the buyer must do next, if anything. A payment is collected in the background, so poll this rather than reading the create response. Accepts either a secret key or the payment's own `client_secret`, so the surface collecting the payment can poll it directly.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveStatusPaymentsRequest{
    PaymentID: "payment_id",
}
client.Payments.RetrieveStatus(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentID:** `string` — The unique identifier of the payment.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## PayoutAccounts
<details><summary><code>client.PayoutAccounts.Retrieve(ID) -> *whopsdk.PayoutAccount</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing payout account.

Required permissions:
 - `payout:account:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrievePayoutAccountsRequest{
    ID: "poact_xxxxxxxxxxxx",
}
client.PayoutAccounts.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the payout account to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## PayoutMethods
<details><summary><code>client.PayoutMethods.ListPayoutMethod() -> *whopsdk.ListPayoutMethodResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of active payout methods configured for a company, ordered by most recently created.

Required permissions:
 - `payout:destination:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListPayoutMethodRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    CompanyID: "biz_xxxxxxxxxxxxxx",
}
client.PayoutMethods.ListPayoutMethod(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `string` — The unique identifier of the company to list payout methods for.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PayoutMethods.RetrievePayoutMethod(ID) -> *whopsdk.PayoutMethod</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing payout method.

Required permissions:
 - `payout:destination:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrievePayoutMethodRequest{
    ID: "potk_xxxxxxxxxxxxx",
}
client.PayoutMethods.RetrievePayoutMethod(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the payout method to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Payouts
<details><summary><code>client.Payouts.List() -> *whopsdk.ListPayoutsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists an account's or user's payouts, newest first.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListPayoutsRequest{}
client.Payouts.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The owning account ID (a biz_ identifier). Provide this or user_id.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — The owning user ID (a user_ identifier). Provide this or account_id.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*string` — Optional currency code filter, for example `usd`.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.ListPayoutsRequestStatus` — Filter to payouts whose `status` reads this word, matching exactly what this version displays — `reversed` finds settled payouts the bank later returned. Requires Api-Version-Date 2026-08-21 or later.
    
</dd>
</dl>

<dl>
<dd>

**source:** `*whopsdk.ListPayoutsRequestSource` — Filter by how the payout was created. Payouts created before source tracking or through internal tooling carry no source and never match.
    
</dd>
</dl>

<dl>
<dd>

**payoutMethodID:** `*string` — Filter to payouts sent to one saved payout method (a pytk_ identifier). An unknown id matches nothing.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — Only payouts created before this ISO 8601 time (exclusive).
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — Only payouts created at or after this ISO 8601 time (inclusive).
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Number of payouts to return from the start of the window.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor to fetch the page after (from page_info.end_cursor).
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Number of payouts to return from the end of the window.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Cursor to fetch the page before (from page_info.start_cursor).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payouts.Create(request) -> *whopsdk.CreatePayoutsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Sends money from an account or user balance to a saved payout method for that owner.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreatePayoutsRequestBody{
    Unknown: map[string]any{
        "key": "value",
    },
}
client.Payouts.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*whopsdk.CreatePayoutsRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payouts.CreateQuote(request) -> *whopsdk.CreateQuotePayoutsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a short-lived, provider-backed quote for a payout. No funds move until the returned quote_token is submitted to POST /payouts. An Idempotency-Key header is required.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateQuotePayoutsRequest{
    Amount: 6762.41,
    PayoutMethodID: "potk_xxxxxxxxxxxxxx",
}
client.Payouts.CreateQuote(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — Account to pay out from, prefixed `biz_`. Provide exactly one of `account_id` or `user_id`.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `float64` — The amount to pay out in the specified currency.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*string` — The balance currency to pay out.
    
</dd>
</dl>

<dl>
<dd>

**payoutMethodID:** `string` — The saved payout method to quote (a potk_ identifier).
    
</dd>
</dl>

<dl>
<dd>

**platformCoversFees:** `*bool` — Whether the parent platform covers the payout fee instead of the account being paid out.
    
</dd>
</dl>

<dl>
<dd>

**speed:** `*whopsdk.CreateQuotePayoutsRequestSpeed` — How fast the funds should arrive.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — User to pay out from, prefixed `user_`. Provide exactly one of `account_id` or `user_id`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payouts.Retrieve(ID) -> *whopsdk.RetrievePayoutsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Fetches one payout by its `wdrl_` ID, or by the `cofr_` conversion request ID a stablecoin payout carries as `payout_request_id` — both ids answer with the same payout object.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrievePayoutsRequest{
    ID: "id",
}
client.Payouts.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Payout ID, prefixed `wdrl_` for a payout returned by `GET /payouts` or `cofr_` for the payout request returned by `POST /payouts`.
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `*string` — Owning account ID, prefixed `biz_`. Provide exactly one of `account_id` or `user_id`.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — Owning user ID, prefixed `user_`. Provide exactly one of `account_id` or `user_id`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payouts.Cancel(ID) -> *whopsdk.CancelPayoutsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancels a payout that is still in review and returns the funds, fees included, to the balance. A payout can be canceled while its status is `in_review`. A `requested` payout is still being prepared (its funds may be converting) and answers 409 until it reaches review; from `processing` on, the money is on its way and the answer is 409 with error type `not_cancelable`. Canceling a payout that is already canceled succeeds and returns it unchanged.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CancelPayoutsRequest{
    ID: "id",
}
client.Payouts.Cancel(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Payout ID, prefixed `wdrl_`, or the `cofr_` payout request ID returned by `POST /payouts` — both cancel the same payout.
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `*string` — Owning account ID, prefixed `biz_`. Provide exactly one of `account_id` or `user_id`.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — Owning user ID, prefixed `user_`. Provide exactly one of `account_id` or `user_id`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## People
<details><summary><code>client.People.List() -> *whopsdk.ListPeopleResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the people (visitors and customers) of an account: the identity-linked person profiles aggregated from every pixel, payment, and platform event — identities, purchases and LTV, geo/device profile, traffic sources, and first/last marketing touches.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListPeopleRequest{
    Source: []*string{
        whopsdk.String(
            "direct",
        ),
    },
    EventName: []*string{
        whopsdk.String(
            "payment.completed",
        ),
    },
}
client.People.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — Account ID, prefixed `biz_`. Optional for account API keys; required for credentials that can access multiple accounts.
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` — Search people by name, email, phone, or whop user ID (case-insensitive substring match).
    
</dd>
</dl>

<dl>
<dd>

**source:** `*string` — Only include people acquired from any of these sources — canonical paths (whop:<campaign>:<group>:<ad>, ext:<platform>:..., referrer:<domain>, direct, other), exact or with a trailing :* prefix. The same vocabulary the events / people metrics use.
    
</dd>
</dl>

<dl>
<dd>

**attributionModel:** `*whopsdk.ListPeopleRequestAttributionModel` — Attribution model the source filter matches against (defaults to last_touch).
    
</dd>
</dl>

<dl>
<dd>

**eventName:** `*string` — Only include people who fired any of these events, e.g. payment.completed or page.checkout.view.
    
</dd>
</dl>

<dl>
<dd>

**customEvent:** `*string` — Only include people who fired this custom pixel event.
    
</dd>
</dl>

<dl>
<dd>

**eventFrom:** `*time.Time` — With event_to plus an event or source filter, switches to exact-population mode: person ids are resolved and paginated on the events side within this window (the same query the people metric counts), then hydrated per page.
    
</dd>
</dl>

<dl>
<dd>

**eventTo:** `*time.Time` — The inclusive end of the event window for exact-population mode.
    
</dd>
</dl>

<dl>
<dd>

**audienceID:** `*string` — Only include people in this audience. An audience that keeps itself up to date resolves to the People filters that define it, so this always reflects who matches now; uploaded lists and point-in-time snapshots match their recorded members.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — Only include the person linked to this whop user ID.
    
</dd>
</dl>

<dl>
<dd>

**email:** `*string` — Only include the person linked to this email address.
    
</dd>
</dl>

<dl>
<dd>

**phone:** `*string` — Only include the person linked to this phone number.
    
</dd>
</dl>

<dl>
<dd>

**country:** `*string` — Only include people whose most recent visit came from this ISO 3166-1 alpha-2 country code.
    
</dd>
</dl>

<dl>
<dd>

**hasPurchased:** `*bool` — true for customers only, false for people who have never purchased.
    
</dd>
</dl>

<dl>
<dd>

**contactable:** `*bool` — true for people who have an email address or phone number — the ones an ad platform can match.
    
</dd>
</dl>

<dl>
<dd>

**firstSeenWithinDays:** `*int` — Only include people first seen within this many days, as a rolling window.
    
</dd>
</dl>

<dl>
<dd>

**lastSeenWithinDays:** `*int` — Only include people last seen within this many days, as a rolling window.
    
</dd>
</dl>

<dl>
<dd>

**firstSeenAfter:** `*time.Time` — Only include people first seen at or after this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**firstSeenBefore:** `*time.Time` — Only include people first seen before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**lastSeenAfter:** `*time.Time` — Only include people last seen at or after this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**lastSeenBefore:** `*time.Time` — Only include people last seen before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of people to return (default 100, max 100).
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor for fetching people after a previous page.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — A cursor for fetching people before a later page.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListPeopleRequestOrder` — Column to sort by. Defaults to last_seen_at.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListPeopleRequestDirection` — Sort direction. Defaults to desc.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.People.Retrieve(ID) -> *whopsdk.RetrievePeopleResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves one person for an account. The identifier can be a person ID (prefixed `prsn_`), a user ID (prefixed `user_`), an email address, or a phone number — merged people resolve to the surviving profile.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrievePeopleRequest{
    ID: "id",
}
client.People.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The person ID, user ID, email address, or phone number to look up.
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `*string` — Account ID, prefixed `biz_`. Optional for account API keys; required for credentials that can access multiple accounts.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Permissions
<details><summary><code>client.Permissions.List() -> *whopsdk.ListPermissionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists permission actions and whether the calling credential is granted each one for a resource. Answers for whichever identity authenticated the request — a user session, an OAuth token, or an account or app API key — so it never describes who else can reach the resource.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListPermissionsRequest{
    ResourceID: "resource_id",
}
client.Permissions.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**resourceID:** `string` — Tag of the resource to check against: an account (`biz_`), product (`prod_`), experience (`exp_`), or app (`app_`). A resource the credential cannot see is reported as granted nothing rather than as an error.
    
</dd>
</dl>

<dl>
<dd>

**actions:** `*string` — Comma-separated permission actions to check, for example `stats:read,payment:basic:read`. Every action is returned when omitted.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Plans
<details><summary><code>client.Plans.List() -> *whopsdk.ListPlansResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of plans. Omit `account_id` and pass `product_ids` to list a product's public buyable plans.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListPlansRequest{
    ReleaseMethods: []*string{
        whopsdk.String(
            "buy_now",
        ),
    },
    Visibilities: []*string{
        whopsdk.String(
            "visible",
        ),
    },
    PlanTypes: []*string{
        whopsdk.String(
            "renewal",
        ),
    },
    ProductIDs: []*string{
        whopsdk.String(
            "prod_xxxxxxxxxxxxxx",
        ),
    },
}
client.Plans.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The unique identifier of the account to list plans for. Required unless `product_ids` is provided for a public product-plan read.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListPlansRequestDirection` — The sort direction for results. Defaults to descending.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListPlansRequestOrder` — The field to sort results by. Defaults to created_at.
    
</dd>
</dl>

<dl>
<dd>

**releaseMethods:** `*string` — Filter to only plans matching these release methods.
    
</dd>
</dl>

<dl>
<dd>

**visibilities:** `*string` — Filter to only plans matching these visibility states.
    
</dd>
</dl>

<dl>
<dd>

**planTypes:** `*string` — Filter to only plans matching these billing types.
    
</dd>
</dl>

<dl>
<dd>

**productIDs:** `*string` — Filter to only plans belonging to these product identifiers. When `account_id` is omitted, this is required and the response is publicly readable: only visible, non-invoice plans are returned.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Only return plans created before this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Only return plans created after this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of plans to return (default and max 100).
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor; returns plans after this position.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — The number of plans to return from the end of the range.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — A cursor; returns plans before this position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Plans.Create(request) -> *whopsdk.Plan</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new pricing plan for a product. The plan defines the billing interval, price, and availability for customers.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreatePlansRequest{}
client.Plans.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The unique identifier of the account to create this plan for. Defaults to the caller's account.
    
</dd>
</dl>

<dl>
<dd>

**adaptivePricingEnabled:** `*bool` — Whether this plan accepts local currency payments via adaptive pricing.
    
</dd>
</dl>

<dl>
<dd>

**billingPeriod:** `*int` — Recurring billing interval in days, such as 30 for monthly or 365 for annual.
    
</dd>
</dl>

<dl>
<dd>

**checkoutStyling:** `map[string]any` — Checkout styling overrides for this plan.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*string` — The three-letter ISO currency code for the plan's pricing. Defaults to USD.
    
</dd>
</dl>

<dl>
<dd>

**customFields:** `[]*whopsdk.CreatePlansRequestCustomFieldsItem` — An array of custom field definitions to collect from customers at checkout. Omitting this field clears existing custom fields.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — A text description of the plan displayed to customers on the product page.
    
</dd>
</dl>

<dl>
<dd>

**expirationDays:** `*int` — Access duration in days before the membership expires.
    
</dd>
</dl>

<dl>
<dd>

**image:** `*whopsdk.CreatePlansRequestImage` — An image displayed on the product page to represent this plan.
    
</dd>
</dl>

<dl>
<dd>

**initialPrice:** `*float64` — Initial amount charged in the plan's currency, e.g. 10.43 for $10.43.
    
</dd>
</dl>

<dl>
<dd>

**internalNotes:** `*string` — Private notes visible only to the account owner. Not shown to customers.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Custom key-value pairs to store on the plan. Included in webhook payloads for payment and membership events. Max 50 keys, 100 chars per key, 500 chars per string value. The reserved keys `custom_cta` (a checkout call-to-action button label — one of the product custom CTA values, e.g. `subscribe`, `get_offer`) and `custom_cta_url` (a URL the button links to; web or `tel:`) override the product's call to action for this plan and are validated on save.
    
</dd>
</dl>

<dl>
<dd>

**overrideTaxType:** `*string` — Override the default tax classification for this specific plan.
    
</dd>
</dl>

<dl>
<dd>

**paymentMethodConfiguration:** `*whopsdk.CreatePlansRequestPaymentMethodConfiguration` — Explicit payment method configuration for the plan. When not provided, the account's defaults apply.
    
</dd>
</dl>

<dl>
<dd>

**planType:** `*string` — Plan billing type, such as `one_time` or `renewal`.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `*string` — The unique identifier of the product to attach this plan to.
    
</dd>
</dl>

<dl>
<dd>

**releaseMethod:** `*string` — Sales method for this plan.
    
</dd>
</dl>

<dl>
<dd>

**renewalPrice:** `*float64` — The amount charged each billing period for recurring plans, in the plan's currency.
    
</dd>
</dl>

<dl>
<dd>

**splitPayRequiredPayments:** `*int` — Installment payments required before the subscription pauses.
    
</dd>
</dl>

<dl>
<dd>

**stock:** `*int` — The maximum number of units available for purchase. Ignored when unlimited_stock is true.
    
</dd>
</dl>

<dl>
<dd>

**threeDsLevel:** `*whopsdk.CreatePlansRequestThreeDsLevel` — 3D Secure behavior for this plan. Send `null` to inherit the account default.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The display name of the plan shown to customers on the product page.
    
</dd>
</dl>

<dl>
<dd>

**trialPeriodDays:** `*int` — Free trial duration before the first recurring charge.
    
</dd>
</dl>

<dl>
<dd>

**unlimitedStock:** `*bool` — Whether the plan has unlimited stock. When true, the stock field is ignored.
    
</dd>
</dl>

<dl>
<dd>

**visibility:** `*string` — Whether the plan is visible to customers or hidden from public view.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Plans.Retrieve(ID) -> *whopsdk.Plan</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing plan.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrievePlansRequest{
    ID: "id",
}
client.Plans.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Plan ID, prefixed `plan_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Plans.Delete(ID) -> *whopsdk.DeletePlansResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Permanently delete a plan from a product. Existing memberships on this plan will not be affected.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeletePlansRequest{
    ID: "id",
}
client.Plans.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Plan ID, prefixed `plan_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Plans.Update(ID, request) -> *whopsdk.Plan</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a plan's pricing, billing interval, visibility, stock, and other settings.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdatePlansRequest{
    ID: "id",
}
client.Plans.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Plan ID, prefixed `plan_`.
    
</dd>
</dl>

<dl>
<dd>

**adaptivePricingEnabled:** `*bool` — Whether this plan accepts local currency payments via adaptive pricing.
    
</dd>
</dl>

<dl>
<dd>

**billingPeriod:** `*int` — Recurring billing interval in days, such as 30 for monthly or 365 for annual.
    
</dd>
</dl>

<dl>
<dd>

**cancelDiscountIntervals:** `*int` — How many renewals the retention discount applies to. Required when `offer_cancel_discount` is true.
    
</dd>
</dl>

<dl>
<dd>

**cancelDiscountPercentage:** `*int` — Percentage taken off each discounted renewal. Required when `offer_cancel_discount` is true.
    
</dd>
</dl>

<dl>
<dd>

**checkoutStyling:** `map[string]any` — Checkout styling overrides for this plan.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*string` — The three-letter ISO currency code for the plan's pricing. Defaults to USD.
    
</dd>
</dl>

<dl>
<dd>

**customFields:** `[]*whopsdk.UpdatePlansRequestCustomFieldsItem` — An array of custom field definitions to collect from customers at checkout. Omitting this field clears existing custom fields.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — A text description of the plan displayed to customers on the product page.
    
</dd>
</dl>

<dl>
<dd>

**expirationDays:** `*int` — Access duration in days before the membership expires.
    
</dd>
</dl>

<dl>
<dd>

**image:** `*whopsdk.UpdatePlansRequestImage` — An image displayed on the product page to represent this plan.
    
</dd>
</dl>

<dl>
<dd>

**initialPrice:** `*float64` — Initial amount charged in the plan's currency, e.g. 10.43 for $10.43.
    
</dd>
</dl>

<dl>
<dd>

**internalNotes:** `*string` — Private notes visible only to the account owner. Not shown to customers.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Custom key-value pairs to store on the plan. Included in webhook payloads for payment and membership events. Max 50 keys, 100 chars per key, 500 chars per string value. The reserved keys `custom_cta` (a checkout call-to-action button label — one of the product custom CTA values, e.g. `subscribe`, `get_offer`) and `custom_cta_url` (a URL the button links to; web or `tel:`) override the product's call to action for this plan and are validated on save.
    
</dd>
</dl>

<dl>
<dd>

**offerCancelDiscount:** `*bool` — Whether to offer a retention discount when a customer attempts to cancel.
    
</dd>
</dl>

<dl>
<dd>

**overrideTaxType:** `*string` — Override the default tax classification for this specific plan.
    
</dd>
</dl>

<dl>
<dd>

**paymentMethodConfiguration:** `*whopsdk.UpdatePlansRequestPaymentMethodConfiguration` — Explicit payment method configuration for the plan. When not provided, the account's defaults apply.
    
</dd>
</dl>

<dl>
<dd>

**releaseMethod:** `*string` — Sales method for this plan.
    
</dd>
</dl>

<dl>
<dd>

**renewalPrice:** `*float64` — The amount charged each billing period for recurring plans, in the plan's currency.
    
</dd>
</dl>

<dl>
<dd>

**stock:** `*int` — The maximum number of units available for purchase. Ignored when unlimited_stock is true.
    
</dd>
</dl>

<dl>
<dd>

**strikeThroughInitialPrice:** `*float64` — A comparison price displayed with a strikethrough for the initial price.
    
</dd>
</dl>

<dl>
<dd>

**strikeThroughRenewalPrice:** `*float64` — A comparison price displayed with a strikethrough for the renewal price.
    
</dd>
</dl>

<dl>
<dd>

**threeDsLevel:** `*whopsdk.UpdatePlansRequestThreeDsLevel` — 3D Secure behavior for this plan. Send `null` to inherit the account default.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The display name of the plan shown to customers on the product page.
    
</dd>
</dl>

<dl>
<dd>

**trialPeriodDays:** `*int` — Free trial duration before the first recurring charge.
    
</dd>
</dl>

<dl>
<dd>

**unlimitedStock:** `*bool` — Whether the plan has unlimited stock. When true, the stock field is ignored.
    
</dd>
</dl>

<dl>
<dd>

**visibility:** `*string` — Whether the plan is visible to customers or hidden from public view.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Plans.CalculateTax(ID, request) -> *whopsdk.CalculateTaxPlansResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Previews tax for a plan before checkout, based on the buyer's location.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CalculateTaxPlansRequest{
    ID: "id",
    Address: &whopsdk.CalculateTaxPlansRequestAddress{
        Country: "DE",
        PostalCode: whopsdk.String(
            "10115",
        ),
    },
}
client.Plans.CalculateTax(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Plan ID, prefixed `plan_`.
    
</dd>
</dl>

<dl>
<dd>

**address:** `*whopsdk.CalculateTaxPlansRequestAddress` — Buyer billing address used for tax calculation. Provide either `address.country` or `ip_address`; include state and postal code when available for more accurate results.
    
</dd>
</dl>

<dl>
<dd>

**ipAddress:** `*string` — Buyer IP address used to infer location when no billing address is provided.
    
</dd>
</dl>

<dl>
<dd>

**taxIDs:** `[]*whopsdk.CalculateTaxPlansRequestTaxIDsItem` — Optional buyer tax ID for B2B exemptions. At most one entry is supported.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Products
<details><summary><code>client.Products.List() -> *whopsdk.ListProductsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of products. Omit `account_id` to search the public marketplace.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListProductsRequest{
    Visibilities: []*string{
        whopsdk.String(
            "visible",
        ),
    },
    AccessPassTypes: []*string{
        whopsdk.String(
            "regular",
        ),
    },
}
client.Products.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The unique identifier of the account to list products for. Omit to search the public marketplace.
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` — Ranked search against product title and headline. Omit to browse by recency.
    
</dd>
</dl>

<dl>
<dd>

**marketplaceCategoryRoute:** `*string` — Only return marketplace products assigned to this category route, such as `trading`.
    
</dd>
</dl>

<dl>
<dd>

**planTypes:** `*whopsdk.ListProductsRequestPlanTypesItem` — Filter to products with a buyable plan of these billing models, such as `one_time` or `renewal`.
    
</dd>
</dl>

<dl>
<dd>

**priceMinimum:** `*float64` — Only return products whose advertised buyable plan has a displayed price of at least this amount. Recurring plans use renewal price.
    
</dd>
</dl>

<dl>
<dd>

**priceMaximum:** `*float64` — Only return products whose advertised buyable plan has a displayed price of at most this amount. Recurring plans use renewal price.
    
</dd>
</dl>

<dl>
<dd>

**visibilities:** `*string` — Filter to only products matching these visibility states. Ignored on the public marketplace list, which only returns visible products.
    
</dd>
</dl>

<dl>
<dd>

**accessPassTypes:** `*string` — Filter to only products matching these types.
    
</dd>
</dl>

<dl>
<dd>

**labels:** `*string` — Filter to only products carrying all of these labels. Labels are matched lowercased.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListProductsRequestDirection` — The sort direction for results. Defaults to descending.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — The field to sort results by. Account lists default to `created_at`. Marketplace lists default to `discoverable_at` and accept `created_at` or `discoverable_at`. Cannot be combined with `query`.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of products to return (default and max 100).
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor; returns products after this position.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — The number of products to return from the end of the range.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — A cursor; returns products before this position.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Only return products created after this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Only return products created before this ISO 8601 timestamp.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Products.Create(request) -> *whopsdk.Product</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new product for an account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateProductsRequest{
    Title: "Interior Deep Clean",
}
client.Products.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The unique identifier of the account to create this product for.
    
</dd>
</dl>

<dl>
<dd>

**collectShippingAddress:** `*bool` — Whether to collect a shipping address at checkout.
    
</dd>
</dl>

<dl>
<dd>

**customCta:** `*whopsdk.CreateProductsRequestCustomCta` — The call-to-action button label.
    
</dd>
</dl>

<dl>
<dd>

**customCtaURL:** `*string` — A URL the call-to-action button links to.
    
</dd>
</dl>

<dl>
<dd>

**customStatementDescriptor:** `*string` — Custom bank statement descriptor. Must start with WHOP*.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — A written description displayed on the product page.
    
</dd>
</dl>

<dl>
<dd>

**globalAffiliatePercentage:** `*float64` — The commission rate affiliates earn.
    
</dd>
</dl>

<dl>
<dd>

**globalAffiliateStatus:** `*whopsdk.CreateProductsRequestGlobalAffiliateStatus` — The enrollment status in the global affiliate program.
    
</dd>
</dl>

<dl>
<dd>

**headline:** `*string` — A short marketing headline for the product page.
    
</dd>
</dl>

<dl>
<dd>

**labels:** `[]string` — Labels used to group products into collections. Stored lowercased and de-duplicated. Maximum 20 labels, 50 characters each.
    
</dd>
</dl>

<dl>
<dd>

**memberAffiliatePercentage:** `*float64` — The commission rate members earn.
    
</dd>
</dl>

<dl>
<dd>

**memberAffiliateStatus:** `*whopsdk.CreateProductsRequestMemberAffiliateStatus` — The enrollment status in the member affiliate program.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Custom key-value pairs to store on the product.
    
</dd>
</dl>

<dl>
<dd>

**productTaxCodeID:** `*string` — The unique identifier of the tax classification code. See the available [product categories](https://docs.numeral.com/essentials/product-categories).
    
</dd>
</dl>

<dl>
<dd>

**redirectPurchaseURL:** `*string` — A URL to redirect the customer to after purchase.
    
</dd>
</dl>

<dl>
<dd>

**route:** `*string` — The URL slug for the product's public link.
    
</dd>
</dl>

<dl>
<dd>

**sendWelcomeMessage:** `*bool` — Whether to send an automated welcome message via support chat when a user joins this product. Defaults to true.
    
</dd>
</dl>

<dl>
<dd>

**title:** `string` — The display name of the product. Maximum 80 characters.
    
</dd>
</dl>

<dl>
<dd>

**visibility:** `*string` — Whether the product is visible to customers.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Products.Retrieve(ID) -> *whopsdk.Product</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a product. Public — no credentials.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveProductsRequest{
    ID: "id",
}
client.Products.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the product.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Products.Delete(ID) -> *whopsdk.DeleteProductsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a product. Only products with no memberships, entries, reviews, or invoices can be deleted.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteProductsRequest{
    ID: "id",
}
client.Products.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the product.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Products.Update(ID, request) -> *whopsdk.Product</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an existing product.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateProductsRequest{
    ID: "id",
}
client.Products.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the product.
    
</dd>
</dl>

<dl>
<dd>

**bannerImage:** `*whopsdk.UpdateProductsRequestBannerImage` — A wide image for the product, shown on the product page and on listing cards. Pass `{ id }` for an existing attachment or `{ direct_upload_id }` for a completed direct upload; `null` removes it.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — A written description displayed on the product page.
    
</dd>
</dl>

<dl>
<dd>

**headline:** `*string` — A short marketing headline for the product page.
    
</dd>
</dl>

<dl>
<dd>

**labels:** `[]string` — Labels used to group products into collections. Replaces the existing labels. Send an empty array to clear them.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Custom key-value pairs to store on the product.
    
</dd>
</dl>

<dl>
<dd>

**productTaxCodeID:** `*string` — The unique identifier of the tax classification code. See the available [product categories](https://docs.numeral.com/essentials/product-categories).
    
</dd>
</dl>

<dl>
<dd>

**sendWelcomeMessage:** `*bool` — Whether to send an automated welcome message via support chat when a user joins this product.
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — The display name of the product.
    
</dd>
</dl>

<dl>
<dd>

**visibility:** `*string` — Whether the product is visible to customers.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Products.Publish(ID) -> *whopsdk.Product</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Submits a product to the whop.com marketplace for review. The product moves to `pending_review`; a Whop reviewer approves it before it goes live.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.PublishProductsRequest{
    ID: "id",
}
client.Products.Publish(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the product, prefixed `prod_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Products.Unpublish(ID) -> *whopsdk.Product</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Removes a product from the whop.com marketplace. The product moves to `not_available`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UnpublishProductsRequest{
    ID: "id",
}
client.Products.Unpublish(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the product, prefixed `prod_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Promo Codes
<details><summary><code>client.PromoCodes.List() -> *whopsdk.ListPromoCodesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists promo codes for an account with cursor pagination, filters, and sorting.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListPromoCodesRequest{
    AccountID: "account_id",
    ProductIDs: []*string{
        whopsdk.String(
            "prod_xxxxxxxxxxxxxx",
        ),
    },
    PlanIDs: []*string{
        whopsdk.String(
            "plan_xxxxxxxxxxxxxx",
        ),
    },
}
client.PromoCodes.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string` — Account whose promo codes are listed (`biz_` tag).
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.ListPromoCodesRequestStatus` — Promo-code status. `expired` groups inactive and archived codes.
    
</dd>
</dl>

<dl>
<dd>

**productIDs:** `*string` — Only promo codes scoped to these product IDs.
    
</dd>
</dl>

<dl>
<dd>

**planIDs:** `*string` — Only promo codes scoped to these plan IDs.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — Only promo codes created before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — Only promo codes created after this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListPromoCodesRequestOrder` — Sort field.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListPromoCodesRequestDirection` — Sort direction.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Number of promo codes to return from the start of the window.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor to paginate forwards from.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Number of promo codes to return from the end of the window.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Cursor to paginate backwards from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PromoCodes.Create(request) -> *whopsdk.PromoCode</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a promo code for an account. First-party sessions may attach an affiliate.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreatePromoCodesRequest{
    AccountID: "biz_xxxxxxxxxxxxxx",
    AmountOff: 25,
    BaseCurrency: whopsdk.CreatePromoCodesRequestBaseCurrencyUsd,
    Code: "AFFILIATE25",
    NewUsersOnly: true,
    PromoDurationMonths: 3,
    PromoType: whopsdk.CreatePromoCodesRequestPromoTypePercentage,
}
client.PromoCodes.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**amountOff:** `float64` 
    
</dd>
</dl>

<dl>
<dd>

**baseCurrency:** `*whopsdk.CreatePromoCodesRequestBaseCurrency` 
    
</dd>
</dl>

<dl>
<dd>

**churnedUsersOnly:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**code:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**existingMembershipsOnly:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**expiresAt:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**newUsersOnly:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**onePerCustomer:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**planIDs:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**productID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**promoDurationMonths:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**promoType:** `*whopsdk.CreatePromoCodesRequestPromoType` 
    
</dd>
</dl>

<dl>
<dd>

**stock:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**unlimitedStock:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PromoCodes.Retrieve(ID) -> *whopsdk.PromoCode</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a promo code by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrievePromoCodesRequest{
    ID: "id",
}
client.PromoCodes.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Promo code ID (`promo_` tag).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PromoCodes.Delete(ID) -> *whopsdk.DeletePromoCodesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Archives a promo code so it cannot be used in future checkouts.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeletePromoCodesRequest{
    ID: "id",
}
client.PromoCodes.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Promo code ID (`promo_` tag).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PromoCodes.Activate(ID) -> *whopsdk.PromoCode</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Turns an inactive promo code back on so it can be redeemed at checkout.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ActivatePromoCodesRequest{
    ID: "id",
}
client.PromoCodes.Activate(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Promo code ID (`promo_` tag).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PromoCodes.Deactivate(ID) -> *whopsdk.PromoCode</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Turns off an active promo code so it can no longer be redeemed at checkout.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeactivatePromoCodesRequest{
    ID: "id",
}
client.PromoCodes.Deactivate(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Promo code ID (`promo_` tag).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Reactions
<details><summary><code>client.Reactions.List() -> *whopsdk.ListReactionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of emoji reactions on a specific message or forum post, sorted by most recent.

Required permissions (one of):
 - `chat:read`
 - `dms:read`
 - `forum:read`
 - `livestream:chat:read`
 - `support_chat:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListReactionsRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    ResourceID: "resource_id",
}
client.Reactions.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**resourceID:** `string` — The unique identifier of the message or forum post to list reactions for.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reactions.Create(request) -> *whopsdk.Reaction</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add an emoji reaction or poll vote to a message or forum post. In forums, the reaction is always a like.

Required permissions (one of):
 - `chat:read`
 - `dms:read`
 - `forum:read`
 - `livestream:chat:read`
 - `support_chat:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateReactionsRequest{
    ResourceID: "resource_id",
}
client.Reactions.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**emoji:** `*string` — The emoji to react with, in shortcode or unicode format. For example, ':heart:' or a unicode emoji. Ignored in forums where reactions are always likes.
    
</dd>
</dl>

<dl>
<dd>

**pollOptionID:** `*string` — The unique identifier of a poll option to vote for. Only valid when the target message or post contains a poll.
    
</dd>
</dl>

<dl>
<dd>

**resourceID:** `string` — The unique identifier of the message or forum post to react to.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reactions.Retrieve(ID) -> *whopsdk.Reaction</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing reaction.

Required permissions (one of):
 - `chat:read`
 - `dms:read`
 - `forum:read`
 - `livestream:chat:read`
 - `support_chat:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveReactionsRequest{
    ID: "reac_xxxxxxxxxxxxxxxxxxxxxx",
}
client.Reactions.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the reaction to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reactions.Delete(ID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove an emoji reaction from a message or forum post. Only the reaction author or a channel admin can remove a reaction.

Required permissions (one of):
 - `chat:read`
 - `dms:read`
 - `forum:read`
 - `livestream:chat:read`
 - `support_chat:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteReactionsRequest{
    ID: "reac_xxxxxxxxxxxxxxxxxxxxxx",
}
client.Reactions.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the reaction to remove, or the identifier of the message or forum post to remove a reaction from. When passing a message or post ID, you must also provide the emoji argument.
    
</dd>
</dl>

<dl>
<dd>

**emoji:** `*string` — The emoji to remove, in shortcode or unicode format. For example, ':heart:' or a unicode emoji. Required when the id refers to a message or post instead of a reaction.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Recommended Actions
<details><summary><code>client.RecommendedActions.List() -> *whopsdk.ListRecommendedActionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the recommended action chains for an account — short sequences of actions (create a product, price it, publish it) the account should run next, gated on what it already has.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListRecommendedActionsRequest{}
client.RecommendedActions.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — Account ID, prefixed `biz_`. Defaults to the API key's own account.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.RecommendedActions.Retrieve(ID) -> *whopsdk.AccountRecommendedActionChain</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a recommended action chain by id, including chains that have already been run. Seeded chains are reconstructed from their hard-coded chain; generated chains are read from the account's stored chain, with each step's filled-in input.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveRecommendedActionsRequest{
    ID: "id",
}
client.RecommendedActions.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Chain ID from the list endpoint, e.g. `rac_seed_start_selling_9f2c1a7b04`.
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `*string` — Account ID, prefixed `biz_`. Defaults to the API key's own account.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.RecommendedActions.Run(ID) -> *whopsdk.RunRecommendedActionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Records that the caller ran a recommended action chain. Nothing is executed server-side yet — the client follows the chain's step CTAs itself; this writes the `recommended_action_chain.executed` analytics event.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RunRecommendedActionsRequest{
    ID: "id",
}
client.RecommendedActions.Run(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Chain ID from the list endpoint, e.g. `rac_seed_start_selling_9f2c1a7b04`.
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `*string` — Account ID, prefixed `biz_`. Defaults to the API key's own account.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.RecommendedActions.ListExecutions(ID) -> *whopsdk.ListExecutionsRecommendedActionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the per-step record of a recommended action chain the server ran — one entry per step in position order, each carrying its current status and, once the step completed, the API response it produced. A chain that was never run server-side returns an empty list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListExecutionsRecommendedActionsRequest{
    ID: "id",
}
client.RecommendedActions.ListExecutions(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Chain ID from the list endpoint.
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `*string` — Account ID, prefixed `biz_`. Defaults to the API key's own account.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Refunds
<details><summary><code>client.Refunds.List() -> *whopsdk.ListRefundsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of refunds, with optional filtering by payment, company, user, and creation date.

Required permissions:
 - `payment:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListRefundsRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    PaymentID: whopsdk.String(
        "pay_xxxxxxxxxxxxxx",
    ),
    CompanyID: whopsdk.String(
        "biz_xxxxxxxxxxxxxx",
    ),
    UserID: whopsdk.String(
        "user_xxxxxxxxxxxxx",
    ),
    CreatedBefore: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
    CreatedAfter: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
}
client.Refunds.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**paymentID:** `*string` — Filter refunds to only those associated with this specific payment.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `*string` — Filter refunds to only those belonging to this company.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — Filter refunds to only those associated with this specific user.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.Direction` 
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — Only return refunds created before this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — Only return refunds created after this timestamp.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Refunds.Retrieve(ID) -> *whopsdk.Refund</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing refund.

Required permissions:
 - `payment:basic:read`
 - `plan:basic:read`
 - `access_pass:basic:read`
 - `member:email:read`
 - `member:basic:read`
 - `member:phone:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveRefundsRequest{
    ID: "rf_xxxxxxxxxxxxxxx",
}
client.Refunds.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the refund.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Resolution Center Cases
<details><summary><code>client.ResolutionCenterCases.List() -> *whopsdk.ListResolutionCenterCasesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists resolution center cases. Without `account_id` you get every case you can read — the ones you opened as a buyer and every account you are a team member of; the filters narrow that list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListResolutionCenterCasesRequest{}
client.ResolutionCenterCases.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — Only cases filed against this account (`biz_` tag). With read access to the account this lists its whole queue; without, only the cases you opened against it.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — Only cases opened by this customer — a `user_` tag, or `me` for the calling user. It narrows what you can already read, so `me` lists the cases you opened without the ones on accounts you are a team member of.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of cases to return (default 20, max 100).
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor; returns cases after this position.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — The number of cases to return from the end of the range.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — A cursor; returns cases before this position.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListResolutionCenterCasesRequestOrder` — The field to sort cases by.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListResolutionCenterCasesRequestDirection` — Sort direction.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.ListResolutionCenterCasesRequestStatusItem` — Only cases in these statuses. Repeat the parameter to pass several — one paginated list covers all of them.
    
</dd>
</dl>

<dl>
<dd>

**reason:** `*whopsdk.ListResolutionCenterCasesRequestReasonItem` — Only cases opened for these reasons. Repeat the parameter to pass several.
    
</dd>
</dl>

<dl>
<dd>

**outcome:** `*whopsdk.ListResolutionCenterCasesRequestOutcomeItem` — Only closed cases that ended these ways. Repeat the parameter to pass several.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Only cases created before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Only cases created after this ISO 8601 timestamp.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ResolutionCenterCases.Create(request) -> *whopsdk.ResolutionCenterCase</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Opens a case, as the customer, against one of your own payments. Provide the payment (`receipt_id`), the `reason`, and a `message`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateResolutionCenterCasesRequest{
    Message: "The mobile detailer never showed up for the Ceramic Coating appointment.",
    Reason: whopsdk.CreateResolutionCenterCasesRequestReasonFraudulent,
    ReceiptID: "pay_xxxxxxxxxxxxxx",
}
client.ResolutionCenterCases.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**attachments:** `[]*whopsdk.CreateResolutionCenterCasesRequestAttachmentsItem` 
    
</dd>
</dl>

<dl>
<dd>

**message:** `string` — The customer's explanation.
    
</dd>
</dl>

<dl>
<dd>

**reason:** `*whopsdk.CreateResolutionCenterCasesRequestReason` — What went wrong. Uses the same vocabulary as `/disputes`.
    
</dd>
</dl>

<dl>
<dd>

**receiptID:** `string` — The payment to open the case against (`pay_` tag).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ResolutionCenterCases.Summary() -> *whopsdk.SummaryResolutionCenterCasesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Aggregates the same cases `GET /resolution_center_cases` lists, using the same filters. Use it to build status tabs and issue filters without paging the whole list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.SummaryResolutionCenterCasesRequest{}
client.ResolutionCenterCases.Summary(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**groups:** `*whopsdk.SummaryResolutionCenterCasesRequestGroupsItem` — Which breakdowns to return, keyed by these names under `groups`. Repeat the parameter to ask for several; omit it for all of them.
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `*string` — The account to summarize cases for (`biz_` tag).
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — Only cases opened by this customer — a `user_` tag, or `me` for the calling user.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.SummaryResolutionCenterCasesRequestStatusItem` — Only cases in these statuses.
    
</dd>
</dl>

<dl>
<dd>

**reason:** `*whopsdk.SummaryResolutionCenterCasesRequestReasonItem` — Only cases opened for these reasons.
    
</dd>
</dl>

<dl>
<dd>

**outcome:** `*whopsdk.SummaryResolutionCenterCasesRequestOutcomeItem` — Only closed cases that ended these ways.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Only count cases created before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Only count cases created after this ISO 8601 timestamp.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ResolutionCenterCases.Retrieve(ID) -> *whopsdk.ResolutionCenterCase</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a single resolution center case with its full event timeline.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveResolutionCenterCasesRequest{
    ID: "id",
}
client.ResolutionCenterCases.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The resolution center case ID (`reso_` tag).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ResolutionCenterCases.Accept(ID, request) -> *whopsdk.ResolutionCenterCase</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Accepts the case in the customer's favor, as the merchant: refunds the payment in full and closes the case.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.AcceptResolutionCenterCasesRequest{
    ID: "id",
}
client.ResolutionCenterCases.Accept(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The resolution center case ID (`reso_` tag).
    
</dd>
</dl>

<dl>
<dd>

**attachments:** `[]*whopsdk.AcceptResolutionCenterCasesRequestAttachmentsItem` — Up to 3 evidence files, by existing file `id` or `direct_upload_id`.
    
</dd>
</dl>

<dl>
<dd>

**message:** `*string` — An optional note to the customer, recorded on the case timeline.
    
</dd>
</dl>

<dl>
<dd>

**terminateMembership:** `*bool` — Whether to also terminate the customer's membership.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ResolutionCenterCases.Appeal(ID, request) -> *whopsdk.ResolutionCenterCase</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Appeals a decision, as the customer, on a case that closed in the merchant's favor. Escalates the case to Whop for platform review. A case can be appealed once.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.AppealResolutionCenterCasesRequest{
    ID: "id",
    Message: "The coating is already flaking on the hood two weeks later.",
}
client.ResolutionCenterCases.Appeal(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The resolution center case ID (`reso_` tag).
    
</dd>
</dl>

<dl>
<dd>

**attachments:** `[]*whopsdk.AppealResolutionCenterCasesRequestAttachmentsItem` — Up to 3 evidence files, by existing file `id` or `direct_upload_id`.
    
</dd>
</dl>

<dl>
<dd>

**message:** `string` — Why you are appealing the decision.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ResolutionCenterCases.Deny(ID, request) -> *whopsdk.ResolutionCenterCase</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Denies the case, as the merchant: rejects the claim and closes the case with no refund.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DenyResolutionCenterCasesRequest{
    ID: "id",
    Message: "The ceramic coating was applied and the vehicle was collected on 2026-01-05.",
}
client.ResolutionCenterCases.Deny(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The resolution center case ID (`reso_` tag).
    
</dd>
</dl>

<dl>
<dd>

**attachments:** `[]*whopsdk.DenyResolutionCenterCasesRequestAttachmentsItem` — Up to 3 evidence files, by existing file `id` or `direct_upload_id`.
    
</dd>
</dl>

<dl>
<dd>

**message:** `string` — Why the claim is being denied. Shown to the customer.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ResolutionCenterCases.Events(ID) -> *whopsdk.EventsResolutionCenterCasesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the case timeline, newest first. Events the viewer is not allowed to see are omitted — a customer reads the customer-visible timeline, the merchant reads the full one.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.EventsResolutionCenterCasesRequest{
    ID: "id",
}
client.ResolutionCenterCases.Events(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The resolution center case ID (`reso_` tag).
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of events to return (default 20, max 100).
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor; returns events after this position.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — The number of events to return from the end of the range.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — A cursor; returns events before this position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ResolutionCenterCases.Reply(ID, request) -> *whopsdk.ResolutionCenterCase</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Replies to an open request for information on the case. As the merchant this answers Whop's request (valid while the case awaits your information); as the customer it provides the information requested from you. The actor is resolved from the credential.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ReplyResolutionCenterCasesRequest{
    ID: "id",
    Message: "Here are the before and after photos from the Burnet Rd bay.",
}
client.ResolutionCenterCases.Reply(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The resolution center case ID (`reso_` tag).
    
</dd>
</dl>

<dl>
<dd>

**attachments:** `[]*whopsdk.ReplyResolutionCenterCasesRequestAttachmentsItem` — Up to 3 evidence files, by existing file `id` or `direct_upload_id`.
    
</dd>
</dl>

<dl>
<dd>

**message:** `string` — The reply to add to the case.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ResolutionCenterCases.RequestInfo(ID, request) -> *whopsdk.ResolutionCenterCase</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Asks the customer for more information, as the merchant. Allowed up to 3 times per case before you must accept or deny it.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RequestInfoResolutionCenterCasesRequest{
    ID: "id",
}
client.ResolutionCenterCases.RequestInfo(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The resolution center case ID (`reso_` tag).
    
</dd>
</dl>

<dl>
<dd>

**attachments:** `[]*whopsdk.RequestInfoResolutionCenterCasesRequestAttachmentsItem` — Up to 3 evidence files, by existing file `id` or `direct_upload_id`.
    
</dd>
</dl>

<dl>
<dd>

**message:** `*string` — What you need from the customer.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ResolutionCenterCases.Withdraw(ID) -> *whopsdk.ResolutionCenterCase</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Withdraws (cancels) a case you opened, as the customer. Only possible while the case is still open.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.WithdrawResolutionCenterCasesRequest{
    ID: "id",
}
client.ResolutionCenterCases.Withdraw(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The resolution center case ID (`reso_` tag).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Reviews
<details><summary><code>client.Reviews.List() -> *whopsdk.ListReviewsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of customer reviews for a specific product, with optional filtering by star rating and creation date.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListReviewsRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    ProductID: "prod_xxxxxxxxxxxxx",
    MinStars: whopsdk.Int(
        42,
    ),
    MaxStars: whopsdk.Int(
        42,
    ),
    CreatedBefore: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
    CreatedAfter: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
}
client.Reviews.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**productID:** `string` — The unique identifier of the product to list reviews for.
    
</dd>
</dl>

<dl>
<dd>

**minStars:** `*int` — The minimum star rating to include in results, from 1 to 5 inclusive.
    
</dd>
</dl>

<dl>
<dd>

**maxStars:** `*int` — The maximum star rating to include in results, from 1 to 5 inclusive.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — Only return reviews created before this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — Only return reviews created after this timestamp.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reviews.Retrieve(ID) -> *whopsdk.Review</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing review.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveReviewsRequest{
    ID: "rev_xxxxxxxxxxxxxx",
}
client.Reviews.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the review to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Setup Intents
<details><summary><code>client.SetupIntents.List() -> *whopsdk.ListSetupIntentsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of setup intents for a company, with optional filtering by creation date. A setup intent securely collects and stores a member's payment method for future use without charging them immediately.

Required permissions:
 - `payment:setup_intent:read`
 - `member:basic:read`
 - `member:email:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListSetupIntentsRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    CompanyID: "biz_xxxxxxxxxxxxxx",
    CreatedBefore: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
    CreatedAfter: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
}
client.SetupIntents.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `string` — The unique identifier of the company to list setup intents for.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.Direction` 
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — Only return setup intents created before this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — Only return setup intents created after this timestamp.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SetupIntents.Create(request) -> *whopsdk.CreateSetupIntentsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Save a buyer's payment method for later without charging it. Provide a confirmation token for a method the buyer just supplied, or an existing payment method to re-verify. The buyer may still have a step to complete — 3D Secure, a hosted enrollment, linking a bank account — so poll the setup intent's status endpoint for what to do next.

Required permissions:
 - `payment:charge`
 - `member:basic:read`
 - `member:email:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateSetupIntentsRequest{
    CreateSetupIntentsRequestConfirmationToken: &whopsdk.CreateSetupIntentsRequestConfirmationToken{
        CompanyID: "biz_xxxxxxxxxxxxxx",
        ConfirmationToken: "ctok_xxxxxxxxxxxxxx",
    },
}
client.SetupIntents.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*whopsdk.CreateSetupIntentsRequest` — Parameters for CreateSetupIntent
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SetupIntents.Retrieve(ID) -> *whopsdk.SetupIntent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing setup intent.

Required permissions:
 - `payment:setup_intent:read`
 - `member:basic:read`
 - `member:email:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveSetupIntentsRequest{
    ID: "sint_xxxxxxxxxxxxx",
}
client.SetupIntents.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the setup intent.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SetupIntents.UpdateReturnURL(SetupIntentID, request) -> *whopsdk.SetupStatus</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Changes where the buyer lands after completing an off-site step, up until they return. Accepts either a secret key or the setup's own `client_secret`, so the surface that knows the final destination can set it.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateReturnURLSetupIntentsRequest{
    SetupIntentID: "setup_intent_id",
    ReturnURL: "https://shinetime.example/checkout/thanks",
}
client.SetupIntents.UpdateReturnURL(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**setupIntentID:** `string` — The unique identifier of the setup intent.
    
</dd>
</dl>

<dl>
<dd>

**returnURL:** `string` — Where the buyer continues after completing an off-site step. Must be an absolute https URL without credentials (http is allowed for localhost), at most 2,048 characters.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SetupIntents.RetrieveStatus(SetupIntentID) -> *whopsdk.SetupStatus</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves how far a setup has got and what the buyer must do next, if anything. Collection runs in the background, so poll this rather than reading the create response. Accepts either a secret key or the setup's own `client_secret`, so the surface collecting the payment method can poll it directly.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveStatusSetupIntentsRequest{
    SetupIntentID: "setup_intent_id",
}
client.SetupIntents.RetrieveStatus(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**setupIntentID:** `string` — The unique identifier of the setup intent.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Shipments
<details><summary><code>client.Shipments.List() -> *whopsdk.ListShipmentsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of shipments for an account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListShipmentsRequest{
    PaymentID: []*string{
        whopsdk.String(
            "pay_xxxxxxxxxxxxxx",
        ),
    },
}
client.Shipments.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The account to list shipments for. Defaults to the acting account.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.ListShipmentsRequestStatus` — Filter to shipments with this delivery status.
    
</dd>
</dl>

<dl>
<dd>

**paymentID:** `*string` — Only shipments fulfilling these payments, each prefixed `pay_`. Repeat the parameter to pass several, up to 100 per request — one paginated list covers all of them.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Return shipments created before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Return shipments created after this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListShipmentsRequestOrder` — The field to sort by.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListShipmentsRequestDirection` — The sort direction.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of shipments to return.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor; returns shipments after this position.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — The number of shipments to return from the end of the range.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — A cursor; returns shipments before this position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Shipments.Create(request) -> *whopsdk.Shipment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Attaches a carrier tracking number to a payment and begins tracking it.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateShipmentsRequest{
    PaymentID: "pay_xxxxxxxxxxxxxx",
    TrackingNumber: "1Z999AA10123456784",
}
client.Shipments.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The unique identifier of the account, prefixed `biz_`.
    
</dd>
</dl>

<dl>
<dd>

**paymentID:** `string` — The payment to attach the shipment to, prefixed `pay_`.
    
</dd>
</dl>

<dl>
<dd>

**trackingNumber:** `string` — The carrier-assigned tracking number.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Shipments.Retrieve(ID) -> *whopsdk.Shipment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a shipment by its id, or by the payment id it fulfills.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveShipmentsRequest{
    ID: "id",
}
client.Shipments.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The shipment id (`ship_`), or the payment id (`pay_`) it fulfills.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Shipments.Update(ID, request) -> *whopsdk.Shipment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a shipment's tracking number and re-tracks it with the carrier.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateShipmentsRequest{
    ID: "id",
    TrackingNumber: "9400111899223456789012",
}
client.Shipments.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The shipment id (`ship_`), or the payment id (`pay_`) it fulfills.
    
</dd>
</dl>

<dl>
<dd>

**trackingNumber:** `string` — The new carrier-assigned tracking number.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Social Accounts
<details><summary><code>client.SocialAccounts.List() -> *whopsdk.ListSocialAccountsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the social accounts linked to an account or user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListSocialAccountsRequest{}
client.SocialAccounts.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The Account that the social accounts are connected to. Provide either this or user_id.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — The User that the social accounts are connected to. Provide either this or account_id.
    
</dd>
</dl>

<dl>
<dd>

**platform:** `*whopsdk.ListSocialAccountsRequestPlatform` — Only return social accounts for the platform that is specified.
    
</dd>
</dl>

<dl>
<dd>

**verified:** `*bool` — Only return social accounts that are verified on the platform.
    
</dd>
</dl>

<dl>
<dd>

**scopes:** `*whopsdk.ListSocialAccountsRequestScopesItem` — Only return social accounts that have these scopes.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of social accounts to return.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor to fetch the page after (from page_info.end_cursor).
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — The number of social accounts to return from the end of the range.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Cursor to fetch the page before (from page_info.start_cursor).
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListSocialAccountsRequestOrder` — The field to sort social accounts by.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListSocialAccountsRequestDirection` — Sort direction.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SocialAccounts.Create(request) -> *whopsdk.SocialAccount</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates or returns a Whop-managed Facebook page for an account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateSocialAccountsRequest{
    Platform: whopsdk.CreateSocialAccountsRequestPlatformFacebook,
}
client.SocialAccounts.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The Account (biz_ identifier) to create the social account for. An account-scoped API key may omit this to default to its own account. Account API keys cannot update their own account's branding through Update Account; use a user-authenticated path.
    
</dd>
</dl>

<dl>
<dd>

**platform:** `*whopsdk.CreateSocialAccountsRequestPlatform` — The platform to create the social account on. `facebook` requires the account's `banner_image`, `logo`, and `description`; configure them with [Update Account](/api-reference/beta/accounts/update-account).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SocialAccounts.Connect(request) -> *whopsdk.ConnectSocialAccountsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Starts an OAuth connection flow and returns an authorize_url where the user can connect a social account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ConnectSocialAccountsRequest{
    Platform: whopsdk.ConnectSocialAccountsRequestPlatformMetaBusiness,
}
client.SocialAccounts.Connect(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The Account (biz_ identifier) to connect the social account for. An account-scoped API key may omit this to default to its own account.
    
</dd>
</dl>

<dl>
<dd>

**platform:** `*whopsdk.ConnectSocialAccountsRequestPlatform` — The platform to connect the social account on. Supported options are `meta_business` and `tiktok`.
    
</dd>
</dl>

<dl>
<dd>

**redirectURL:** `*string` — The Whop URL to redirect the user to after they finish connecting.
    
</dd>
</dl>

<dl>
<dd>

**scopes:** `[]*whopsdk.ConnectSocialAccountsRequestScopesItem` — Capabilities to grant for the connected social account. Use `advertise` when connecting a Meta Business or TikTok account for ads.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SocialAccounts.Delete(ID) -> *whopsdk.DeleteSocialAccountsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Disconnects a social account from an account or user without deleting the underlying platform account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteSocialAccountsRequest{
    ID: "id",
}
client.SocialAccounts.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the social account to disconnect.
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `*string` — The Account that the social account is connected to. Provide either this or user_id.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — The User that the social account is connected to. Provide either this or account_id.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SocialAccounts.LeadForms(ID) -> *whopsdk.LeadFormsSocialAccountsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the active lead (instant) forms that already exist on a connected Facebook page, so an ad can reuse one as its `lead_gen_form_id` instead of authoring a new form. Every active form comes back in a single response — the list is not paginated.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.LeadFormsSocialAccountsRequest{
    ID: "id",
    AccountID: "account_id",
}
client.SocialAccounts.LeadForms(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The social account (a sacc_ identifier) whose lead forms to list.
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `string` — The Account (a biz_ identifier) the social account is connected to.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SocialAccounts.Posts(ID) -> *whopsdk.PostsSocialAccountsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the existing posts of a connected Facebook page, Instagram account, or TikTok account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.PostsSocialAccountsRequest{
    ID: "id",
    AccountID: "account_id",
}
client.SocialAccounts.Posts(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The social account (a sacc_ identifier) whose posts to list.
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `string` — The Account (a biz_ identifier) the social account is connected to.
    
</dd>
</dl>

<dl>
<dd>

**postID:** `*string` — Return only the single post with this platform id, instead of the full list.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of posts to return.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor to fetch the page after (from page_info.end_cursor).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Stats
<details><summary><code>client.Stats.List() -> *whopsdk.ListStatsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists every metric you can query, with its unit and the properties you can filter or break it down by.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Stats.List(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.DescribeStats() -> *whopsdk.DescribeStatsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Describe available stats schema. Without resource returns root nodes and metrics. With resource returns node columns, associations, and available metrics.

Required permissions:
 - `stats:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DescribeStatsRequest{
    CompanyID: whopsdk.String(
        "biz_xxxxxxxxxxxxxx",
    ),
    UserID: whopsdk.String(
        "user_xxxxxxxxxxxxx",
    ),
}
client.Stats.DescribeStats(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**resource:** `*string` — Resource path using : as separator (e.g., 'receipts', 'payments:membership', 'receipts:gross_revenue').
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `*string` — Scope query to a specific company.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — Scope query to a specific user.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.MetricStats() -> *whopsdk.MetricStatsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Query an aggregated metric. Returns data grouped by period with optional breakdowns.

Required permissions:
 - `stats:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.MetricStatsRequest{
    Resource: "resource",
    From: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
    To: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
    CompanyID: whopsdk.String(
        "biz_xxxxxxxxxxxxxx",
    ),
    UserID: whopsdk.String(
        "user_xxxxxxxxxxxxx",
    ),
}
client.Stats.MetricStats(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**resource:** `string` — Metric resource using : as separator (e.g., 'receipts:gross_revenue', 'members:new_users').
    
</dd>
</dl>

<dl>
<dd>

**granularity:** `*string` — Time granularity (daily, weekly, monthly).
    
</dd>
</dl>

<dl>
<dd>

**breakdowns:** `*string` — Columns to break down the metric by.
    
</dd>
</dl>

<dl>
<dd>

**filters:** `map[string]any` — Key-value pairs to filter the data.
    
</dd>
</dl>

<dl>
<dd>

**timeZone:** `*string` — IANA timezone for period bucketing (e.g. 'America/New_York'). Defaults to UTC. Only applies to ClickHouse metrics.
    
</dd>
</dl>

<dl>
<dd>

**from:** `*time.Time` — Start of time range (unix timestamp).
    
</dd>
</dl>

<dl>
<dd>

**to:** `*time.Time` — End of time range (unix timestamp).
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `*string` — Scope query to a specific company.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — Scope query to a specific user.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.RawStats() -> *whopsdk.RawStatsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Query raw data from a resource. Returns paginated rows with all columns.

Required permissions:
 - `stats:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RawStatsRequest{
    Resource: "resource",
    From: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
    To: whopsdk.Time(
        whopsdk.MustParseDateTime(
            "2023-12-01T05:00:00Z",
        ),
    ),
    Limit: whopsdk.Int(
        42,
    ),
    CompanyID: whopsdk.String(
        "biz_xxxxxxxxxxxxxx",
    ),
    UserID: whopsdk.String(
        "user_xxxxxxxxxxxxx",
    ),
}
client.Stats.RawStats(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**resource:** `string` — Resource path using : as separator (e.g., 'members', 'payments:membership').
    
</dd>
</dl>

<dl>
<dd>

**from:** `*time.Time` — Start of time range (unix timestamp).
    
</dd>
</dl>

<dl>
<dd>

**to:** `*time.Time` — End of time range (unix timestamp).
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of records to return (max 10000).
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Pagination cursor for next page.
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*string` — Column to sort by.
    
</dd>
</dl>

<dl>
<dd>

**sortDirection:** `*whopsdk.Direction` 
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `*string` — Scope query to a specific company.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — Scope query to a specific user.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.Retrieve(Metric) -> *whopsdk.RetrieveStatsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a metric as a time series of points for an account or user over a time range. The `market_prices` metric is public and requires no authentication.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveStatsRequest{
    Metric: "metric",
    From: "from",
    To: "to",
    AdCampaignIDs: []*string{
        whopsdk.String(
            "adcamp_xxxxxxxxxxxxxx",
        ),
    },
    AdGroupIDs: []*string{
        whopsdk.String(
            "adgrp_xxxxxxxxxxxxxx",
        ),
    },
    AdIDs: []*string{
        whopsdk.String(
            "ad_xxxxxxxxxxxxxx",
        ),
    },
}
client.Stats.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**metric:** `string` — The metric to retrieve, for example net_revenue. Use GET /stats to see every metric key. The metric sets the unit and the properties you can filter or break down by.
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `*string` — The account this query concerns, for example biz_AbC123.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — The user this query concerns, for example user_AbC123. Available on metrics that support user subjects, such as account_balance.
    
</dd>
</dl>

<dl>
<dd>

**from:** `string` — Start of the range — a date (YYYY-MM-DD), expanded to the start of that day, or an ISO 8601 timestamp (for example 2026-07-16T16:37:00Z), used exactly.
    
</dd>
</dl>

<dl>
<dd>

**to:** `string` — End of the range — a date (YYYY-MM-DD), expanded to the end of that day, or an ISO 8601 timestamp (for example 2026-07-17T16:37:00Z), used exactly.
    
</dd>
</dl>

<dl>
<dd>

**interval:** `*whopsdk.RetrieveStatsRequestInterval` — How wide each point is. Defaults to day. Snapshot metrics are day-only.
    
</dd>
</dl>

<dl>
<dd>

**breakdownBy:** `*string` — Split the metric out by one of its properties — each point gets a breakdown array. For example breakdown_by=currency returns an entry for usd, an entry for eur, and so on.
    
</dd>
</dl>

<dl>
<dd>

**convertTo:** `*string` — Display currency for money metrics — every amount is converted into this ISO currency using the exchange rate on each period's date. Defaults to usd. For the ads metrics (ad_spend, ad_delivery), pass the account's ads reporting currency to match the ad entity endpoints. On transaction metrics, it is ignored when you filter or break down by currency (those report the original transaction currency, unconverted).
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*string` — Select the source currency or asset on metrics that list currency. For transaction metrics, for example currency=eur, values are reported without conversion. For market_prices, use btc or xaut and convert_to=usd. Pair with breakdown_by=currency to split a metric by currency.
    
</dd>
</dl>

<dl>
<dd>

**timeZone:** `*string` — IANA time zone to bucket the series in, for example America/New_York. Defaults to UTC. Not accepted by snapshot metrics, which are UTC only.
    
</dd>
</dl>

<dl>
<dd>

**paymentMethod:** `*string` — Filter to a single payment method, for example card or crypto. Available on metrics that list payment_method.
    
</dd>
</dl>

<dl>
<dd>

**cardNetwork:** `*string` — Filter to a single card brand, for example visa. A refinement of payment_method=card. Available on metrics that list card_network.
    
</dd>
</dl>

<dl>
<dd>

**disputeReason:** `*string` — Filter disputes to a normalized reason, for example product_not_received. Pair with breakdown_by=dispute_reason to split dispute counts by reason.
    
</dd>
</dl>

<dl>
<dd>

**source:** `*string` — Filter to a single GMV source, for example payments — or, on the traffic metrics, a visit source (whop_ads, direct, or a utm_source value). Pair with breakdown_by=source to split by source. Available on metrics that list source.
    
</dd>
</dl>

<dl>
<dd>

**hostname:** `*string` — Filter traffic metrics to one website hostname, for example shop.example.com. Pair with breakdown_by=hostname to split by website.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Filter traffic metrics to one page — a hostname plus normalized path, for example shop.example.com/pricing. Pair with breakdown_by=page to split by page.
    
</dd>
</dl>

<dl>
<dd>

**deviceType:** `*string` — Filter traffic metrics to one device type: desktop, mobile, tablet, or unknown. Pair with breakdown_by=device_type to split by device.
    
</dd>
</dl>

<dl>
<dd>

**countryCode:** `*string` — Filter traffic metrics to one visitor country (uppercase ISO 3166-1 alpha-2, for example US). Pair with breakdown_by=country_code to split by country.
    
</dd>
</dl>

<dl>
<dd>

**eventName:** `*string` — Filter the events metric to one tracked event name, for example pixel.page or pixel.custom. Pair with breakdown_by=event_name to split by event.
    
</dd>
</dl>

<dl>
<dd>

**eventType:** `*whopsdk.RetrieveStatsRequestEventType` — Filter the events metric to a canonical group of events: page_view (pixel page views plus whop.com store views), checkout_start (hosted and embedded checkout views), or other. Pair with breakdown_by=event_type to split by group.
    
</dd>
</dl>

<dl>
<dd>

**customName:** `*string` — Filter the events metric to one merchant-defined custom event name. Only valid alongside event_name=pixel.custom. Pair with breakdown_by=custom_name to split custom events by name.
    
</dd>
</dl>

<dl>
<dd>

**segment:** `*string` — Filter to a single wallet-balance segment, for example available. Pair with breakdown_by=segment to split the balance. Available on metrics that list segment.
    
</dd>
</dl>

<dl>
<dd>

**category:** `*string` — Filter to a single balance-activity category, for example payments. Pair with breakdown_by=category to split the activity. Available on metrics that list category.
    
</dd>
</dl>

<dl>
<dd>

**merchant:** `*string` — Filter to a single cashback merchant bucket, for example whop-ads. Pair with breakdown_by=merchant to split cashback by merchant. Available on metrics that list merchant.
    
</dd>
</dl>

<dl>
<dd>

**feeType:** `*string` — Filter to a single fee type. Pair with breakdown_by=fee_type to split fees by type. Available on metrics that list fee_type.
    
</dd>
</dl>

<dl>
<dd>

**product:** `*string` — Filter to a single product (access pass id), for example prod_AbC123. Pair with breakdown_by=product. Available on metrics that list product.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*string` — Filter to a single membership status. Pair with breakdown_by=status. Available on metrics that list status.
    
</dd>
</dl>

<dl>
<dd>

**accessLevel:** `*string` — Filter to a single access level. Pair with breakdown_by=access_level. Available on metrics that list access_level.
    
</dd>
</dl>

<dl>
<dd>

**mostRecentAction:** `*string` — Filter to a single most-recent member action. Pair with breakdown_by=most_recent_action. Available on metrics that list most_recent_action.
    
</dd>
</dl>

<dl>
<dd>

**referredUserID:** `*string` — Filter a referral metric to the businesses attributed to one person you referred, for example user_AbC123. Available on metrics that list referred_user_id.
    
</dd>
</dl>

<dl>
<dd>

**adCampaignIDs:** `*string` — Ad campaign ids (adcamp_...) to scope the report to; stats are summed across them. Available on metrics that list ad_campaign_ids.
    
</dd>
</dl>

<dl>
<dd>

**adGroupIDs:** `*string` — Ad group ids (adgrp_...) to scope the report to; stats are summed across them. Available on metrics that list ad_group_ids.
    
</dd>
</dl>

<dl>
<dd>

**adIDs:** `*string` — Ad ids (ad_...) to scope the report to; stats are summed across them. Available on metrics that list ad_ids.
    
</dd>
</dl>

<dl>
<dd>

**snapshotWindow:** `*whopsdk.RetrieveStatsRequestSnapshotWindow` — Window used by a snapshot metric. Ordinary snapshots accept 30d as their trailing activity window. Cohorted dispute metrics accept 7d or 28d as the sales-transaction pool; their attribution window is fixed in the metric name. Each metric lists its accepted values in the catalog.
    
</dd>
</dl>

<dl>
<dd>

**event:** `*string` — Filter the events metric to one or more full event names, for example payment.completed or pixel.lead. Comma-separate several to break the metric down by each event. Available on metrics that list event.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## SupportChannels
<details><summary><code>client.SupportChannels.List() -> *whopsdk.ListSupportChannelsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of support channels for a specific company, with optional filtering by resolution status and custom sorting.

Required permissions:
 - `support_chat:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListSupportChannelsRequest{
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
    CompanyID: whopsdk.String(
        "biz_xxxxxxxxxxxxxx",
    ),
}
client.SupportChannels.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `*string` — The unique identifier of the company to list support channels for. Includes channels of child companies. When omitted, returns support channels across all companies the user has access to.
    
</dd>
</dl>

<dl>
<dd>

**view:** `*whopsdk.SupportChannelView` 
    
</dd>
</dl>

<dl>
<dd>

**open:** `*bool` — Whether to filter by open or resolved support channels. Set to true to only return channels awaiting a response, or false for resolved channels.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.Direction` 
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.MessageChannelOrder` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SupportChannels.Create(request) -> *whopsdk.SupportChannel</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Open a new support channel between a company team member and a customer. Returns the existing channel if one already exists for that user.

Required permissions:
 - `support_chat:create`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateSupportChannelsRequest{
    CompanyID: "biz_xxxxxxxxxxxxxx",
    UserID: "user_xxxxxxxxxxxxx",
}
client.SupportChannels.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyID:** `string` — The unique identifier of the company to create the support channel in.
    
</dd>
</dl>

<dl>
<dd>

**customName:** `*string` — Optional custom display name for the support channel.
    
</dd>
</dl>

<dl>
<dd>

**notificationsEnabled:** `*bool` — Whether Whop app notifications are enabled for this support channel. Webhooks still fire.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `string` — The user ID (e.g. 'user_xxxxx') or username of the customer to open a support channel for.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SupportChannels.Retrieve(ID) -> *whopsdk.SupportChannel</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing support channel.

Required permissions:
 - `support_chat:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveSupportChannelsRequest{
    ID: "id",
}
client.SupportChannels.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The unique identifier of the support channel to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Swaps
<details><summary><code>client.Swaps.List() -> *whopsdk.ListSwapsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the account's completed or pending swaps — currently just the latest one.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListSwapsRequest{
    AccountID: "account_id",
}
client.Swaps.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string` — Business or user account ID (biz_* / user_*).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Swaps.Create(request) -> *whopsdk.CreateSwapsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Swaps one token for another from the account's wallet, or converts between fiat currencies in the account's ledger at the mid-market rate. Crypto swaps finish in the background — check the swap for its status.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateSwapsRequest{
    AccountID: "biz_xxxxxxxxxxxxxx",
    FromToken: "usd",
    ToToken: "cad",
}
client.Swaps.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string` — Business or user account ID (biz_* / user_*).
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*string` — Source token amount. Required for crypto swaps. For fiat pairs: the amount of from_token to convert at the mid-market rate; omit (along with to_amount) to repay the full negative to_token balance instead.
    
</dd>
</dl>

<dl>
<dd>

**fromChain:** `*whopsdk.CreateSwapsRequestFromChain` — Source chain name or chain ID. Defaults to the source token's chain when omitted.
    
</dd>
</dl>

<dl>
<dd>

**fromToken:** `string` — Source token contract address or ticker symbol, such as "USDT".
    
</dd>
</dl>

<dl>
<dd>

**slippageBps:** `*int` — Maximum slippage tolerance in basis points.
    
</dd>
</dl>

<dl>
<dd>

**toAmount:** `*string` — Fiat pairs only: sizes a partial repayment of the negative to_token balance, denominated in to_token. Must not exceed the debt. Mutually exclusive with amount.
    
</dd>
</dl>

<dl>
<dd>

**toChain:** `*whopsdk.CreateSwapsRequestToChain` — Destination chain name or chain ID. Defaults to the destination token's chain when omitted.
    
</dd>
</dl>

<dl>
<dd>

**toToken:** `string` — Destination token contract address or ticker symbol, such as "XAUT".
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Swaps.CreateQuote(request) -> *whopsdk.CreateQuoteSwapsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Previews the price of a swap. Fiat pairs quote the in-ledger mid-market conversion — the same rate creating the swap fills at. No funds move and nothing is saved.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateQuoteSwapsRequest{
    Amount: "100",
    FromToken: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    ToToken: "0x1b64b9025eebb9a6239575df9ea4b9ac46d4d193",
}
client.Swaps.CreateQuote(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**amount:** `string` — Source token amount.
    
</dd>
</dl>

<dl>
<dd>

**fromAddress:** `*string` — Source wallet address used for the quote.
    
</dd>
</dl>

<dl>
<dd>

**fromChain:** `*whopsdk.CreateQuoteSwapsRequestFromChain` — Source chain name or chain ID. Defaults to the source token's chain when omitted.
    
</dd>
</dl>

<dl>
<dd>

**fromToken:** `string` — Source token contract address or ticker symbol, such as "USDT".
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Metadata to include with the quote response.
    
</dd>
</dl>

<dl>
<dd>

**slippageBps:** `*int` — Maximum slippage tolerance in basis points.
    
</dd>
</dl>

<dl>
<dd>

**toAddress:** `*string` — Destination wallet address used for the quote.
    
</dd>
</dl>

<dl>
<dd>

**toChain:** `*whopsdk.CreateQuoteSwapsRequestToChain` — Destination chain name or chain ID. Defaults to the destination token's chain when omitted.
    
</dd>
</dl>

<dl>
<dd>

**toToken:** `string` — Destination token contract address or ticker symbol, such as "XAUT".
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Swaps.Retrieve(ID) -> *whopsdk.RetrieveSwapsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a single swap and its status.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveSwapsRequest{
    ID: "id",
}
client.Swaps.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Swap ID returned from POST /swaps.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Team Members
<details><summary><code>client.TeamMembers.List() -> *whopsdk.ListTeamMembersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists an account's team members, including pending invites (`status: "pending"`, `ausri_` ids; `user` is `null` for invites sent to an email with no Whop account yet). For accepted members, `email` requires the `company:authorized_user:email:read` scope and is `null` otherwise. Listing `role=workforce` is also allowed with the `bounty:create` scope.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListTeamMembersRequest{
    AccountID: "account_id",
}
client.TeamMembers.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string` — Account ID, prefixed `biz_`.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*whopsdk.ListTeamMembersRequestStatus` — Only return members with this status: `joined` (accepted members) or `pending` (pending invites). Both are returned by default.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — Only return the membership for this user ID, prefixed `user_`.
    
</dd>
</dl>

<dl>
<dd>

**role:** `*whopsdk.ListTeamMembersRequestRole` — Only return members with this role. `custom` matches members on a dashboard-managed custom role.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Only return members added before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Only return members added after this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListTeamMembersRequestOrder` — Field used to sort members.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListTeamMembersRequestDirection` — Sort direction. Defaults to `desc`.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Number of members to return. Defaults to 20; maximum 100.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor for the next page of members.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Number of members to return from the end of the window.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Cursor to paginate backwards from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TeamMembers.Create(request) -> *whopsdk.TeamMember</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Adds a member to an account's team with a system role. Identify them by exactly one of `user_id` or `email`. If the person has not yet accepted — or the email does not belong to a Whop account yet — an invitation is sent instead and the response is `202` with `{ "object": "team_member_invite", "invitation_sent": true }`. If they already have a pending invite, the request fails with a `400`. Custom roles cannot be granted via the API. Granting the `workforce` role is also allowed with the `bounty:create` scope.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateTeamMembersRequest{
    AccountID: "biz_xxxxxxxxxxxxxx",
    Role: whopsdk.CreateTeamMembersRequestRoleOwner,
}
client.TeamMembers.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string` — Account ID, prefixed `biz_`.
    
</dd>
</dl>

<dl>
<dd>

**email:** `*string` — Email address to invite. Mutually exclusive with `user_id`. If the email already belongs to a Whop account it is treated the same as passing that account's `user_id`; otherwise a pending invite is created for the email.
    
</dd>
</dl>

<dl>
<dd>

**role:** `*whopsdk.CreateTeamMembersRequestRole` — The system role to grant.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — The user to add to the team, prefixed `user_`. Mutually exclusive with `email`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TeamMembers.Retrieve(ID) -> *whopsdk.TeamMember</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a team member by ID. `email` requires the `company:authorized_user:email:read` scope and is `null` otherwise.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveTeamMembersRequest{
    ID: "id",
}
client.TeamMembers.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Team member ID — `ausr_` for accepted members, `ausri_` for pending invites.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TeamMembers.Delete(ID) -> *whopsdk.DeleteTeamMembersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Removes a team member from the account, or revokes a pending invite when given an `ausri_` ID. A user session may delete its own membership to leave the team without the delete scope. Removing a member on the `workforce` role is also allowed with the `bounty:create` scope. The account owner cannot be removed.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteTeamMembersRequest{
    ID: "id",
}
client.TeamMembers.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Team member ID — `ausr_` for accepted members, `ausri_` for pending invites.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TeamMembers.Update(ID, request) -> *whopsdk.TeamMember</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Changes a team member's system role. Requires a user session — account API keys cannot change member roles. The account owner's role cannot be changed, and you cannot change your own role.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateTeamMembersRequest{
    ID: "id",
    Role: whopsdk.UpdateTeamMembersRequestRoleOwner,
}
client.TeamMembers.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Team member ID — `ausr_` for accepted members, `ausri_` for pending invites.
    
</dd>
</dl>

<dl>
<dd>

**role:** `*whopsdk.UpdateTeamMembersRequestRole` — The system role to grant.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Topups
<details><summary><code>client.Topups.Create(request) -> *whopsdk.Topup</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add funds to a company's platform balance by charging a stored payment method. Top-ups have no fees or taxes and do not count as revenue.

Required permissions:
 - `payment:charge`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateTopupsRequest{
    Amount: 6.9,
    CompanyID: "biz_xxxxxxxxxxxxxx",
    Currency: whopsdk.CurrenciesUsd,
    PaymentMethodID: "pmt_xxxxxxxxxxxxxx",
}
client.Topups.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**amount:** `float64` — The amount to add to the balance in the specified currency. For example, 50.00 for $50.00 USD.
    
</dd>
</dl>

<dl>
<dd>

**companyID:** `string` — The unique identifier of the company to add funds to, starting with 'biz_'.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*whopsdk.Currencies` — The currency for the top-up amount, such as 'usd'.
    
</dd>
</dl>

<dl>
<dd>

**paymentMethodID:** `string` — The unique identifier of the stored payment method to charge for the top-up.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Transfers
<details><summary><code>client.Transfers.List() -> *whopsdk.ListTransfersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists an account's transfers.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListTransfersRequest{}
client.Transfers.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**originID:** `*string` — Filter to transfers sent from this account. Provide this or destination_id.
    
</dd>
</dl>

<dl>
<dd>

**destinationID:** `*string` — Filter to transfers received by this account. Provide this or origin_id.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListTransfersRequestOrder` — Sort column. Defaults to created_at.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListTransfersRequestDirection` — Sort direction. Defaults to desc.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Only transfers created strictly before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Only transfers created strictly after this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Number of transfers to return from the start of the window.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor to fetch the page after (from page_info.end_cursor).
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Number of transfers to return from the end of the window.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Cursor to fetch the page before (from page_info.start_cursor).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Transfers.Create(request) -> *whopsdk.CreateTransfersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Moves money between accounts, or into a claim link anyone with the URL can redeem.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateTransfersRequest{
    Amount: 25,
    OriginID: "biz_xxxxxxxxxxxxxx",
}
client.Transfers.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**amount:** `float64` — The amount to move, in the transfer currency. For example 25.00.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*string` — Currency, such as `usd`. Required for ledger transfers.
    
</dd>
</dl>

<dl>
<dd>

**destinationID:** `*string` — The recipient. Required for ledger and wallet_send (a user_/biz_/ldgr_ ID, or — for sends — an email). Omit for claim_link.
    
</dd>
</dl>

<dl>
<dd>

**expiresAt:** `*time.Time` — claim_link only. Link expiry as an ISO 8601 timestamp. Defaults to 24 hours from creation.
    
</dd>
</dl>

<dl>
<dd>

**idempotenceKey:** `*string` — Ledger transfers and wallet sends. A unique key that makes retries safe. Retrying with the same key returns the original transfer, or attaches to the original wallet send, instead of moving money twice.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Ledger transfers only. Custom key-value pairs attached to the transfer. Max 50 keys, 100 chars per key, 500 chars per string value.
    
</dd>
</dl>

<dl>
<dd>

**notes:** `*string` — Ledger transfers only. A short note describing the transfer.
    
</dd>
</dl>

<dl>
<dd>

**originID:** `string` — The account sending the funds. A user ID (user_xxx), account ID (biz_xxx), or ledger account ID (ldgr_xxx).
    
</dd>
</dl>

<dl>
<dd>

**redeemableCount:** `*int` — claim_link only. How many different users can claim the link. Defaults to 1.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*whopsdk.CreateTransfersRequestType` — The kind of money movement, which decides what comes back. Defaults to ledger. `ledger` moves credit between two Whop balances and returns a `transfer`; `wallet_send` sends USDT from the origin account's Ethereum wallet and returns a `send`; `claim_link` funds a shareable link anyone with the URL can redeem and returns a `claim_link`. A `ledger` transfer from a stablecoin-rails account settles on-chain when covered, and still returns a `transfer`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Transfers.ListRecipients() -> *whopsdk.ListRecipientsTransfersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the people and accounts you can send money to.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListRecipientsTransfersRequest{
    OriginID: "origin_id",
}
client.Transfers.ListRecipients(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**originID:** `string` — The account sending the money: a company account ID (`biz_`), or a user ID (`user_`) for that user's own personal balance.
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` — Search anyone on Whop by name or username, plus your own accounts by name or ID. Omit it to get the team around the balance, the people you follow, and your own accounts. The list is the same whether the balance belongs to a company or to you. Searching from a `biz_` origin additionally requires the member:basic:read scope. A credential scoped to a single company is the exception to the search itself: it only ever sees that company's own people. Complete email addresses return no matches.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Number of recipients per page. Search queries preserve the dashboard's 20-result maximum.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor to fetch the page after (from page_info.end_cursor).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Transfers.Retrieve(ID) -> *whopsdk.RetrieveTransfersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a single transfer.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveTransfersRequest{
    ID: "id",
}
client.Transfers.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The transfer ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users
<details><summary><code>client.Users.List() -> *whopsdk.ListUsersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Search for users by name or username, ranked by social proximity to the authenticated user. Returns the user's most recently followed users when no query is given.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListUsersRequest{}
client.Users.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**query:** `*string` — A search term to filter users by name or username.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of users to return (max 50).
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor; returns users after this position.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — The number of users to return from the end of the range.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — A cursor; returns users before this position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Me() -> *whopsdk.User</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the authenticated user — the self view of the user object. Same shape as `GET /users/{id}`, with the self-only fields populated: `email` (email-read scope), `staff` (Whop staff only, staff-read scope), `balance` and `earnings_usd` (balance-read scope), the opt-in `balance_history`, and every linked social account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.MeUsersRequest{}
client.Users.Me(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — When set, returns your account-specific profile overrides for this account.
    
</dd>
</dl>

<dl>
<dd>

**includeBalanceHistory:** `*bool` — Also compute your balance history (opt-in; runs a heavier query). Ignored for callers without balance-read scope.
    
</dd>
</dl>

<dl>
<dd>

**from:** `*string` — Balance-history window start, ISO 8601 date or datetime. Defaults to 30 days ago. Only used with `include_balance_history`.
    
</dd>
</dl>

<dl>
<dd>

**to:** `*string` — Balance-history window end, ISO 8601 date or datetime. Defaults to now. Only used with `include_balance_history`.
    
</dd>
</dl>

<dl>
<dd>

**interval:** `*whopsdk.MeUsersRequestInterval` — Balance-history point granularity. Defaults to `day`. Only used with `include_balance_history`.
    
</dd>
</dl>

<dl>
<dd>

**timeZone:** `*string` — IANA time zone the balance-history points are bucketed in. Defaults to `UTC`. Only used with `include_balance_history`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.UpdateMe(request) -> *whopsdk.User</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the authenticated user's global profile, or their profile override for an account when account_id is given. Not available to API keys.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateMeUsersRequest{}
client.Users.UpdateMe(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — When set, updates the authenticated user's profile override for this account instead of their global profile.
    
</dd>
</dl>

<dl>
<dd>

**banner:** `*whopsdk.UpdateMeUsersRequestBanner` 
    
</dd>
</dl>

<dl>
<dd>

**bio:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**profilePicture:** `*whopsdk.UpdateMeUsersRequestProfilePicture` 
    
</dd>
</dl>

<dl>
<dd>

**username:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Retrieve(ID) -> *whopsdk.User</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a user by `user_` tag or username, or the authenticated user with the reserved id `me`. Profiles include linked social accounts — reading your own profile returns every linked account, other profiles only what is public on Whop (the primary Discord and the X account). The self-only fields are populated only when the id is `me`: `email` (email-read scope), `staff` (Whop staff only, staff-read scope), `balance` and `earnings_usd` (balance-read scope), and the opt-in `balance_history`. They are always `null` when addressing a user by tag or username.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveUsersRequest{
    ID: "id",
}
client.Users.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — User ID (prefixed `user_`), username, or `me` for the authenticated user.
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `*string` — When set, returns the user's account-specific profile overrides for this account.
    
</dd>
</dl>

<dl>
<dd>

**includeBalanceHistory:** `*bool` — Also compute your balance history (opt-in; runs a heavier query). Only applies when the id is `me`; ignored for callers without balance-read scope.
    
</dd>
</dl>

<dl>
<dd>

**from:** `*string` — Balance-history window start, ISO 8601 date or datetime. Defaults to 30 days ago. Only used with `include_balance_history`.
    
</dd>
</dl>

<dl>
<dd>

**to:** `*string` — Balance-history window end, ISO 8601 date or datetime. Defaults to now. Only used with `include_balance_history`.
    
</dd>
</dl>

<dl>
<dd>

**interval:** `*whopsdk.RetrieveUsersRequestInterval` — Balance-history point granularity. Defaults to `day`. Only used with `include_balance_history`.
    
</dd>
</dl>

<dl>
<dd>

**timeZone:** `*string` — IANA time zone the balance-history points are bucketed in. Defaults to `UTC`. Only used with `include_balance_history`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Update(ID, request) -> *whopsdk.User</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a user, addressed by `user_` tag, username, or the reserved id `me` for the authenticated user. A user token updates their own global profile; an API key updates the user's account-specific profile override (account_id required).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateUsersRequest{
    ID: "id",
}
client.Users.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — User ID (prefixed `user_`), username, or `me` for the authenticated user.
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `*string` — The account whose profile override to update. Required for API key callers.
    
</dd>
</dl>

<dl>
<dd>

**banner:** `*whopsdk.UpdateUsersRequestBanner` 
    
</dd>
</dl>

<dl>
<dd>

**bio:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**profilePicture:** `*whopsdk.UpdateUsersRequestProfilePicture` 
    
</dd>
</dl>

<dl>
<dd>

**username:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.CheckAccess(ID, ResourceID) -> *whopsdk.CheckAccessUsersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Checks whether a user has access to an account, product, or experience the caller can reach.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CheckAccessUsersRequest{
    ID: "id",
    ResourceID: "resource_id",
}
client.Users.CheckAccess(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The user_ tag or username to check access for.
    
</dd>
</dl>

<dl>
<dd>

**resourceID:** `string` — An account (biz_), product (prod_), or experience (exp_) ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.RecommendActions(ID) -> *whopsdk.RecommendActionsUsersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the recommended actions computed for the user: personal suggestions (e.g. start a business or become an affiliate) pooled with the highest-impact actions across the accounts the user owns. Business actions are tagged with their `account_id`/`account_name`; personal actions leave those `null`. Self-only: `id` must be `me` or the authenticated user's own tag/username.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RecommendActionsUsersRequest{
    ID: "id",
}
client.Users.RecommendActions(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — `me`, or the authenticated user's own `user_` tag or username.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Verifications
<details><summary><code>client.Verifications.List() -> *whopsdk.ListVerificationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns verifications for an account, including their status and any required actions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListVerificationsRequest{
    AccountID: "account_id",
}
client.Verifications.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string` — Account or user ID whose verifications you want to list. Use a `biz_` account ID, or the caller's `user_` ID for personal verifications.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*whopsdk.ListVerificationsRequestOrder` — Field used to sort returned verifications.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*whopsdk.ListVerificationsRequestDirection` — Sort direction for returned verifications.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Verifications.Create(request) -> *whopsdk.CreateVerificationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Starts a hosted verification session for an account or user, or returns the active session when one already exists. Any fields you include in the request body are used to prefill the session. Send `documents` (with `document_type`) to instead verify the person from identity documents included in this request — no hosted session involved. Send `share_token` to reuse a verification another Sumsub account has already completed for this person, instead of verifying them again. If the account already has an `approved` verification the request is rejected; unlink it first to start a new one.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateVerificationsRequest{
    AccountID: "account_id",
    Body: &whopsdk.CreateVerificationsRequestBody{
        Individual: &whopsdk.CreateVerificationsRequestBodyIndividual{},
    },
}
client.Verifications.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string` — Account or user ID whose identity you want to verify. Use a `biz_` account ID for account verifications, or the caller's `user_` ID for personal verification.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*whopsdk.CreateVerificationsRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Verifications.Retrieve(ID) -> *whopsdk.RetrieveVerificationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns verifications for an account, including their status and any required actions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveVerificationsRequest{
    ID: "id",
}
client.Verifications.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Verification profile ID, prefixed `idpf_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Verifications.Update(ID, request) -> *whopsdk.UpdateVerificationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates editable profile details or submits answers for items returned in `requested_information`. Once a verification is `approved` its profile details are locked and can no longer be edited.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateVerificationsRequest{
    ID: "id",
    Body: &whopsdk.UpdateVerificationsRequestBody{
        UpdateVerificationsRequestBodyPersonalAddress: &whopsdk.UpdateVerificationsRequestBodyPersonalAddress{},
    },
}
client.Verifications.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Verification profile ID, prefixed `idpf_`.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*whopsdk.UpdateVerificationsRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Webhooks
<details><summary><code>client.Webhooks.List() -> *whopsdk.ListWebhooksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of webhook endpoints configured for an account, ordered by most recently created.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListWebhooksRequest{
    AccountID: "account_id",
}
client.Webhooks.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string` — The unique identifier of the account to list webhooks for.
    
</dd>
</dl>

<dl>
<dd>

**appID:** `*string` — Only return webhooks attached to this app. Omit to list the account's own webhooks.
    
</dd>
</dl>

<dl>
<dd>

**includeAppWebhooks:** `*bool` — Also return webhooks attached to the account's apps, not just the account's own. Cannot be combined with `app_id`.
    
</dd>
</dl>

<dl>
<dd>

**hasFailures:** `*bool` — Only return webhooks whose endpoint is currently failing — every delivery since the current failure streak began has been rejected. Clears as soon as a delivery succeeds.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of webhooks to return (default 20, max 100).
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor; returns webhooks after this position.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — The number of webhooks to return from the end of the range.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — A cursor; returns webhooks before this position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Create(request) -> *whopsdk.Webhook</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a webhook endpoint that receives event notifications via HTTP POST.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.CreateWebhooksRequest{
    URL: "https://example.com/hooks",
}
client.Webhooks.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiVersionDate:** `*string` — The dated API version (Api-Version-Date) to pin this webhook's payloads to. Omit to leave the webhook unpinned, tracking the current payload shape.
    
</dd>
</dl>

<dl>
<dd>

**childResourceEvents:** `*bool` — Whether to send events for child resources. For example, if the webhook is created for an account, enabling this sends events only from its connected accounts.
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` — Whether or not the webhook is enabled. Defaults to `true`.
    
</dd>
</dl>

<dl>
<dd>

**events:** `[]*whopsdk.CreateWebhooksRequestEventsItem` — The events to send the webhook for, in dot form (for example `payment.succeeded`).
    
</dd>
</dl>

<dl>
<dd>

**resourceID:** `*string` — The account or app to create the webhook for. Defaults to the current account.
    
</dd>
</dl>

<dl>
<dd>

**url:** `string` — The URL to send the webhook to.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Retrieve(ID) -> *whopsdk.Webhook</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of an existing webhook.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.RetrieveWebhooksRequest{
    ID: "id",
}
client.Webhooks.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Webhook ID, prefixed `hook_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Delete(ID) -> *whopsdk.DeleteWebhooksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Permanently deletes a webhook endpoint. Returns `true` on success, matching the legacy proxy response.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeleteWebhooksRequest{
    ID: "id",
}
client.Webhooks.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Webhook ID, prefixed `hook_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Update(ID, request) -> *whopsdk.Webhook</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a webhook endpoint's URL, subscribed events, pinned payload version, or enabled state.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.UpdateWebhooksRequest{
    ID: "id",
}
client.Webhooks.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Webhook ID, prefixed `hook_`.
    
</dd>
</dl>

<dl>
<dd>

**apiVersionDate:** `*string` — The dated API version (Api-Version-Date) to pin this webhook's payloads to. Only valid for `v1` webhooks. Omit to leave the current pin unchanged, or pass `null` to unpin and track the current payload shape.
    
</dd>
</dl>

<dl>
<dd>

**childResourceEvents:** `*bool` — Whether or not to send events for child resources.
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` — Whether or not the webhook is enabled.
    
</dd>
</dl>

<dl>
<dd>

**events:** `[]*whopsdk.UpdateWebhooksRequestEventsItem` — The events to send the webhook for, in dot form (for example `payment.succeeded`).
    
</dd>
</dl>

<dl>
<dd>

**url:** `*string` — The URL to send the webhook to.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.ListDeliveries(ID) -> *whopsdk.ListDeliveriesWebhooksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of delivery attempts for a webhook, ordered by most recent first. Includes the request payload, response body, response code, and timing for each attempt.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ListDeliveriesWebhooksRequest{
    ID: "id",
}
client.Webhooks.ListDeliveries(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Webhook ID, prefixed `hook_`.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of deliveries to return (default 50, max 100).
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor; returns deliveries after this position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.ReplayDelivery(ID, DeliveryID, request) -> *whopsdk.ReplayDeliveryWebhooksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Re-sends the exact payload of a past delivery to the webhook's current URL and returns the delivery result. By default the replay keeps the original `webhook-id`, so consumers that deduplicate on it can drop events they already processed. Pass `regenerate_id` to re-send under a freshly generated `webhook-id` instead, so a deduplicating consumer processes the replay as a new message. Only available for enabled webhooks on API version v1; deliveries are retained for 30 days.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ReplayDeliveryWebhooksRequest{
    ID: "id",
    DeliveryID: "delivery_id",
}
client.Webhooks.ReplayDelivery(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Webhook ID, prefixed `hook_`.
    
</dd>
</dl>

<dl>
<dd>

**deliveryID:** `string` — Delivery ID, prefixed `whdel_`, from the List Deliveries endpoint.
    
</dd>
</dl>

<dl>
<dd>

**regenerateID:** `*bool` — Re-send the delivery under a freshly generated `webhook-id` (in both the envelope and the signed headers) instead of the original one. Defaults to false. Use this when your endpoint deduplicates on `webhook-id` and you want it to process the replay as a new message.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Replay(ID, request) -> *whopsdk.ReplayWebhooksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Re-sends the webhook's past deliveries within a time window, optionally limited to specific events or to messages whose most recent delivery attempt failed. Fire and forget: nothing about the replay is stored, and each re-send appears as a new entry in the webhook's delivery log. Each matching message is re-sent once, by default with its original `webhook-id`, so consumers that deduplicate are unaffected; pass `regenerate_ids` to re-send under freshly generated ids instead. Only available for enabled webhooks on API version v1; deliveries are retained for 30 days.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.ReplayWebhooksRequest{
    ID: "id",
    SentAfter: "2026-01-01T12:00:00.000Z",
}
client.Webhooks.Replay(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Webhook ID, prefixed `hook_`.
    
</dd>
</dl>

<dl>
<dd>

**events:** `[]string` — Only replay these event types, in dot form (for example `payment.succeeded`). Omit to include every event.
    
</dd>
</dl>

<dl>
<dd>

**failedOnly:** `*bool` — Only replay messages whose most recent delivery attempt in the window failed. Defaults to false. Best-effort: a message whose attempts span processing batches can still be re-sent — replays keep the original `webhook-id` by default, so consumers that deduplicate are unaffected.
    
</dd>
</dl>

<dl>
<dd>

**regenerateIDs:** `*bool` — Re-send each replayed message under a freshly generated `webhook-id` (in both the envelope and the signed headers) instead of its original one. Defaults to false. Use this when your endpoint deduplicates on `webhook-id` and you want it to process the replays as new messages.
    
</dd>
</dl>

<dl>
<dd>

**sentAfter:** `string` — Start of the delivery window to replay, as an ISO 8601 timestamp. Clamped to the 30-day delivery retention.
    
</dd>
</dl>

<dl>
<dd>

**sentBefore:** `*string` — End of the delivery window to replay, as an ISO 8601 timestamp. Defaults to now.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Test(ID, request) -> *whopsdk.TestWebhooksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Sends a sample payload for the given event to the webhook's URL and returns the delivery result.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.TestWebhooksRequest{
    ID: "id",
    Event: "payment.succeeded",
}
client.Webhooks.Test(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Webhook ID, prefixed `hook_`.
    
</dd>
</dl>

<dl>
<dd>

**event:** `string` — The event to test the webhook for, in dot form (for example `payment.succeeded`).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.DeliveriesWebhook(WebhookID) -> *whopsdk.DeliveriesWebhookResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of delivery attempts for a webhook, ordered by most recent first. Includes the request payload, response body, response code, and timing for each attempt.

Required permissions:
 - `developer:manage_webhook`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &whopsdk.DeliveriesWebhookRequest{
    WebhookID: "webhook_id",
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
}
client.Webhooks.DeliveriesWebhook(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookID:** `string` — The unique identifier of the webhook to list deliveries for.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounts Preferences
<details><summary><code>client.Accounts.Preferences.Retrieve(AccountID) -> *accounts.RetrievePreferencesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the account's preferences: a singleton settings document keyed by preference name.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounts.RetrievePreferencesRequest{
    AccountID: "account_id",
}
client.Accounts.Preferences.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string` — Account ID, prefixed `biz_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.Preferences.Update(AccountID, request) -> *accounts.UpdatePreferencesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the account's preferences. Each top-level key present in the body is replaced as a whole; omitted keys are left untouched. `ads_triple_whale_integration` takes the Data-In API key to connect with, or `null` to disconnect. `ads_payment_methods` always requires a `primary` entry. `backup` is optional and any pairing is allowed — two cards, `card`+`platform_balance`, or a single method — so a card-only advertiser can fund ads without a platform balance. The `primary` and `backup` must be different sources. A `platform_balance` entry may omit `id` to use the account's default Whop balance. Configuring a `card` requires a user token; account API keys can set up platform-balance billing only.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounts.UpdatePreferencesRequest{
    AccountID: "account_id",
}
client.Accounts.Preferences.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string` — Account ID, prefixed `biz_`.
    
</dd>
</dl>

<dl>
<dd>

**adsPaymentMethods:** `*accounts.UpdatePreferencesRequestAdsPaymentMethods` — How the account pays for Whop Ads spend. `primary` is charged first; `backup` covers the charge when the primary fails.
    
</dd>
</dl>

<dl>
<dd>

**adsReportingCurrency:** `*string` — Lowercase ISO currency code, such as `usd` or `eur`, used to display ad spend and stats. Defaults to `usd`.
    
</dd>
</dl>

<dl>
<dd>

**adsSchedulingTimezone:** `*string` — IANA timezone (e.g. `America/New_York`) used to interpret campaign start/end times and to bucket reports. Cannot be cleared once set — pass a new value to change it.
    
</dd>
</dl>

<dl>
<dd>

**adsTripleWhaleIntegration:** `*accounts.UpdatePreferencesRequestAdsTripleWhaleIntegration` — Connects or disconnects the Triple Whale integration. Requires a connected Shopify store, since Triple Whale keys spend records by Shopify shop.
    
</dd>
</dl>

<dl>
<dd>

**cardsAutoTopUp:** `*bool` — Whether incoming funds are automatically moved to the account's cards balance. Requires a cards balance on the account.
    
</dd>
</dl>

<dl>
<dd>

**disputeFighterEnabled:** `*bool` — Whether Whop assembles and files the evidence response when this account's payments are disputed. Off by default; enabling it also opts the account into the success fee charged only on disputes it wins.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounts Reserves
<details><summary><code>client.Accounts.Reserves.List(AccountID) -> *accounts.ListReservesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists what the account's held balance is made of, one entry per currency: the total held, why each part is held, and the days it unlocks.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounts.ListReservesRequest{
    AccountID: "account_id",
}
client.Accounts.Reserves.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string` — Account ID, prefixed `biz_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Affiliates Overrides
<details><summary><code>client.Affiliates.Overrides.List(ID) -> *affiliates.ListOverridesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of overrides for an affiliate.

Required permissions:
 - `affiliate:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affiliates.ListOverridesRequest{
    ID: "aff_xxxxxxxxxxxxxx",
    First: whopsdk.Int(
        42,
    ),
    Last: whopsdk.Int(
        42,
    ),
}
client.Affiliates.Overrides.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The affiliate ID.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Returns the elements in the list that come after the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Returns the elements in the list that come before the specified cursor.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Returns the first _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Returns the last _n_ elements from the list.
    
</dd>
</dl>

<dl>
<dd>

**overrideType:** `*whopsdk.AffiliateOverrideRoles` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Affiliates.Overrides.Create(ID, request) -> *affiliates.CreateOverridesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a commission override for an affiliate.

Required permissions:
 - `affiliate:create`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affiliates.CreateOverridesRequest{
    ID: "aff_xxxxxxxxxxxxxx",
    Body: &affiliates.CreateOverridesRequestBody{
        Standard: &affiliates.CreateOverridesRequestBodyStandard{
            CommissionValue: 6.9,
            ID: "id",
            PlanID: "plan_xxxxxxxxxxxxx",
        },
    },
}
client.Affiliates.Overrides.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The affiliate ID.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*affiliates.CreateOverridesRequestBody` — Parameters for CreateAffiliateOverride
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Affiliates.Overrides.Retrieve(ID, OverrideID) -> *affiliates.RetrieveOverridesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the details of a specific affiliate override.

Required permissions:
 - `affiliate:basic:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affiliates.RetrieveOverridesRequest{
    ID: "aff_xxxxxxxxxxxxxx",
    OverrideID: "override_id",
}
client.Affiliates.Overrides.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The affiliate ID.
    
</dd>
</dl>

<dl>
<dd>

**overrideID:** `string` — The override ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Affiliates.Overrides.Delete(ID, OverrideID) -> bool</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes an affiliate override.

Required permissions:
 - `affiliate:update`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affiliates.DeleteOverridesRequest{
    ID: "aff_xxxxxxxxxxxxxx",
    OverrideID: "override_id",
}
client.Affiliates.Overrides.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The affiliate ID.
    
</dd>
</dl>

<dl>
<dd>

**overrideID:** `string` — The override ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Affiliates.Overrides.Update(ID, OverrideID, request) -> *affiliates.UpdateOverridesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an existing affiliate override.

Required permissions:
 - `affiliate:update`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affiliates.UpdateOverridesRequest{
    ID: "aff_xxxxxxxxxxxxxx",
    OverrideID: "override_id",
}
client.Affiliates.Overrides.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The affiliate ID.
    
</dd>
</dl>

<dl>
<dd>

**overrideID:** `string` — The override ID.
    
</dd>
</dl>

<dl>
<dd>

**appliesToPayments:** `*whopsdk.AffiliateAppliesToPayments` — Whether commission applies to first payment or all payments (standard only).
    
</dd>
</dl>

<dl>
<dd>

**commissionType:** `*whopsdk.AffiliatePayoutTypes` — The commission type (percentage or flat_fee).
    
</dd>
</dl>

<dl>
<dd>

**commissionValue:** `*float64` — The commission value (percentage 1-100 or flat fee in dollars).
    
</dd>
</dl>

<dl>
<dd>

**revenueBasis:** `*whopsdk.AffiliateRevenueBases` — The revenue calculation basis (rev-share only).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Bounties Submissions
<details><summary><code>client.Bounties.Submissions.List(BountyID) -> *bounties.ListSubmissionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists a bounty's publicly visible work — submitted, approved, and denied submissions in the reduced public shape. Authentication is optional; a bounty that is not publicly visible returns `404`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bounties.ListSubmissionsRequest{
    BountyID: "bounty_id",
}
client.Bounties.Submissions.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**bountyID:** `string` — The bounty whose public submissions to list (`bnty_` tag).
    
</dd>
</dl>

<dl>
<dd>

**status:** `*bounties.ListSubmissionsRequestStatus` — Filter by lifecycle state.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Only submissions created after this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Only submissions created before this ISO 8601 timestamp.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*bounties.ListSubmissionsRequestOrder` — Sort field.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*bounties.ListSubmissionsRequestDirection` — Sort direction.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Number of submissions to return from the start of the window.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor to paginate forwards from.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Number of submissions to return from the end of the window.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Cursor to paginate backwards from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bounties.Submissions.Retrieve(BountyID, ID) -> *whopsdk.PublicBountySubmission</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves one of a bounty's publicly visible submissions in the reduced public shape — the read behind a shared proof link, whose submission is usually outside the bounty page's capped preview. Authentication is optional; a bounty that is not publicly visible, and a submission that is not publicly visible work on it, both return `404`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bounties.RetrieveSubmissionsRequest{
    BountyID: "bounty_id",
    ID: "id",
}
client.Bounties.Submissions.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**bountyID:** `string` — The bounty the submission belongs to (`bnty_` tag).
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` — The submission to retrieve (`btys_` tag).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Members Logs
<details><summary><code>client.Members.Logs.List(ID) -> *members.ListLogsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists activity for a member and all of their non-drafted memberships, most recent first.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &members.ListLogsRequest{
    ID: "id",
}
client.Members.Logs.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Member ID (`mber_` tag).
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Number of log entries to return from the start of the window.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor to paginate forwards from.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Number of log entries to return from the end of the window.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Cursor to paginate backwards from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Notifications Topics
<details><summary><code>client.Notifications.Topics.List() -> *notifications.ListTopicsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the platform's visible notification topics — the categories users can set notification preferences on. App-created topics are internal and not returned.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &notifications.ListTopicsRequest{}
client.Notifications.Topics.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**topicType:** `*notifications.ListTopicsRequestTopicType` — Only return topics of this scope: `user` (member notifications) or `account_team` (team notifications).
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of topics to return (default 20, max 100).
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor; returns topics after this position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Partners Businesses
<details><summary><code>client.Partners.Businesses.List() -> *partners.ListBusinessesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the businesses the authenticated user referred onto Whop, most recent first.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &partners.ListBusinessesRequest{}
client.Partners.Businesses.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**status:** `*partners.ListBusinessesRequestStatus` — Filter by referral status.
    
</dd>
</dl>

<dl>
<dd>

**hasEarnings:** `*bool` — When true, only businesses with pending or completed earnings paid to the caller.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Number of partner businesses to return from the start of the window.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor to fetch the page after (from page_info.end_cursor).
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Number of partner businesses to return from the end of the window.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Cursor to fetch the page before (from page_info.start_cursor).
    
</dd>
</dl>

<dl>
<dd>

**order:** `*partners.ListBusinessesRequestOrder` — The field to sort partner businesses by.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*partners.ListBusinessesRequestDirection` — Sort direction.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Only return partner businesses created before this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Only return partner businesses created after this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**referredUserID:** `*string` — Filter to referrals attributed to this user. For first-tier referrals, this is the referred account owner; for second-tier referrals, this is the partner you recruited.
    
</dd>
</dl>

<dl>
<dd>

**referredUsername:** `*string` — Filter by the referred user's exact username. Ignored when `referred_user_id` is present.
    
</dd>
</dl>

<dl>
<dd>

**tier:** `*partners.ListBusinessesRequestTier` — Filter to referrals from a single tier: first, second, or blueprint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Partners.Businesses.Retrieve(ID) -> *partners.RetrieveBusinessesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a single referred business and its referral terms.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &partners.RetrieveBusinessesRequest{
    ID: "id",
}
client.Partners.Businesses.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The partner business ID (a coma_ identifier).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Partners Businesses Earnings
<details><summary><code>client.Partners.Businesses.Earnings.List(ID) -> *businesses.ListEarningsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the earnings Whop pays out for one referred business's activity, most recent first.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &businesses.ListEarningsRequest{
    ID: "id",
}
client.Partners.Businesses.Earnings.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The partner business ID (a coma_ identifier).
    
</dd>
</dl>

<dl>
<dd>

**status:** `*businesses.ListEarningsRequestStatus` — Filter by earning status.
    
</dd>
</dl>

<dl>
<dd>

**incomeSource:** `*businesses.ListEarningsRequestIncomeSourceItem` — Filter to earnings from these income sources. Repeat the parameter for each one (income_source=sales&income_source=ad_spend).
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**order:** `*businesses.ListEarningsRequestOrder` — The field to sort earnings by.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*businesses.ListEarningsRequestDirection` — Sort direction.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — Only return earnings created before this timestamp.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` — Only return earnings created after this timestamp.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Payouts Methods
<details><summary><code>client.Payouts.Methods.List() -> *payouts.ListMethodsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the bank accounts, wallets, and crypto addresses an account or user can pay out to, newest first.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payouts.ListMethodsRequest{}
client.Payouts.Methods.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The owning account ID (a biz_ identifier). Provide this or user_id.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — The owning user ID (a user_ identifier). Provide this or account_id.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*payouts.ListMethodsRequestStatus` — Optional status filter. `created` means saved but unused, `active` means a payout through it succeeded, `broken` means the last payout failed and the method needs fixing.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*float64` — Optional payout amount in whole currency units, for example `250.00`. When provided, each method includes a quote with the estimated fee, amount received, and delivery date for that amount.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*string` — Currency code of the amount, for example `usd`. Only meaningful with amount or include_limits.
    
</dd>
</dl>

<dl>
<dd>

**includeLimits:** `*bool` — When true, the response also carries limits — the live per-speed payout caps the account's payout requests are validated against, in the requested currency. Requires the payout:withdrawal:read scope.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Number of payout methods to return from the start of the window.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor to fetch the page after (from page_info.end_cursor).
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Number of payout methods to return from the end of the window.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Cursor to fetch the page before (from page_info.start_cursor).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payouts.Methods.Create(request) -> *payouts.CreateMethodsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Saves a new place an account or user can pay out to. Sensitive details are vaulted in transit and never stored raw.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payouts.CreateMethodsRequest{
    SupportedPayoutMethodID: "podst_xxxxxxxxxxxxxx",
}
client.Payouts.Methods.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The account to add the payout method for, prefixed `biz_`. Provide this or `user_id`.
    
</dd>
</dl>

<dl>
<dd>

**destinationCurrency:** `*string` — Currency the supported payout method delivers payouts in.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `map[string]string` — The supported payout method's required field values, keyed by field id — list them with `GET /payouts/supported_methods?supported_payout_method_id=...`. Field ids are stable `fld_` identifiers you may hardcode; they never change for a given field. A Basis Theory token id may be passed in place of a raw value. For a U.S. bank routing-number field, a raw nine-digit value must also pass the ABA checksum. A validation failure returns the method's full required_fields schema alongside the error. Required whenever the account details are supplied directly.
    
</dd>
</dl>

<dl>
<dd>

**isDefault:** `*bool` — Whether to make this the account's default payout method.
    
</dd>
</dl>

<dl>
<dd>

**nickname:** `*string` — A label for the payout method, unique per destination.
    
</dd>
</dl>

<dl>
<dd>

**supportedPayoutMethodID:** `string` — The supported payout method to save (a podst_ identifier from a previous listing).
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — The user to add the payout method for, prefixed `user_`. Provide this or `account_id`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payouts.Methods.Delete(ID) -> *payouts.DeleteMethodsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a saved payout method so it can no longer receive payouts.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payouts.DeleteMethodsRequest{
    ID: "id",
}
client.Payouts.Methods.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Payout method ID, prefixed `potk_`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Payouts.Methods.Update(ID, request) -> *payouts.UpdateMethodsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Changes the label used to identify a saved payout method.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payouts.UpdateMethodsRequest{
    ID: "id",
    Nickname: "Primary checking",
}
client.Payouts.Methods.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Payout method ID, prefixed `potk_`.
    
</dd>
</dl>

<dl>
<dd>

**nickname:** `string` — New label for the payout method, with at least one non-whitespace character and a maximum of 100 characters.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Payouts SupportedMethods
<details><summary><code>client.Payouts.SupportedMethods.List() -> *payouts.ListSupportedMethodsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the payout methods an account or user is eligible to add.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payouts.ListSupportedMethodsRequest{}
client.Payouts.SupportedMethods.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — The owning account ID (a biz_ identifier). Provide this or user_id.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — The owning user ID (a user_ identifier). Provide this or account_id.
    
</dd>
</dl>

<dl>
<dd>

**country:** `*string` — ISO 3166-1 alpha-2 country code for the bank account or wallet, such as `US`. Defaults to the country of supported_payout_method_id when one is given, otherwise the payout account's country.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*float64` — Optional payout amount in whole currency units, for example `250.00`. When provided, each destination includes per-currency fee and delivery quotes.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*string` — Currency code of the amount, for example `usd`. Only meaningful with amount.
    
</dd>
</dl>

<dl>
<dd>

**supportedPayoutMethodID:** `*string` — Narrows the list to one supported payout method (a podst_ identifier) and includes the required_fields needed to save it as a payout method.
    
</dd>
</dl>

<dl>
<dd>

**destinationCurrency:** `*string` — Currency the supported payout method would deliver payouts in. Only meaningful with supported_payout_method_id; required fields vary by destination currency.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — Number of supported payout methods to return from the start of the window.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — Cursor to fetch the page after (from page_info.end_cursor).
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — Number of supported payout methods to return from the end of the window.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — Cursor to fetch the page before (from page_info.start_cursor).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users OauthGrants
<details><summary><code>client.Users.OauthGrants.List() -> *users.ListOauthGrantsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the authenticated user's own OAuth grants — one per app they have authorized, per account they authorized it for. The list is always the caller's own; there is no parameter for reading another user's grants. Requires a user session: an API key or an OAuth token is refused, so an app can never enumerate the other apps a user has authorized.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &users.ListOauthGrantsRequest{}
client.Users.OauthGrants.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**appID:** `*string` — Only return grants for this app, prefixed `app_`. An app the user has never authorized returns an empty list.
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of grants to return (default 20, max 100).
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor; returns grants after this position.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — The number of grants to return from the end of the range.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — A cursor; returns grants before this position.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*users.ListOauthGrantsRequestOrder` — The field to sort grants by.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*users.ListOauthGrantsRequestDirection` — Sort direction.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.OauthGrants.Create(request) -> *whopsdk.OauthGrant</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Completes the OAuth authorization step for the authenticated user: records their consent for the scopes an app asked for and mints the authorization code to hand back to it. Returns the grant, plus a `redirect_url` carrying that code — the one and only time it is returned. Exchange the code at `POST /oauth/token` with the verifier for `code_challenge`. Requires a user session, because consent has to come from the account holder: an API key or an OAuth token is refused, so an app can never authorize itself. Send an `Idempotency-Key` to make a retry safe — a replay returns the original `redirect_url` and its code rather than issuing a second one.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &users.CreateOauthGrantsRequest{
    ClientID: "app_xxxxxxxxxxxxxx",
    CodeChallenge: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    CodeChallengeMethod: users.CreateOauthGrantsRequestCodeChallengeMethodS256,
    RedirectURI: "https://Booking.Shinetime.example:8443/oauth/Callback/",
    RequestedScopes: []string{
        "profile",
    },
}
client.Users.OauthGrants.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `*string` — Authorize the app for one of the user's accounts rather than for the user alone, prefixed `biz_`. The user must have access to it.
    
</dd>
</dl>

<dl>
<dd>

**clientID:** `string` — The app being authorized, prefixed `app_`.
    
</dd>
</dl>

<dl>
<dd>

**codeChallenge:** `string` — The PKCE code challenge: the base64url-encoded SHA-256 of your code verifier, without padding.
    
</dd>
</dl>

<dl>
<dd>

**codeChallengeMethod:** `*users.CreateOauthGrantsRequestCodeChallengeMethod` — How `code_challenge` was derived. Only `S256` is accepted.
    
</dd>
</dl>

<dl>
<dd>

**consentShown:** `*bool` — Whether the consent UI listed these scopes for the user. Sending `false` succeeds only when the user has already granted every scope requested.
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*string` — OIDC nonce, echoed into the resulting ID token. Required when `requested_scopes` includes `openid`.
    
</dd>
</dl>

<dl>
<dd>

**redirectURI:** `string` — Where to send the user once they have consented. Must match one of the app's registered redirect URIs exactly — it is compared as a string, not normalized.
    
</dd>
</dl>

<dl>
<dd>

**requestedScopes:** `[]string` — The permissions the app is asking for, for example `member:basic:read`. `GET /api_keys/permissions` names and describes each one. Granting adds to whatever the user already granted this app rather than replacing it.
    
</dd>
</dl>

<dl>
<dd>

**responseType:** `*users.CreateOauthGrantsRequestResponseType` — The OAuth response type. Only `code` is accepted; defaults to `code`.
    
</dd>
</dl>

<dl>
<dd>

**state:** `*string` — Opaque value appended to `redirect_url` unchanged, for the client to correlate the response with its request.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users Passkeys
<details><summary><code>client.Users.Passkeys.List() -> *users.ListPasskeysResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the authenticated user's own passkeys, newest first. The list is always the caller's own; there is no parameter for reading another user's passkeys. Requires a user session: an API key or an OAuth token is refused, because a passkey confirms the account holder before a sensitive action and no app may enumerate one.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &users.ListPasskeysRequest{}
client.Users.Passkeys.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**first:** `*int` — The number of passkeys to return (default 20, max 100).
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor; returns passkeys after this position.
    
</dd>
</dl>

<dl>
<dd>

**last:** `*int` — The number of passkeys to return from the end of the range.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` — A cursor; returns passkeys before this position.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*users.ListPasskeysRequestOrder` — The field to sort passkeys by.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*users.ListPasskeysRequestDirection` — Sort direction.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Passkeys.Create(request) -> *whopsdk.Passkey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Registers a passkey for the authenticated user from the attestation a browser produced for a `registration` challenge. Mint that challenge first with `POST /users/me/passkeys/challenge`; it is single-use and expires 5 minutes after it is issued. Requires a user session.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &users.CreatePasskeysRequest{
    AttestationObject: "YXR0ZXN0YXRpb24",
    ClientDataJSON: "Y2xpZW50LWRhdGE",
    CredentialID: "bmV3LWNyZWRlbnRpYWw",
    Nickname: "Work laptop",
}
client.Users.Passkeys.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**attestationObject:** `string` — The `attestationObject` from the WebAuthn attestation response, base64url-encoded.
    
</dd>
</dl>

<dl>
<dd>

**clientDataJSON:** `string` — The `clientDataJSON` from the WebAuthn attestation response, base64url-encoded.
    
</dd>
</dl>

<dl>
<dd>

**credentialID:** `string` — The WebAuthn credential ID the authenticator returned, base64url-encoded.
    
</dd>
</dl>

<dl>
<dd>

**nickname:** `string` — A name for this passkey, usually the device it lives on. 255 characters or fewer.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Passkeys.Challenge(request) -> *users.ChallengePasskeysResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Mints the challenge a browser needs to run a WebAuthn ceremony against the authenticated user's own passkeys. A `registration` challenge enrolls a new passkey; a `deletion` challenge is bound to the one passkey named by `passkey_id` and proves the user still holds it. Challenges are single-use and expire 5 minutes after they are issued, so send a fresh `Idempotency-Key` per ceremony — a replayed key returns the original challenge, which may already have expired. Requires a user session.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &users.ChallengePasskeysRequest{
    ChallengeType: users.ChallengePasskeysRequestChallengeTypeRegistration,
}
client.Users.Passkeys.Challenge(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**challengeType:** `*users.ChallengePasskeysRequestChallengeType` — The ceremony this challenge is for.
    
</dd>
</dl>

<dl>
<dd>

**passkeyID:** `*string` — The passkey the ceremony targets, prefixed `wcred_`. Required when `challenge_type` is `deletion`, ignored otherwise.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Passkeys.Delete(ID, request) -> *users.DeletePasskeysResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes one of the authenticated user's own passkeys. The request body carries a WebAuthn assertion from the passkey being deleted, so possession of the credential is proven before it is removed: mint a `deletion` challenge for it first, run the ceremony with that passkey, and send the result here. Deleting the user's last passkey is allowed — their other step-up factors remain. Requires a user session.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &users.DeletePasskeysRequest{
    ID: "id",
    AuthenticatorData: "YXV0aGVudGljYXRvci1kYXRh",
    ClientDataJSON: "Y2xpZW50LWRhdGE",
    Signature: "c2lnbmF0dXJl",
}
client.Users.Passkeys.Delete(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Passkey ID, prefixed `wcred_`.
    
</dd>
</dl>

<dl>
<dd>

**authenticatorData:** `string` — The `authenticatorData` from the WebAuthn assertion, base64url-encoded.
    
</dd>
</dl>

<dl>
<dd>

**clientDataJSON:** `string` — The `clientDataJSON` from the WebAuthn assertion, base64url-encoded.
    
</dd>
</dl>

<dl>
<dd>

**signature:** `string` — The `signature` from the WebAuthn assertion, base64url-encoded.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users Preferences
<details><summary><code>client.Users.Preferences.Retrieve() -> *whopsdk.UserPreferences</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the authenticated user's settings document. Addressed only as `me` — the document always belongs to the session user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Preferences.Retrieve(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Preferences.Update(request) -> *whopsdk.UserPreferences</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the authenticated user's settings document. Replaces the top-level keys it is given and leaves the rest untouched.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &users.UpdatePreferencesRequest{}
client.Users.Preferences.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**bountyWorkerOnboardingDismissed:** `*bool` — Whether the user has dismissed the first-time bounty worker onboarding. Set to `false` to show it again.
    
</dd>
</dl>

<dl>
<dd>

**investigationEnabled:** `*bool` — Whether investigation mode is enabled for the user. Only meaningful for staff users with investigation access.
    
</dd>
</dl>

<dl>
<dd>

**termsAccepted:** `*bool` — Records the user's acceptance of Whop's terms and policies. Only `true` is accepted — the server stamps `terms_accepted_at` and acceptance cannot be withdrawn here.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users Preferences Notifications
<details><summary><code>client.Users.Preferences.Notifications.Set(request) -> *preferences.SetNotificationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Sets the authenticated user's notification preferences. Each preference is addressed by `scope`, not by id, so a scope read back from either list endpoint can be sent straight here.

A scope naming an experience with no topic sets that experience's level, and accepts all three levels. Any other scope sets a topic override, which is binary — `all` or `nothing` — and requires a `channel`.

`level: null` clears the preference. Preferences are stored as overrides, so clearing one means the scope inherits its default again rather than being switched off.

The batch is applied in one transaction: if any entry is rejected, none are written. Experience levels are applied before topic overrides, because setting a level replaces every topic preference for that experience — so an override sent alongside a level wins. The response reports what each scope now resolves to, in the order the entries were sent.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &preferences.SetNotificationsRequest{
    Preferences: []*preferences.SetNotificationsRequestPreferencesItem{
        &preferences.SetNotificationsRequestPreferencesItem{
            Scope: &preferences.SetNotificationsRequestPreferencesItemScope{},
        },
    },
}
client.Users.Preferences.Notifications.Set(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**preferences:** `[]*preferences.SetNotificationsRequestPreferencesItem` — The preferences to set, at most 100 per request.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users Preferences Notifications Experiences
<details><summary><code>client.Users.Preferences.Notifications.Experiences.List() -> *notifications.ListExperiencesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the authenticated user's per-experience notification levels. Experiences the user never set a level for are omitted — their effective level is `all`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &notifications.ListExperiencesRequest{}
client.Users.Preferences.Notifications.Experiences.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**first:** `*int` — The number of preferences to return.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor; returns preferences after this position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users Preferences Notifications Topics
<details><summary><code>client.Users.Preferences.Notifications.Topics.List() -> *notifications.ListTopicsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the authenticated user's topic-scoped notification preferences, plus user-agnostic platform defaults. Each filter matches preferences scoped to its value or not narrowed on that dimension. Per-experience levels are listed separately, by `GET /users/me/preferences/notifications/experiences`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &notifications.ListTopicsRequest{}
client.Users.Preferences.Notifications.Topics.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**channel:** `*notifications.ListTopicsRequestChannel` — Only return preferences for this delivery channel (or not narrowed to a channel).
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `*string` — Only return preferences scoped to this account's member notifications (`biz_` tag).
    
</dd>
</dl>

<dl>
<dd>

**teamAccountID:** `*string` — Only return preferences scoped to this account's team notifications (`biz_` tag).
    
</dd>
</dl>

<dl>
<dd>

**experienceID:** `*string` — Only return preferences scoped to this experience (`exp_` tag).
    
</dd>
</dl>

<dl>
<dd>

**topicID:** `*string` — Only return preferences scoped to this notification topic (`topic_` tag).
    
</dd>
</dl>

<dl>
<dd>

**first:** `*int` — The number of preferences to return.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` — A cursor; returns preferences after this position.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

