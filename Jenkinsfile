pipeline {
    agent any

    stages {
        stage('Build') {
            stages {
                stage('Compile') {
                    steps {
                        echo 'Compiling...'
                        sleep 10
                    }
                }
                stage('Package') {
                    steps {
                        echo 'Packaging...'
                        sleep 5
                    }
                }
            }
        }
        

        stage('New stage 3-edited') {
            steps {
                echo 'Deploying...'
                sleep 3
            }
        }

        stage('Registering build artifact') {
            steps {
                echo 'Registering the metadata'
                echo 'Another echo to make the pipeline a bit more complex'
                registerBuildArtifactMetadata(
                    name: "Internal-demo-runs-BT",
                    version: "1.0.1",
                    type: "docker",
                    url: "http://localhost:4001",
                    digest: "6f637064707039346163663761",
                    label: "artifact"
                )
            }
        }

         stage('New stage 2') {
            steps {
                echo 'Deploying...'
                sleep 2
            }
        }

        stage('Test') {
            steps {
                echo 'Running Unit Tests...'
                sleep 10
                echo 'Running Integration Tests...'
                sleep 5
            }
        }

        stage('Deploy') {
            steps {
                echo 'Deploying...'
                sleep 5
            }
        }

        stage('New stage 1 - edits') {
            steps {
                echo 'Deploying...'
                sleep 1
            }
        }

        stage('Final Status') {
            steps {
                script {
                    echo "Marking build as UNSTABLE."
                    currentBuild.result = 'UNSTABLE'
                }
            }
        }
    }
}
