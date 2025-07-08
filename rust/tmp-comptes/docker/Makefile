build:
	docker build -t pim .

run:
	docker run -d --name pcon -p 9001:5432 pim

run_with_vol:
	docker run -d --name pcon -p 9001:5432 -v my_postgres -v my_postgres_volume:/var/lib/postgresql/data/ pim

rm:
	docker stop pcon && docker rm pcon
