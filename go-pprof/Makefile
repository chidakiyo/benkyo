build:
	go build -o cut

run:
	time ./cut -f 3 -d ',' 50000_target.csv

pprof:
	go tool pprof -http=:8888 profile