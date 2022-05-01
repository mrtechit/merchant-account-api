# merchant-account-api

## Overview
This service is a CRUD application for managing merchant account . This service is based on hexagonal architecture with golang

### Install Go
This project is implemented using Go. If Go is not yet installed, please download and install from [here](https://golang.org/doc/install)

## Getting started
Build service using `docker-compose build`
Run service using `docker-compose up`

## API spec
Can be found under `/docs folder` or accessible at http://localhost:3000/swagger/index.html

## Authentication
All merchant and team member's API need JWT token to be sent in header

## Endpoints
merchant-acount-api service has ten endpoints.
1. `GetMerchant` 
2. `CreateMerchant` 
3. `UpdateMerchant`
4. `DeleteMerchant`
5. `GetTeamMember`
6. `CreateTeamMember`
7. `UpdateTeamMember`
8. `DeleteTeamMember`
9. `GetTeamMembersByMerchantCode`
10. `GetAccessToken`

The list of endpoints, including request and response payloads, can be found in `http://localhost:3000/swagger/index.html`.
E.g. The request payload of `GetMerchant` endpoint is a get request. 
