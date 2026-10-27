# ShiftGroupApi

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**createShiftGroup**](#createshiftgroup) | **POST** /organs/{organId}/shift-groups | Create a new shift group|
|[**getShiftGroup**](#getshiftgroup) | **GET** /shift-groups/{shiftGroupId} | Get a specific shift group by ID|
|[**getShiftGroupPriorities**](#getshiftgrouppriorities) | **GET** /shift-groups/{shiftGroupId}/priorities | Get a shift group priorities for a shift group|
|[**getShiftGroups**](#getshiftgroups) | **GET** /organs/{organId}/shift-groups | Get all shift groups for an organ|
|[**pushUserToBottom**](#pushusertobottom) | **POST** /shift-groups/{shiftGroupId}/push-to-bottom | Manually push a user to the bottom of a shift group\&#39;s ordering, as if just assigned|
|[**updateShiftGroupPriority**](#updateshiftgrouppriority) | **PUT** /shift-groups/{shiftGroupId}/priorities | Update a shift group priority|

# **createShiftGroup**
> ShiftGroup createShiftGroup(params)


### Example

```typescript
import {
    ShiftGroupApi,
    Configuration,
    ShiftGroupCreateRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new ShiftGroupApi(configuration);

let organId: number; //Organ ID (default to undefined)
let params: ShiftGroupCreateRequest; //Shift Group Details

const { status, data } = await apiInstance.createShiftGroup(
    organId,
    params
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **params** | **ShiftGroupCreateRequest**| Shift Group Details | |
| **organId** | [**number**] | Organ ID | defaults to undefined|


### Return type

**ShiftGroup**

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

# **getShiftGroup**
> ShiftGroup getShiftGroup()


### Example

```typescript
import {
    ShiftGroupApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ShiftGroupApi(configuration);

let shiftGroupId: number; //Shift Group ID (default to undefined)

const { status, data } = await apiInstance.getShiftGroup(
    shiftGroupId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **shiftGroupId** | [**number**] | Shift Group ID | defaults to undefined|


### Return type

**ShiftGroup**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**404** | Not Found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getShiftGroupPriorities**
> Array<ShiftGroupPriority> getShiftGroupPriorities()


### Example

```typescript
import {
    ShiftGroupApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ShiftGroupApi(configuration);

let shiftGroupId: number; //ShiftGroup ID (default to undefined)

const { status, data } = await apiInstance.getShiftGroupPriorities(
    shiftGroupId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **shiftGroupId** | [**number**] | ShiftGroup ID | defaults to undefined|


### Return type

**Array<ShiftGroupPriority>**

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getShiftGroups**
> Array<ShiftGroup> getShiftGroups()


### Example

```typescript
import {
    ShiftGroupApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ShiftGroupApi(configuration);

let organId: number; //Organ ID (default to undefined)

const { status, data } = await apiInstance.getShiftGroups(
    organId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **organId** | [**number**] | Organ ID | defaults to undefined|


### Return type

**Array<ShiftGroup>**

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

# **pushUserToBottom**
> string pushUserToBottom(params)


### Example

```typescript
import {
    ShiftGroupApi,
    Configuration,
    PushToBottomRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new ShiftGroupApi(configuration);

let shiftGroupId: number; //Shift Group ID (default to undefined)
let params: PushToBottomRequest; //User to push to the bottom

const { status, data } = await apiInstance.pushUserToBottom(
    shiftGroupId,
    params
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **params** | **PushToBottomRequest**| User to push to the bottom | |
| **shiftGroupId** | [**number**] | Shift Group ID | defaults to undefined|


### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**400** | Invalid request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **updateShiftGroupPriority**
> ShiftGroupPriority updateShiftGroupPriority(updateParams)


### Example

```typescript
import {
    ShiftGroupApi,
    Configuration,
    GroupPriorityUpdateParam
} from './api';

const configuration = new Configuration();
const apiInstance = new ShiftGroupApi(configuration);

let shiftGroupId: number; //ShiftGroup ID (default to undefined)
let updateParams: GroupPriorityUpdateParam; //Update parameters

const { status, data } = await apiInstance.updateShiftGroupPriority(
    shiftGroupId,
    updateParams
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateParams** | **GroupPriorityUpdateParam**| Update parameters | |
| **shiftGroupId** | [**number**] | ShiftGroup ID | defaults to undefined|


### Return type

**ShiftGroupPriority**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**400** | Invalid request |  -  |
|**404** | SavedShift not found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

