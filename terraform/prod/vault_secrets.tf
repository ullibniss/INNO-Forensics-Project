resource "vault_mount" "secret" {
  path = "secret"
  type = "kv"

  options = {
    version = "2"
  }
}

resource "random_password" "web_server_password" {
  length           = 24
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>?"
}

resource "vault_kv_secret_v2" "web_server" {
  mount = vault_mount.secret.path
  name  = "prod/web_server"

  data_json = jsonencode({
    ip       = twc_floating_ip.web-floating-ip.ip
    user     = "ullibniss"
    password = random_password.web_server_password.result
  })
}

