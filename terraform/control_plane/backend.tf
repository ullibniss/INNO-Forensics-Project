terraform {
  backend "s3" {
    endpoints = {
      s3 = "https://s3.twcstorage.ru"
    }
    region = "ru-1"
    bucket = "0c688cf3-90f7-48f1-a9d4-ca190df6c131"
    key = "control-plane.tfstate"
    access_key = "EGF2YZ54L4K0KQCEYN01"
    secret_key = "sQIIZZs4P25Ovctm9v7kNjyABD1XAVCV8j0bILN1"
    skip_region_validation = true
    skip_credentials_validation = true
    skip_metadata_api_check = true
    skip_requesting_account_id = true
  }
}