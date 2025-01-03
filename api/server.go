package api

import (
	"log"
	"net/http"

	"github.com/auth0-community/go-auth0"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/patrickchap/clipsapi/api/controllers"
	db "github.com/patrickchap/clipsapi/db/sqlc"
	"github.com/patrickchap/clipsapi/util"
	jose "gopkg.in/square/go-jose.v2"
)

type Server struct {
	store db.Store
	Router *gin.Engine
}

var (
    audience string
    domain   string
)

//Creates a new Server instance
func NewServer(store db.Store) *Server {
	configEnv, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	
	audience = configEnv.Auth0Identifier
	domain = configEnv.Auth0Domain

	server := &Server{store: store}
	router := gin.Default()	
	router.SetTrustedProxies(nil)
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{configEnv.BaseUrl} 
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	router.Use(cors.New(config))


	//controllers
	commentController := controllers.NewCommentController(store)
	videoController := controllers.NewVideoController(store)
	userController := controllers.NewUserController(store)

	v1 := router.Group("/api/v1")
	{
		/* VIDEOS */
		v1.GET("/video/:id", videoController.GetVideo)
		v1.GET("/video",videoController.GetListVideo)
		v1.GET("/video/user/:user_id", videoController.GetUserVideoList)
		v1.POST("/video", authRequired(), videoController.CreateVideo)
		v1.PUT("/video/:id", authRequired(), videoController.UpdateVideo)

		/* USER */
		v1.POST("/user", userController.AddUser)

		/* COMMENTS */
		v1.GET("/videos/:video_id/comments", commentController.GetVideoComments)
		v1.POST("/videos/:video_id/comments", authRequired(), commentController.AddVideoComments)
		v1.DELETE("/videos/:video_id/comments/:comment_id", commentController.DeleteVideoComments)
/* 		v1.PUT("/videos/:video_id/comments/:comment_id", server.) */

		/* LIKES */
		/* v1.GET("/videos/:video_id/likes", server.)
		v1.POST("/videos/:video_id/likes", server.)
		v1.PUT("/videos/:video_id/likes/:like_id", server.)
		v1.DELETE("/videos/:video_id/likes/:like_id", server.) */
	}

	server.Router = router 
	return server
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}


func ValidateClaims(ctx *gin.Context, userID string) bool {
	claims, exists := ctx.Get("claims")
	if !exists {
		ctx.JSON(http.StatusForbidden, "Claims not found in context")
		return false
	}	

	claimsMap, ok := claims.(map[string]interface{})
	if !ok {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Invalid format for claims"})
		return false
	}

	subClaim, ok := claimsMap["sub"].(string)
	if !ok {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Sub claim not found or not a string"})
		return false
	}

	if subClaim != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Only owners of this video can update"})
		return false
	}
	
	return true
}

func authRequired() gin.HandlerFunc {
    return func(c *gin.Context) {

        var auth0Domain = "https://" + domain + "/"
        client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: auth0Domain + ".well-known/jwks.json"}, nil)
        configuration := auth0.NewConfiguration(client, []string{audience}, auth0Domain, jose.RS256)
        validator := auth0.NewValidator(configuration, nil)

        token, err := validator.ValidateRequest(c.Request)

        if err != nil {
            log.Println(err)
            terminateWithError(http.StatusUnauthorized, "token is not valid", c)
            return
        }

	claims := map[string]interface{}{}
	err = validator.Claims(c.Request, token, &claims)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
		c.Abort()
		log.Println("Invalid claims:", err)
		return
	}
	

	c.Set("claims", claims)

        c.Next()
    }
}

func terminateWithError(statusCode int, message string, c *gin.Context) {
    c.JSON(statusCode, gin.H{"error": message})
    c.Abort()
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
