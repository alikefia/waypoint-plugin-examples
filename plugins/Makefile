all:

	@echo ""
	@echo "### Build Go Builder Plugin"
	$(MAKE) -C gobuilder_final
	
	@echo ""
	@echo ""
	@echo "### Build Filepath Builder Plugin"
	$(MAKE) -C filepath

	@echo ""
	@echo ""
	@echo "### Install Plugins"
	mkdir -p ./example_app/.waypoint/plugins

	mkdir -p  ${HOME}/.config/waypoint/plugins
	cp ./gobuilder_final/bin/* ${HOME}/.config/waypoint/plugins
	cp ./filepath/bin/* ${HOME}/.config/waypoint/plugins