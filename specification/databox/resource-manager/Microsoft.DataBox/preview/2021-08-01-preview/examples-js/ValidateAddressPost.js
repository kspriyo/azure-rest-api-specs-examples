const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to [DEPRECATED NOTICE: This operation will soon be removed]. This method validates the customer shipping address and provide alternate addresses if any.
 *
 * @summary [DEPRECATED NOTICE: This operation will soon be removed]. This method validates the customer shipping address and provide alternate addresses if any.
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/ValidateAddressPost.json
 */
async function validateAddressPost() {
  const subscriptionId = "fa68082f-8ff7-4a25-95c7-ce9da541242f";
  const location = "westus";
  const validateAddress = {
    deviceType: "DataBox",
    shippingAddress: {
      addressType: "Commercial",
      city: "San Francisco",
      companyName: "Microsoft",
      country: "US",
      postalCode: "94107",
      stateOrProvince: "CA",
      streetAddress1: "16 TOWNSEND ST",
      streetAddress2: "Unit 1",
    },
    validationType: "ValidateAddress",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const result = await client.service.validateAddress(location, validateAddress);
  console.log(result);
}

validateAddressPost().catch(console.error);
