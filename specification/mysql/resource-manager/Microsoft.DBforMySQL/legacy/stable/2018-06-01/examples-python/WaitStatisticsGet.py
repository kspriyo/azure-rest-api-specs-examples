from azure.identity import DefaultAzureCredential
from azure.mgmt.rdbms.mysql import MySQLManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-rdbms
# USAGE
    python wait_statistics_get.py

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

    response = client.wait_statistics.get(
        resource_group_name="testResourceGroupName",
        server_name="testServerName",
        wait_statistics_id="636927606000000000-636927615000000000-send-wait/io/socket/sql/client_connection-2--0",
    )
    print(response)


# x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2018-06-01/examples/WaitStatisticsGet.json
if __name__ == "__main__":
    main()
