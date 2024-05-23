all :: gen dep build

build ::
	go build -x -trimpath -o mully ./cmd/mully/...

dep ::
	go mod tidy -v

pkg/mullvad_mgmt/management_interface.pb.go ::
	cd pkg/mullvad_mgmt; \
		make

pkg/mullvad_api/api_default.go ::
	cd pkg/mullvad_api; \
		make

gen :: pkg/mullvad_mgmt/management_interface.pb.go pkg/mullvad_api/api_default.go

clean ::
	rm -v mully 2>/dev/null || true
	cd pkg/mullvad_mgmt; \
		make clean
	cd pkg/mullvad_api; \
		make clean
