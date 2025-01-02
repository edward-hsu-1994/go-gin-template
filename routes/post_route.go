package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-fiber-template/domain"
	_ "go-fiber-template/domain"
	"go-fiber-template/helpers"
	"go-fiber-template/services"
)

type PostRouter struct {
	_postService *services.PostService
}

func NewPostRouter(
	postService *services.PostService) *PostRouter {
	return &PostRouter{
		_postService: postService,
	}
}

func (r *PostRouter) ConfigureRoutes(app *gin.Engine) {
	routes := app.Group("/api/v1/posts")

	routes.GET("/", r.ListPosts)
	routes.GET("/:postId", r.GetPostById)
}

// ListPosts godoc
// @Summary List posts
// @Description get posts paging list
// @Tags posts
// @Produce json
// @Success 200 {object} domain.Paging[domain.PostSummary]
// @Router /api/v1/posts [get]
func (r *PostRouter) ListPosts(c *gin.Context) {
	posts, err := r._postService.ListPosts(0, 10)

	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	c.JSON(200, posts)
}

// GetPostById godoc
// @Summary Get post by id
// @Description get post by id
// @Tags posts
// @Produce json
// @Param postId path string true "Post ID"
// @Success 200 {object} domain.Post
// @Router /api/v1/posts/{postId} [get]
func (r *PostRouter) GetPostById(c *gin.Context) {
	postId := c.Param("postId")

	post, err := r._postService.GetPostById(postId)

	if post == nil {
		helpers.ErrorResponse(c, domain.ErrorPostNotFound.New(fmt.Sprintf("Post with ID %s not found", postId)))
		return
	}

	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	c.JSON(200, post)
}
