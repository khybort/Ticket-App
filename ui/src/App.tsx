import React from 'react';
import AppRoutes from './routes/AppRoutes';
import Footer from './components/Footer/Footer';

const App: React.FC = () => {
  return (
    <div className="flex flex-col min-h-screen">
      <AppRoutes />
      <Footer />
    </div>
  );
};

export default App;
