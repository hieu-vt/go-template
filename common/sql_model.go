package common

import "time"

type SqlModel struct {
	Id        int       `json:"-" gorm:"column:id;"`
	FakeId    *UID      `json:"id" gorm:"-"`
	Status    int       `json:"status" gorm:"column:status;default=1"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (m *SqlModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.Id), dbType, 1)
	m.FakeId = &uid
}

func (sqlModel *SqlModel) PrepareForInsert() {
	now := time.Now().UTC()
	sqlModel.Id = 0
	sqlModel.Status = 1
	sqlModel.CreatedAt = now
	sqlModel.UpdatedAt = now
}

func (sqlModel *SqlModel) GetRealId() {
	if sqlModel.FakeId == nil {
		return
	}

	sqlModel.Id = int(sqlModel.FakeId.GetLocalID())
}
