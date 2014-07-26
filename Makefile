PROGRAM   := pushme-server
GO_FILES  := $(wildcard *.go)
PROTO_DIR := ./k12http_proto

k12http: $(GO_FILES) 
	    @echo Building "$(PROGRAM)"
		@go build
		./$(PROGRAM)

clean:
	    rm -f $(PROGRAM) $(PROTO_DIR)/$(PROGRAM).pb.go *.log

.PNONY: clean prepare deploy
