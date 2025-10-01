import React from 'react';

const Logo: React.FC<{ size?: 'small' | 'medium' | 'large' }> = ({ size = 'medium' }) => {
  const sizeClasses = {
    small: 'text-2xl',
    medium: 'text-4xl',
    large: 'text-6xl'
  };

  return (
    <div className={`inline-block ${sizeClasses[size]}`}>
      <span className="font-black uppercase">
        FinanceBroke
      </span>
    </div>
  );
};

export default Logo;