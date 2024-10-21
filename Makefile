STEAMPIPE_INSTALL_DIR ?= ~/.steampipe
BUILD_TAGS = netgo
install:
	go build -o $(STEAMPIPE_INSTALL_DIR)/plugins/hub.steampipe.io/plugins/turbot/env0@latest/steampipe-plugin-env0.plugin -tags "${BUILD_TAGS}" *.go
