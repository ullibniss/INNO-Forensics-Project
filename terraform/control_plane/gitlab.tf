resource "twc_server" "gitlab-server" {
  name = "Gitlab server"
  os_id = data.twc_os.os.id
  preset_id = data.twc_presets.preset-2-4-50.id
  cloud_init = file("./cloud-init/gitlab.yaml")

  local_network {
    id = twc_vpc.project-vpc.id
    ip = "192.168.0.8"
  }
}

resource "twc_floating_ip" "gitlab-floating-ip" {
  availability_zone = "spb-3"

  comment = "Gitlab floating ip"

  resource {
    type = "server"
    id   = twc_server.gitlab-server.id
  }
}