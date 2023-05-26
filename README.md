#gin-api backend for a blog web app

Restful CRUD API Inplemented through GIN, POSTGRES, godotenv, GORM, Gomail

#how to use this backend repository (not on docker)

Step #1: create your own .evn file, and set up below environemnt variables

DB_URL (for your own POSTGRES DB_URL, will be used for the connection of the GROM )
SECRET (your secret string, will be used for JWT to encode your jwt token and then be sent to client end as httponly cookie)
EMAIL_PD (your email application password, will be used to access your email provider SMTP server)
EMAIL (your email address to send emails to your subscribers of your blog website)

Step #2: RUN "go mod download" to install any needed dependencies

Step #3: RUN "go run migrate/migrate.go" to migrate the database

Step #4: RUN "go run main.go" to start your backend server, the server will start on port 8080

#how to use this backend repository ( on docker)

Step #1: create your own .evn file, and set up below environemnt variables

DB_URL (for your own POSTGRES DB_URL, will be used for the connection of the GROM )
SECRET (your secret string, will be used for JWT to encode your jwt token and then be sent to client end as httponly cookie)
EMAIL_PD (your email application password, will be used to access your email provider SMTP server)
EMAIL (your email address to send emails to your subscribers of your blog website)

Step #2: RUN "docker build -t my-blog-backend ." to build docker image

Step #3: Run "docker run -p 8080:8080 -d my-blog-backend" to run the docker image