resource "twc_server" "web-server" {
  name = "Web server"
  os_id = data.twc_os.os.id
  preset_id = data.twc_presets.preset-1-2-30.id
  cloud_init = templatefile("./cloud-init/server.yaml", {
    password = random_password.web_server_password.result
  })


}

resource "twc_floating_ip" "web-floating-ip" {
  availability_zone = "spb-3"

  comment = "Web floating ip"

  resource {
    type = "server"
    id   = twc_server.web-server.id
  }
}