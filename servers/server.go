package servers

import (
	"dns/domains/maths"
	"dns/requests"
	"dns/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ms maths.IMathService

// Server a basic server configuration with its router embedded
type Server struct {
	Router *gin.Engine
}

// New creates a new instance of the server with the routers configured.
func New(sectorID float64) *Server {
	ms = maths.New(sectorID)
	router := gin.Default()
	router.POST("calculate", calculate)

	return &Server{
		Router: router,
	}
}

func calculate(ctx *gin.Context) {
	var req *requests.Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	respType := ctx.Query("resp")
	domainReq, err := req.ToDomainRequest()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	result := ms.Calculate(ctx.Request.Context(), domainReq)
	var res interface{}
	switch respType {
	case "mom":
		res = responses.MomCorp{
			Location: result,
		}

	default:
		res = responses.Basic{
			Location: result,
		}
	}

	ctx.JSON(http.StatusOK, res)
}
