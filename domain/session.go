package domain

type (
	Session struct {
		ID        int
		Raw       string
		CreatedAt int64
		DeletedAt int64
	}

	ISessionRepository interface {
		IReadOneRepository[Session]
		ICreateRepository[Session]
		IDeleteRepository[Session]
	}
)
