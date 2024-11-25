export interface SystemStats {
    cpu: CPUStats;
    memory: MemoryStats;
    disk: DiskStats;
    network: NetworkStats;
    system_info: SystemInfo;
  }
  
  export interface CPUStats {
    usage_percent: number;
    cores: number;
  }
  
  export interface MemoryStats {
    total: number;
    used: number;
    free: number;
    used_percent: number;
  }
  
  export interface DiskStats {
    total: number;
    used: number;
    free: number;
  }
  
  export interface NetworkStats {
    upload_speed: number;
    download_speed: number;
  }
  
  export interface SystemInfo {
    os: string;
    uptime: string;
    hostname: string;
  }
  