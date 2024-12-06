export interface Metric {
  original: number;
  value: string;
  unit: string;
}

export interface CPU {
  usage_percent: Metric;
  cores: Metric;
}

export interface Memory {
  total: Metric;
  used: Metric;
  free: Metric;
  used_percent: Metric;
}

export interface Disk {
  total: Metric;
  used: Metric;
  free: Metric;
}

export interface Network {
  upload_speed: Metric;
  download_speed: Metric;
}

export interface SystemInfo {
  os: string;
  uptime: string;
  hostname: string;
}

export interface Metrics {
  cpu: CPU;
  memory: Memory;
  disk: Disk;
  network: Network;
  system_info: SystemInfo;
}