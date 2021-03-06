package webserver

import "github.com/kataras/iris/v12"

// Instance of webserver
type Instance struct {
	App *iris.Application
}

// URLRequest router method on app
func (objInstance Instance) URLRequest(passedPath string, resHandler iris.Handler) {
	objInstance.App.Get(passedPath, resHandler)
}

// StartInstance prepare instance data, before running
func (objInstance *Instance) StartInstance() {
	objInstance.App = iris.New()
}

// Run iris instance
func (objInstance *Instance) Run() {
	objInstance.App.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}
