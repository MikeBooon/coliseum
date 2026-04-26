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
		_, err = db.NewCreateTable().Model((*dao.Tenant)(nil)).Exec(ctx)
		if err != nil {
			return err
		}
		_, err = db.NewCreateTable().Model((*dao.User)(nil)).
			ForeignKey(`(tenant_id) REFERENCES tenants(id) ON DELETE CASCADE`).Exec(ctx)
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
