package container

import (
	"fmt"
	"reflect"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/heyubani/go-template/core/base"
	"github.com/heyubani/go-template/interfaces"
	"github.com/heyubani/go-template/logger"
)

// ControllerInvoker exposes functions that can be used to interact with the new instances of a controller created
// at each request
type ControllerInvoker interface {
	Call(method string) gin.HandlerFunc
}

var loggerType = reflect.TypeOf(new(interfaces.ILogger)).Elem()
var ginType = reflect.TypeOf(new(gin.Context)).Elem()

type controllerContainer struct {
	// the controller init function
	controllerInit any
	// the name of the controller init function
	controllerInitName string
	// the methods registered under the controller
	methods []string
}

func (c *controllerContainer) Call(method string) gin.HandlerFunc {
	foundMethod := false
	for i := 0; i < len(c.methods); i++ {
		if c.methods[i] == method {
			foundMethod = true
		}
	}

	if !foundMethod {
		panic(fmt.Sprintf("method `%s` not found on controller `%s`", method, c.controllerInitName))
	}

	return func(context *gin.Context) {
		// create logger
		requestId := context.GetHeader(base.REQUEST_ID_HEADER_NAME)
		logger := logger.NewLogger(map[string]interface{}{base.REQUEST_ID_NAME: requestId})

		valueOfInitializer := reflect.ValueOf(c.controllerInit)

		// create the controller
		controller := valueOfInitializer.Call([]reflect.Value{reflect.ValueOf(logger)})[0]

		// invoke the controller method to serve the request
		controller.MethodByName(method).Call([]reflect.Value{reflect.ValueOf(context)})
	}
}

// CreateControllerInvoker will create a transient instance of a controller that is invoked anytime a new
// http request is received. This is important because of dependency injection.
// Each controller instantiator needs a logger that has the request ID of the current request as part if it's field.
// Thus during each request flight, a new logger and controller is created.
func CreateControllerInvoker(init any) ControllerInvoker {
	fncType := reflect.TypeOf(init)

	if fncType.Kind() != reflect.Func {
		panic("the controller initializer must be a function, instead '" + fncType.Kind().String() + "' received")
	}

	if fncType.NumIn() != 1 {
		panic("the controller initializer must take in just one parameter which is the `logger.LoggerInterface`")
	}

	// check if the first input implements the logger interface
	if !fncType.In(0).Implements(loggerType) {
		panic("the input parameter to the contoller initializer must implement the `logger.LoggerInterface` interface")
	}

	initName := runtime.FuncForPC(reflect.ValueOf(init).Pointer()).Name()

	// get the return type of the controller initializer: this should be the contoller itself
	returnType := fncType.Out(0)
	var methods []string

	// Inspect the methods of the controller return type
	for i := 0; i < returnType.NumMethod(); i++ {
		method := returnType.Method(i)

		// inputs are three and have this format:
		//	`func(*controllers.TemplateController,  *gin.Context)`
		// Since Golang makes sure returned Method's Type and Func fields describe a function whose first argument is the receiver,

		if method.Type.NumIn() == 2 {
			firstParam := method.Type.In(1)
			// checking if the function implements `type HandlerFunc func( *gin.Context)``
			if firstParam.Kind() == reflect.Pointer && firstParam.Elem() == ginType {
				methods = append(methods, method.Name)
			}

		}
	}

	if len(methods) == 0 {
		panic("no registered 'http.HandlerFunc' on controller `" + returnType.String() + "`")
	}

	return &controllerContainer{
		controllerInit:     init,
		controllerInitName: initName,
		methods:            methods,
	}
}
