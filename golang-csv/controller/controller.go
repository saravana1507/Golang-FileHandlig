package controller

import (
	"golang-csv/modal"
	"golang-csv/service"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Userservice service.Userservice
}

func NewController(c service.Userservice) UserController {
	return UserController{Userservice: c}
}

var ErrGroup = map[string]error{
	"fullDetails":     nil,
	"employeeDetails": nil,
	"managerDetails":  nil,
}

func (control *UserController) FindAll(ctx *gin.Context) {
	var wg sync.WaitGroup
	for key, value := range modal.Data {
		wg.Add(1)
		go func(key string, value modal.Responder) {
			defer wg.Done()
			err := control.Userservice.Getvalue(key, value)
			ErrGroup[key] = err
		}(key, value)
	}
	for name, val := range ErrGroup {
		if val == nil {
			ctx.JSON(http.StatusOK, gin.H{name: "file retrive successfully"})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{name: val.Error()})
		}
	}
	wg.Wait()
}

func (control *UserController) Save(ctx *gin.Context) {
	var wg sync.WaitGroup
	for key, value := range modal.Data {
		wg.Add(1)
		go func(key string, value modal.Responder) {
			defer wg.Done()
			err := control.Userservice.Setvalue(key, value)
			ErrGroup[key] = err

		}(key, value)
	}
	for name, val := range ErrGroup {
		if val == nil {
			ctx.JSON(http.StatusOK, gin.H{name: "file upload successfully"})

		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{name: val.Error()})

		}
	}
	wg.Wait()

}
