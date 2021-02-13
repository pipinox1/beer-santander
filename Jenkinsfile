pipeline {
    agent {
            docker {
                image 'golang:1.15-alpine'
            }
    }
    environment {
           GOCACHE = "/tmp"
       }
   stages {
        stage('Install Dependencies') {
           steps {
               sh 'go mod tidy'
           }
       }
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