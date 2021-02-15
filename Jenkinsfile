pipeline {
    agent any
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
                sh '$PAHT'
                sh 'docker ps'
            }
        }
        stage('Install Dependencies') {
        agent{
         docker {
                        image 'golang:1.15-alpine'
                    }
        }
           steps {
               sh 'pwd'
               sh 'ls'
               sh 'go mod tidy'
           }
       }
       stage('Build') {
               agent{
                docker {
                               image 'golang:1.15-alpine'
                           }
               }
           steps {
               sh 'go build'
           }
       }
       stage('Test') {
               agent{
                docker {
                               image 'golang:1.15-alpine'
                           }
               }
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