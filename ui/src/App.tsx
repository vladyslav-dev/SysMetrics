import React, { useEffect, useState } from 'react';
import { SystemStats } from './types/stats.d';

const App: React.FC = () => {
  const [stats, setStats] = useState<SystemStats | null>(null);

  useEffect(() => {
    window.electronAPI.onSystemStats((data: any) => {
      if (stats === null) {
        console.log("data", JSON.parse(data.data))
        const parsedData: SystemStats = JSON.parse(data.data);
        setStats(parsedData);
      }
     
    });
  }, []);

  if (!stats) {
    return <div>Loading system stats...</div>;
  }

  return (
    <div>
      <h1>System Monitor</h1>
      {/* <CPUChart usage={stats.cpu.usage_percent} cores={stats.cpu.cores} />
      <RAMChart stats={stats.memory} />
      <DiskChart stats={stats.disk} />
      <NetworkChart stats={stats.network} /> */}
      <div>
        <h2>System Info</h2>
        <p>OS: {stats.system_info.os}</p>
        <p>Uptime: {stats.system_info.uptime}</p>
        <p>Hostname: {stats.system_info.hostname}</p>
      </div>
    </div>
  );
};

export default App;
