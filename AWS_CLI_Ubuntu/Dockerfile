# Base Image
FROM ubuntu:18.04

# Owner details
LABEL maintainer='manoj.meparambu@gmail.com'
LABEL build_date='2019-03-11'

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

# Set environment variables
ENV HOME /root
