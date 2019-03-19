# Base Image
FROM ubuntu:18.04

# Owner details
LABEL maintainer='manoj.meparambu@gmail.com'
LABEL build_date='2019-11-03'

# Install OS necessary, python3, pip and awscli.
RUN \
  apt-get update && apt-get install -y \
#  build-essential \
  apt-utils \
  bash \
  git \
  curl

RUN \
  apt-get install -y software-properties-common

# Install Anisble
Run \
  apt-add-repository ppa:ansible/ansible

RUN \
  apt-get update && apt-get install -y ansible
