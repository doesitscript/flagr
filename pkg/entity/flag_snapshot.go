//go:generate goqueryset -in flag_snapshot.go

package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// FlagSnapshot is the snapshot of a flag
// Any change of the flag will create a new snapshot
// gen:qs
type FlagSnapshot struct {
	gorm.Model
	FlagID    uint `gorm:"index:idx_flagsnapshot_flagid"`
	UpdatedBy string
	Flag      *Flag `sql:"type:text"`
}

// Scan implements scanner interface
func (f *Flag) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	if b, ok := value.([]byte); ok {
		err := json.Unmarshal(b, f)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("Cannot scan %v into Flag type", value)
}

// Value implements valuer interface
func (f *Flag) Value() (driver.Value, error) {
	bytes, err := json.Marshal(f)
	if err != nil {
		return nil, err
	}
	return string(bytes), nil
}

// SaveFlagSnapshot saves the Flag Snapshot
func SaveFlagSnapshot(db *gorm.DB, flagID uint, updatedBy string) {
	tx := db.Begin()

	f := &Flag{}
	q := NewFlagQuerySet(tx).IDEq(flagID)
	if err := q.One(f); err != nil {
		logrus.WithFields(logrus.Fields{
			"err":    err,
			"flagID": flagID,
		}).Error("failed to find the flag when SaveFlagSnapshot")
		return
	}
	f.Preload(tx)

	fs := FlagSnapshot{FlagID: f.ID, UpdatedBy: updatedBy, Flag: f}
	if err := tx.Create(&fs).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"err":    err,
			"flagID": f.ID,
		}).Error("failed to save FlagSnapshot")
		tx.Rollback()
		return
	}

	if err := q.GetUpdater().SetUpdatedBy(updatedBy).SetSnapshotID(fs.ID).Update(); err != nil {
		logrus.WithFields(logrus.Fields{
			"err":            err,
			"flagID":         f.ID,
			"flagSnapshotID": fs.ID,
		}).Error("failed to save Flag's UpdatedBy and SnapshotID")
		tx.Rollback()
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
	}
}
