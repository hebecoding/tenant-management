#!/bin/bash

# Set environment based on git branch name
if [[ $GITHUB_REF_NAME == "main" ]]; then
  echo "Setting environment to dev"
  export ENVIRONMENT="dev"
  echo "Environment set to dev"
elif [[ $GITHUB_REF_NAME == "test" ]]; then
  echo "Setting environment to test"
  export ENVIRONMENT="test"
  echo "Environment set to test"
elif [[ $GITHUB_REF_NAME == "stage" ]]; then
  echo "Setting environment to stage"
  export ENVIRONMENT="stage"
  echo "Environment set to stage"
elif [[ $GITHUB_REF_NAME == "prod" ]]; then
  echo "Setting environment to prod"
  export ENVIRONMENT="prod"
  echo "Environment set to prod"
else
  export ENVIRONMENT="local"
fi