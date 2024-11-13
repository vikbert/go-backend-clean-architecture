package bootstrap

import "github.com/amitshekhariitbhu/go-backend-clean-architecture/mongo"

type Application struct {
	Env   *Env
	Mongo mongo.Client
}

// App initializes and returns a new Application instance.
// The function sets up the environment and MongoDB connection for the application.
//
// The function performs the following steps:
// 1. Creates a new Application instance.
// 2. Initializes the environment using NewEnv().
// 3. Establishes a MongoDB connection using NewMongoDatabase() with the initialized environment.
// 4. Returns the initialized Application instance.
func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
