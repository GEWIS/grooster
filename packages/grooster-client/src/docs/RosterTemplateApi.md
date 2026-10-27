# RosterTemplateApi

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**createRosterTemplate**](#createrostertemplate) | **POST** /organs/{organId}/templates | Creates a template of a roster by defining the name of the shifts|
|[**createRosterTemplateShiftPreference**](#createrostertemplateshiftpreference) | **POST** /templates/{templateId}/preferences | Creates a roster template shift preference|
|[**deleteRosterTemplate**](#deleterostertemplate) | **DELETE** /templates/{templateId} | Deletes a roster template by ID|
|[**getRosterTemplate**](#getrostertemplate) | **GET** /templates/{templateId} | Get a roster template by ID|
|[**getRosterTemplateShiftPreferences**](#getrostertemplateshiftpreferences) | **GET** /templates/{templateId}/preferences | Gets shift preferences of a template filtered by user|
|[**getRosterTemplates**](#getrostertemplates) | **GET** /organs/{organId}/templates | Get all roster templates of an organ|
|[**updateRosterTemplate**](#updaterostertemplate) | **PUT** /templates/{templateId} | Updates a roster template by ID|
|[**updateRosterTemplateShift**](#updaterostertemplateshift) | **PATCH** /templates/{templateId}/shifts/{shiftId} | Updates a roster template shift by ID|
|[**updateRosterTemplateShiftPreference**](#updaterostertemplateshiftpreference) | **PATCH** /templates/{templateId}/preferences/{preferenceId} | Updates a roster template shift preference by ID|

# **createRosterTemplate**
> Array<RosterTemplate> createRosterTemplate()


### Example

```typescript
import {
    RosterTemplateApi,
    Configuration,
    TemplateCreateRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new RosterTemplateApi(configuration);

let organId: number; //Organ ID (default to undefined)
let params: TemplateCreateRequest; //Template Params (optional)

const { status, data } = await apiInstance.createRosterTemplate(
    organId,
    params
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **params** | **TemplateCreateRequest**| Template Params | |
| **organId** | [**number**] | Organ ID | defaults to undefined|


### Return type

**Array<RosterTemplate>**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Created Template |  -  |
|**400** | Invalid request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **createRosterTemplateShiftPreference**
> RosterTemplateShiftPreference createRosterTemplateShiftPreference(params)


### Example

```typescript
import {
    RosterTemplateApi,
    Configuration,
    TemplateShiftPreferenceCreateRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new RosterTemplateApi(configuration);

let templateId: number; //Template ID (default to undefined)
let params: TemplateShiftPreferenceCreateRequest; //Creation params

const { status, data } = await apiInstance.createRosterTemplateShiftPreference(
    templateId,
    params
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **params** | **TemplateShiftPreferenceCreateRequest**| Creation params | |
| **templateId** | [**number**] | Template ID | defaults to undefined|


### Return type

**RosterTemplateShiftPreference**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**201** | Created |  -  |
|**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteRosterTemplate**
> string deleteRosterTemplate()


### Example

```typescript
import {
    RosterTemplateApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new RosterTemplateApi(configuration);

let templateId: number; //Template ID (default to undefined)

const { status, data } = await apiInstance.deleteRosterTemplate(
    templateId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **templateId** | [**number**] | Template ID | defaults to undefined|


### Return type

**string**

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

# **getRosterTemplate**
> RosterTemplate getRosterTemplate()


### Example

```typescript
import {
    RosterTemplateApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new RosterTemplateApi(configuration);

let templateId: number; //Template ID (default to undefined)

const { status, data } = await apiInstance.getRosterTemplate(
    templateId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **templateId** | [**number**] | Template ID | defaults to undefined|


### Return type

**RosterTemplate**

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

# **getRosterTemplateShiftPreferences**
> Array<RosterTemplateShiftPreference> getRosterTemplateShiftPreferences()


### Example

```typescript
import {
    RosterTemplateApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new RosterTemplateApi(configuration);

let templateId: number; //Template ID (default to undefined)
let userId: number; //User ID (default to undefined)

const { status, data } = await apiInstance.getRosterTemplateShiftPreferences(
    templateId,
    userId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **templateId** | [**number**] | Template ID | defaults to undefined|
| **userId** | [**number**] | User ID | defaults to undefined|


### Return type

**Array<RosterTemplateShiftPreference>**

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getRosterTemplates**
> Array<RosterTemplate> getRosterTemplates()


### Example

```typescript
import {
    RosterTemplateApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new RosterTemplateApi(configuration);

let organId: number; //Organ ID (default to undefined)

const { status, data } = await apiInstance.getRosterTemplates(
    organId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **organId** | [**number**] | Organ ID | defaults to undefined|


### Return type

**Array<RosterTemplate>**

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **updateRosterTemplate**
> RosterTemplate updateRosterTemplate()


### Example

```typescript
import {
    RosterTemplateApi,
    Configuration,
    TemplateUpdateParams
} from './api';

const configuration = new Configuration();
const apiInstance = new RosterTemplateApi(configuration);

let templateId: number; //Template ID (default to undefined)
let params: TemplateUpdateParams; //Update params (optional)

const { status, data } = await apiInstance.updateRosterTemplate(
    templateId,
    params
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **params** | **TemplateUpdateParams**| Update params | |
| **templateId** | [**number**] | Template ID | defaults to undefined|


### Return type

**RosterTemplate**

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

# **updateRosterTemplateShift**
> RosterTemplateShift updateRosterTemplateShift(params)


### Example

```typescript
import {
    RosterTemplateApi,
    Configuration,
    TemplateShiftUpdateRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new RosterTemplateApi(configuration);

let templateId: number; //Template ID (default to undefined)
let shiftId: number; //Template Shift ID (default to undefined)
let params: TemplateShiftUpdateRequest; //Update params

const { status, data } = await apiInstance.updateRosterTemplateShift(
    templateId,
    shiftId,
    params
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **params** | **TemplateShiftUpdateRequest**| Update params | |
| **templateId** | [**number**] | Template ID | defaults to undefined|
| **shiftId** | [**number**] | Template Shift ID | defaults to undefined|


### Return type

**RosterTemplateShift**

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

# **updateRosterTemplateShiftPreference**
> RosterTemplateShiftPreference updateRosterTemplateShiftPreference(params)


### Example

```typescript
import {
    RosterTemplateApi,
    Configuration,
    TemplateShiftPreferenceUpdateRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new RosterTemplateApi(configuration);

let templateId: number; //Template ID (default to undefined)
let preferenceId: number; //Preference ID (default to undefined)
let params: TemplateShiftPreferenceUpdateRequest; //Update params

const { status, data } = await apiInstance.updateRosterTemplateShiftPreference(
    templateId,
    preferenceId,
    params
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **params** | **TemplateShiftPreferenceUpdateRequest**| Update params | |
| **templateId** | [**number**] | Template ID | defaults to undefined|
| **preferenceId** | [**number**] | Preference ID | defaults to undefined|


### Return type

**RosterTemplateShiftPreference**

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

