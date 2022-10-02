import com.azure.core.util.Context;

/** Samples for VendorSkus Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/VendorSkuDelete.json
     */
    /**
     * Sample code: Delete the sku of vendor resource.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void deleteTheSkuOfVendorResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.vendorSkus().delete("TestVendor", "TestSku", Context.NONE);
    }
}
