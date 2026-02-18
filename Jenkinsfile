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

        stage('Test') {
            steps {
                echo 'Running Unit Tests...'
                sleep 10
                echo 'Running Integration Tests...'
                sleep 5
            }
        }

         stage('New stage') {
            steps {
                echo 'Deploying...'
                sleep 3
            }
        }

        stage('Deploy') {
            steps {
                echo 'Deploying...'
                sleep 5
            }
        }

         stage('New stage 2') {
            steps {
                echo 'Deploying...'
                sleep 7
            }
        }

        stage('Final Status') {
            steps {
                script {
                    echo "Forcing FAILURE."
                    error("Build failed intentionally.")
                }
            }
        }
    }
}
