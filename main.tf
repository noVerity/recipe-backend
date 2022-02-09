variable "name" {
  type = string
}

variable "FOODDATA_TOKEN" {
  type = string
}

variable "heroku_region" {
  type    = string
  default = "eu"
}

locals {
  recipe_app_name = "${var.name}-recipe-app"
}

terraform {
  backend "pg" {}
  required_providers {
    heroku = {
      source  = "heroku/heroku"
      version = "4.6.0"
    }
  }
}

provider "heroku" {
}

resource "random_password" "password" {
  length           = 16
  special          = true
  override_special = "_%@"
}

resource "random_password" "forager_token" {
  length           = 16
  special          = true
  override_special = "_%@"
}

resource "heroku_app" "gateway" {
  name   = "${local.recipe_app_name}-gateway"
  region = var.heroku_region

  config_vars = {
    GIN_MODE         = "release",
    APP_USER_SERVICE = heroku_app.user.web_url,
    APP_RECIPE_SHARDS = jsonencode({
      "shards" = [
        {
          "name" : "one",
          "url" : heroku_app.recipe_one.web_url
        },
        {
          "name" : "two",
          "url" : heroku_app.recipe_two.web_url
        }
      ]
    }),
  }
  sensitive_config_vars = {
    JWT_SECRET = random_password.password.result
  }
}

resource "heroku_app" "user" {
  name   = "${local.recipe_app_name}-user"
  region = var.heroku_region

  config_vars = {
    GIN_MODE = "release",
    APP_RECIPE_SHARDS = jsonencode({
      "shards" = [
        {
          "name" : "one",
          "url" : heroku_app.recipe_one.web_url
        },
        {
          "name" : "two",
          "url" : heroku_app.recipe_two.web_url
        }
      ]
    }),
  }
  sensitive_config_vars = {
    JWT_SECRET = random_password.password.result
  }
}

resource "heroku_app" "recipe_one" {
  name   = "${local.recipe_app_name}-recipe-one"
  region = var.heroku_region

  config_vars = {
    GIN_MODE        = "release",
    "APP_IN_QUEUE"  = "ingredients_results",
    "APP_OUT_QUEUE" = "ingredients_lookup"
  }
  sensitive_config_vars = {
    JWT_SECRET = random_password.password.result
  }
}

resource "heroku_app" "recipe_two" {
  name   = "${local.recipe_app_name}-recipe-two"
  region = var.heroku_region

  config_vars = {
    GIN_MODE        = "release",
    "APP_IN_QUEUE"  = "ingredients_results",
    "APP_OUT_QUEUE" = "ingredients_lookup"
  }
  sensitive_config_vars = {
    JWT_SECRET = random_password.password.result
  }
}

resource "heroku_addon" "user_db" {
  app  = heroku_app.user.name
  plan = "heroku-postgresql:hobby-dev"
}

resource "heroku_addon" "recipe_one_db" {
  app  = heroku_app.recipe_one.name
  plan = "heroku-postgresql:hobby-dev"
}

resource "heroku_addon" "recipe_two_db" {
  app  = heroku_app.recipe_two.name
  plan = "heroku-postgresql:hobby-dev"
}

resource "heroku_addon" "mq" {
  app  = heroku_app.recipe_one.name
  plan = "cloudamqp:lemur"
}

resource "heroku_build" "gateway" {
  app        = heroku_app.gateway.id
  buildpacks = ["https://github.com/heroku/heroku-buildpack-go.git"]

  source {
    path = "gateway"
  }
}

resource "heroku_build" "user" {
  app        = heroku_app.user.id
  buildpacks = ["https://github.com/heroku/heroku-buildpack-go.git"]

  source {
    path = "user"
  }
}

resource "heroku_build" "recipe_one" {
  app        = heroku_app.recipe_one.id
  buildpacks = ["https://github.com/heroku/heroku-buildpack-go.git"]

  source {
    path = "api"
  }
}

resource "heroku_build" "recipe_two" {
  app        = heroku_app.recipe_two.id
  buildpacks = ["https://github.com/heroku/heroku-buildpack-go.git"]

  source {
    path = "api"
  }
}

resource "heroku_formation" "gateway" {
  app        = heroku_app.gateway.id
  type       = "web"
  quantity   = 1
  size       = "free"
  depends_on = [heroku_build.gateway]
}

resource "heroku_formation" "user" {
  app        = heroku_app.user.id
  type       = "web"
  quantity   = 1
  size       = "free"
  depends_on = [heroku_build.user]
}

resource "heroku_formation" "recipe_one" {
  app        = heroku_app.recipe_one.id
  type       = "web"
  quantity   = 1
  size       = "free"
  depends_on = [heroku_build.recipe_one]
}

resource "heroku_formation" "recipe_two" {
  app        = heroku_app.recipe_two.id
  type       = "web"
  quantity   = 1
  size       = "free"
  depends_on = [heroku_build.recipe_two]
}

resource "heroku_addon_attachment" "mq_recipe_two" {
  app_id   = heroku_app.recipe_two.id
  addon_id = heroku_addon.mq.id
}

resource "heroku_app" "forager" {
  name   = "${local.recipe_app_name}-forager"
  region = var.heroku_region

  config_vars = {
    "PYTHON_RUNTIME_VERSION" = "3.8.11",
    "POETRY_VERSION"         = "1.1.0",
    "WEB_CONCURRENCY"        = "3",
    "APP_OUT_QUEUE"          = "ingredients_results",
    "APP_IN_QUEUE"           = "ingredients_lookup"
  }
  sensitive_config_vars = {
    "APP_TOKEN"          = random_password.forager_token.result,
    "APP_FOODDATA_TOKEN" = var.FOODDATA_TOKEN
  }
}

resource "heroku_addon_attachment" "mq" {
  app_id   = heroku_app.forager.id
  addon_id = heroku_addon.mq.id
}

resource "heroku_build" "forager" {
  app        = heroku_app.forager.id
  buildpacks = ["https://github.com/moneymeets/python-poetry-buildpack.git", "https://github.com/heroku/heroku-buildpack-python.git"]

  source {
    path = "forager"
  }
}

resource "heroku_formation" "forager" {
  app        = heroku_app.forager.id
  type       = "web"
  quantity   = 1
  size       = "free"
  depends_on = [heroku_build.forager]
}