PROGRAM   := pushme-server
GO_FILES  := $(wildcard *.go)
PROTO_DIR := ./k12http_proto

$(PROGRAM): $(GO_FILES) 
	    @echo Building "$(PROGRAM)"
		@go build
		./$(PROGRAM)

deploy: $(GO_FILES)
	    @echo Building "$(PROGRAM)"
		@go build
		@cp $(PROGRAM) deploy
		+@killall $(PROGRAM) &
		@cd deploy; ./$(PROGRAM) &

clean:
	    rm -f $(PROGRAM) $(PROTO_DIR)/$(PROGRAM).pb.go *.log

.PNONY: clean prepare deploy
