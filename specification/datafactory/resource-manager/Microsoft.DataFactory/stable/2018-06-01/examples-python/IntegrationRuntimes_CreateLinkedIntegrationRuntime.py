from azure.identity import DefaultAzureCredential
from azure.mgmt.datafactory import DataFactoryManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datafactory
# USAGE
    python integration_runtimes_create_linked_integration_runtime.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataFactoryManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="12345678-1234-1234-1234-12345678abc",
    )

    response = client.integration_runtimes.create_linked_integration_runtime(
        resource_group_name="exampleResourceGroup",
        factory_name="exampleFactoryName",
        integration_runtime_name="exampleIntegrationRuntime",
        create_linked_integration_runtime_request={
            "dataFactoryLocation": "West US",
            "dataFactoryName": "e9955d6d-56ea-4be3-841c-52a12c1a9981",
            "name": "bfa92911-9fb6-4fbe-8f23-beae87bc1c83",
            "subscriptionId": "061774c7-4b5a-4159-a55b-365581830283",
        },
    )
    print(response)


# x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_CreateLinkedIntegrationRuntime.json
if __name__ == "__main__":
    main()
