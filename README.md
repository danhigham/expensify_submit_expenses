# handy curls

## get policy list for user

```
curl -X POST 'https://integrations.expensify.com/Integration-Server/ExpensifyIntegrations' \
    -d 'requestJobDescription={
        "type":"get",
        "credentials":{
            "partnerUserID": "aa_dan_higham_sysdig_com",
            "partnerUserSecret": "075c46ff080abd3c17c077368393396b9d225bd1"
        },
        "inputSettings":{
            "type":"policyList",
            "adminOnly":false,
            "userEmail":"dan.higham@sysdig.com"
        }
    }' | jq
```

## get category names for policy

```
curl -X POST 'https://integrations.expensify.com/Integration-Server/ExpensifyIntegrations' \
    -d 'requestJobDescription={
        "type":"get",
        "credentials":{
            "partnerUserID": "aa_dan_higham_sysdig_com",
            "partnerUserSecret": "075c46ff080abd3c17c077368393396b9d225bd1"
        },
        "inputSettings":{
            "type":"policy",
            "fields": ["categories"],
            "policyIDList": ["AA4D777BBAA4523A"],
            "userEmail":"dan.higham@sysdig.com"
        }
    }' | jq -r ".policyInfo.AA4D777BBAA4523A.categories[].name"
```