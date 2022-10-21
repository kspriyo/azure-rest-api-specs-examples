from azure.identity import DefaultAzureCredential
from azure.mgmt.cdn import CdnManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cdn
# USAGE
    python endpoints_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret 
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CdnManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.endpoints.begin_create(
        resource_group_name="RG",
        profile_name="profile1",
        endpoint_name="endpoint1",
        endpoint={
            "location": "WestUs",
            "properties": {
                "contentTypesToCompress": ["text/html", "application/octet-stream"],
                "defaultOriginGroup": {
                    "id": "/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/originGroups/originGroup1"
                },
                "deliveryPolicy": {
                    "description": "Test description for a policy.",
                    "rules": [
                        {
                            "actions": [
                                {
                                    "name": "CacheExpiration",
                                    "parameters": {
                                        "cacheBehavior": "Override",
                                        "cacheDuration": "10:10:09",
                                        "cacheType": "All",
                                        "typeName": "DeliveryRuleCacheExpirationActionParameters",
                                    },
                                },
                                {
                                    "name": "ModifyResponseHeader",
                                    "parameters": {
                                        "headerAction": "Overwrite",
                                        "headerName": "Access-Control-Allow-Origin",
                                        "typeName": "DeliveryRuleHeaderActionParameters",
                                        "value": "*",
                                    },
                                },
                                {
                                    "name": "ModifyRequestHeader",
                                    "parameters": {
                                        "headerAction": "Overwrite",
                                        "headerName": "Accept-Encoding",
                                        "typeName": "DeliveryRuleHeaderActionParameters",
                                        "value": "gzip",
                                    },
                                },
                            ],
                            "conditions": [
                                {
                                    "name": "RemoteAddress",
                                    "parameters": {
                                        "matchValues": ["192.168.1.0/24", "10.0.0.0/24"],
                                        "negateCondition": True,
                                        "operator": "IPMatch",
                                        "typeName": "DeliveryRuleRemoteAddressConditionParameters",
                                    },
                                }
                            ],
                            "name": "rule1",
                            "order": 1,
                        }
                    ],
                },
                "isCompressionEnabled": True,
                "isHttpAllowed": True,
                "isHttpsAllowed": True,
                "originGroups": [
                    {
                        "name": "originGroup1",
                        "properties": {
                            "healthProbeSettings": {
                                "probeIntervalInSeconds": 120,
                                "probePath": "/health.aspx",
                                "probeProtocol": "Http",
                                "probeRequestType": "GET",
                            },
                            "origins": [
                                {
                                    "id": "/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/origin1"
                                },
                                {
                                    "id": "/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/origin2"
                                },
                            ],
                            "responseBasedOriginErrorDetectionSettings": {
                                "responseBasedDetectedErrorTypes": "TcpErrorsOnly",
                                "responseBasedFailoverThresholdPercentage": 10,
                            },
                        },
                    }
                ],
                "originHostHeader": "www.bing.com",
                "originPath": "/photos",
                "origins": [
                    {
                        "name": "origin1",
                        "properties": {
                            "enabled": True,
                            "hostName": "www.someDomain1.net",
                            "httpPort": 80,
                            "httpsPort": 443,
                            "originHostHeader": "www.someDomain1.net",
                            "priority": 1,
                            "weight": 50,
                        },
                    },
                    {
                        "name": "origin2",
                        "properties": {
                            "enabled": True,
                            "hostName": "www.someDomain2.net",
                            "httpPort": 80,
                            "httpsPort": 443,
                            "originHostHeader": "www.someDomain2.net",
                            "priority": 2,
                            "weight": 50,
                        },
                    },
                ],
                "queryStringCachingBehavior": "BypassCaching",
            },
            "tags": {"key1": "value1"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Endpoints_Create.json
if __name__ == "__main__":
    main()
