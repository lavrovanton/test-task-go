package controller

import (
	"net/http"
	"slices"
	"strconv"
	"test-task-go/internal/controller/request"
	"test-task-go/internal/model"
	"test-task-go/pkg"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type ServiceRepository interface {
	Fetch(pagination *request.PaginationService) error
	GetById(id uint64) (model.Service, error)
	Store(m *model.Service) error
	DeleteById(id uint64) error
}

type ServiceController struct {
	repo ServiceRepository
}

var defaultPagination request.PaginationService

func init() {
	defaultPagination.AddSort("id")
	defaultPagination.AddSort("created_at")
	defaultPagination.AddFilter("type", request.ServiceTypes[:])

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("type_validation", func(fl validator.FieldLevel) bool {
			value := fl.Field().Interface().(string)
			return slices.Contains(request.ServiceTypes[:], value)
		})

		v.RegisterValidation("payment_type_validation", func(fl validator.FieldLevel) bool {
			value := fl.Field().Interface().(string)
			return slices.Contains(request.ServicePaymentTypes[:], value)
		})
	}
}

func NewServiceController(repo ServiceRepository) *ServiceController {
	return &ServiceController{repo}
}

// @Summary get services
// @Tags services
// @Accept json
// @Produce json
// @Param limit query integer false "limit"
// @Param page query integer false "page"
// @Param sort_field query string false "sort field" Enums(id, created_at)
// @Param sort_order query string false "sort order" Enums(asc, desc)
// @Param filter_field query string false "filter field" Enums(type)
// @Param filter_value query string false " filter value" Enums(VDS, Dedicated_Server, Hosting))
// @Success 200 {array} request.PaginationService
// @Failure 400 "{ "message": "param is not valid" }"
// @Failure 401 "{ "message": "authorization failed" }"
// @Router /services [get]
func (c *ServiceController) Index(ctx *gin.Context) {
	pagination := defaultPagination

	err := ctx.BindQuery(&pagination)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": pkg.ErrBadParamInput.Error()})
		return
	}

	err = c.repo.Fetch(&pagination)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": pkg.ErrBadParamInput.Error()})
		return
	}

	ctx.JSON(http.StatusOK, pagination)
}

// @Summary get service by id
// @Tags services
// @Accept json
// @Produce json
// @Param id path integer true "id"
// @Success 200 {array} model.Service
// @Failure 400 "{ "message": "param is not valid" }"
// @Failure 401 "{ "message": "authorization failed" }"
// @Failure 404 "{ "message": "resource not found" }"
// @Router /services/{id} [get]
func (c *ServiceController) Get(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": pkg.ErrBadParamInput.Error()})
		return
	}

	service, err := c.repo.GetById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, service)
}

// @Summary create service
// @Tags services
// @Accept json
// @Produce json
// @Param request body request.CreateService true "service creation request"
// @Success 200 ""
// @Failure 400 "{ message: "param is not valid" }"
// @Failure 401 "{ "message": "authorization failed" }"
// @Router /services [post]
func (c *ServiceController) Create(ctx *gin.Context) {
	request := request.CreateService{}

	err := ctx.BindJSON(&request)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	model := request.ToModel()

	if err = c.repo.Store(&model); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary delete service by id
// @Tags services
// @Accept json
// @Produce json
// @Param id path integer true "id"
// @Success 200 ""
// @Failure 400 "{ message: "param is not valid" }"
// @Failure 401 "{ "message": "authorization failed" }"
// @Failure 404 "{ "message": "resource not found" }"
// @Router /services/{id} [delete]
func (c *ServiceController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": pkg.ErrBadParamInput.Error()})
		return
	}

	if err = c.repo.DeleteById(id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
