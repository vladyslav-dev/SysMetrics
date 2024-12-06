import React, { useEffect, useState } from 'react';
import { Metrics } from '../types/stats.d'
import Dashboard from './components/Dashboard';

const App: React.FC = () => {
  const [stats, setStats] = useState<Metrics | null>(null);

  useEffect(() => {
    window.electronAPI.onSystemStats((data: any) => {
      if (stats === null) {
        const parsedData: Metrics = JSON.parse(data.data);

        // Update state for SystemStats
        setStats(parsedData);
      }
     
    });
  }, []);

  if (!stats) {
    return <div>Loading system metrics...</div>;
  }

  return (
    <div className='test'>
       <h1>System Metrics</h1>
       <Dashboard data={stats} />
    </div>
  );
};

export default App;
