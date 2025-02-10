package config

import (
	"application/controller"
	"application/interfaces"
	"application/repository"
	"application/service"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Server estrutura principal do servidor HTTP
type Server struct {
	server       *http.Server
	Logger       *log.Logger
	userService  interfaces.UserService
	authService  interfaces.AuthService
	tokenService interfaces.TokenService
}

func initializeLogger() *log.Logger {
	return log.New(log.Writer(), "app: ", log.LstdFlags)
}

func initializeDatabase(logger *log.Logger) *pgxpool.Pool {
	connString := "host=postgres port=5432 user=bem-te-vi password=bemtevi dbname=bem-te-vi sslmode=disable"
	dbPool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		logger.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	return dbPool
}

func InitializeServer() *Server {

	logger := initializeLogger()

	dbPool := initializeDatabase(logger)

	userRepo := repository.NewUserRepository(dbPool)

	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userService)
	tokenService := service.NewTokenService()

	return NewServer(":8080", logger, userService, authService, tokenService)
}

// NewServer construtor que inicializa um novo servidor HTTP
func NewServer(
	addr string,
	logger *log.Logger,
	userService interfaces.UserService,
	authService interfaces.AuthService,
	tokenService interfaces.TokenService,
) *Server {

	// Criar o roteador chi
	router := chi.NewRouter()

	// Basic CORS settings
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:9000"}, // Or "*"
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Adicionar middlewares do chi
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Criar o servidor HTTP com o roteador chi
	serve := &Server{
		server: &http.Server{
			Addr:         addr,   // e.g. ":8080"
			Handler:      router, // roteador chi
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
		},
		Logger:       logger,
		userService:  userService,
		authService:  authService,
		tokenService: tokenService,
	}

	// Adicionar rotas
	serve.registerRoutes(router)

	return serve
}

// registerRoutes adiciona as rotas ao roteador
func (s *Server) registerRoutes(router *chi.Mux) {
	controller.NewUserController(router, s.userService)
	controller.NewAuthController(router, s.authService, s.tokenService)
}

// Start inicia o servidor HTTP
func (s *Server) Start() error {
	s.Logger.Printf("Iniciando o servidor na porta %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.Logger.Fatalf("Erro ao iniciar o servidor: %v", err)
		return err
	}
	return nil
}

// Stop finaliza o servidor HTTP
func (s *Server) Stop(ctx context.Context) error {
	s.Logger.Println("Parando o servidor")
	return s.server.Shutdown(ctx)
}
