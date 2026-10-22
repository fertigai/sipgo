FROM debian:13.7-slim
LABEL maintainer="Razvan Crainea <razvan@opensips.org>"

USER root

# Set Environment Variables
ENV DEBIAN_FRONTEND=noninteractive

ARG OPENSIPS_VERSION=4.0
ARG OPENSIPS_BUILD=releases

#install basic components
RUN apt-get -y update -qq && apt-get -y install ca-certificates

#add signing key, repository
ADD --chmod=644 https://apt.opensips.org/opensips-org.gpg /usr/share/keyrings/opensips-org.gpg
RUN echo "deb [signed-by=/usr/share/keyrings/opensips-org.gpg] https://apt.opensips.org trixie ${OPENSIPS_VERSION}-${OPENSIPS_BUILD}" >/etc/apt/sources.list.d/opensips.list

RUN apt-get -y update -qq && apt-get -y install opensips

ARG OPENSIPS_CLI=false
RUN if [ ${OPENSIPS_CLI} = true ]; then \
    apt-get -y install opensips-cli \
    ;fi

ARG OPENSIPS_EXTRA_MODULES
RUN if [ -n "${OPENSIPS_EXTRA_MODULES}" ]; then \
    apt-get -y install ${OPENSIPS_EXTRA_MODULES} \
    ;fi

RUN rm -rf /var/lib/apt/lists/*
RUN sed -i "s/^\(socket\|listen\)=udp.*5060/\1=udp:eth0:5060/g" /etc/opensips/opensips.cfg
# OpenSIPS 4.0 has no -E flag: log to stderr through the config instead
RUN sed -i "s/stderror_enabled=no/stderror_enabled=yes/g; s/syslog_enabled=yes/syslog_enabled=no/g" /etc/opensips/opensips.cfg

EXPOSE 5060/udp

ENTRYPOINT ["/usr/sbin/opensips", "-F"]
