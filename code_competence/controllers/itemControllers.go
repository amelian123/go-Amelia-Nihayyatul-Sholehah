package controllers

import (
	"code/models/payload"
	"code/usecase"
	"fmt"
	"strconv"

	"github.com/labstack/echo"
)

type ItemController interface {
	GetAllItemController(c echo.Context) error
	GetItemByIdController(c echo.Context) error
	GetItemByCategoryIdController(c echo.Context) error
	CreateItemController(c echo.Context) error
	UpdateItemByIdController(c echo.Context)
	DeleteItemByIdController(c echo.Context) error
}

type itemController struct {
	itemUsecase usecase.ItemUsecase
}

func NewItemController(itemUsecase usecase.ItemUsecase) *itemController {
	return &itemController{itemUsecase}
}

func (i *itemController) GetAllItemController(c echo.Context) error {
	item, err := i.itemUsecase.GetAllitems()
	if err != nil {
		return err
	}

	return c.JSON(200, payload.Response{
		Message: "Success get all item",
		Data:    item,
	})
}

func (i *itemController) GetItemByIdController(c echo.Context) error {
	id := c.Param("id")

	item, err := i.itemUsecase.GetItemById(id)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: fmt.Sprintf("item %s", item.Name),
		Data:    item,
	})
}

func (i *itemController) GetItemByNameController(c echo.Context) error {
	name := c.QueryParam("name")

	item, err := i.itemUsecase.GetItemByName(name)
	if err != nil {
		return err
	}

	return c.JSON(200, payload.Response{
		Message: "Success get items",
		Data:    item,
	})
}

func (i *itemController) GetItemByCategoryIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("category_id"))

	item, err := i.itemUsecase.GetItemByCategoryId(id)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success get item by id category",
		Data:    item,
	})
}

func (i *itemController) CreateItemController(c echo.Context) error {
	req := payload.CreateItem{}

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(400, "Field Cannot be empty")
	}

	item, err := i.itemUsecase.CreateItem(&req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Create Item",
		Data:    item,
	})
}

func (i *itemController) UpdateItemByIdController(c echo.Context) error {
	req := payload.UpdateItem{}

	id := c.Param("id")

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(400, "Field cannot be empty")
	}

	item, err := i.itemUsecase.UpdateItemById(id, &req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Update item",
		Data:    item,
	})
}

func (i *itemController) DeleteItemByIdController(c echo.Context) error {
	id := c.Param("id")

	_, err := i.itemUsecase.DeleteItemById(id)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	return c.JSON(200, "Success delete item")
}
