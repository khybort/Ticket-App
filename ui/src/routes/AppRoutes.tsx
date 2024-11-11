import React from 'react';
import { Routes, Route } from 'react-router-dom';
import Home from '../pages/Home';
import CreateTicket from '../pages/CreateTicket';
import TicketListPage from '../pages/TicketListPage';
import ErrorPage from '../pages/ErrorPage';

const AppRoutes: React.FC = () => {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/create-ticket" element={<CreateTicket />} />
      <Route path="/ticket-list" element={<TicketListPage />} />
      <Route path="*" element={<ErrorPage />} />
    </Routes>
  );
};

export default AppRoutes;
