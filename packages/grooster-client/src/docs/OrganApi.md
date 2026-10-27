# OrganApi

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**getMemberSettings**](#getmembersettings) | **GET** /organs/{organId}/members/{userId} | Get settings for a user within an organ|
|[**getMembersSettings**](#getmemberssettings) | **GET** /organs/{organId}/members | Get settings for all members within an organ|
|[**updateMemberRole**](#updatememberrole) | **PATCH** /organs/{organId}/members/{userId}/role | Update the role of a user within a specific organ|
|[**updateMemberSettings**](#updatemembersettings) | **PATCH** /organs/{organId}/members/{userId} | Update settings for a user within an organ|

# **getMemberSettings**
> UserOrgan getMemberSettings()

Get organ-specific settings like nickname/username

### Example

```typescript
import {
    OrganApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new OrganApi(configuration);

let organId: number; //Organ ID (default to undefined)
let userId: number; //User ID (default to undefined)

const { status, data } = await apiInstance.getMemberSettings(
    organId,
    userId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **organId** | [**number**] | Organ ID | defaults to undefined|
| **userId** | [**number**] | User ID | defaults to undefined|


### Return type

**UserOrgan**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**400** | Bad Request |  -  |
|**404** | Not Found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getMembersSettings**
> Array<UserOrgan> getMembersSettings()

Get organ-specific settings like nickname/username for all its members

### Example

```typescript
import {
    OrganApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new OrganApi(configuration);

let organId: number; //Organ ID (default to undefined)

const { status, data } = await apiInstance.getMembersSettings(
    organId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **organId** | [**number**] | Organ ID | defaults to undefined|


### Return type

**Array<UserOrgan>**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**400** | Bad Request |  -  |
|**404** | Not Found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **updateMemberRole**
> UserOrgan updateMemberRole(updateParams)

Update a users role within a specific organ

### Example

```typescript
import {
    OrganApi,
    Configuration,
    UpdateMemberRoleParams
} from './api';

const configuration = new Configuration();
const apiInstance = new OrganApi(configuration);

let organId: number; //Organ ID (default to undefined)
let userId: number; //User ID (default to undefined)
let updateParams: UpdateMemberRoleParams; //Settings input

const { status, data } = await apiInstance.updateMemberRole(
    organId,
    userId,
    updateParams
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateParams** | **UpdateMemberRoleParams**| Settings input | |
| **organId** | [**number**] | Organ ID | defaults to undefined|
| **userId** | [**number**] | User ID | defaults to undefined|


### Return type

**UserOrgan**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**400** | Bad Request |  -  |
|**404** | Not Found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **updateMemberSettings**
> UserOrgan updateMemberSettings(updateParams)

Update organ-specific settings like nickname/username

### Example

```typescript
import {
    OrganApi,
    Configuration,
    UpdateMemberSettingsParams
} from './api';

const configuration = new Configuration();
const apiInstance = new OrganApi(configuration);

let organId: number; //Organ ID (default to undefined)
let userId: number; //User ID (default to undefined)
let updateParams: UpdateMemberSettingsParams; //Settings input

const { status, data } = await apiInstance.updateMemberSettings(
    organId,
    userId,
    updateParams
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateParams** | **UpdateMemberSettingsParams**| Settings input | |
| **organId** | [**number**] | Organ ID | defaults to undefined|
| **userId** | [**number**] | User ID | defaults to undefined|


### Return type

**UserOrgan**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

