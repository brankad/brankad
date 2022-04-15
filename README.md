To run project - enter directory bdankad
docker build . -t brankad_api:v0.0.1
docker run -d --rm -p 8080:8080 --name="brankad_api" brankad_api:v0.0.1

Go to browser and type
http://localhost:8080/

To run unit test do
go test -v ./api