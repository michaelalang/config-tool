all: 

generate: 
	go run ./utils/generate/gen.go

install:

	go install ./...

clean: 
	rm ./pkg/lib/fieldgroups/*