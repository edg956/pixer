package uow

type UnitOfWork interface {
	begin()
	commit()
	rollback()
}
