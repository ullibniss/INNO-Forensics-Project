data "twc_presets" "preset-2-4-50" {
  location = "ru-1"
  disk_type = "nvme"
  disk = 1024 * 80
  cpu = 4
  cpu_frequency = 3.3
  ram = 1024 * 8
}

data "twc_os" "os" {
  name = "ubuntu"
  version = "22.04"
}