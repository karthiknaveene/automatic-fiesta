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
                stage('Package-edited') {
                    steps {
                        echo 'Packaging...'
                        sleep 2
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

        stage('Test- edited') {
            steps {
                echo 'Running Unit Tests...'
                sleep 10
                echo 'Running Integration Tests...'
                sleep 2
            }
        }

        stage('Edit stage -1 ') {
            steps {
                echo 'Deploying...'
                sleep 1
            }
        }

        stage('Final Status (Aborted)') {
            steps {
                timeout(time: 5, unit: 'SECONDS') {
                    echo "This stage will run too long and trigger ABORT."
                    sh 'sleep 20'
                }
            }
        }
    }
}
