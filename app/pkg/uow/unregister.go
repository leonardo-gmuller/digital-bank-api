package uow

func (u *Uow) Unregister(name string) {
	delete(u.Repositories, name)
}
