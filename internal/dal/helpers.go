package dal

import (
	"context"
	"fmt"
	"reflect"
	"washbonus/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
	"go.uber.org/zap"
)

func LogOptionalError(l *zap.SugaredLogger, module string, err error, additionalInfo ...any) {
	if err != nil {
		l.Named("dal").Named(module).Errorw(fmt.Sprintf("failed to perform database action, error : %s", err), additionalInfo...)
	}
}

func ConstructUpdateMap(model interface{}) map[string]interface{} {
	updateMap := make(map[string]interface{})
	v := reflect.ValueOf(model)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i).Interface()
		fieldType := t.Field(i)
		dbTag := fieldType.Tag.Get("db")

		if dbTag == "" {
			continue
		}

		if pointer, ok := fieldValue.(*string); ok {
			if pointer != nil {
				updateMap[dbTag] = *pointer
			}
		} else if pointer, ok := fieldValue.(*int64); ok {
			if pointer != nil {
				updateMap[dbTag] = *pointer
			}
		}
	}
	return updateMap
}

func GetUserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value("userID").(string)
	return userID, ok
}

func WriteAuditLog(ctx context.Context, tx *dbr.Tx, resource dbmodels.ResourceType, entID, action string, ent interface{}) error {
	op := "failed to write audit log: %w"

	performedUser, ok := GetUserIDFromContext(ctx)
	if !ok {
		return dbmodels.ErrBadRequest
	}

	log, err := dbmodels.BuildAuditLog(resource, entID, action, performedUser, ent)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	columns := []string{"resource", "entity_id", "action", "user_performing_action"}
	values := []interface{}{log.Resource, log.EntityID, log.Action, log.UserID}

	if log.Data != nil && len(*log.Data) > 0 {
		columns = append(columns, "data")
		values = append(values, *log.Data)
	}

	_, err = tx.InsertInto("audit_logs").
		Columns(columns...).
		Values(values...).
		ExecContext(ctx)

	if err != nil {
		return fmt.Errorf(op, err)
	}

	return nil
}
