# Base Image
FROM ubuntu:18.04

# Owner details
LABEL maintainer='manoj.meparambu@gmail.com'
LABEL build_date='2018-10-03'

# Install OS necessary, python3, pip and awscli.
RUN \
  apt-get update && apt-get install -y \
  build-essential \
  apt-utils \
  bash \
  git \
  curl \
  python3 \
  python3-pip \
  python-setuptools \
  && pip3 --no-cache-dir install --upgrade awscli

# Install npm
RUN \
  apt-get install -y npm

# Install npm
RUN \
  npm i npm@latest -g

RUN \
    curl -sL https://deb.nodesource.com/setup_10.x | bash -

# Install NodeJS
RUN \
  apt-get install -y nodejs

# Install angular-Cli
RUN \
   npm install -g @angular/cli


# create angular project folder
RUN \
    ng new oouve

# Add material package
 RUN \
    cd oouve \
    && npm install bootstrap@4.3.1

# Set environment variables
ENV HOME /root

# Set path for nodejs
ENV PATH $PATH:/nodejs/bin

WORKDIR /oouve/src

EXPOSE 4200
