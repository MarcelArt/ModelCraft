# ModelCraft
 
## How to Use
1. Click "Use this template" on github
2. Clone repository to your PC
3. Refactor module name
`github.com/MarcelArt/ModelCraft` to whatever module name you want
4. If there is any error on main.go run `make swag` to generate swagger documentation
5. Create .env file based from example.env
6. Try to run project by running command `make run`

## Automatic CRUD Scaffolding
1. You can automatically scaffold basic CRUD operations by running this command
`make scaffolder model=<Your model name>`
2. Make sure model name is in TitleCase without space or underscore
3. It will crete the neccesary files
4. Next add the new SetupRoutes function to routes/route.go file

`api_routes.SetupUserRoutes(api, authMiddleware)`

Modify it to your actual generated setup route function

5. And also add your new models to function MigrateDB and DropDB in database/postgres.go file
6. Then you can migrate or drop db using `make migrate-up` and `make migrate-down`

## Automatic CRUD Scaffolding using UI
1. Enter the url in the browser http://localhost:`<PORT>`/dev-tools
2. You can also do scaffolding and migration there