import React from 'react';

type CircleChartProps = {
  percent: number;
};

const CircleChart: React.FC<CircleChartProps> = ({ percent }) => {
    return (
      <div className="circle-chart" style={{ "--percent": percent } as React.CSSProperties}>
        <svg>
          {/* Background circle */}
          <circle className="background" cx="60" cy="60" r="50"></circle>
          {/* Progress circle */}
          <circle className="progress" cx="60" cy="60" r="50"></circle>
        </svg>
        <span>{percent}%</span>
      </div>
    );
  };
  

export default CircleChart;