package api

import (
	"log"
	"net/http"

	"github.com/auth0-community/go-auth0"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/patrickchap/clipsapi/db/sqlc"
	"github.com/patrickchap/clipsapi/util"
	jose "gopkg.in/square/go-jose.v2"
)

type Server struct {
	store db.Store
	router *gin.Engine
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


	v1 := router.Group("/api/v1")
	{

		/* VIDEOS */
		v1.GET("/video/:id", server.getVideo)
		v1.GET("/video", server.getListVideo)
		v1.GET("/video/user/:user_id", server.getUserVideoList)
		v1.POST("/video", authRequired(), server.createVideo)
		v1.POST("/user", server.addUser)

		/* COMMENTS */
		/* v1.GET("/videos/:video_id/comments", server.)
		v1.POST("/videos/:video_id/comments", server.)
		v1.PUT("/videos/:video_id/comments/:comment_id", server.)
		v1.DELETE("/videos/:video_id/comments/:comment_id", server.) */

		/* LIKES */
		/* v1.GET("/videos/:video_id/likes", server.)
		v1.POST("/videos/:video_id/likes", server.)
		v1.PUT("/videos/:video_id/likes/:like_id", server.)
		v1.DELETE("/videos/:video_id/likes/:like_id", server.) */
	}

	server.router = router 
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}


func authRequired() gin.HandlerFunc {
    return func(c *gin.Context) {

        var auth0Domain = "https://" + domain + "/"
        client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: auth0Domain + ".well-known/jwks.json"}, nil)
        configuration := auth0.NewConfiguration(client, []string{audience}, auth0Domain, jose.RS256)
        validator := auth0.NewValidator(configuration, nil)

        _, err := validator.ValidateRequest(c.Request)

        if err != nil {
            log.Println(err)
            terminateWithError(http.StatusUnauthorized, "token is not valid", c)
            return
        }
        c.Next()
    }
}

func terminateWithError(statusCode int, message string, c *gin.Context) {
    c.JSON(statusCode, gin.H{"error": message})
    c.Abort()
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
