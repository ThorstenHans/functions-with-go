#!/bin/bash
# login
az login

# list subscriptions
az account list -o table

# set active subscription
az account set --subscription <SUBSCRIPTION_ID>

# create an Azure Resource Group 
az group create -n rg-functions-with-go \
  -l germanywestcentral

# create an Azure Storage Account (required for Azure Functions App)
az storage account create -n safnwithgo2021 \
  -g rg-functions-with-go \
  -l germanywestcentral

# create an Azure Functions App
az functionapp create -n fn-with-go \
  -g rg-functions-with-go \
  --consumption-plan-location germanywestcentral \
  --os-type Linux \
  --runtime custom \
  --functions-version 3 \
  --storage-account safnwithgo2021
