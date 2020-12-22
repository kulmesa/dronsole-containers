FROM gazebo:gzserver11-focal as builder
RUN apt-get update -y && apt-get install -y --no-install-recommends \
    build-essential \
    ca-certificates \
    cmake \
    git-core \
    imagemagick \
    libgazebo11-dev \
    libjansson-dev \
    npm \
    pkg-config \
    python-is-python2 \
    && rm -rf /var/lib/apt/lists/*
RUN git clone --depth 1 --branch gzweb_1.4.1 https://github.com/osrf/gzweb.git /gzweb
WORKDIR /gzweb
RUN sed -i 's/"node-gyp": "",/"node-gyp": "3.8.0",/g' package.json
# Clear the modelList
RUN perl -i -p0e 's/var modelList.*?];/var modelList = [];/s' gz3d/src/gzgui.js
#RUN sed -i 's/var modelList = .*];/var modelList = [];/g' gz3d/src/gzgui.js
RUN . /usr/share/gazebo/setup.sh \
    && npm run deploy
# organize file structure for cleaner container
RUN sed -i "s/const staticBasePath.*$/const staticBasePath\ =\ '.\/static';/g" gzbridge/server.js
RUN sed -i "s/\.\/build\/Debug\/gzbridge/\.\/gzbridge/g" gzbridge/server.js

FROM ubuntu:20.04
# Setup timezone
RUN echo 'Etc/UTC' > /etc/timezone \
    && ln -s /usr/share/zoneinfo/Etc/UTC /etc/localtime \
    && apt-get update && apt-get install -q -y tzdata \
    && rm -rf /var/lib/apt/lists/*
# Install tools to add apt repository
RUN apt-get update && apt-get install -y --no-install-recommends \
    gnupg2 \
    lsb-release \
    libjansson4 \
    npm \
    && rm -rf /var/lib/apt/lists/*
# Install Gazebo
RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys D2486D2DD83DB69272AFE98867170598AF249743
RUN . /etc/os-release \
    && echo "deb http://packages.osrfoundation.org/gazebo/$ID-stable `lsb_release -sc` main" > /etc/apt/sources.list.d/gazebo-latest.list
RUN apt-get update -y && apt-get install -y --no-install-recommends \
    libgazebo11 \
    && rm -rf /var/lib/apt/lists/*

# Copy static files
COPY --from=builder /gzweb/http/client /gzweb/static
# Copy backend files
COPY --from=builder /gzweb/gzbridge/server.js /gzweb/server.js
COPY --from=builder /gzweb/gzbridge/build/Debug/gzbridge.node /gzweb/gzbridge.node
# TODO: copy only relevant parts
# currently node_modules contains also all build dependencies
COPY --from=builder /gzweb/node_modules /gzweb/node_modules

COPY gzweb_entrypoint.sh /
ENTRYPOINT ["/gzweb_entrypoint.sh"]

EXPOSE 8080
