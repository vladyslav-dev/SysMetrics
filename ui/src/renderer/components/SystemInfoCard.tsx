import React from 'react';

type SystemInfoProps = {
  os: string;
  uptime: string;
  hostname: string;
};

const SystemInfo: React.FC<SystemInfoProps> = ({ os, uptime, hostname }) => {
  return (
    <div className="system-info">
      <div><span>OS:</span> <span>{os}</span></div>
      <div><span>Uptime:</span> <span>{uptime}</span></div>
      <div><span>Hostname:</span> <span>{hostname}</span></div>
    </div>
  );
};

export default SystemInfo;