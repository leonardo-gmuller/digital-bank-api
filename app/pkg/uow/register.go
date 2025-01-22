package uow

func (u *Uow) Register(name string, fc RepositoryFactory) {
	u.Repositories[name] = fc
}
