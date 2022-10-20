package middlewares

//func ProductAuthorization() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		db := database.ConnectDB()
//		productId, err := strconv.Atoi(c.Param("productId"))
//		if err != nil {
//			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
//				"error":   "Bad Request",
//				"message": "Invalid parameter",
//			})
//			return
//		}
//		userData := c.MustGet("userData").(jwt.MapClaims)
//		userID := uint(userData["id"].(float64))
//		Product := models.Product{}
//		err = db.Select("user_id").First(&Product, uint(productId)).Error
//
//		if err != nil {
//			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
//				"error":   "Bad Request",
//				"message": "data doesn't exist",
//			})
//			return
//		}
//		if Product.UserID != userID {
//			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
//				"error":   "Unauthorized",
//				"message": "you are not allowed to access this data",
//			})
//			return
//		}
//		c.Next()
//	}
//}
