from azure.identity import DefaultAzureCredential
from azure.mgmt.rdbms.mysql import MySQLManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-rdbms
# USAGE
    python server_admin_create_update.py

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

    response = client.server_administrators.begin_create_or_update(
        resource_group_name="testrg",
        server_name="mysqltestsvc4",
        properties={
            "properties": {
                "administratorType": "ActiveDirectory",
                "login": "bob@contoso.com",
                "sid": "c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c",
                "tenantId": "c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2017-12-01/examples/ServerAdminCreateUpdate.json
if __name__ == "__main__":
    main()
