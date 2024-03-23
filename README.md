# Go GIN API with Mongo

Boilerplate microservice API with GIN

# Folder structure


| Folder & Files                    | Description                                                                           |
| --------------------------------- | ------------------------------------------------------------------------------------- |
| /src/routers                      | AKA endpoints                                                                         |
| /src/routers/users                | router folder example                                                                 | 
| /src/routers/users/handler.go     | Handler functions                                                                     |
| /src/routers/users/router.go      | Router                                                                                |
| /src/routers/users/validations.go | Endpoint Request Validations                                                          |
| /src/middlewares                  | Custom Middlewares                                                                    |
| /src/services                     | Business logic, including internal or external services such as SDK (service layer)   |
| /src/configs                      | Environment variables and configuration-related things                                |
