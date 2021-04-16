# share common command
GOPATH              := $(or $(GOPATH), $(HOME)/go)
WIRE                := $(GOPATH)/bin/wire
GENNY               := $(GOPATH)/bin/genny
GOSTATIC            := CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s"
GO_TEST_PARALLEL    := go test -parallel $(shell nproc) -count=1
GO_TEST_INTEGRATION := go test -tags integration -p 1 -count=1
SOPS_DWH            := sops --encrypt --gcp-kms sops --encrypt --gcp-kms projects/tiki-dwh/locations/global/keyRings/dinf-sops/cryptoKeys/dinf-sops-key


$(WIRE):
	GOPATH=$(GOPATH) go install -mod=mod github.com/google/wire/cmd/wire

$(GENNY):
	cd /tmp && GOPATH=$(GOPATH) go get github.com/cheekybits/genny
