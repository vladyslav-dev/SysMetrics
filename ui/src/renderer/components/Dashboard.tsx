import React from 'react';
import Card from './Card';
import CircleChart from './CircleChart';
import ProgressBar from './ProgressBar';
import SystemInfo from './SystemInfoCard';

type DashboardProps = {
  data: any;
};

const Dashboard: React.FC<DashboardProps> = ({ data }) => {
  return (
    <div className="dashboard">
      <Card title="CPU">
        <CircleChart percent={data.cpu.usage_percent.original} />
        <div className="stats">
          <div><span>Cores:</span> <span>{data.cpu.cores.value}</span></div>
          <div><span>Usage:</span> <span>{data.cpu.usage_percent.value}%</span></div>
        </div>
      </Card>

      <Card title="Memory">
        <CircleChart percent={data.memory.used_percent.original} />
        <div className="stats">
          <div><span>Total:</span> <span>{data.memory.total.value} {data.memory.total.unit}</span></div>
          <div><span>Used:</span> <span>{data.memory.used.value} {data.memory.used.unit}</span></div>
          <div><span>Free:</span> <span>{data.memory.free.value} {data.memory.free.unit}</span></div>
        </div>
      </Card>

      <Card title="Disk">
        <CircleChart percent={23} />
        <div className="stats">
          <div><span>Total:</span> <span>{data.disk.total.value} {data.disk.total.unit}</span></div>
          <div><span>Used:</span> <span>{data.disk.used.value} {data.disk.used.unit}</span></div>
          <div><span>Free:</span> <span>{data.disk.free.value} {data.disk.free.unit}</span></div>
        </div>
      </Card>

      <Card title="Network">
        <ProgressBar percent={18} />
        <div className="stats network-stats">
          <div><span>Upload:</span> <span>{data.network.upload_speed.value} {data.network.upload_speed.unit}</span></div>
          <div><span>Download:</span> <span>{data.network.download_speed.value} {data.network.download_speed.unit}</span></div>
        </div>
      </Card>

      <Card title="System Info">
        <SystemInfo 
          os={data.system_info.os} 
          uptime={data.system_info.uptime} 
          hostname={data.system_info.hostname} 
        />
      </Card>
    </div>
  );
};

export default Dashboard;