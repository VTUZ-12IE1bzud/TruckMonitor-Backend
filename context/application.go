package context

type ApplicationContext interface {
	DbContext() DbContext
	DaoContext() DaoContext
	ServiceContext() ServiceContext
}

type applicationContext struct {
	dbContext  DbContext
	daoContext DaoContext
	serviceContext ServiceContext
}

func (c *applicationContext) DbContext() DbContext {
	return c.dbContext
}

func (c *applicationContext) DaoContext() DaoContext {
	return c.daoContext
}

func (c *applicationContext) ServiceContext() ServiceContext {
	return c.serviceContext
}

func NewApplicationContext(configuration Configuration) (ApplicationContext, error) {
	db, err := NewDbContext(configuration)
	if err != nil {
		return nil, err
	}
	daoContext := NewDaoContext(db)
	return &applicationContext{
		dbContext:  db,
		daoContext: daoContext,
		serviceContext: NewServiceContext(configuration.ServerConfiguration.TokenKey, daoContext),
	}, nil
}
