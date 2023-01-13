export GO111MODULE=on

DOORAY_BOOT_IMAGE_NAME="deepdiveinwinter/dooraybot"
DOORAY_BOOT_IMAGE_VERSION="latest"
DOORAY_HOOK_URL="https://hook.dooray.com/services/{{SERVICE_HOOK}}"

.PHONY: build
build:
	docker build -t ${DOORAY_BOOT_IMAGE_NAME}:${DOORAY_BOOT_IMAGE_VERSION} .

.PHONY: run
run:
	docker run --rm -e DOORAY_HOOK_URL=${DOORAY_HOOK_URL} --name dooray-bot ${DOORAY_BOOT_IMAGE_NAME}:${DOORAY_BOOT_IMAGE_VERSION}
