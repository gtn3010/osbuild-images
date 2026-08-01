ARG BUILDERIMAGE="golang:1.25.3"
ARG BASEIMAGE="registry.redhat.io/rhel10/bootc-image-builder:10.2-1785716433"

FROM $BUILDERIMAGE as builder

WORKDIR /build
ENV TAGS="containers_image_openpgp,exclude_graphdriver_btrfs,exclude_graphdriver_devicemapper"

COPY . .
RUN mkdir /blueprint && git clone https://github.com/gtn3010/blueprint.git --depth 1 /blueprint && \
    mkdir /fastami && git clone https://github.com/gtn3010/fastami.git --depth 1 /fastami && \
    ls && go mod tidy && VERSION="$(git rev-parse --short HEAD)" && mkdir /build/bin && \
    CGO_ENABLED=0 go build -tags "${TAGS}" -ldflags="-X main.version=${VERSION}" -o /build/bin/bootc-image-builder ./cmd/image-builder && \
    cd /fastami && go mod tidy && \
    CGO_ENABLED=0 go build -o fastami .


FROM $BASEIMAGE as base
COPY custom-osbuild-stages/org.osbuild.oscap.remediation /usr/lib/osbuild/stages/org.osbuild.oscap.remediation
COPY custom-osbuild-stages/org.osbuild.oscap.remediation.meta.json /usr/lib/osbuild/stages/org.osbuild.oscap.remediation.meta.json
RUN dnf install -y libxcrypt-compat zip openssl && \
    curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip && \
    ./aws/install
COPY --from=builder /build/bin/bootc-image-builder /usr/bin/bootc-image-builder
COPY --from=builder /fastami/fastami /usr/bin/fastami
