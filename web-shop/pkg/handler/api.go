package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get all item groups
func (h *Handler) getAllItemGroups(c *gin.Context) {
	if itemGroups, err := h.services.ItemGroup.GetAllItemGroups(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error msg": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, itemGroups)
	}
}

// Get item group by id
func (h *Handler) getItemGroupById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	if itemGroup, err := h.services.ItemGroup.GetItemGroupById(id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error msg": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, itemGroup)
	}
}

// Get item groups by category id
func (h *Handler) getItemGroupsByCategoryId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error msg": err.Error()})
		return
	}

	if itemGroups, err := h.services.ItemGroup.GetItemGroupsByCategoryId(id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error msg": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, itemGroups)
	}
}

// Get all items
func (h *Handler) getAllItems(c *gin.Context) {
	if items, err := h.services.Item.GetAllItems(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error msg": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, items)
	}
}

// Get item by id
func (h *Handler) getItemById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error msg": err.Error()})
		return
	}

	if item, err := h.services.Item.GetItemById(id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error msg": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, item)
	}
}

// Get all items by item group id
func (h *Handler) getItemsByItemGroupId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error msg": err.Error()})
		return
	}

	if items, err := h.services.Item.GetAllItemsByItemGroupId(id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error msg": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, items)
	}
}

// Get all categories
func (h *Handler) getAllCategories(c *gin.Context) {
	if categories, err := h.services.Category.GetAllCategories(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error msg": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, categories)
	}
}

// Get category by id
func (h *Handler) getCategoryById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error msg": err.Error()})
		return
	}

	if categories, err := h.services.Category.GetCategoryById(id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error msg": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, categories)
	}
}
