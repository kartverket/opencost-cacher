package database

import (
	"gorm.io/gorm"
	"time"
)

// AfterCreate trigger to update labels
func (r *Report) AfterCreate(tx *gorm.DB) error {
	var reports []Report
	err := tx.Where("name = ? AND cluster = ? AND type = ?", r.Name, r.Cluster, r.Type).
		Order("start DESC").
		Find(&reports).Error

	if err != nil {
		return err
	}

	if len(reports) == 0 {
		// No existing reports found; possibly the first one
		return nil
	}

	if r.Start.Before(reports[0].Start) {
		r.NamespaceLabels = reports[0].NamespaceLabels
		r.TeamLabel = reports[0].TeamLabel
		r.DivisionLabel = reports[0].DivisionLabel
		if err = tx.Updates(r).Error; err != nil {
			return err
		}
		return nil
	}

	updates := map[string]interface{}{
		"namespace_labels": r.NamespaceLabels,
		"team_label":       r.TeamLabel,
		"division_label":   r.DivisionLabel,
		"updated_at":       time.Now(),
	}

	if err = tx.Model(&Report{}).
		Where("name = ? AND cluster = ? AND type = ?", r.Name, r.Cluster, r.Type).
		UpdateColumns(updates).
		Error; err != nil {
		return err
	}

	return nil
}
