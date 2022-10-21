from azure.identity import DefaultAzureCredential
from azure.mgmt.elasticsan import ElasticSanManagement

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-elasticsan
# USAGE
    python volume_groups_update_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret 
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ElasticSanManagement(
        credential=DefaultAzureCredential(),
        subscription_id="aaaaaaaaaaaaaaaaaa",
    )

    response = client.volume_groups.begin_update(
        resource_group_name="rgelasticsan",
        elastic_san_name="ti7q-k952-1qB3J_5",
        volume_group_name="u_5I_1j4t3",
        parameters={
            "properties": {
                "encryption": "EncryptionAtRestWithPlatformKey",
                "networkAcls": {"virtualNetworkRules": [{"action": "Allow", "id": "aaaaaaaaaaaaaaaa"}]},
                "protocolType": "Iscsi",
            },
            "tags": {"key7542": "aaaaaaaaaaaaaaaaaaaa"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/VolumeGroups_Update_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
