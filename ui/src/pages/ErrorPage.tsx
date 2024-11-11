import React from 'react';

const ErrorPage: React.FC = () => {
  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100 text-center">
      <div className="bg-white p-6 rounded-lg shadow-md">
        <h1 className="text-4xl font-semibold text-red-500 mb-4">
          Oops! Something went wrong.
        </h1>
        <p className="text-lg text-gray-700">
          <p>This isn&apos;t the page you&apos;re looking for.</p>
        </p>
      </div>
    </div>
  );
};

export default ErrorPage;
