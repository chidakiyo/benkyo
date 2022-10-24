
ent/gen:
	go generate ./ent

connect:
	#psql -h localhost -p 15432 -U postgres
	psql "sslmode=disable dbname=example host=127.0.0.1 port=15432 user=postgres password=pass"

m/01:
	docker run \
        --name atlas-db \
        -e POSTGRES_DB=example \
        -e POSTGRES_USER=postgres \
        -e POSTGRES_PASSWORD=pass \
        -p 15432:5432 \
        -d postgres

m/02:
	atlas schema inspect -u "postgres://postgres:pass@localhost:15432/example?sslmode=disable" > schema.hcl