package user

import (
	"github.com/li1553770945/openmcp-gateway/biz/internal/do"
	"github.com/li1553770945/openmcp-gateway/biz/internal/domain"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	err := db.AutoMigrate(&do.UserDO{})
	if err != nil {
		panic("迁移用户模型失败：" + err.Error())
	}
	return &UserRepoImpl{
		DB: db,
	}
}

func (Repo *UserRepoImpl) FindUserByUsername(username string) (*domain.UserEntity, error) {
	var user do.UserDO
	err := Repo.DB.Where("username = ?", username).Limit(1).Find(&user).Error
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, nil
	}
	return DoToEntity(&user), nil
}

func (Repo *UserRepoImpl) FindUserById(userId int64) (*domain.UserEntity, error) {
	var user do.UserDO
	err := Repo.DB.Where("id = ?", userId).Limit(1).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return DoToEntity(&user), nil
}

func (Repo *UserRepoImpl) SaveUser(userEntity *domain.UserEntity) error {
	userDO := EntityToDo(userEntity)
	if userDO.ID == 0 {
		err := Repo.DB.Create(&userDO).Error
		return err
	} else {
		err := Repo.DB.Omit("CreatedAt", "DeletedAt").Save(&userDO).Error
		return err
	}
}
