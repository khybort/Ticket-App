import React from 'react';
import TicketList from '../components/TicketList/TicketList';
import Navbar from '../components/Navbar/Navbar';

const TicketListPage: React.FC = () => {
  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar />
      <TicketList />
    </div>
  );
};

export default TicketListPage;
