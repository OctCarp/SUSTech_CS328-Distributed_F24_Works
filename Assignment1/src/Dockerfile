# Use the official Ubuntu base image
# Latest version of Ubuntu is 24.04 in 2024 fall, so we specify it as 22.04
FROM ubuntu:22.04

# Set environment variables to prevent interactive prompts during package installations
ENV DEBIAN_FRONTEND=noninteractive

# Update the package list and install necessary packages

# Change the source of apt-get to Tsinghua University
RUN sed -i s:/archive.ubuntu.com:/mirrors.tuna.tsinghua.edu.cn/ubuntu:g /etc/apt/sources.list && \
    sed -i s:/security.ubuntu.com:/mirrors.tuna.tsinghua.edu.cn/ubuntu:g /etc/apt/sources.list && \
    cat /etc/apt/sources.list && \
    apt-get clean && \
    apt-get -y update --fix-missing

# Install the necessary packages
RUN apt-get install -y \
    build-essential \
    openssh-server \
    openmpi-bin \
    openmpi-common \
    libopenmpi-dev \
    nano \
    sudo && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# set root user passwd to 'root'
RUN echo 'root:root' | chpasswd

# Create the mpi user with password 'mpi' (mpi does not like running as root so we need a user)
RUN useradd -rm -d /home/mpi -s /bin/bash -g root -G sudo -u 1001 mpi && \
    echo 'mpi:mpi' | chpasswd

# Set up SSH configuration for the mpi user (used for mpi to connect to other contaiers for distributed execution)
RUN mkdir -p /home/mpi/.ssh && \
    ssh-keygen -q -t rsa -N '' -f /home/mpi/.ssh/id_rsa && \
    cat /home/mpi/.ssh/id_rsa.pub >> /home/mpi/.ssh/authorized_keys && \
    # Set ownership to mpi user and correct permissions
    chown -R mpi:root /home/mpi/.ssh && \
    chmod 700 /home/mpi/.ssh && \
    chmod 600 /home/mpi/.ssh/authorized_keys && \
    echo "StrictHostKeyChecking no" >> /home/mpi/.ssh/config

# Set the working directory
WORKDIR /home/mpi 
