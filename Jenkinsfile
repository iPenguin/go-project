pipeline {
    agent {
        kubernetes {
            label 'go-jenkins-worker'
            idleMinutes 5
            yamlFile 'go-pod.yml'
        }
    }
    stages {
        stage('Test') {
            steps {
                sh 'go test .'
            }
        }
        stage('Build Application') {
            steps {
                sh 'go build .'
            }
        }
        stage('Build Docker Image') {
            steps {
                sh 'echo "Building docker image"'
            }
        }
        stage('Deploy') {
            steps {
                sh 'echo "Deploying to live server"'
            }
        }
    }
}
