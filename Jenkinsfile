pipeline {
    agent {
            docker {
                image 'golang:1.15-alpine'
            }
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
               sh 'go test ./... -v'
           }
       }
    }
}