## pubsub-example

Example of how to use GCP Pub/Sub.

## Usage

To publish one message:

    godotenv go run pub/pub.go

To start pull subscriber in background:

    godotenv go run sub-pull/pull.go &

To start push subscriber in background:

    godotenv go run sub-push/push.go &

To deploy push subscriber to GAE:

    cd sub-push
    gcloud app deploy

You can now access the application at https://[YOUR_PROJECT_ID].appspot.com
