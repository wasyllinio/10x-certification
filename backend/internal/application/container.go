package application

import (
	"10x-certification/internal/config"
	authCommand "10x-certification/internal/domain/auth/command"
	authQuery "10x-certification/internal/domain/auth/query"
	chargersCommand "10x-certification/internal/domain/chargers/command"
	chargersQuery "10x-certification/internal/domain/chargers/query"
	locationsCommand "10x-certification/internal/domain/locations/command"
	locationsQuery "10x-certification/internal/domain/locations/query"
	locationsService "10x-certification/internal/domain/locations/service"
	"10x-certification/internal/infrastructure/auth/jwt"
	"10x-certification/internal/infrastructure/auth/password"
	"10x-certification/internal/infrastructure/http/handlers"
	"10x-certification/internal/infrastructure/persistence/postgres"
	"10x-certification/internal/infrastructure/persistence/postgres/repository"
)

// Container - DI Container
//
// Ten plik odpowiada za:
// 1. Inicjalizację wszystkich dependencies w odpowiedniej kolejności
// 2. Wiring up wszystkich komponentów aplikacji
// 3. Dependency injection dla wszystkich warstw
type Container struct {
	// Infrastructure dependencies
	DB             *postgres.Connection // *gorm.DB - będzie dodane po implementacji GORM
	JWTService     *jwt.TokenService
	PasswordHasher *password.Hasher

	// Repositories
	// UserRepository     interface{} // domain.UserRepository
	// ChargerRepository  interface{} // domain.ChargerRepository
	// LocationRepository interface{} // domain.LocationRepository
	// AuditRepository    interface{} // domain.AuditRepository

	// Command Handlers (Auth)
	RegisterUserHandler *authCommand.RegisterUserHandler
	LoginUserHandler    *authCommand.LoginUserHandler

	// Query Handlers (Auth)
	GetUserByIDHandler    *authQuery.GetUserByIDHandler
	GetUserByEmailHandler *authQuery.GetUserByEmailHandler

	// Command Handlers (Chargers)
	CreateChargerHandler   *chargersCommand.CreateChargerHandler
	UpdateChargerHandler   *chargersCommand.UpdateChargerHandler
	DeleteChargerHandler   *chargersCommand.DeleteChargerHandler
	AddConnectorHandler    *chargersCommand.AddConnectorHandler
	UpdateConnectorHandler *chargersCommand.UpdateConnectorHandler
	DeleteConnectorHandler *chargersCommand.DeleteConnectorHandler

	// Query Handlers (Chargers)
	GetChargerByIDHandler *chargersQuery.GetChargerByIDHandler
	ListChargersHandler   *chargersQuery.ListChargersHandler
	SearchChargersHandler *chargersQuery.SearchChargersHandler

	// Command Handlers (Locations)
	CreateLocationHandler *locationsCommand.CreateLocationHandler
	UpdateLocationHandler *locationsCommand.UpdateLocationHandler
	DeleteLocationHandler *locationsCommand.DeleteLocationHandler
	AssignChargerHandler  *locationsCommand.AssignChargerHandler
	DetachChargerHandler  *locationsCommand.DetachChargerHandler

	// Query Handlers (Locations)
	GetLocationByIDHandler     *locationsQuery.GetLocationByIDHandler
	ListLocationsHandler       *locationsQuery.ListLocationsHandler
	GetLocationChargersHandler *locationsQuery.GetLocationChargersHandler
	GetLocationEVSEHandler     *locationsQuery.GetLocationEVSEHandler

	// HTTP Handlers
	AuthHandler     *handlers.AuthHandler
	ChargerHandler  *handlers.ChargerHandler
	LocationHandler *handlers.LocationHandler
	HealthHandler   *handlers.HealthHandler
}

// NewContainer creates and initializes DI Container
func NewContainer(cfg *config.Config) *Container {
	// 1. Infrastructure layer
	jwtService := jwt.NewTokenService(cfg.JWTSecret)
	passwordHasher := password.NewHasher()

	db := postgres.Connect(cfg)

	// 2. Repository implementations
	userRepo := repository.NewUserRepository(db.DB())
	chargerRepo := repository.NewChargerRepository(db.DB())
	locationRepo := repository.NewLocationRepository(db.DB())
	evseRepo := repository.NewEVSERepository(db.DB())
	connectorRepo := repository.NewConnectorRepository(db.DB())
	// auditRepo := postgres.NewAuditRepository(db)

	// 3. Domain services
	evseGeneratorService := locationsService.NewEVSEGeneratorService(evseRepo)

	// 3. Domain handlers (Auth)
	registerUserHandler := authCommand.NewRegisterUserHandler(
		userRepo, // userRepo - będzie dodane po implementacji
		passwordHasher,
		jwtService,
	)
	loginUserHandler := authCommand.NewLoginUserHandler(
		userRepo, // userRepo - będzie dodane po implementacji
		passwordHasher,
		jwtService,
	)
	getUserByIDHandler := authQuery.NewGetUserByIDHandler(nil)
	getUserByEmailHandler := authQuery.NewGetUserByEmailHandler(nil)

	// 4. Domain handlers (Chargers)
	createChargerHandler := chargersCommand.NewCreateChargerHandler(chargerRepo, connectorRepo)
	updateChargerHandler := chargersCommand.NewUpdateChargerHandler(chargerRepo)
	deleteChargerHandler := chargersCommand.NewDeleteChargerHandler(chargerRepo)
	addConnectorHandler := chargersCommand.NewAddConnectorHandler(connectorRepo)
	updateConnectorHandler := chargersCommand.NewUpdateConnectorHandler(connectorRepo)
	deleteConnectorHandler := chargersCommand.NewDeleteConnectorHandler(connectorRepo)

	getChargerByIDHandler := chargersQuery.NewGetChargerByIDHandler(chargerRepo)
	listChargersHandler := chargersQuery.NewListChargersHandler(chargerRepo)
	searchChargersHandler := chargersQuery.NewSearchChargersHandler(chargerRepo)

	// 5. Domain handlers (Locations)
	createLocationHandler := locationsCommand.NewCreateLocationHandler(locationRepo)
	updateLocationHandler := locationsCommand.NewUpdateLocationHandler(locationRepo)
	deleteLocationHandler := locationsCommand.NewDeleteLocationHandler(locationRepo)
	assignChargerHandler := locationsCommand.NewAssignChargerHandler(locationRepo, chargerRepo, evseGeneratorService)
	detachChargerHandler := locationsCommand.NewDetachChargerHandler(locationRepo, chargerRepo)

	getLocationByIDHandler := locationsQuery.NewGetLocationByIDHandler(locationRepo)
	listLocationsHandler := locationsQuery.NewListLocationsHandler(locationRepo)
	getLocationChargersHandler := locationsQuery.NewGetLocationChargersHandler(locationRepo)
	getLocationEVSEHandler := locationsQuery.NewGetLocationEVSEHandler(evseRepo)

	// 6. HTTP Handlers
	authHandler := handlers.NewAuthHandler(
		registerUserHandler,
		loginUserHandler,
		getUserByIDHandler,
		getUserByEmailHandler,
	)
	chargerHandler := handlers.NewChargerHandler(
		createChargerHandler,
		updateChargerHandler,
		deleteChargerHandler,
		addConnectorHandler,
		updateConnectorHandler,
		deleteConnectorHandler,
		getChargerByIDHandler,
		listChargersHandler,
		searchChargersHandler,
	)
	locationHandler := handlers.NewLocationHandler(
		createLocationHandler,
		updateLocationHandler,
		deleteLocationHandler,
		assignChargerHandler,
		detachChargerHandler,
		getLocationByIDHandler,
		listLocationsHandler,
		getLocationChargersHandler,
		getLocationEVSEHandler,
	)
	healthHandler := handlers.NewHealthHandler()

	return &Container{
		DB:             db,
		JWTService:     jwtService,
		PasswordHasher: passwordHasher,

		// Auth handlers
		RegisterUserHandler:   registerUserHandler,
		LoginUserHandler:      loginUserHandler,
		GetUserByIDHandler:    getUserByIDHandler,
		GetUserByEmailHandler: getUserByEmailHandler,

		// Charger handlers
		CreateChargerHandler:   createChargerHandler,
		UpdateChargerHandler:   updateChargerHandler,
		DeleteChargerHandler:   deleteChargerHandler,
		AddConnectorHandler:    addConnectorHandler,
		UpdateConnectorHandler: updateConnectorHandler,
		DeleteConnectorHandler: deleteConnectorHandler,
		GetChargerByIDHandler:  getChargerByIDHandler,
		ListChargersHandler:    listChargersHandler,
		SearchChargersHandler:  searchChargersHandler,

		// Location handlers
		CreateLocationHandler:      createLocationHandler,
		UpdateLocationHandler:      updateLocationHandler,
		DeleteLocationHandler:      deleteLocationHandler,
		AssignChargerHandler:       assignChargerHandler,
		DetachChargerHandler:       detachChargerHandler,
		GetLocationByIDHandler:     getLocationByIDHandler,
		ListLocationsHandler:       listLocationsHandler,
		GetLocationChargersHandler: getLocationChargersHandler,
		GetLocationEVSEHandler:     getLocationEVSEHandler,

		// HTTP handlers
		AuthHandler:     authHandler,
		ChargerHandler:  chargerHandler,
		LocationHandler: locationHandler,
		HealthHandler:   healthHandler,
	}
}
