const { RecoveryServicesClient } = require("@azure/arm-recoveryservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Uploads a certificate for a resource.
 *
 * @summary Uploads a certificate for a resource.
 * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2022-04-01/examples/PUTVaultCred.json
 */
async function downloadVaultCredentialFile() {
  const subscriptionId = "77777777-d41f-4550-9f70-7708a3a2283b";
  const resourceGroupName = "BCDRIbzRG";
  const vaultName = "BCDRIbzVault";
  const certificateName =
    "BCDRIbzVault77777777-d41f-4550-9f70-7708a3a2283b-12-18-2017-vaultcredentials";
  const certificateRequest = {
    properties: {
      authType: "AAD",
      certificate: Buffer.from(
        "MTTC3TCCAcWgAwIBAgIQEj9h+ZLlXK9KrqZX9UkAnzANBgkqhkiG9w0BAQUFADAeMRwwGgYDVQQDExNXaW5kb3dzIEF6dXJlIFRvb2xzMB4XDTE3MTIxODA5MTc1M1oXDTE3MTIyMzA5Mjc1M1owHjEcMBoGA1UEAxMTV2luZG93cyBBenVyZSBUb29sczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAK773/eZZ69RbZZAT05r9MjUxu9y1L1Pn1EgPk62IPJyHlO3OZA922eSBahhP4bgmFljN4LVReqQ5eT/wqO0Zhc+yFkUy4U4RdbQLeUZt2W7yy9XLXgVvqeYDgsjg/QhHetgHArQBW+tlQq5+zPdU7zchI4rbShSJrWhLrZFWiOyFPsuAE4joUQHNlRifdCTsBGKk8HRCY3j1S3c4bfEn3zxlrvrXXssRuW5mJM95rMk0tskoRxXSCi6i9bnlki2Cs9mpVMmBFeofs41KwzlWU0TgpdD8s1QEdvfGB5NbByfetPX7MfJaTBeHZEGbv/Iq8l72u8sPBoOhcaH7qDE/mECAwEAAaMXMBUwEwYDVR0lBAwwCgYIKwYBBQUHAwIwDQYJKoZIhvcNAQEFBQADggEBAILfgHluye1Q+WelhgWhpBBdIq2C0btfV8eFsZaTlBUrM0fwpxQSlAWc2oYHVMQI4A5iUjbDOY35O4yc+TnWKDBKf+laqDP+yos4aiUPuadGUZfvDk7kuw7xeECs64JpHAIEKdRHFW9rD3gwG+nIWaDnEL/7rTyhL3kXrRW2MSUAL8g3GX8Z45c+MQY0jmASIqWdhGn1vpAGyA9mKkzsqg7FXjg8GZb24tGl5Ky85+ip4dkBfXinDD8WwaGyjhGGK97ErvNmN36qly/H0H1Qngiovg1FbHDmkcFO5QclnEJsFFmcO2CcHp5Fqh2wXn5O1cQaxCIRTpQ/uXRpDjl2wKs="
      ),
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesClient(credential, subscriptionId);
  const result = await client.vaultCertificates.create(
    resourceGroupName,
    vaultName,
    certificateName,
    certificateRequest
  );
  console.log(result);
}

downloadVaultCredentialFile().catch(console.error);
