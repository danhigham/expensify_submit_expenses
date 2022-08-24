# handy curls

## get policy list for user

```
curl -X POST 'https://integrations.expensify.com/Integration-Server/ExpensifyIntegrations' \
    -d 'requestJobDescription={
        "type":"get",
        "credentials":{
            "partnerUserID": "aa_john_smith",
            "partnerUserSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
        },
        "inputSettings":{
            "type":"policyList",
            "adminOnly":false,
            "userEmail":"john.smith@exameple.com"
        }
    }' | jq
```

## get category names for policy

```
curl -X POST 'https://integrations.expensify.com/Integration-Server/ExpensifyIntegrations' \
    -d 'requestJobDescription={
        "type":"get",
        "credentials":{
            "partnerUserID": "aa_john_smith",
            "partnerUserSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
        },
        "inputSettings":{
            "type":"policy",
            "fields": ["categories"],
            "policyIDList": ["XXXXXXXXXXXXXXXX"],
            "userEmail":"john.smith@exameple.com"
        }
    }' | jq -r ".policyInfo.XXXXXXXXXXXXXXXX.categories[].name"
```