from azure.identity import DefaultAzureCredential
from azure.mgmt.rdbms.mysql import MySQLManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-rdbms
# USAGE
    python configurations_update_by_server.py

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

    response = client.server_parameters.begin_list_update_configurations(
        resource_group_name="testrg",
        server_name="mysqltestsvc1",
        value=[
            {"name": "event_scheduler", "properties": {"value": "OFF"}},
            {"name": "div_precision_increment", "properties": {"value": "4"}},
        ],
    ).result()
    print(response)


# x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2017-12-01/examples/ConfigurationsUpdateByServer.json
if __name__ == "__main__":
    main()
