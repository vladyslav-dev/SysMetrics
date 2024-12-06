import React from 'react';

type ProgressBarProps = {
  percent: number;
};

const ProgressBar: React.FC<ProgressBarProps> = ({ percent }) => {
  return (
    <div className="bar">
      <div className="bar-fill" style={{ width: `${percent}%` }}></div>
    </div>
  );
};

export default ProgressBar;

