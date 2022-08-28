package cases

import (
	unitOfWork "github.com/edg956/pixer/internal/infrastructure/uow"
)

type command interface{}

type result interface{}

type useCase[C command, R result] func(cmd C, uow unitOfWork.UnitOfWork) (R, error)

func executeInTransaction[C command, R result, U useCase[C, R]](uc U, cmd C, uow unitOfWork.UnitOfWork) (R, error) {
	uow.Begin()
	res, err := uc(cmd, uow)

	if err != nil {
		uow.Rollback()
		return res, err
	}

	uow.Commit()

	return res, nil
}
