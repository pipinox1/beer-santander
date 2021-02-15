pipeline {
    agent {
            docker {
                image 'golang:1.15-alpine'
            }
    }
    environment {
           GOCACHE  = '/tmp'
           CGO_ENABLED= 0
           registry = "pipinox1/beer-santander"
           registryCredential = 'dockerhub'
            dockerHome = tool 'docker'
            PATH = "${docker}/bin:${PATH}"
    }
   stages {
        stage('Initialize'){
          steps {
           sh 'whereis docker'
                sh 'docker ps'
            }
        }
        stage('Install Dependencies') {
           steps {
               sh 'pwd'
               sh 'ls'
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
stage('Building image') {
      steps{
        script {
          dockerImage = docker.build registry + ":$BUILD_NUMBER"
        }
      }
    }
    stage('Deploy Image') {
      steps{
        script {
          docker.withRegistry( '', registryCredential ) {
            dockerImage.push()
          }
        }
      }
    }
    }
}