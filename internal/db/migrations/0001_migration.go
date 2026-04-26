package migrations

import (
	"context"

	"github.com/MikeBooon/coliseum/internal/db/dao"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		// Migration up
		var err error

		// Enums
		db.Exec(`CREATE TYPE user_type AS ENUM ('tenant', 'client')`)

		// RBAC
		_, err = db.NewCreateTable().Model((*dao.Permission)(nil)).Exec(ctx)
		if err != nil {
			return err
		}
		_, err = db.NewCreateTable().Model((*dao.Role)(nil)).Exec(ctx)
		if err != nil {
			return err
		}
		_, err = db.NewCreateTable().Model((*dao.RolePermission)(nil)).
			ForeignKey(`(role_id) REFERENCES roles(id) ON DELETE CASCADE`).
			ForeignKey(`(permission_id) REFERENCES permissions(id) ON DELETE CASCADE`).Exec(ctx)

		// Tenant
		_, err = db.NewCreateTable().Model((*dao.Tenant)(nil)).Exec(ctx)
		if err != nil {
			return err
		}

		// User
		_, err = db.NewCreateTable().Model((*dao.User)(nil)).
			ForeignKey(`(tenant_id) REFERENCES tenants(id) ON DELETE CASCADE`).
			ForeignKey(`(role_id) REFERENCES roles(id) ON DELETE CASCADE`).Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		// Migration down
		_, err := db.NewDropTable().Model((*dao.User)(nil)).Exec(ctx)
		_, err = db.NewDropTable().Model((*dao.Tenant)(nil)).Exec(ctx)
		return err
	})
}
