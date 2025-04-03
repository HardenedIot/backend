package db

import (
	"log"

	"gorm.io/gorm"
)

func CreateInstance(db *gorm.DB, model interface{}) error {
	if err := db.Create(model).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateInstance(db *gorm.DB, model interface{}) error {
	if err := db.Save(model).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteInstance(db *gorm.DB, model interface{}) error {
	if err := db.Delete(model).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func FindInstance(db *gorm.DB, model interface{}, key string, value interface{}) (interface{}, error) {
	if err := db.Where(key+" = ?", value).First(model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}

	return model, nil

}
