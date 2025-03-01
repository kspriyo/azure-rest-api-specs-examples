package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceTimeZoneListByLocation.json
func ExampleTimeZonesClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTimeZonesClient().NewListByLocationPager("canadaeast", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.TimeZoneListResult = armsql.TimeZoneListResult{
		// 	Value: []*armsql.TimeZone{
		// 		{
		// 			Name: to.Ptr("Afghanistan Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Afghanistan Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+04:30) Kabul"),
		// 				TimeZoneID: to.Ptr("Afghanistan Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Alaskan Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Alaskan Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-09:00) Alaska"),
		// 				TimeZoneID: to.Ptr("Alaskan Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Aleutian Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Aleutian Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-10:00) Aleutian Islands"),
		// 				TimeZoneID: to.Ptr("Aleutian Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Altai Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Altai Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+07:00) Barnaul, Gorno-Altaysk"),
		// 				TimeZoneID: to.Ptr("Altai Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Arab Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Arab Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+03:00) Kuwait, Riyadh"),
		// 				TimeZoneID: to.Ptr("Arab Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Arabian Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Arabian Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+04:00) Abu Dhabi, Muscat"),
		// 				TimeZoneID: to.Ptr("Arabian Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Arabic Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Arabic Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+03:00) Baghdad"),
		// 				TimeZoneID: to.Ptr("Arabic Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Argentina Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Argentina Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-03:00) City of Buenos Aires"),
		// 				TimeZoneID: to.Ptr("Argentina Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Astrakhan Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Astrakhan Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+04:00) Astrakhan, Ulyanovsk"),
		// 				TimeZoneID: to.Ptr("Astrakhan Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Atlantic Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Atlantic Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-04:00) Atlantic Time (Canada)"),
		// 				TimeZoneID: to.Ptr("Atlantic Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("AUS Central Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/AUS Central Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+09:30) Darwin"),
		// 				TimeZoneID: to.Ptr("AUS Central Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Aus Central W. Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Aus Central W. Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+08:45) Eucla"),
		// 				TimeZoneID: to.Ptr("Aus Central W. Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("AUS Eastern Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/AUS Eastern Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+10:00) Canberra, Melbourne, Sydney"),
		// 				TimeZoneID: to.Ptr("AUS Eastern Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Azerbaijan Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Azerbaijan Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+04:00) Baku"),
		// 				TimeZoneID: to.Ptr("Azerbaijan Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Azores Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Azores Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-01:00) Azores"),
		// 				TimeZoneID: to.Ptr("Azores Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Bahia Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Bahia Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-03:00) Salvador"),
		// 				TimeZoneID: to.Ptr("Bahia Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Bangladesh Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Bangladesh Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+06:00) Dhaka"),
		// 				TimeZoneID: to.Ptr("Bangladesh Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Belarus Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Belarus Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+03:00) Minsk"),
		// 				TimeZoneID: to.Ptr("Belarus Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Bougainville Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Bougainville Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+11:00) Bougainville Island"),
		// 				TimeZoneID: to.Ptr("Bougainville Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Canada Central Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Canada Central Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-06:00) Saskatchewan"),
		// 				TimeZoneID: to.Ptr("Canada Central Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Cape Verde Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Cape Verde Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-01:00) Cabo Verde Is."),
		// 				TimeZoneID: to.Ptr("Cape Verde Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Caucasus Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Caucasus Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+04:00) Yerevan"),
		// 				TimeZoneID: to.Ptr("Caucasus Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Cen. Australia Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Cen. Australia Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+09:30) Adelaide"),
		// 				TimeZoneID: to.Ptr("Cen. Australia Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Central America Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Central America Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-06:00) Central America"),
		// 				TimeZoneID: to.Ptr("Central America Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Central Asia Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Central Asia Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+06:00) Nur-Sultan"),
		// 				TimeZoneID: to.Ptr("Central Asia Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Central Brazilian Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Central Brazilian Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-04:00) Cuiaba"),
		// 				TimeZoneID: to.Ptr("Central Brazilian Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Central Europe Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Central Europe Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+01:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague"),
		// 				TimeZoneID: to.Ptr("Central Europe Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Central European Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Central European Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+01:00) Sarajevo, Skopje, Warsaw, Zagreb"),
		// 				TimeZoneID: to.Ptr("Central European Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Central Pacific Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Central Pacific Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+11:00) Solomon Is., New Caledonia"),
		// 				TimeZoneID: to.Ptr("Central Pacific Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Central Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Central Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-06:00) Central Time (US & Canada)"),
		// 				TimeZoneID: to.Ptr("Central Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Central Standard Time (Mexico)"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Central Standard Time (Mexico)"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-06:00) Guadalajara, Mexico City, Monterrey"),
		// 				TimeZoneID: to.Ptr("Central Standard Time (Mexico)"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Chatham Islands Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Chatham Islands Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+12:45) Chatham Islands"),
		// 				TimeZoneID: to.Ptr("Chatham Islands Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("China Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/China Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+08:00) Beijing, Chongqing, Hong Kong, Urumqi"),
		// 				TimeZoneID: to.Ptr("China Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Cuba Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Cuba Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-05:00) Havana"),
		// 				TimeZoneID: to.Ptr("Cuba Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Dateline Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Dateline Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-12:00) International Date Line West"),
		// 				TimeZoneID: to.Ptr("Dateline Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("E. Africa Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/E. Africa Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+03:00) Nairobi"),
		// 				TimeZoneID: to.Ptr("E. Africa Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("E. Australia Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/E. Australia Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+10:00) Brisbane"),
		// 				TimeZoneID: to.Ptr("E. Australia Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("E. Europe Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/E. Europe Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+02:00) Chisinau"),
		// 				TimeZoneID: to.Ptr("E. Europe Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("E. South America Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/E. South America Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-03:00) Brasilia"),
		// 				TimeZoneID: to.Ptr("E. South America Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Easter Island Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Easter Island Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-06:00) Easter Island"),
		// 				TimeZoneID: to.Ptr("Easter Island Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Eastern Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Eastern Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-05:00) Eastern Time (US & Canada)"),
		// 				TimeZoneID: to.Ptr("Eastern Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Eastern Standard Time (Mexico)"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Eastern Standard Time (Mexico)"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-05:00) Chetumal"),
		// 				TimeZoneID: to.Ptr("Eastern Standard Time (Mexico)"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Egypt Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Egypt Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+02:00) Cairo"),
		// 				TimeZoneID: to.Ptr("Egypt Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Ekaterinburg Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Ekaterinburg Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+05:00) Ekaterinburg"),
		// 				TimeZoneID: to.Ptr("Ekaterinburg Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Fiji Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Fiji Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+12:00) Fiji"),
		// 				TimeZoneID: to.Ptr("Fiji Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("FLE Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/FLE Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+02:00) Helsinki, Kyiv, Riga, Sofia, Tallinn, Vilnius"),
		// 				TimeZoneID: to.Ptr("FLE Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Georgian Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Georgian Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+04:00) Tbilisi"),
		// 				TimeZoneID: to.Ptr("Georgian Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("GMT Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/GMT Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+00:00) Dublin, Edinburgh, Lisbon, London"),
		// 				TimeZoneID: to.Ptr("GMT Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Greenland Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Greenland Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-03:00) Greenland"),
		// 				TimeZoneID: to.Ptr("Greenland Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Greenwich Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Greenwich Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+00:00) Monrovia, Reykjavik"),
		// 				TimeZoneID: to.Ptr("Greenwich Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("GTB Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/GTB Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+02:00) Athens, Bucharest"),
		// 				TimeZoneID: to.Ptr("GTB Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Haiti Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Haiti Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-05:00) Haiti"),
		// 				TimeZoneID: to.Ptr("Haiti Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Hawaiian Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Hawaiian Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-10:00) Hawaii"),
		// 				TimeZoneID: to.Ptr("Hawaiian Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("India Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/India Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+05:30) Chennai, Kolkata, Mumbai, New Delhi"),
		// 				TimeZoneID: to.Ptr("India Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Iran Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Iran Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+03:30) Tehran"),
		// 				TimeZoneID: to.Ptr("Iran Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Israel Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Israel Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+02:00) Jerusalem"),
		// 				TimeZoneID: to.Ptr("Israel Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Jordan Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Jordan Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+02:00) Amman"),
		// 				TimeZoneID: to.Ptr("Jordan Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Kaliningrad Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Kaliningrad Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+02:00) Kaliningrad"),
		// 				TimeZoneID: to.Ptr("Kaliningrad Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Kamchatka Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Kamchatka Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+12:00) Petropavlovsk-Kamchatsky - Old"),
		// 				TimeZoneID: to.Ptr("Kamchatka Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Korea Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Korea Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+09:00) Seoul"),
		// 				TimeZoneID: to.Ptr("Korea Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Libya Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Libya Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+02:00) Tripoli"),
		// 				TimeZoneID: to.Ptr("Libya Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Line Islands Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Line Islands Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+14:00) Kiritimati Island"),
		// 				TimeZoneID: to.Ptr("Line Islands Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Lord Howe Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Lord Howe Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+10:30) Lord Howe Island"),
		// 				TimeZoneID: to.Ptr("Lord Howe Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Magadan Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Magadan Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+11:00) Magadan"),
		// 				TimeZoneID: to.Ptr("Magadan Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Magallanes Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Magallanes Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-03:00) Punta Arenas"),
		// 				TimeZoneID: to.Ptr("Magallanes Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Marquesas Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Marquesas Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-09:30) Marquesas Islands"),
		// 				TimeZoneID: to.Ptr("Marquesas Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Mauritius Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Mauritius Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+04:00) Port Louis"),
		// 				TimeZoneID: to.Ptr("Mauritius Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Mid-Atlantic Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Mid-Atlantic Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-02:00) Mid-Atlantic - Old"),
		// 				TimeZoneID: to.Ptr("Mid-Atlantic Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Middle East Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Middle East Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+02:00) Beirut"),
		// 				TimeZoneID: to.Ptr("Middle East Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Montevideo Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Montevideo Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-03:00) Montevideo"),
		// 				TimeZoneID: to.Ptr("Montevideo Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Morocco Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Morocco Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+01:00) Casablanca"),
		// 				TimeZoneID: to.Ptr("Morocco Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Mountain Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Mountain Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-07:00) Mountain Time (US & Canada)"),
		// 				TimeZoneID: to.Ptr("Mountain Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Mountain Standard Time (Mexico)"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Mountain Standard Time (Mexico)"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-07:00) Chihuahua, La Paz, Mazatlan"),
		// 				TimeZoneID: to.Ptr("Mountain Standard Time (Mexico)"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Myanmar Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Myanmar Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+06:30) Yangon (Rangoon)"),
		// 				TimeZoneID: to.Ptr("Myanmar Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("N. Central Asia Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/N. Central Asia Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+07:00) Novosibirsk"),
		// 				TimeZoneID: to.Ptr("N. Central Asia Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Namibia Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Namibia Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+02:00) Windhoek"),
		// 				TimeZoneID: to.Ptr("Namibia Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Nepal Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Nepal Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+05:45) Kathmandu"),
		// 				TimeZoneID: to.Ptr("Nepal Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("New Zealand Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/New Zealand Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+12:00) Auckland, Wellington"),
		// 				TimeZoneID: to.Ptr("New Zealand Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Newfoundland Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Newfoundland Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-03:30) Newfoundland"),
		// 				TimeZoneID: to.Ptr("Newfoundland Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Norfolk Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Norfolk Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+11:00) Norfolk Island"),
		// 				TimeZoneID: to.Ptr("Norfolk Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("North Asia East Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/North Asia East Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+08:00) Irkutsk"),
		// 				TimeZoneID: to.Ptr("North Asia East Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("North Asia Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/North Asia Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+07:00) Krasnoyarsk"),
		// 				TimeZoneID: to.Ptr("North Asia Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("North Korea Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/North Korea Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+09:00) Pyongyang"),
		// 				TimeZoneID: to.Ptr("North Korea Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Omsk Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Omsk Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+06:00) Omsk"),
		// 				TimeZoneID: to.Ptr("Omsk Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Pacific SA Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Pacific SA Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-04:00) Santiago"),
		// 				TimeZoneID: to.Ptr("Pacific SA Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Pacific Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Pacific Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-08:00) Pacific Time (US & Canada)"),
		// 				TimeZoneID: to.Ptr("Pacific Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Pacific Standard Time (Mexico)"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Pacific Standard Time (Mexico)"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-08:00) Baja California"),
		// 				TimeZoneID: to.Ptr("Pacific Standard Time (Mexico)"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Pakistan Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Pakistan Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+05:00) Islamabad, Karachi"),
		// 				TimeZoneID: to.Ptr("Pakistan Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Paraguay Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Paraguay Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-04:00) Asuncion"),
		// 				TimeZoneID: to.Ptr("Paraguay Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Qyzylorda Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Qyzylorda Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+05:00) Qyzylorda"),
		// 				TimeZoneID: to.Ptr("Qyzylorda Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Romance Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Romance Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+01:00) Brussels, Copenhagen, Madrid, Paris"),
		// 				TimeZoneID: to.Ptr("Romance Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Russia Time Zone 10"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Russia Time Zone 10"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+11:00) Chokurdakh"),
		// 				TimeZoneID: to.Ptr("Russia Time Zone 10"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Russia Time Zone 11"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Russia Time Zone 11"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+12:00) Anadyr, Petropavlovsk-Kamchatsky"),
		// 				TimeZoneID: to.Ptr("Russia Time Zone 11"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Russia Time Zone 3"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Russia Time Zone 3"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+04:00) Izhevsk, Samara"),
		// 				TimeZoneID: to.Ptr("Russia Time Zone 3"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Russian Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Russian Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+03:00) Moscow, St. Petersburg"),
		// 				TimeZoneID: to.Ptr("Russian Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SA Eastern Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/SA Eastern Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-03:00) Cayenne, Fortaleza"),
		// 				TimeZoneID: to.Ptr("SA Eastern Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SA Pacific Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/SA Pacific Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-05:00) Bogota, Lima, Quito, Rio Branco"),
		// 				TimeZoneID: to.Ptr("SA Pacific Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SA Western Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/SA Western Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-04:00) Georgetown, La Paz, Manaus, San Juan"),
		// 				TimeZoneID: to.Ptr("SA Western Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Saint Pierre Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Saint Pierre Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-03:00) Saint Pierre and Miquelon"),
		// 				TimeZoneID: to.Ptr("Saint Pierre Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Sakhalin Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Sakhalin Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+11:00) Sakhalin"),
		// 				TimeZoneID: to.Ptr("Sakhalin Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Samoa Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Samoa Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+13:00) Samoa"),
		// 				TimeZoneID: to.Ptr("Samoa Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Sao Tome Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Sao Tome Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+00:00) Sao Tome"),
		// 				TimeZoneID: to.Ptr("Sao Tome Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Saratov Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Saratov Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+04:00) Saratov"),
		// 				TimeZoneID: to.Ptr("Saratov Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SE Asia Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/SE Asia Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+07:00) Bangkok, Hanoi, Jakarta"),
		// 				TimeZoneID: to.Ptr("SE Asia Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Singapore Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Singapore Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+08:00) Kuala Lumpur, Singapore"),
		// 				TimeZoneID: to.Ptr("Singapore Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("South Africa Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/South Africa Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+02:00) Harare, Pretoria"),
		// 				TimeZoneID: to.Ptr("South Africa Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Sri Lanka Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Sri Lanka Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+05:30) Sri Jayawardenepura"),
		// 				TimeZoneID: to.Ptr("Sri Lanka Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Sudan Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Sudan Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+02:00) Khartoum"),
		// 				TimeZoneID: to.Ptr("Sudan Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Syria Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Syria Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+02:00) Damascus"),
		// 				TimeZoneID: to.Ptr("Syria Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Taipei Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Taipei Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+08:00) Taipei"),
		// 				TimeZoneID: to.Ptr("Taipei Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Tasmania Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Tasmania Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+10:00) Hobart"),
		// 				TimeZoneID: to.Ptr("Tasmania Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Tocantins Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Tocantins Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-03:00) Araguaina"),
		// 				TimeZoneID: to.Ptr("Tocantins Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Tokyo Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Tokyo Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+09:00) Osaka, Sapporo, Tokyo"),
		// 				TimeZoneID: to.Ptr("Tokyo Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Tomsk Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Tomsk Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+07:00) Tomsk"),
		// 				TimeZoneID: to.Ptr("Tomsk Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Tonga Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Tonga Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+13:00) Nuku'alofa"),
		// 				TimeZoneID: to.Ptr("Tonga Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Transbaikal Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Transbaikal Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+09:00) Chita"),
		// 				TimeZoneID: to.Ptr("Transbaikal Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Turkey Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Turkey Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+03:00) Istanbul"),
		// 				TimeZoneID: to.Ptr("Turkey Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Turks And Caicos Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Turks And Caicos Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-05:00) Turks and Caicos"),
		// 				TimeZoneID: to.Ptr("Turks And Caicos Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Ulaanbaatar Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Ulaanbaatar Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+08:00) Ulaanbaatar"),
		// 				TimeZoneID: to.Ptr("Ulaanbaatar Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("US Eastern Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/US Eastern Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-05:00) Indiana (East)"),
		// 				TimeZoneID: to.Ptr("US Eastern Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("US Mountain Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/US Mountain Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-07:00) Arizona"),
		// 				TimeZoneID: to.Ptr("US Mountain Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("UTC"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/UTC"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC) Coordinated Universal Time"),
		// 				TimeZoneID: to.Ptr("UTC"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("UTC+12"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/UTC+12"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+12:00) Coordinated Universal Time+12"),
		// 				TimeZoneID: to.Ptr("UTC+12"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("UTC+13"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/UTC+13"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+13:00) Coordinated Universal Time+13"),
		// 				TimeZoneID: to.Ptr("UTC+13"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("UTC-02"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/UTC-02"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-02:00) Coordinated Universal Time-02"),
		// 				TimeZoneID: to.Ptr("UTC-02"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("UTC-08"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/UTC-08"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-08:00) Coordinated Universal Time-08"),
		// 				TimeZoneID: to.Ptr("UTC-08"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("UTC-09"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/UTC-09"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-09:00) Coordinated Universal Time-09"),
		// 				TimeZoneID: to.Ptr("UTC-09"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("UTC-11"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/UTC-11"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-11:00) Coordinated Universal Time-11"),
		// 				TimeZoneID: to.Ptr("UTC-11"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Venezuela Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Venezuela Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC-04:00) Caracas"),
		// 				TimeZoneID: to.Ptr("Venezuela Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Vladivostok Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Vladivostok Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+10:00) Vladivostok"),
		// 				TimeZoneID: to.Ptr("Vladivostok Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Volgograd Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Volgograd Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+04:00) Volgograd"),
		// 				TimeZoneID: to.Ptr("Volgograd Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("W. Australia Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/W. Australia Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+08:00) Perth"),
		// 				TimeZoneID: to.Ptr("W. Australia Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("W. Central Africa Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/W. Central Africa Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+01:00) West Central Africa"),
		// 				TimeZoneID: to.Ptr("W. Central Africa Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("W. Europe Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/W. Europe Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+01:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna"),
		// 				TimeZoneID: to.Ptr("W. Europe Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("W. Mongolia Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/W. Mongolia Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+07:00) Hovd"),
		// 				TimeZoneID: to.Ptr("W. Mongolia Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("West Asia Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/West Asia Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+05:00) Ashgabat, Tashkent"),
		// 				TimeZoneID: to.Ptr("West Asia Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("West Bank Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/West Bank Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+02:00) Gaza, Hebron"),
		// 				TimeZoneID: to.Ptr("West Bank Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("West Pacific Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/West Pacific Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+10:00) Guam, Port Moresby"),
		// 				TimeZoneID: to.Ptr("West Pacific Standard Time"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Yakutsk Standard Time"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/timeZones"),
		// 			ID: to.Ptr("/subscriptions/37d5e605-6142-4d79-b564-28b6dbfeec0f/providers/Microsoft.Sql/locations/onebox/timeZones/Yakutsk Standard Time"),
		// 			Properties: &armsql.TimeZoneProperties{
		// 				DisplayName: to.Ptr("(UTC+09:00) Yakutsk"),
		// 				TimeZoneID: to.Ptr("Yakutsk Standard Time"),
		// 			},
		// 	}},
		// }
	}
}
