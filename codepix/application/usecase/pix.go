package usecase

import (
	"github.com/MelkdeSousa/codepix/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (pix *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := pix.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(account, kind, key)

	if err != nil {
		return nil, err
	}

	pix.PixKeyRepository.RegisterPixKey(pixKey)

	if pixKey.ID == "" {
		return nil, err
	}

	return pixKey, nil
}

func (pix *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := pix.PixKeyRepository.FindKeyByKind(key, kind)

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
