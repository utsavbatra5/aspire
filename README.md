# aspire
The aim of the project is to enable users to create loan requests and pay emis for them

# Running the project
go build . -t aspire

go run aspire

Also, Dockerfile is provided to run the code on docker.

# Database Choice
I have used postgres DB. Since our application needs high level of data consistency, we needed ACID properties of SQL db to support that. 

# Authentication
User authentication is handled by JWT token. User ID is encoded in JWT token for security purpose. Following are few tokens for testing APIs:

Token for UserID 1:

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJc3N1ZXIiOiJBc3BpcmUiLCJJc3N1ZWRBdCI6IjIwMjItMDUtMDMgMTQ6MzQ6MjAuNTI0Mzk3IiwiRXhwaXJhdGlvbiI6IjIwMjItMDYtMDIgMTQ6MzQ6MjAuNTI0Mzk3IiwiVXNlcklEIjoiMSJ9.svnIdxJsjx1kNSRSv1FvZVM_klLUTOJW0FmRYzgov4w


Token for UserID 2:

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJc3N1ZXIiOiJBc3BpcmUiLCJJc3N1ZWRBdCI6IjIwMjItMDUtMDMgMTQ6MzQ6MjAuNTI0Mzk3IiwiRXhwaXJhdGlvbiI6IjIwMjItMDYtMDIgMTQ6MzQ6MjAuNTI0Mzk3IiwiVXNlcklEIjoiMiJ9._UfJxCuaqLfyF5qdS7u1FDkcytx16TUHbjmaTncJL4I


Token for UserID 3:

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJc3N1ZXIiOiJBc3BpcmUiLCJJc3N1ZWRBdCI6IjIwMjItMDUtMDMgMTQ6MzQ6MjAuNTI0Mzk3IiwiRXhwaXJhdGlvbiI6IjIwMjItMDYtMDIgMTQ6MzQ6MjAuNTI0Mzk3IiwiVXNlcklEIjoiMyJ9.XCtwutK1q-ZKD6k3Fy474nC_lm8_EC124a-CsV-2J9c


Token for UserID 4:

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJc3N1ZXIiOiJBc3BpcmUiLCJJc3N1ZWRBdCI6IjIwMjItMDUtMDMgMTQ6MzQ6MjAuNTI0Mzk3IiwiRXhwaXJhdGlvbiI6IjIwMjItMDYtMDIgMTQ6MzQ6MjAuNTI0Mzk3IiwiVXNlcklEIjoiNCJ9.9TvnPL57lDDldGb4wqtJDZGHVtoD30Vz4he8STcJqLk



# API details
POSTMAN collection is shared over email.

# Unit testing
All functions are covered by unit test cases.
