pb:
	  go get -u github.com/golang/protobuf/protoc-gen-go
		@echo "pb Start"
		cd configure && make pb
asset:
	mkdir assets
	cd assets;curl https://cdn.rawgit.com/v2ray/v2ray-core/e60de73c704d46d91633035e6b06184f7186a4e0/tools/release/config/geosite.dat > geosite.dat
	cd assets;curl https://cdn.rawgit.com/v2ray/v2ray-core/1777540e3d9eb7429c1ba72a93d8ef6c426bda13/release/config/geoip.dat > geoip.dat

shippedBinary:
	cd shippedBinarys; $(MAKE) shippedBinary

fetchDep:
	-go get  github.com/xiaokangwang/V2RayConfigureFileUtil
	-cd $(GOPATH)/src/github.com/xiaokangwang/V2RayConfigureFileUtil;$(MAKE) all
	go get  github.com/xiaokangwang/V2RayConfigureFileUtil
	-go get  github.com/xiaokangwang/AndroidLibV2ray
	-cd $(GOPATH)/src/github.com/xiaokangwang/libV2RayAuxiliaryURL; $(MAKE) all
	-go get  github.com/xiaokangwang/AndroidLibV2ray
	-cd $(GOPATH)/src/github.com/xiaokangwang/waVingOcean/configure; $(MAKE) pb
	go get github.com/xiaokangwang/AndroidLibV2ray

ANDROID_HOME=$(HOME)/android-sdk-linux
export ANDROID_HOME
PATH:=$(PATH):$(GOPATH)/bin
export PATH
downloadGoMobile:
	go get golang.org/x/mobile/cmd/...
	cd ~ ; curl -L https://raw.githubusercontent.com/0ee010f29436e240703012ce/AndroidLibV2ray/master/ubuntu_ndk.sh | sudo bash
	ls ~
	ls ~/android-sdk-linux/
	gomobile init -ndk ~/android-ndk-r15c;gomobile bind -v  -tags json github.com/xiaokangwang/AndroidLibV2ray
	
BuildMobile:
	@echo Stub

all: asset pb shippedBinary fetchDep
	@echo DONE
