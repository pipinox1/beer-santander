pipeline {
    agent {
            docker {
                image 'golang:1.15-alpine'
            }
    }
    environment {
           GOCACHE  = '/tmp'
           CGO_ENABLED= 0
           GO111MODULE='off'
      }
   stages {
       stage('Build') {
           steps {
               sh 'go build'
           }
       }
       stage('Test') {
           steps {
               sh 'go clean -cache'
               sh 'go test ./... -v'
           }
       }
    }
}