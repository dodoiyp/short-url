package mysqldb

type SequenceImp interface {
	CreateSequenceAndGetID() (int, error)
}

type Sequence struct {
	ID int `gorm:"size:20;uniqueIndex"`
}

func (mydb *mysqlDBObj) CreateSequenceAndGetID() (int, error) {
	seq := Sequence{}

	err := mydb.DB.Create(&seq).Error
	if err != nil {
		return 0, err
	}
	return seq.ID, nil
}
