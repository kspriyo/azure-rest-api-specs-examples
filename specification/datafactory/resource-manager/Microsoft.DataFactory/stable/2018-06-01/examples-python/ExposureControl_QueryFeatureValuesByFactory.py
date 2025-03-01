from azure.identity import DefaultAzureCredential
from azure.mgmt.datafactory import DataFactoryManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datafactory
# USAGE
    python exposure_control_query_feature_values_by_factory.py

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

    response = client.exposure_control.query_feature_values_by_factory(
        resource_group_name="exampleResourceGroup",
        factory_name="exampleFactoryName",
        exposure_control_batch_request={
            "exposureControlRequests": [
                {"featureName": "ADFIntegrationRuntimeSharingRbac", "featureType": "Feature"},
                {"featureName": "ADFSampleFeature", "featureType": "Feature"},
            ]
        },
    )
    print(response)


# x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ExposureControl_QueryFeatureValuesByFactory.json
if __name__ == "__main__":
    main()
