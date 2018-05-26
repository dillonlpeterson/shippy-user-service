package go_micro_srv_user

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// BeforeCreate Hooks into GORMS event lifecycle (Called before new entity is saved).
// Notice, we don't need to manually manage database connections here.
func (model *User) BeforeCreate(scope *gorm.Scope) error {
	randUUID := uuid.Must(uuid.NewV4())
	/*if err != nil {
		return err
	}*/
	return scope.SetColumn("Id", randUUID.String())
}
