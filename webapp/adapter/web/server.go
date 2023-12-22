package web

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/waffleboot/app/port/usecase"
)

type Option interface {
	apply(*config) error
}

type optionFunc func(*config) error

func (f optionFunc) apply(s *config) error {
	return f(s)
}

type Server struct {
	srv *http.Server
}

type config struct {
	router *gin.Engine
	port   int
}

func NewServer(opts ...Option) (*Server, error) {
	c := &config{router: gin.New()}

	for _, opt := range opts {
		err := opt.apply(c)
		if err != nil {
			return nil, err
		}
	}

	s := &Server{srv: &http.Server{
		Addr:    net.JoinHostPort("", strconv.Itoa(c.port)),
		Handler: c.router,
	}}

	return s, nil
}

func WithHttpPort(port int) Option {
	return optionFunc(func(s *config) error {
		if port <= 0 {
			return fmt.Errorf("bad http port port=%d", port)
		}
		s.port = port
		return nil
	})
}

func WithCreateUseCase(svc usecase.CreateFileUseCase) Option {
	return optionFunc(func(c *config) error {
		c.router.POST("/upload", func(c *gin.Context) {
			data, err := io.ReadAll(c.Request.Body)
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}

			id, err := svc.CreateFile(c.Request.Context(), data)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}

			c.JSON(http.StatusOK, gin.H{"id": id})
		})
		return nil
	})
}

func WithGetFileUseCase(svc usecase.GetFileUseCase) Option {
	return optionFunc(func(c *config) error {
		c.router.GET("/file/:id", func(c *gin.Context) {
			strID := c.Param("id")

			id, err := uuid.Parse(strID)
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}

			data, err := svc.GetFile(c.Request.Context(), id)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}

			_, _ = c.Writer.Write(data)
		})
		return nil
	})
}

func WithGetFilesUseCase(svc usecase.ListFilesUseCase) Option {
	return optionFunc(func(c *config) error {
		c.router.GET("/", func(c *gin.Context) {
			ids, err := svc.GetFiles(c.Request.Context())
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}
			fmt.Fprint(c.Writer, `<html><body><ul><li><a href="test">test</a></li>`)
			for _, id := range ids {
				fmt.Fprintf(c.Writer, `<li><a href="file/%s">%s</a></li>`, id, id)
			}
			fmt.Fprint(c.Writer, "</ul></body></html>")
		})
		return nil
	})
}

func WithTestUseCase(svc usecase.TestUseCase) Option {
	return optionFunc(func(c *config) error {
		c.router.GET("/test", func(c *gin.Context) {
			err := svc.Test(c.Request.Context())
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}
			fmt.Fprint(c.Writer, "ok")
		})
		return nil
	})
}

func (s *Server) Start() error {
	err := s.srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and server: %w", err)
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	err := s.srv.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}
	return nil
}
