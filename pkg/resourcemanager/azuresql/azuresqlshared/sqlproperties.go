// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlshared

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
)

// SQLServerProperties contains values needed for adding / updating SQL servers,
// wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql#Server
// also wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql#ServerProperties
type SQLServerProperties struct {

	// AdministratorLogin - Administrator username for the server. Once created it cannot be changed.
	AdministratorLogin *string

	// AdministratorLoginPassword - The administrator login password (required for server creation).
	AdministratorLoginPassword *string
}

// SQLDatabaseProperties contains values needed for adding / updating SQL servers,
// wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql#Database
// also wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql#DatabaseProperties
type SQLDatabaseProperties struct {

	// DatabaseName is the name of the database
	DatabaseName string

	// Edition - The edition of the database. The DatabaseEditions enumeration contains all the valid editions. If createMode is NonReadableSecondary or OnlineSecondary, this value is ignored.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
	//
	// ```azurecli
	// az sql db list-editions -l <location> -o table
	// ````
	//
	// ```powershell
	// Get-AzSqlServerServiceObjective -Location <location>
	// ````
	// . Possible values include: 'Web', 'Business', 'Basic', 'Standard', 'Premium', 'PremiumRS', 'Free', 'Stretch', 'DataWarehouse', 'System', 'System2', 'GeneralPurpose', 'BusinessCritical', 'Hyperscale'
	Edition v1alpha1.DBEdition
}

// SQLFailoverGroupProperties contains values needed for adding / updating SQL failover groups,
// wraps: https://github.com/Azure/azure-sdk-for-go/blob/master/services/preview/sql/mgmt/2015-05-01-preview/sql/failovergroups.go#L53
type SQLFailoverGroupProperties struct {

	// FailoverPolicy can be Automatic or Manual
	FailoverPolicy v1alpha1.ReadWriteEndpointFailoverPolicy

	// Read/Write Grace Period in minutes
	FailoverGracePeriod int32

	// Secondary server to failover to (should be in a different region)
	SecondaryServer string

	// Resource Group for the Secondary server
	SecondaryServerResourceGroup string

	// Names of Databases to add to the failover group
	DatabaseList []string
}

// SQLServerPropertiesToServer translates SQLServerProperties to ServerProperties
func SQLServerPropertiesToServer(properties SQLServerProperties) (result sql.ServerProperties) {

	result = sql.ServerProperties{
		AdministratorLogin:         properties.AdministratorLogin,
		AdministratorLoginPassword: properties.AdministratorLoginPassword,
	}

	return result
}

// SQLDatabasePropertiesToDatabase translates SQLDatabaseProperties to DatabaseProperties
func SQLDatabasePropertiesToDatabase(properties SQLDatabaseProperties) (result sql.DatabaseProperties) {

	result = sql.DatabaseProperties{
		Edition: translateDBEdition(properties.Edition),
	}

	return result
}

// translateDBEdition translates enums
func translateDBEdition(in v1alpha1.DBEdition) (result sql.DatabaseEdition) {
	switch in {
	case 0:
		result = sql.Basic
	case 1:
		result = sql.Business
	case 2:
		result = sql.BusinessCritical
	case 3:
		result = sql.DataWarehouse
	case 4:
		result = sql.Free
	case 5:
		result = sql.GeneralPurpose
	case 6:
		result = sql.Hyperscale
	case 7:
		result = sql.Premium
	case 8:
		result = sql.PremiumRS
	case 9:
		result = sql.Standard
	case 10:
		result = sql.Stretch
	case 11:
		result = sql.System
	case 12:
		result = sql.System2
	case 13:
		result = sql.Web
	default:
		result = sql.Free
	}

	return result
}

// translateFailoverPolicy translates the enum
func TranslateFailoverPolicy(in v1alpha1.ReadWriteEndpointFailoverPolicy) (result sql.ReadWriteEndpointFailoverPolicy) {
	switch in {
	case v1alpha1.FailoverPolicyAutomatic:
		result = sql.Automatic
	case v1alpha1.FailoverPolicyManual:
		result = sql.Manual
	default:
		result = sql.Automatic
	}

	return result
}
