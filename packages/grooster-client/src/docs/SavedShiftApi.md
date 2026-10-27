# SavedShiftApi

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**getSavedShifts**](#getsavedshifts) | **GET** /rosters/{rosterId}/saved-shifts | Get all saved shifts for a specific roster|
|[**lockRoster**](#lockroster) | **POST** /rosters/{rosterId}/lock | Lock a specific roster, storing its current assignments as saved shifts|
|[**updateSavedShift**](#updatesavedshift) | **PATCH** /rosters/{rosterId}/saved-shifts/{savedShiftId} | Update a specific saved shift|

# **getSavedShifts**
> SavedShiftResponse getSavedShifts()


### Example

```typescript
import {
    SavedShiftApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new SavedShiftApi(configuration);

let rosterId: number; //Roster ID (default to undefined)

const { status, data } = await apiInstance.getSavedShifts(
    rosterId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **rosterId** | [**number**] | Roster ID | defaults to undefined|


### Return type

**SavedShiftResponse**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Saved Shift Response |  -  |
|**400** | Invalid request |  -  |
|**404** | SavedShift not found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **lockRoster**
> string lockRoster()


### Example

```typescript
import {
    SavedShiftApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new SavedShiftApi(configuration);

let rosterId: number; //Roster ID (default to undefined)

const { status, data } = await apiInstance.lockRoster(
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

# **updateSavedShift**
> SavedShift updateSavedShift(updateParams)


### Example

```typescript
import {
    SavedShiftApi,
    Configuration,
    SavedShiftUpdateRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new SavedShiftApi(configuration);

let rosterId: number; //Roster ID (default to undefined)
let savedShiftId: number; //SavedShift ID (default to undefined)
let updateParams: SavedShiftUpdateRequest; //Update data

const { status, data } = await apiInstance.updateSavedShift(
    rosterId,
    savedShiftId,
    updateParams
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateParams** | **SavedShiftUpdateRequest**| Update data | |
| **rosterId** | [**number**] | Roster ID | defaults to undefined|
| **savedShiftId** | [**number**] | SavedShift ID | defaults to undefined|


### Return type

**SavedShift**

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

