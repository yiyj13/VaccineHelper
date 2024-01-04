package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"v-helper/internal/model"
	"v-helper/internal/service"
	"v-helper/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileService *service.ProfileService
}

func NewProfileHandler(profileService *service.ProfileService) *ProfileHandler {
	return &ProfileHandler{profileService: profileService}
}

func (h *ProfileHandler) HandleCreateProfile(c *gin.Context) {
	// 读取加密的字符串数据
	var encryptedData struct {
		Data string `json:"data"`
	}
	if err := c.ShouldBindJSON(&encryptedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// 解密数据
	decryptedData, err := utils.Decrypt(encryptedData.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decrypt data"})
		return
	}

	// 解析JSON到Profile结构体
	var profile model.Profile
	if err := json.Unmarshal([]byte(decryptedData), &profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile data"})
		return
	}

	// 创建Profile
	if err := h.profileService.CreateProfile(profile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("Profile created successfully: ", profile)
	c.JSON(http.StatusOK, gin.H{"message": "profile created successfully"})
}

func (h *ProfileHandler) HandleGetAllProfiles(c *gin.Context) {
	profiles, err := h.profileService.GetAllProfiles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	EncryptedJSON(c, http.StatusOK, profiles)
}

func (h *ProfileHandler) HandleGetProfileByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	profile, err := h.profileService.GetProfileByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	EncryptedJSON(c, http.StatusOK, profile)
}

func (h *ProfileHandler) HandleUpdateProfileByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, exists := c.Get("parsedData")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Profile data not found"})
		return
	}

	if err := h.profileService.UpdateProfileByID(uint(id), profile.(model.Profile)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println("profile updated successfully: ", profile)
	c.JSON(http.StatusOK, gin.H{"message": "profile updated successfully"})
}

func (h *ProfileHandler) HandleDeleteProfileByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.profileService.DeleteProfileByID(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println("profile deleted successfully: ", id)
	c.JSON(http.StatusOK, gin.H{"message": "profile deleted successfully"})
}

// GetProfilesByUserID 处理通过 UserID 获取 Profile 的请求
func (h *ProfileHandler) HandleGetProfilesByUserID(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	profiles, err := h.profileService.GetProfilesByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching profiles"})
		return
	}

	EncryptedJSON(c, http.StatusOK, profiles)
}
