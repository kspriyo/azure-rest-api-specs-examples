from azure.identity import DefaultAzureCredential
from azure.mgmt.rdbms.mysql import MySQLManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-rdbms
# USAGE
    python server_update.py

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

    response = client.servers.begin_update(
        resource_group_name="testrg",
        server_name="mysqltestsvc4",
        parameters={
            "properties": {"administratorLoginPassword": "<administratorLoginPassword>", "sslEnforcement": "Disabled"}
        },
    ).result()
    print(response)


# x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2017-12-01/examples/ServerUpdate.json
if __name__ == "__main__":
    main()
