#!/bin/sh

##########
# Secrets file
#
# !!! Attention !!! Please do not commit your plain secrets.sh file!
# The k8s-jcasc.sh script supports encrypting and decrypting, which makes it much safer to use such
# a configuration.
#
# This file was written as shell script, because it is not possible to create a secrets.yaml file for
# multiple secrets. Also it is much easier to create/maintain the docker credentials like here.
##########

##########
# Docker registry credentials
# If you need to configure a private docker registry in Kubernetes, that the Jenkins is able to
# download its own containers, you can add the credentials like here.
# Please do not delete the delete command, because this kind of secret can not be updated easily.
# If no credentials are stored in the secrets map, it will show you an error, which is absolute fine!
#
# Credentials ID for containers in this example: docker-registry-credentialsId
##########
kubectl -n ${NAMESPACE} delete secret docker-registry-credentialsid
kubectl -n ${NAMESPACE} \
                create secret docker-registry \
                docker-registry-credentialsid \
                --docker-server=hub.docker.com \
                --docker-username=myUsername \
                --docker-password=myPassword \
                --docker-email=username@company.com

##########
# Credentials for the VCS access (e.g. Github/Bitbucket...)
# Later the Jenkins / your pipeline will use the metadata.name credentials Id to checkout.
##########
cat <<EOF | kubectl -n ${NAMESPACE} apply -f -
apiVersion: v1
kind: Secret
metadata:
  # this is the jenkins credentials id
  name: "github-credentialsid"
  labels:
    "jenkins.io/credentials-type": "basicSSHUserPrivateKey"
  annotations:
    "jenkins.io/credentials-description" : "basic user private key credential from Kubernetes"
type: Opaque
stringData:
  username: myGitHubUsername
  privateKey: |
    -----BEGIN RSA PRIVATE KEY-----
    MIIEpAIBAAKCAQEApYQNqZINxC/JgTS/JBklzbZITG0YxFhu97JkMoUb9GpxEuTG
    iHN2wG+eyDGpbhxFFthyIv9wzndw7eexOl9qa+E2wvxe4S+PMUhxjT3mgocav/D6
    uZbpIMQiDdg5dwqtRYtbnDGQR/rZnlz8/htqfz9L1B0rWmkGb7o7qmAQCReuUThu
    myzTGjMXA5TXd7dvq6LwVAeVxCY5gMq7xnB0ahE0dMwLHZJFsTLponpSyZCcp7/p
    EPBcNz9j2ZXLtY0EjJmWG5uAFpE08YUcPmsqT30HBTLglmRLHCrRbImYCZAy8STu
    QdhecJIuz0WWEkfaP4vQEXyAV+Z1q0hpc/vQiwIDAQABAoIBAQCVqHOaEf/0lqkC
    9SingT2XZey1fifY0YV03o6Ox6DhPaWESevh0VVc5xCRP6ZNc65c29VII+RiK8mm
    s8qiV0gox3j2ka9Queoly7Uw2vmkqHjdeQ2b5gJhqDaKgipjovnNWha5gm9NRlqj
    QL1ZHn0LFbaA0ucyVqiFOcVdZoLZVGWKt44RfbunsxOSHAv/Eo330sKAvteTi7E8
    Kt95Vrdgr2XgiRqSRXDAx0aK6bKLyCyjuRdVSDK+ibHpNs1l+COyMZ/GB1QZRZ+B
    ILXqAcuSaNXxD+zHeBMKlrAESndXVaTOVMB7L14PUije0AgSxFQt1omnAXLzveIi
    HGgibaXRAoGBANhKgJc5ofh3l7PQDaQP8VVlTWLHXq6f16rIbHerXDj1Q6dqM4ka
    GmY6Z2Tc3SjB7s2WVBqY6VgfzhiVXvDWMAMcTPi2ODCfEMfGtIzxZuDIY0kbT2Hr
    6jgzrC+Fk8Tel7Zw1+gVCRiWvq7oj5mfF6pL95me8DqONh1TuYN+L5a3AoGBAMPn
    KNElnWzssYJwZl4eGvHatXvFJLuM6HhpbljS8zBodh5VwtcJIKVU+ilJ8ij17xFL
    pgLIlysvastuVzoixqjDVZnLakPlHvHMqWOqMZkg/TXf1ze/+oRBgcrIhwZV/GHo
    I/iaAUSqCU9SaCkBAafkz04bMxa8pa8DAMZcAODNAoGAe+BXu8UPZj4gjaTIW0Gi
    R/WIF9319W+o1rCJpxRm8lxOjjD+KTThD9G9bAAvTmucOPUzYDRZ2NYGdP//61HR
    F8b6squyjO5dbv34ZIzSDkXWz4UrtvqmH+BAastHccbG/3+ruMlrd0DHH2gk7qg2
    pptxyPNFxVOz3KIaKxx3ZwsCgYEAqocV6LktiBuhiUH+Wf8qxUz0nYDGsNu/oNFl
    1LwMJR9Jcq6EpFq1qDWIbViJC07Jg+yt3c5uiJEGDX9HPrv24gDnCrEfF2rivOjC
    qpcEBZ/JypPG7CiZEXdUXAiiQMmooDFK3qRwZiz9XacGNGtD3bo3Gm5i0m/0aZvb
    mM+NlCECgYAolfzE8wJ/Y5sqXe0jskei+DU0hFiRUCSDLyllv3+zps9dJtwJW1cv
    BFwhpzXWD7YOoFPhAmJbWn10C9rjCNxSClL0ddceVUfq65AveKZr4PH1UlGSlMNo
    zyFSdpC6EfLRWEn1S0TuHuD5J/116G8JYCR2oUee2vkN0omzOJU6cQ==
    -----END RSA PRIVATE KEY-----
EOF

#########
# VCS Notification
# If you have Bitbucket or Github, you can use such an technical user to notify the system about the build.
#########
cat <<EOF | kubectl -n ${NAMESPACE} apply -f -
apiVersion: v1
kind: Secret
metadata:
  # this is the jenkins credentials id.
  name: "vcs-notification-credentialsid"
  labels:
    "jenkins.io/credentials-type": "usernamePassword"
  annotations:
    "jenkins.io/credentials-description" : "credentials from Kubernetes"
type: Opaque
stringData:
  username: vcs-notification-service-account
  password: password-for-the-service-account
EOF

#########
# Nexus/jFrog Artifactory/NPM/Azure Repositories/... credentialsId
# If you have a private repository, you can also add the credentials here and access them later with the credentialsId.
# Create as much credentialsIds as you need (e.g. read-only, read-write...)
#########
cat <<EOF | kubectl -n ${NAMESPACE} apply -f -
apiVersion: v1
kind: Secret
metadata:
  # this is the jenkins credentials id.
  name: "repository-credentialsid"
  labels:
    "jenkins.io/credentials-type": "usernamePassword"
  annotations:
    "jenkins.io/credentials-description" : "credentials from Kubernetes"
type: Opaque
stringData:
  username: repository-service-account
  password: password-for-the-service-account
EOF

#########
# Docker repository credentialsId.
#
# If you want to build containers for docker and push them (or maybe in case of a
# private repository read FROM then), you can also add these credentials as
# usernamePassword credentials type.
# The credentials here will be used while docker build or docker push commands.
# Do not confuse the Docker Registry with the Docker Repository!
#########
cat <<EOF | kubectl -n ${NAMESPACE} apply -f -
apiVersion: v1
kind: Secret
metadata:
  # this is the jenkins credentials id.
  name: "docker-repository-credentialsid"
  labels:
    "jenkins.io/credentials-type": "usernamePassword"
  annotations:
    "jenkins.io/credentials-description" : "credentials from Kubernetes"
type: Opaque
stringData:
  username: docker-repository-service-account
  password: password-for-the-service-account
EOF
