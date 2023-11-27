package migration_20231107_0000

import (
	"github.com/google/uuid"
	. "github.com/nexodus-io/nexodus/internal/database/migrations"
)

type Organization struct {
}

type VPC struct {
	SecurityGroupId uuid.UUID
}

type SecurityGroup struct {
	VpcId uuid.UUID
}

func init() {
	migrationId := "20231107-0000"
	CreateMigrationFromActions(migrationId,
		AddTableColumnsAction(&SecurityGroup{}),
		// TODO
		// This is destructive, as it's not copying the default sec group from the org to the VPC.
		// We do not plan to deploy anything prior to this to prod, so it's fine. We should probably
		// just flatten the migrations one more time before we promote to prod.
		DropTableColumnAction(&Organization{}, "security_group_id"),
		AddTableColumnsAction(&VPC{}),
	)
}
