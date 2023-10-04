
# API Spec GateWatchApps

## 1.1 MCU

### 1.1.1 Get Settings MCU

Request :
- Method : POST
- Endpoint : `api/getwatchapp/v1/mcu/getSettingMcu`
- Header :
    - Content-Type : application/x-www-form-urlencoded
    - x-api-key : 
- Body (form-data: x-www-form-urlencoded) :
    - token : string, required
- Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer",
    },
        "data":{
            "keypad_password": "string",
            "delay": "string",
            "door_name_mcu": "string"  
        }
 }
```
## 1.2 Registry or Validation Card
### 1.1.1 Registry or Validation Card and Add Log to DB

Request :
- Method : POST
- Endpoint : `/api/getwatchapp/v1/registryandvalidationcard/`
- Header : 
    - Content-Type : application/x-www-form-urlencoded
    - x-api-key : 
- Body Body (form-data: x-www-form-urlencoded) :
    - no_card : string, required
    - door_token : string, required
- Response :

```json 
{
    "meta": {
        "message": "string",
        "code": "integer",
    },
        "data" : {
            "log_status": "integer",
            "message": "string"
        }
}
```





