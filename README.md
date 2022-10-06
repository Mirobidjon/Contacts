# Contacts

Go RESTful API Contacts
-------------------------

**_This project includes the following_**
  1. RESTful endpoints in the widely accepted format
  2. Standard CRUD operations of a database table
  3. JWT-based authentication
  4. Environment dependent application configuration management
  5. Structured logging with contextual information
  6. Error handling with proper error response generation
  7. Database migration
  8. Data validation 
 <br>
<p> About Project
<br>
 <br> Everyone create new account and sign-in project
 <br>  All users 
   <br>  1. Create
    <br> 2. Read
 <br>    3. Update
  <br>   4. Delete 
   <br> themselves contacts
 
 </p>
  
## Deployment

The application can be run as a docker container. You can use `make build-image` to build the application 
into a docker image. 

```shell
make build-image PROJECT_NAME=${PROJECT_NAME} APP=${SERVICE_NAME} TAG=${TAG}
```

If you want to push this image to registry you can use following command
```shell
make push-image
```

You can also run `go mod vendor && make build` to build an executable binary named `contact`. Then start the API server using the following
command,

```shell
./bin/contact
```

In additional you can deploy to GOOGLE CLOUD APP ENGINE

```shell
make deploy PROJECT_NAME=${PROJECT_NAME} ENV_TAG=${TAG}
```

### And you can use the online on my gcp application [link](https://contact-list-dot-learn-cloud-0809.uw.r.appspot.com/swagger/index.html)