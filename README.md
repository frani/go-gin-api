# Go GIN API with Mongo

Boilerplate microservice API with GIN

# Folder structure

```sh
| Folder               | description                                                                            |
| -------------------- | -------------------------------------------------------------------------------------- |
| /presentations       | Group of endpoints. here can find folders with endpoints.                              |
| - /:example          | example folder for endpoint                                                            |
| - / - /handler.go    | Handler functions                                                                      |
| - / - /routes.go     | Router configuration                                                                   |
| - / - /validator.go  | Endpoint Request Validator                                                             |
| /middlewares         | Custom gin middlewares                                                                 |
| /services            | Business logic, incluye intertal or externaa servicies as SDK (service layer)          |
| /config              | Environment variables and configuration related things                                 |
```
