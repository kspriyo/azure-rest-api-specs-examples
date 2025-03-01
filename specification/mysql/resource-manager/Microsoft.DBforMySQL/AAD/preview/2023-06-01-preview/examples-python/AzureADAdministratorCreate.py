from azure.identity import DefaultAzureCredential
from azure.mgmt.rdbms.mysql_flexibleservers import MySQLManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-rdbms
# USAGE
    python azure_ad_administrator_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MySQLManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="ffffffff-ffff-ffff-ffff-ffffffffffff",
    )

    response = client.azure_ad_administrators.begin_create_or_update(
        resource_group_name="testrg",
        server_name="mysqltestsvc4",
        administrator_name="ActiveDirectory",
        parameters={
            "properties": {
                "administratorType": "ActiveDirectory",
                "identityResourceId": "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/test-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-umi",
                "login": "bob@contoso.com",
                "sid": "c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c",
                "tenantId": "c12b7025-bfe2-46c1-b463-993b5e4cd467",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/AAD/preview/2023-06-01-preview/examples/AzureADAdministratorCreate.json
if __name__ == "__main__":
    main()
