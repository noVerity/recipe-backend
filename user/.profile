mkdir -p $(dirname $GOOGLE_APPLICATION_CREDENTIALS) && touch $GOOGLE_APPLICATION_CREDENTIALS;
echo $GCP_SECRET_BASE64 | base64 --decode > $GOOGLE_APPLICATION_CREDENTIALS