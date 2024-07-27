gen-api:
	@echo "Generating API..."
	@protoc --go_out=pkg/grpc --go_opt=paths=source_relative \
				--go-grpc_out=pkg/grpc --go-grpc_opt=paths=source_relative \
				api/*.proto
	@echo "API generated successfully"

image:
	docker build -f build/Dockerfile -t pokecalc .

run:
	docker run -d -p 8080:8080 --name pokecalc pokecalc

try:
	@grpcurl -plaintext -d '{"name": "ピカチュウ"}' localhost:8080 damage.DamageCalc.Attack
