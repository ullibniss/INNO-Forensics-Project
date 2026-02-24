resource "twc_server" "vault-server" {
  name = "Vault server"
  os_id = data.twc_os.os.id
  preset_id = data.twc_presets.preset-2-4-50.id
  cloud_init = file("./cloud-init/vault.yml")

  local_network {
    id = twc_vpc.project-vpc.id
    ip = "192.168.0.9"
  }
}

resource "twc_floating_ip" "vault-floating-ip" {
  availability_zone = "spb-3"

  comment = "Vault floating ip"

  resource {
    type = "server"
    id   = twc_server.vault-server.id
  }
}