data "twc_presets" "preset-1-2-30" {
  location = "ru-1"
  disk_type = "nvme"
  disk = 1024 * 30
  cpu = 1
  cpu_frequency = 3.3
  ram = 1024 * 2
}

data "twc_os" "os" {
  name = "ubuntu"
  version = "22.04"
}