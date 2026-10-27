# RosterApi

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**createRoster**](#createroster) | **POST** /organs/{organId}/rosters | CreateRoster a new roster|
|[**createRosterResponsible**](#createrosterresponsible) | **POST** /rosters/{rosterId}/responsibles | Create a roster responsible|
|[**deleteRoster**](#deleteroster) | **DELETE** /rosters/{rosterId} | DeleteRoster a roster|
|[**deleteRosterResponsible**](#deleterosterresponsible) | **DELETE** /rosters/{rosterId}/responsibles/{userId} | Delete a roster responsible|
|[**fillRoster**](#fillroster) | **POST** /rosters/{rosterId}/fill | Fills a roster with the linked user template preferences|
|[**getRoster**](#getroster) | **GET** /rosters/{rosterId} | Get a specific roster by id|
|[**getRosterResponsibles**](#getrosterresponsibles) | **GET** /rosters/{rosterId}/responsibles | Get roster responsibles|
|[**getRosters**](#getrosters) | **GET** /organs/{organId}/rosters | Get all rosters of an organ, optionally filtered by date or archived state|
|[**updateRoster**](#updateroster) | **PATCH** /rosters/{rosterId} | Update a roster|

# **createRoster**
> Roster createRoster(createParams)


### Example

```typescript
import {
    RosterApi,
    Configuration,
    RosterCreateRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new RosterApi(configuration);

let organId: number; //Organ ID (default to undefined)
let createParams: RosterCreateRequest; //Roster input

const { status, data } = await apiInstance.createRoster(
    organId,
    createParams
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createParams** | **RosterCreateRequest**| Roster input | |
| **organId** | [**number**] | Organ ID | defaults to undefined|


### Return type

**Roster**

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

# **createRosterResponsible**
> RosterResponsible createRosterResponsible(params)

Assigns a user as responsible for a given roster. Requires admin level or higher in the roster\'s organ.

### Example

```typescript
import {
    RosterApi,
    Configuration,
    RosterResponsibleCreateRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new RosterApi(configuration);

let rosterId: number; //Roster ID (default to undefined)
let params: RosterResponsibleCreateRequest; //User to assign as responsible

const { status, data } = await apiInstance.createRosterResponsible(
    rosterId,
    params
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **params** | **RosterResponsibleCreateRequest**| User to assign as responsible | |
| **rosterId** | [**number**] | Roster ID | defaults to undefined|


### Return type

**RosterResponsible**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**201** | Created |  -  |
|**400** | Invalid request |  -  |
|**403** | Insufficient organ permissions |  -  |
|**404** | Roster not found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteRoster**
> string deleteRoster()


### Example

```typescript
import {
    RosterApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new RosterApi(configuration);

let rosterId: number; //Roster ID (default to undefined)

const { status, data } = await apiInstance.deleteRoster(
    rosterId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **rosterId** | [**number**] | Roster ID | defaults to undefined|


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

# **deleteRosterResponsible**
> { [key: string]: string; } deleteRosterResponsible()

Removes a user\'s responsible assignment for a given roster. Requires admin level or higher in the roster\'s organ.

### Example

```typescript
import {
    RosterApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new RosterApi(configuration);

let rosterId: number; //Roster ID (default to undefined)
let userId: number; //User ID (default to undefined)

const { status, data } = await apiInstance.deleteRosterResponsible(
    rosterId,
    userId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **rosterId** | [**number**] | Roster ID | defaults to undefined|
| **userId** | [**number**] | User ID | defaults to undefined|


### Return type

**{ [key: string]: string; }**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**400** | Invalid request |  -  |
|**403** | Insufficient organ permissions |  -  |
|**404** | Responsible not found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **fillRoster**
> Array<RosterAnswer> fillRoster()


### Example

```typescript
import {
    RosterApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new RosterApi(configuration);

let rosterId: number; //Roster ID (default to undefined)

const { status, data } = await apiInstance.fillRoster(
    rosterId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **rosterId** | [**number**] | Roster ID | defaults to undefined|


### Return type

**Array<RosterAnswer>**

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

# **getRoster**
> Roster getRoster()


### Example

```typescript
import {
    RosterApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new RosterApi(configuration);

let rosterId: number; //Roster ID (default to undefined)

const { status, data } = await apiInstance.getRoster(
    rosterId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **rosterId** | [**number**] | Roster ID | defaults to undefined|


### Return type

**Roster**

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

# **getRosterResponsibles**
> Array<RosterResponsible> getRosterResponsibles()

Returns the responsibles for a given roster. Callable by any member of the roster\'s organ.

### Example

```typescript
import {
    RosterApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new RosterApi(configuration);

let rosterId: number; //Roster ID (default to undefined)

const { status, data } = await apiInstance.getRosterResponsibles(
    rosterId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **rosterId** | [**number**] | Roster ID | defaults to undefined|


### Return type

**Array<RosterResponsible>**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**400** | Invalid roster ID |  -  |
|**403** | Insufficient organ permissions |  -  |
|**404** | Roster not found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getRosters**
> Array<Roster> getRosters()


### Example

```typescript
import {
    RosterApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new RosterApi(configuration);

let organId: number; //Organ ID (default to undefined)
let archived: boolean; // (optional) (default to undefined)
let date: string; // (optional) (default to undefined)
let id: number; // (optional) (default to undefined)

const { status, data } = await apiInstance.getRosters(
    organId,
    archived,
    date,
    id
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **organId** | [**number**] | Organ ID | defaults to undefined|
| **archived** | [**boolean**] |  | (optional) defaults to undefined|
| **date** | [**string**] |  | (optional) defaults to undefined|
| **id** | [**number**] |  | (optional) defaults to undefined|


### Return type

**Array<Roster>**

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

# **updateRoster**
> Roster updateRoster(updateParams)


### Example

```typescript
import {
    RosterApi,
    Configuration,
    RosterUpdateRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new RosterApi(configuration);

let rosterId: number; //Roster ID (default to undefined)
let updateParams: RosterUpdateRequest; //Roster input

const { status, data } = await apiInstance.updateRoster(
    rosterId,
    updateParams
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateParams** | **RosterUpdateRequest**| Roster input | |
| **rosterId** | [**number**] | Roster ID | defaults to undefined|


### Return type

**Roster**

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

