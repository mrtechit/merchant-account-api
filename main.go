package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	_ "github.com/mrtechit/merchant-account-api/docs"
	"log"
	"merchant-account-api/adapters/api"
	"merchant-account-api/adapters/sql_repository"
	"merchant-account-api/core/services"
	"merchant-account-api/middleware"
)

// @title Merchant account management
// @version 1.0
// @description System to manage Merchant and Team Member accounts.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api/v1
func main() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	//Can use embedded db if needed
	/*postgres := embeddedpostgres.NewDatabase()
	defer postgres.Stop()

	err = postgres.Start()
	if err != nil {
		fmt.Println("test", err)
	}*/

	//init DB
	fmt.Println("connecting to DB")
	db, err := sql_repository.NewPostgreSQLRepository()

	if err != nil {
		log.Fatalf("shutdown")
	}

	merchantService := services.NewMerchantService(db)
	merchantAdapter := api.NewMerchantAdapter(merchantService)
	teamMemberService := services.NewTeamMembersService(db)
	teamMemberAdapter := api.NewTeamMemberAdapter(teamMemberService)

	app := fiber.New()
	route := app.Group("api/v1")

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// Route for access token
	route.Get("/token/new", middleware.GetNewAccessToken) // create a new access tokens

	route.Get("/merchant/:merchant_code", middleware.JWTProtected(), merchantAdapter.Get)       // retrieve a merchant
	route.Post("/merchant/", middleware.JWTProtected(), merchantAdapter.Post)                   // update a merchant
	route.Put("/merchant/", middleware.JWTProtected(), merchantAdapter.Put)                     // create a merchant
	route.Delete("/merchant/:merchant_code", middleware.JWTProtected(), merchantAdapter.Delete) // delete a merchant

	route.Get("/teams/:id", middleware.JWTProtected(), teamMemberAdapter.Get)       // retrieve a team member
	route.Post("/teams/", middleware.JWTProtected(), teamMemberAdapter.Post)        // update a team member
	route.Put("/teams/", middleware.JWTProtected(), teamMemberAdapter.Put)          // create a team member
	route.Delete("/teams/:id", middleware.JWTProtected(), teamMemberAdapter.Delete) // delete a team member
	route.Get("/teams/:merchant_code/:page/:page_size", middleware.JWTProtected(),
		teamMemberAdapter.GetTeamMembersByMerchantCode) // paginated team member by merchant code

	//setupRoutes(app)
	app.Listen(":3000")

}
